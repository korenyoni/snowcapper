package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Packages []Package
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Packages, validation.Length(1, 0)),
	)
}

func New(configYaml []byte) (config Config, err error) {
	config = Config{}
	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		return config, err
	}
	err = config.Validate()
	if err != nil {
		return config, err
	}
	return config, nil
}
