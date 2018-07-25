package config

import (
	"errors"
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
)

const (
	Command string = "command"
	OpenRC  string = "openrc"
)

type Init struct {
	Type    string	`yaml:"type"`
	Content string	`yaml:"content"`
}

func (i Init) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Type, validation.Required, validation.By(checkType)),
		validation.Field(&i.Content),
	)
}

func checkType(value interface{}) error {
	s, _ := value.(string)
	if (s != Command) && (s != OpenRC) {
		return errors.New(fmt.Sprintf("command must be one of %s, %s", Command, OpenRC))
	}
	return nil
}
