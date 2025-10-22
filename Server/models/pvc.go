package models

// PVC represents a PersistentVolumeClaim in Kubernetes
type PVCInfo struct {
	Name         string
	Namespace    string
	Status       string
	StorageClass string
	Capacity     string
	AccessModes  []string
	VolumeName   string
	CreationTime string
}

var Pvcs []PVCInfo

func GetPVC() []PVCInfo {
	return Pvcs
}

func SetPVC(pvcs []PVCInfo) {
	Pvcs = pvcs
}
