package interfaces

type PluginID string

type Deployment string

const (
	MICROSERVICE Deployment = "microservice"
	MONOLITHIC   Deployment = "monolithic"
)

type PluginData struct {
	Name       string `yaml:"name"`
	Interface  string `yaml:"interface"`
	Deployment string `yaml:"deployment"`
	Source     string `yaml:"source,omitempty"`
}
