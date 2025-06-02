package interfaces

type PluginID string

type PluginData struct {
	Name      string `yaml:"name"`
	Interface string `yaml:"interface"`
}
