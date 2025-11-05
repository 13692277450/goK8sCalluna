package kubernetsServ

import (
	"context"
	"fmt"
	"gok8s/models"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PVCList []models.PVCInfo

// var pvcController PVCController
var Ctx = context.Background()
var Namespace = "default"

func GetPVCList(namespace string) ([]models.PVCInfo, error) {

	pvcs, err := Clientset.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get PVC list: %v", err)
	}

	PVCList = []models.PVCInfo{} // clean
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
			AccessModes: convertAccessModes(pvc.Spec.AccessModes),
			// ResourceVersion: pvc.ResourceVersion,
			// APIVersion:      pvc.APIVersion,
			// GenerateName:    pvc.GenerateName,
			// Labels:          convertLabelsToString(pvc.Labels),
			UID: string(pvc.UID),
			// Annotations:     convertAnnotationsToString(pvc.Annotations),
			// OwnerReferences: convertOwnerReferencesToString(pvc.OwnerReferences),
			// Finalizers:      convertFinalizersToString(pvc.Finalizers),
			// Selector:        convertLabelSelectorToString(pvc.Spec.Selector), // 添加Selector字段
			CreationTime: pvc.CreationTimestamp.String(),
		})
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
//func GetPVCList() {
// ctx := context.Background()
// namespace := "default"

// 添加类型转换辅助函数
func convertLabelsToString(labels map[string]string) string {
	if labels == nil {
		return ""
	}
	var labelStrings []string
	for k, v := range labels {
		labelStrings = append(labelStrings, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(labelStrings, ",")
}

func convertAnnotationsToString(annotations map[string]string) string {
	if annotations == nil {
		return ""
	}
	var annoStrings []string
	for k, v := range annotations {
		annoStrings = append(annoStrings, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(annoStrings, ",")
}

func convertOwnerReferencesToString(ownerRefs []metav1.OwnerReference) string {
	if ownerRefs == nil {
		return ""
	}
	var ownerStrings []string
	for _, owner := range ownerRefs {
		ownerStrings = append(ownerStrings, fmt.Sprintf("%s/%s", owner.Kind, owner.Name))
	}
	return strings.Join(ownerStrings, ",")
}

func convertFinalizersToString(finalizers []string) string {
	if finalizers == nil {
		return ""
	}
	return strings.Join(finalizers, ",")
}

func convertLabelSelectorToString(selector *metav1.LabelSelector) string {
	if selector == nil {
		return ""
	}
	// 处理MatchLabels
	if selector.MatchLabels != nil {
		var labelStrings []string
		for k, v := range selector.MatchLabels {
			labelStrings = append(labelStrings, fmt.Sprintf("%s=%s", k, v))
		}
		return strings.Join(labelStrings, ",")
	}
	// 未来可以添加对MatchExpressions的处理
	return ""
}
