package models

type PodsLogs struct {
	LogString string
}

var getPodLogs PodsLogs

func GetPodLogs() PodsLogs {
	return getPodLogs
}

func SetPodLogs(podLogs PodsLogs) {
	getPodLogs = podLogs
}
