package config

import (
	"os"

	"github.com/go-ozzo/ozzo-validation"
)

type File struct {
	Path    string		`yaml:"path"`
	Content string		`yaml:"content"`
	Mode    os.FileMode	`yaml:"mode"`
}

func (f File) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Path, validation.Required),
		validation.Field(&f.Content, validation.Required),
		validation.Field(&f.Mode, validation.Required),
	)
}
