package kubernetsServ

import (
	"gok8s/models"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

// var KubeconfigResources *string
// var ClientsetResouces *kubernetes.Clientset
// var ConfigR *rest.Config // 资源获取专用配置

func GetK8sResources() []models.ResourcesInfo {
	// 检查ConfigR是否已初始化
	if ConfigR == nil {
		return nil
	}
	// 创建DiscoveryClient
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(ConfigR)
	if err != nil {
		return nil
	}

	// 获取API资源列表
	_, APIResourceList, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		return nil
	}

	var resourcesInfos []models.ResourcesInfo

	for _, list := range APIResourceList {
		// 跳过空列表
		if len(list.APIResources) == 0 {
			continue
		}

		GV, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			return nil
		}

		for _, Resource := range list.APIResources {
			// 创建资源信息对象
			resourcesInfo := models.ResourcesInfo{
				Name:               Resource.Name,
				Kind:               Resource.Kind,
				SingularName:       Resource.SingularName,
				StorageVersionHash: Resource.StorageVersionHash,
				ResourcesVersion:   Resource.Version,
				Categories: func() string {
					if len(Resource.Categories) > 0 {
						return strings.Join(Resource.Categories, ",")
					}
					return ""
				}(),
				ShortNames: func() string {
					if len(Resource.ShortNames) > 0 {
						return strings.Join(Resource.ShortNames, ",")
					}
					return ""
				}(),
				Verbs:         strings.Join(Resource.Verbs, ","),
				DeepCopyGroup: Resource.DeepCopy().Group,
				DeepCopyName:  Resource.DeepCopy().Name,
				Namespaced:    Resource.Namespaced,
				ResouceSize:   strconv.Itoa(Resource.Size()),
				Group:         GV.Group,
				Version:       GV.Version,
			}
			resourcesInfos = append(resourcesInfos, resourcesInfo)
		}
	}

	// 更新资源信息到模型
	models.SetResources(resourcesInfos)
	return resourcesInfos
}
