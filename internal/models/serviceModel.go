package models

type EngineType string

const (
	EngineJava8  EngineType = "java8"
	EngineNodeJs EngineType = "nodejs"
)

type ServiceStatus int

const (
	ServiceStoped     ServiceStatus = 0
	ServiceStarting   ServiceStatus = 1
	ServiceStarted    ServiceStatus = 2
	ServiceInDownload ServiceStatus = 3
)

type ServiceModel struct {
	Id          int           `json:"id"`
	Tag         string        `json:"tag"`
	PrefixUrl   string        `json:"prefixUrl"`
	DefaultPort int           `json:"defaultPort"`
	DisplayName string        `json:"displayName"`
	FileName    string        `json:"fileName"`
	EngineType  EngineType    `json:"engineType"`
	Status      ServiceStatus `json:"status"`
}
