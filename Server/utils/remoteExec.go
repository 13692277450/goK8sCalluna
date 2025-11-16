package utils

import (
	"bytes"
	"context"
	"fmt"
	"gok8s/kubernetsServ"
	"io"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/homedir"
)

var WG sync.WaitGroup

// GetNodeLogs 在指定节点上执行journalctl命令
func GetNodeLogs(nodeName string) (string, error) {
	fmt.Println("GetNodeLogs started.......................")
	// 创建Kubernetes配置
	//Kubeconfig = flag.String("kubeconfig", "kubeconfig", "")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	config, err := rest.InClusterConfig()
	if err != nil {
		return "", fmt.Errorf("failed to get in-cluster config: %v", err)
	}

	// 创建客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", fmt.Errorf("failed to create clientset: %v", err)
	}

	// 创建执行命令的请求
	Req := clientset.CoreV1().RESTClient().Post().
		Resource("nodes").
		Name(nodeName).
		SubResource("proxy").
		Suffix("/exec").
		Param("command", "journalctl").
		Param("command", "-u").
		Param("command", "kubelet").
		Param("command", "-n").
		Param("command", "100").
		Param("command", "--no-pager")
	fmt.Println(Req)
	// 创建执行器
	executor, err := remotecommand.NewSPDYExecutor(config, "POST", Req.URL())
	if err != nil {
		return "", fmt.Errorf("failed to create executor: %v", err)
	}

	// 执行命令并捕获输出
	var stdout, stderr bytes.Buffer
	err = executor.StreamWithContext(ctx, remotecommand.StreamOptions{ //context.TODO(),
		Stdout: &stdout,
		Stderr: &stderr,
		Tty:    false,
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, stderr: %s", err, stderr.String())
	}
	fmt.Printf("Remote comand:================\n", stdout.String())
	return stdout.String(), nil
}

func CaptureExecOutput(useInCluster bool, kubeconfigPath string, namespace, podName, containerName string, command []string) (string, string, error) {
	var config *rest.Config
	var err error

	if useInCluster {
		// 集群内模式
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		kubeClient, err := kubernetes.NewForConfig(config)
		fmt.Println(kubeClient)
		config, err = rest.InClusterConfig()
		if err != nil {
			return "", "", fmt.Errorf("in-cluster config failed: %v", err)
		}
	} else {
		// 集群外模式
		if kubeconfigPath == "" {
			return "", "", fmt.Errorf("kubeconfig path is required for external mode")
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			return "", "", fmt.Errorf("kubeconfig load failed: %v", err)
		}
	}
	fmt.Println("CaptureExecOutput started.......................")

	// 创建管道
	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader, stderrWriter := io.Pipe()
	fmt.Println("stdoutReader is waiting for done............................")

	// 检查配置是否有效
	if config == nil {
		return "", "", fmt.Errorf("kubernetes config is not initialized")
	}
	fmt.Println("config is waiting for done............................")

	// 准备请求
	Req := kubernetsServ.Clientset.CoreV1().RESTClient().Post().
		Resource("nodes"). //pods
		Name("k8s-master01").
		Namespace(namespace).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Container: containerName,
			Command:   command,
			Stdin:     false,
			Stdout:    true,
			Stderr:    true,
			TTY:       false,
		}, scheme.ParameterCodec)

	// 创建执行器
	executor, err := remotecommand.NewSPDYExecutor(config, "POST", Req.URL())
	if err != nil {
		return "", "", err
	}
	fmt.Println("executor is waiting for done............................")

	// 启动goroutine读取输出
	var stdoutBuf, stderrBuf bytes.Buffer
	var stdoutErr, stderrErr error
	WG.Add(2)

	go func() {
		defer WG.Done()
		_, stdoutErr = io.Copy(&stdoutBuf, stdoutReader)
	}()

	go func() {
		defer WG.Done()
		_, stderrErr = io.Copy(&stderrBuf, stderrReader)
	}()

	// 执行命令 - 使用推荐的 StreamWithContext 方法
	ctx := context.Background()
	err = executor.StreamWithContext(ctx, remotecommand.StreamOptions{
		Stdout: stdoutWriter,
		Stderr: stderrWriter,
	})

	// 关闭管道以触发读取完成
	stdoutWriter.Close()
	stderrWriter.Close()
	fmt.Println("WG is waiting for done............................")
	WG.Wait()

	if err != nil {
		return "", "", fmt.Errorf("stream execution failed: %v", err)
	}
	if stdoutErr != nil {
		return "", "", fmt.Errorf("stdout copy failed: %v", stdoutErr)
	}
	if stderrErr != nil {
		return "", "", fmt.Errorf("stderr copy failed: %v", stderrErr)
	}
	fmt.Printf("stdout: %s\nstderr: %s\n", stdoutBuf.String(), stderrBuf.String())
	return stdoutBuf.String(), stderrBuf.String(), nil
}

func GetNodeLogs1(nodeName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "kubectl", "debug",
		"node/"+nodeName,
		"--image=busybox",
		"--quiet", // 减少不必要的输出
		"--",
		"sh", "-c", "journalctl -u kubelet -n 100 --no-pager")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Printf("Executing kubectl debug on node: %s\n", nodeName)

	err := cmd.Run()
	if err != nil {
		// 检查是否是超时错误
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("command timed out after 60 seconds")
		}

		// 检查 kubectl 是否可用
		if exitErr, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("kubectl exited with code %d: %s", exitErr.ExitCode(), stderr.String())
		}

		return "", fmt.Errorf("failed to execute kubectl: %v, stderr: %s", err, stderr.String())
	}

	output := stdout.String()
	if output == "" {
		return "", fmt.Errorf("no output received from command")
	}

	return output, nil
}

