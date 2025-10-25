package kubernetsServ

import (
	"context"
	"gok8s/models"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodesInfo() []models.NodesInfo { //[]models.ClusterInfo
	api := Clientset.CoreV1()
	nodeList, err := api.Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	var nodesInfos []models.NodesInfo
	for _, node := range nodeList.Items {
		var addressesStr string
		for i, addr := range node.Status.Addresses {
			if i > 0 {
				addressesStr += ", "
			}
			addressesStr += addr.Address
		}
		nodesInfo := models.NodesInfo{
			Name:            node.Name,
			Phase:           string(node.Status.Phase),
			Addresses:       []string{addressesStr},
			OSImage:         node.Status.NodeInfo.OSImage,
			KubeletVersion:  node.Status.NodeInfo.KubeletVersion,
			OperatingSystem: node.Status.NodeInfo.OperatingSystem,
			Architecture:    node.Status.NodeInfo.Architecture,
		}
		nodesInfos = append(nodesInfos, nodesInfo)
		models.SetNodes(nodesInfos)
		// fmt.Println(nodesInfo)
	}
	return nodesInfos
}
