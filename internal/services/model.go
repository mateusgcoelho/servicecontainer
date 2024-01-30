package services

type engineType string

const (
	Java8  engineType = "java8"
	NodeJs engineType = "nodejs"
)

type ServiceModel struct {
	Id          int        `json:"id"`
	Tag         string     `json:"tag"`
	PrefixUrl   string     `json:"prefixUrl"`
	DefaultPort int        `json:"defaultPort"`
	DisplayName string     `json:"displayName"`
	FileName    string     `json:"fileName"`
	EngineType  engineType `json:"engineType"`
}
