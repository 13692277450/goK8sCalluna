package models

type ResourcesInfo struct {
	Name               string
	Kind               string
	SingularName       string
	StorageVersionHash string
	ResourcesVersion   string
	ShortNames         string
	Categories         string
	DeepCopyGroup      string
	DeepCopyName       string
	Verbs              string
	Namespaced         bool
	ResouceSize        string
	Group              string
	Version            string
}

var getResourcesInfo []ResourcesInfo

func GetResources() []ResourcesInfo {
	return getResourcesInfo
}

func SetResources(resourcesInfo []ResourcesInfo) {
	getResourcesInfo = resourcesInfo
}
