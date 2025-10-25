package models

type NameSpaces struct {
	Name     string `json:"name"`
	Creation string `json:"creation"`
	Status   string `json:"status"`
}

var nameSpaces []NameSpaces

func GetNameSpaces() []NameSpaces {
	return nameSpaces
}

func SetNameSpaces(NameSpaces []NameSpaces) {
	nameSpaces = NameSpaces
}
