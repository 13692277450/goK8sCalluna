package handlers

import (
	"context"
	"fmt"
	"io"
	"os/exec"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// 1. 从Pod获取日志
func getPodLogs(clientset *kubernetes.Clientset, namespace, podName string) (string, error) {
	logOptions := v1.PodLogOptions{}
	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, &logOptions)
	reader, err := req.Stream(context.TODO())
	if err != nil {
		return "", err
	}
	defer reader.Close()

	logs, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(logs), nil
}

// 2. 从特定系统组件获取日志（如etcd）
func GetSystemComponentLogs(clientset *kubernetes.Clientset, componentName, namespace string) (string, error) {
	// 通过kubectl logs命令获取
	cmd := fmt.Sprintf("kubectl logs -n %s %s", namespace, componentName)
	return ExecCommand(cmd)
}

// 3. 使用systemd/journalctl获取节点日志
func GetNodeLogs(nodeName string) (string, error) {
	cmd := "journalctl -u kubelet -n 100 --no-pager"
	return ExecCommand(cmd)
}

// 4. 通用命令执行函数
func ExecCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
