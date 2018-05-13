package config

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	Packages []Package
}

func New(configYaml []byte) (config Config, err error) {
	config = Config{}
	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
