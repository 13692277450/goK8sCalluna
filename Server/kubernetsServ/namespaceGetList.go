package kubernetsServ

import (
	"context"
	"fmt"
	"gok8s/models"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// corev1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNameSpaceList() []models.NameSpaces {
	// Get namespace list
	namespaceList, err := Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// Print namespace info
	var nameSpaces []models.NameSpaces
	for _, namespace := range namespaceList.Items {
		nsp := models.NameSpaces{
			Name:     namespace.Name,
			Creation: fmt.Sprintf("%v", namespace.CreationTimestamp),
			Status:   string(namespace.Status.Phase),
		}
		nameSpaces = append(nameSpaces, nsp)
		models.SetNameSpaces(nameSpaces)
	}
	fmt.Println("Namespaces: ", nameSpaces)
	return nameSpaces
}
