package kubernetsServ

import (
	"gok8s/models"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

// var KubeconfigResources *string
// var ClientsetResouces *kubernetes.Clientset
var ConfigR *rest.Config

func GetK8sResources() []models.ResourcesInfo {
	//Kubeconfig := flag.String("kubeconfig", "kubeconfig", "")
	// ConfigResources, err := clientcmd.BuildConfigFromFlags("", *Kubeconfig)
	// if err != nil {
	// 	fmt.Println("BuildConfigFromFlags error: ", err)
	// 	panic(err.Error())
	// }
	// //fmt.Println("config: ", config)
	// Clientset, err = kubernetes.NewForConfig(ConfigResources)
	// if err != nil {
	// 	fmt.Println("Clientset  error: ", err, Clientset)
	// 	log.Fatal(err)
	// }
	//fmt.Println("Configure$$$$$$$$$$$$$$$$$$$4: ", ConfigR)
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(ConfigR)
	if err != nil {

	}
	_, APIResourceList, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {

	}
	var resourcesInfos []models.ResourcesInfo

	for _, list := range APIResourceList {
		GV, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			break
		}
		for _, Resource := range list.APIResources {
			//fmt.Printf("nameXXXXXXXXXX: %v, group: %v, version: %v", resource.Name, gv.Group, gv.Version)
			GV = GV
			Resource = Resource

			resourcesInfo := models.ResourcesInfo{
				Name:               Resource.Name,
				Kind:               Resource.Kind,
				SingularName:       Resource.SingularName,
				StorageVersionHash: Resource.StorageVersionHash,
				ResourcesVersion:   Resource.Version,
				Categories: func() string {
					if strings.Join(Resource.Categories, "") != "" {
						return strings.Join(Resource.Categories, "")
					}
					return ""
				}(),
				ShortNames: func() string {
					if strings.Join(Resource.ShortNames, "") != "" {
						return strings.Join(Resource.ShortNames, "")
					}
					return ""

				}(),
				Verbs:         Resource.Verbs.String(),
				DeepCopyGroup: Resource.DeepCopy().Group,
				DeepCopyName:  Resource.DeepCopy().Name,
				Namespaced:    Resource.Namespaced,
				ResouceSize:   strconv.Itoa(Resource.Size()),

				Group:   GV.Group,
				Version: GV.Version,
			}
			resourcesInfos = append(resourcesInfos, resourcesInfo)
			//fmt.Printf("name*************: %v, group: %v, version: %v", resourcesInfo.Name, resourcesInfo.Group, resourcesInfo.Version)

		}
		models.SetResources(resourcesInfos)
	}
	return resourcesInfos
}
