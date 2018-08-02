package config

import (
	"gopkg.in/yaml.v2"

	"github.com/go-ozzo/ozzo-validation"
)

type Config struct {
	Extends  []Extend	`yaml:"extends"`
	Packages []Package 	`yaml:"packages"`
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Extends),
		validation.Field(&c.Packages, validation.Length(1, 0), validation.Required),
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
