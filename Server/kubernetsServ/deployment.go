package kubernetsServ

import (
	"context"
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//var kubeconfig2 *string

func Deployment() {
	// kubeconfig2 = flag.String("kubeconfig", "./kubeconfig", "path to kubeconfig file")
	// // }
	// flag.Parse()

	// fmt.Println("kubeconfig:", *kubeconfig2)

	config, err := clientcmd.BuildConfigFromFlags("", "./kubeconfig")
	if err != nil {
		fmt.Println("BuildConfigFromFlags error: ", err)
		log.Fatal(err)
	}
	Clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Clientset error: ", err)
		log.Fatal(err)
	}
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "example-name77",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "example-app11"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": "example-Labels11"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "example-container11",
							Image: "nginx:latest",
						},
					},
				},
			},
		},
	}
	result, err := Clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("Create deployment error: ", err)
		log.Fatal(err)
	}
	fmt.Printf("Deployment %q was created\n", result.GetObjectMeta().GetName())

}
func int32Ptr(i int32) *int32 { return &i }
