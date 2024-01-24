package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// TargetConfig はYAMLファイルの構造を定義します。
type TargetConfig struct {
	Targets map[string]struct {
		Namespace string `yaml:"Namespace"`
		Name      string `yaml:"Name"`
		Help      string `yaml:"Help"`
		URL       string `yaml:"URL"`
	} `yaml:"targets"`
}

func GetTargets() TargetConfig {
	ymlPath := "./config/targets.yml"

	ymlData, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		log.Printf("error: Failed to read YAML file: %v", err)
	}

	var targetConfig TargetConfig
	if err := yaml.Unmarshal(ymlData, &targetConfig); err != nil {
		log.Printf("error: Failed to unmarshal YAML data: %v", err)
	}

	return targetConfig
}