func CaptureNodeExecOutput1(useInCluster bool, kubeconfigPath string, nodeName string, command []string) (string, string, error) {
	var config *rest.Config
	var err error

	if useInCluster {
		config, err = rest.InClusterConfig()
		if err != nil {
			return "", "", fmt.Errorf("in-cluster config failed: %v", err)
		}
	} else {
		if kubeconfigPath == "" {
			if home := homedir.HomeDir(); home != "" {
				kubeconfigPath = filepath.Join(home, ".kube", "config")
			} else {
				return "", "", fmt.Errorf("kubeconfig path is required for external mode")
			}
		}
		config, err = clientcmd.BuildConfigFromFlags("", "./kubeconfig")
		if err != nil {
			return "", "", fmt.Errorf("kubeconfig load failed: %v", err)
		}
	}
	// var Kubeconfig2 *string
	// var Config2 *rest.Config
	// Kubeconfig2 = flag.String("kubeconfig", "kubeconfig", "")
	// //}
	// flag.Parse() // Ensure parsing command line arguments

	// // Load main configuration
	// Config2, err = clientcmd.BuildConfigFromFlags("", *Kubeconfig2)

	// 创建新的 Clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", "", fmt.Errorf("failed to create clientset: %v", err)
	}
	fmt.Println(clientset)

	// 创建管道
	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader, stderrWriter := io.Pipe()

	// 关键修改：正确的节点执行 URL 构建方式
	Req := kubernetsServ.Clientset.CoreV1().RESTClient().Post().
		Resource("nodes").
		Name(nodeName).
		SubResource("proxy").
		Suffix("/exec") // 注意这里只是 /exec，不是完整的命令

	// 添加命令参数 - 通过 URL 参数传递
	for _, cmd := range command {
		Req = Req.Param("command", cmd)
	}

	// 或者使用 Param 方法设置其他参数
	Req = Req.Param("stdin", "false").
		Param("stdout", "true").
		Param("stderr", "true").
		Param("tty", "false")

	fmt.Printf("Node exec request URL: %s\n", Req.URL().String())

	// 创建执行器
	executor, err := remotecommand.NewSPDYExecutor(config, "POST", Req.URL())
	if err != nil {
		return "", "", fmt.Errorf("failed to create executor: %v", err)
	}

	// 启动goroutine读取输出
	var stdoutBuf, stderrBuf bytes.Buffer
	var stdoutErr error
	var stderrErr error
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		_, stdoutErr = io.Copy(&stdoutBuf, stdoutReader)
	}()

	go func() {
		defer wg.Done()
		_, stderrErr = io.Copy(&stderrBuf, stderrReader)
	}()

	// 执行命令
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Printf("开始执行命令，超时时间: 30秒\n")
	err = executor.StreamWithContext(ctx, remotecommand.StreamOptions{
		Stdout: stdoutWriter,
		Stderr: stderrWriter,
		Tty:    false,
	})
	fmt.Printf("执行器返回，错误: %v\n", err)

	// 关闭管道
	fmt.Println("关闭stdout管道...")
	stdoutWriter.Close()
	fmt.Println("关闭stderr管道...")
	stderrWriter.Close()
	fmt.Println("等待goroutine完成...")
	wg.Wait()
	fmt.Printf("goroutine完成，stdout错误: %v, stderr错误: %v\n", stdoutErr, stderrErr)

	if err != nil {
		// 更详细的错误处理
		fmt.Printf("执行器错误详情:\n")

		// 尝试转换为 StatusError 获取更多信息
		if statusErr, ok := err.(*errors.StatusError); ok {
			if statusErr.ErrStatus.Message != "" {
				fmt.Printf("API 服务器返回的错误消息: %s\n", statusErr.ErrStatus.Message)
			}
			if statusErr.ErrStatus.Reason != "" {
				fmt.Printf("错误原因: %s\n", statusErr.ErrStatus.Reason)
			}
			if statusErr.ErrStatus.Details != nil {
				fmt.Printf("错误详情: %+v\n", statusErr.ErrStatus.Details)
			}
		}

		// 检查是否是超时错误
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("执行器错误: 命令执行超时")
		}

		// 打印输出缓冲区内容，即使出错
		fmt.Printf("stdout内容长度: %d, 内容: %q\n", stdoutBuf.Len(), stdoutBuf.String())
		fmt.Printf("stderr内容长度: %d, 内容: %q\n", stderrBuf.Len(), stderrBuf.String())

		// 返回更友好的错误信息
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Message != "" {
			return stdoutBuf.String(), stderrBuf.String(),
				fmt.Errorf("Kubernetes API 错误: %s (stdout: %q, stderr: %q)",
					statusErr.ErrStatus.Message, stdoutBuf.String(), stderrBuf.String())
		}

		return stdoutBuf.String(), stderrBuf.String(),
			fmt.Errorf("命令执行失败: %v (stdout: %q, stderr: %q)",
				err, stdoutBuf.String(), stderrBuf.String())
	}

	if stdoutErr != nil {
		fmt.Printf("stdout复制错误: %v\n", stdoutErr)
		return stdoutBuf.String(), stderrBuf.String(),
			fmt.Errorf("stdout copy failed: %v", stdoutErr)
	}

	if stderrErr != nil {
		fmt.Printf("stderr复制错误: %v\n", stderrErr)
		return stdoutBuf.String(), stderrBuf.String(),
			fmt.Errorf("stderr copy failed: %v", stderrErr)
	}

	fmt.Printf("命令执行成功，stdout: %d字节, stderr: %d字节\n", stdoutBuf.Len(), stderrBuf.Len())
	return stdoutBuf.String(), stderrBuf.String(), nil
}
