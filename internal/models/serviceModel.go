package models

type engineType string

const (
	EngineJava8  engineType = "java8"
	EngineNodeJs engineType = "nodejs"
)

type status int

const (
	ServiceStoped     status = 0
	ServiceStarting   status = 1
	ServiceStarted    status = 2
	ServiceInDownload status = 3
)

type ServiceModel struct {
	Id          int        `json:"id"`
	Tag         string     `json:"tag"`
	PrefixUrl   string     `json:"prefixUrl"`
	DefaultPort int        `json:"defaultPort"`
	DisplayName string     `json:"displayName"`
	FileName    string     `json:"fileName"`
	EngineType  engineType `json:"engineType"`
	Status      status     `json:"status"`
}
