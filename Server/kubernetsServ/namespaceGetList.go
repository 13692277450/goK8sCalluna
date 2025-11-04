package kubernetsServ

import (
	"context"
	"fmt"
	"gok8s/models"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// corev1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NameSpacesTotal []models.NameSpaces

func GetNameSpacesTotal() []models.NameSpaces {
	return NameSpacesTotal
}
func GetNameSpaceList() []models.NameSpaces {
	NameSpacesTotal = []models.NameSpaces{}
	// Get namespace list
	namespaceList, err := Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// Print namespace info
	for _, namespace := range namespaceList.Items {
		nsp := models.NameSpaces{
			Name:     namespace.Name,
			Creation: fmt.Sprintf("%v", namespace.CreationTimestamp),
			Status:   string(namespace.Status.Phase),
		}
		NameSpacesTotal = append(NameSpacesTotal, nsp)
		models.SetNameSpaces(NameSpacesTotal)
	}
	fmt.Println("Namespaces: ", NameSpacesTotal)
	return NameSpacesTotal
}
