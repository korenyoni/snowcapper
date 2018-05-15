package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"os"
)

type Binary struct {
	Name   string
	Src    string
	Format string
	Mode   os.FileMode
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
