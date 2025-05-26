package interfaces

import (
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type PluginID string

type Deployment string

const (
	MICROSERVICE Deployment = "microservice"
	MONOLITHIC   Deployment = "monolithic"
)

type PluginConfig struct {
	Plugins []PluginData `yaml:"plugins"`
}

type PluginData struct {
	Name       string `yaml:"name"`
	Interface  string `yaml:"interface"`
	Deployment string `yaml:"deployment"`
	Source     string `yaml:"source,omitempty"`
}

func ParsePluginYaml(src string) *PluginConfig {
	currDir, _ := filepath.Abs(".")
	filePath := path.Join(currDir, src)

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	result := PluginConfig{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		panic(err)
	}

	return &result
}
