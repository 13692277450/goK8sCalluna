package kubernetsServ

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetLogsFromMultiPods(nameSpace, LabelSelector string) {
	Kubeconfig := flag.String("kubeconfig", "kubeconfig", "")
	Config, err := clientcmd.BuildConfigFromFlags("", *Kubeconfig)
	if err != nil {
		fmt.Println("BuildConfigFromFlags error: ", err)
		panic(err.Error())
	}
	//fmt.Println("Configure======:", Config)
	//fmt.Println("config: ", config)
	Clientset, err = kubernetes.NewForConfig(Config)
	if err != nil {
		fmt.Println("Clientset  error: ", err, Clientset)
		log.Fatal(err)
	}
	pods, err := Clientset.CoreV1().Pods(nameSpace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LabelSelector,
	})
	if err != nil {
		return
	}
	for _, pod := range pods.Items {
		fmt.Printf("*** Logs for Pod %s *** \n", pod.Name)
		GetPodLogs(nameSpace, pod.Name)
		fmt.Println()
	}
}

func GetPodLogs(nameSpace, podName string) {
	// Kubeconfig := flag.String("kubeconfig", "kubeconfig", "")
	// Config, err := clientcmd.BuildConfigFromFlags("", *Kubeconfig)
	// if err != nil {
	// 	fmt.Println("BuildConfigFromFlags error: ", err)
	// 	panic(err.Error())
	// }
	// //fmt.Println("Configure======:", Config)
	// //fmt.Println("config: ", config)
	// Clientset, err = kubernetes.NewForConfig(Config)
	// if err != nil {
	// 	fmt.Println("Clientset  error: ", err, Clientset)
	// 	log.Fatal(err)
	// }
	podLogOpts := corev1.PodLogOptions{
		Container: "example-container",
		Follow:    false,
		Previous:  false,
	}

	req := Clientset.CoreV1().Pods(nameSpace).GetLogs(podName, &podLogOpts)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return
	}
	defer podLogs.Close()
	logs, err := io.ReadAll(podLogs)
	if err != nil {
		return
	}
	fmt.Println("Logs.......:", string(logs))
}

func logsGetPodsList() {
	//GetPodLogs("default", "example-deployment-6fdbfc7c54-cdwtt")
	GetLogsFromMultiPods("default", "app=example-app")
	fmt.Println("Logs collect start...")
}

// /////////////////////////////////////////////////////////////
func TryGetPodLogs(namespace string, podName string, containerName string, follow bool) error {
	count := int64(100)
	podLogOptions := v1.PodLogOptions{
		Container: containerName,
		Follow:    follow,
		TailLines: &count,
	}

	podLogRequest := Clientset.CoreV1().
		Pods(namespace).
		GetLogs(podName, &podLogOptions)
	stream, err := podLogRequest.Stream(context.TODO())
	if err != nil {
		return err
	}
	defer stream.Close()

	for {
		buf := make([]byte, 2000)
		numBytes, err := stream.Read(buf)
		if numBytes == 0 {
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		message := string(buf[:numBytes])
		fmt.Print(message)
	}
	return nil
}
