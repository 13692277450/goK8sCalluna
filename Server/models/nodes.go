package models

type NodesInfo struct {
	Name            string   `json:"name"`
	Phase           string   `json:"phase"`
	OSImage         string   `json:"osimage"`
	KubeletVersion  string   `json:"kubeletversion"`
	OperatingSystem string   `json:"operatingsystem"`
	Architecture    string   `json:"architecture"`
	Addresses       []string `json:"addresses"`
}

var nodes []NodesInfo

func GetNodess() []NodesInfo {
	return nodes
}

func SetNodes(Nodes []NodesInfo) {
	nodes = Nodes
}
