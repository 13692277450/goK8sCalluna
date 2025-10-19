package kubernetsServ

import (
	"context"
	"fmt"
	"gok8s/models"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var PVCList []models.PVCInfo

type PVCController struct {
	clientSet *kubernetes.Clientset
}

func PvcInit() (*PVCController, error) {
	config, err := clientcmd.BuildConfigFromFlags("", "kubeconfig")
	if err != nil {
		return nil, fmt.Errorf("failed to build kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %v", err)
	}

	return &PVCController{
		clientSet: clientset,
	}, nil
}

func (c *PVCController) GetPVCList(ctx context.Context, namespace string) ([]models.PVCInfo, error) {

	pvcs, err := c.clientSet.CoreV1().PersistentVolumeClaims(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get PVC list: %v", err)
	}

	PVCList = []models.PVCInfo{} // 清空之前的列表
	for _, pvc := range pvcs.Items {
		PVCList = append(PVCList, models.PVCInfo{
			Name:      pvc.Name,
			Namespace: pvc.Namespace,
			Status:    string(pvc.Status.Phase),
			StorageClass: func() string {
				if pvc.Spec.StorageClassName != nil {
					return *pvc.Spec.StorageClassName
				}
				return ""
			}(),
			Capacity: func() string {
				if pvc.Spec.Resources.Requests != nil {
					return pvc.Spec.Resources.Requests.Storage().String()
				}
				return ""
			}(),
			AccessModes:  convertAccessModes(pvc.Spec.AccessModes),
			VolumeName:   pvc.Spec.VolumeName,
			CreationTime: pvc.CreationTimestamp.String(),
		})
	}
	fmt.Println("PVCList##############:", PVCList)
	for i := range PVCList {
		fmt.Println("PVCList[i]##############:", PVCList[i])
	}
	// 将获取的PVC数据设置到模型中
	models.SetPVC(PVCList)
	return PVCList, nil
}

func convertAccessModes(modes []v1.PersistentVolumeAccessMode) []string {
	var result []string
	for _, mode := range modes {
		result = append(result, string(mode))
	}
	return result
}

// GetPVCList 包级别函数，获取默认命名空间的PVC列表
func GetPVCList() {
	ctx := context.Background()
	namespace := "default"

	pvcController, err := PvcInit()
	if err != nil {
		fmt.Printf("Failed to initialize PVC controller: %v\n", err)
		return
	}
	if pvcController == nil {
		fmt.Println("PVC controller is nil")
		return
	}

	pvcs, err := pvcController.GetPVCList(ctx, namespace)
	if err != nil {
		fmt.Printf("Failed to get PVC list: %v\n", err)
		return
	}

	fmt.Printf("Found %d PVCs in namespace %s:\n", len(pvcs), namespace)
	for _, pvc := range pvcs {
		fmt.Printf("- %s (Status: %s, StorageClass: %s)\n", pvc.Name, pvc.Status, pvc.StorageClass)
	}
}
