package config

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/go-ozzo/ozzo-validation"
)

type Config struct {
	Packages []Package
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Packages, validation.Length(1, 0), validation.Required, validation.By(checkFilesOrServices)),
	)
}

func checkFilesOrServices(value interface{}) error {
	for _, p := range value.([]Package) {
		if p.Files == nil && p.Services == nil {
			return errors.New(fmt.Sprintf("Either Files or Services must be present"))
		}
	}
	return nil
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
