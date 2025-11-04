package kubernetsServ

import (
	"context"
	"encoding/json"
	"fmt"
	"gok8s/models"
	"io"
	"strings"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodLogEntry struct {
	PodName string `json:"podName"`
	Logs    string `json:"logs"`
}

func GetLogsFromMultiPods(NameSpacesTotal []models.NameSpaces, LabelSelector string) string {
	var allLogs []PodLogEntry

	for _, ns := range NameSpacesTotal {
		pods, err := Clientset.CoreV1().Pods(ns.Name).List(context.TODO(), metav1.ListOptions{
			LabelSelector: LabelSelector,
		})
		if err != nil {
			fmt.Printf("Error listing pods in namespace %s: %v\n", ns, err)
			continue
		}

		for _, pod := range pods.Items {
			podLogs := GetPodLogs(pod.Namespace, pod.Name)
			allLogs = append(allLogs, PodLogEntry{
				PodName: fmt.Sprintf("%s/%s", pod.Namespace, pod.Name),
				Logs:    podLogs,
			})
		}
	}

	logData, err := json.Marshal(allLogs)
	if err != nil {
		return fmt.Sprintf(`{"error": "Failed to marshal logs: %v"}`, err)
	}
	return string(logData)
}

//GetLogsFromMultiPods

// func GetLogsFromMultiPods(nameSpace, LabelSelector string) string {
// 	var totalPodsLogs string

// 	for nameSpacei := range NameSpacesTotal {

// 		pods, err := Clientset.CoreV1().Pods(NameSpacesTotal[nameSpacei].Name).List(context.TODO(), metav1.ListOptions{
// 			LabelSelector: LabelSelector,
// 		})
// 		if err != nil {
// 			fmt.Println("List logs error: ", err)
// 			//log.Fatal(err)
// 		}
// 		fmt.Printf("Namespace is ===================== %v\n\n", NameSpacesTotal[nameSpacei].Name)
// 		//var podsLogsResult []string

// 		for _, pod := range pods.Items {
// 			go func() {

// 				var podLogs = GetPodLogs(NameSpacesTotal[nameSpacei].Name, pod.Name)
// 				// 将pod.Name作为标题，podLogs作为内容
// 				totalPodsLogs += "TITLE:" + pod.Name + "\nCONTENT:" + podLogs + "\n"
// 			}()
// 			//fmt.Println(totalPodsLogs)
// 			fmt.Printf("Pod  %v name  is ===================== %v\n\n", pod.Name, NameSpacesTotal[nameSpacei].Name)

// 		}

// 	}
// 	//fmt.Println("totoalpodsLogs: ***********************", totalPodsLogs)
// 	return totalPodsLogs
// }

func GetPodLogs(nameSpace, podName string) (LogStrings string) {

	podLogOpts := corev1.PodLogOptions{
		//Container: "example-container",
		Follow:   false,
		Previous: true,
	}

	req := Clientset.CoreV1().Pods(nameSpace).GetLogs(podName, &podLogOpts)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return fmt.Sprintf("Failure to get logs: %v", err)
	}
	defer podLogs.Close()
	logs, err := io.ReadAll(podLogs)
	if err != nil {
		return fmt.Sprintf("Failure to read log: %v", err)
	}
	logContent := string(logs)
	// 提取第一行作为标题，其余作为内容
	logLines := strings.Split(logContent, "\n")
	if len(logLines) > 0 {
		title := strings.TrimSuffix(strings.TrimPrefix(logLines[0], podName+" Logs.......:"), ":")
		content := strings.Join(logLines[1:], "\n")
		return fmt.Sprintf("%s\n %s", title, content)
	}
	return logContent
}

func logsGetPodsList() {
	//GetPodLogs("default", "example-deployment-6fdbfc7c54-cdwtt")
	//GetLogsFromMultiPods("default", "app=example-app")
	fmt.Println("Logs collect start...")
}

// /////////////////////////////////////////////////////////////did not use so far.
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
