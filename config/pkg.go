package config

import (
	"errors"
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
)

type Package struct {
	Name     string
	Binaries []Binary
	Files    []File
	Services []Service
	Inits    []Init
}

func (p Package) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Binaries, validation.Length(1, 0), validation.Required),
		validation.Field(&p.Files, validation.Length(1, 0)),
		validation.Field(&p.Services, validation.Length(1, 0)),
		validation.Field(&p.Inits, validation.Length(0, 0), validation.By(checkInitContent)),
	)
}

func checkInitContent(value interface{}) error {
	for _, i := range value.([]Init) {
		if (i.Type == Command) && (validation.ValidateStruct(&i,
			validation.Field(&i.Type, validation.Required),
			validation.Field(&i.Content, validation.Required)) != nil) {
			return errors.New(fmt.Sprintf("Content must be supplied to command"))
		}
	}
	return nil
}
