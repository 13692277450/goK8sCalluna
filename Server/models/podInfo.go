package models

type PodInfo struct {
	Name      string
	Status    string
	NodeName  string
	HostIP    string
	PodIP     string
	StartTime string
	Namespace string
}

var Request struct {
	PodName   string `json:"podname"`
	NameSpace string `json:"namespace"`
}
var Pods []PodInfo

func GetPods() []PodInfo {
	return Pods
}

func SetPods(pods []PodInfo) {
	Pods = pods
}
