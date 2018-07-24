package config

import (
	"os"

	"github.com/go-ozzo/ozzo-validation"
)

type Binary struct {
	Name   string 		`yaml:"name"`
	Src    string 		`yaml:"src"`
	Format string 		`yaml:"format"`
	Mode   os.FileMode 	`yaml:"mode"`
}

func (b Binary) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required),
		validation.Field(&b.Src, validation.Required),
		validation.Field(&b.Format, validation.Required),
		validation.Field(&b.Mode, validation.Required),
	)
}

func (b *Binary) GetBinaryPath() string {
	return "/usr/bin/" + b.Name
}

func (b *Binary) GetDownloadPath() string {
	return "/tmp/" + b.Name + "." + b.Format
}
