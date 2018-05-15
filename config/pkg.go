package config

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Package struct {
	Name     string
	Binaries []Binary
	Files    []File
}

func (p Package) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Binaries, validation.Required),
		validation.Field(&p.Files, validation.Required),
	)
}
