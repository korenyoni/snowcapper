package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"os"
)

type File struct {
	Path    string
	Content string
	Mode    os.FileMode
}

func (f File) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Path, validation.Required),
		validation.Field(&f.Content, validation.Required),
		validation.Field(&f.Mode, validation.Required),
	)
}
