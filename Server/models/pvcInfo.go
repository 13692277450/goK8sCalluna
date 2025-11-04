package models

// PVC represents a PersistentVolumeClaim in Kubernetes
type PVCInfo struct {
	Name         string
	Namespace    string
	Status       string
	StorageClass string
	Capacity     string
	AccessModes  []string
	CreationTime string
	// ResourceVersion string
	// APIVersion      string
	// GenerateName    string
	// Labels          string
	UID string
	// Annotations     string
	// OwnerReferences string
	// Finalizers      string
	// Selector        string
}

var Pvcs []PVCInfo

func GetPVC() []PVCInfo {
	return Pvcs
}

func SetPVC(pvcs []PVCInfo) {
	Pvcs = pvcs
}
