package config

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Service struct {
	Binary Binary
	Args   []string
}

func (s Service) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Binary, validation.Required),
		validation.Field(&s.Args),
	)
}
