package config

import (
	"errors"
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
)

type packagePointer struct {
	Pointer *Package
}

type Package struct {
	Name     string    `yaml:"name"`
	Binaries []Binary  `yaml:"binaries"`
	Files    []File    `yaml:"files"`
	Services []Service `yaml:"services"`
	Inits    []Init    `yaml:"inits"`
}

func (p Package) Validate() error {
	err := validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Binaries, validation.Length(1, 0)),
		validation.Field(&p.Files, validation.Length(1, 0)),
		validation.Field(&p.Services, validation.Length(1, 0)),
		validation.Field(&p.Inits, validation.Length(1, 0), validation.By(checkInitContent)),
	)
	if err != nil {
		return err
	}
	return validation.Validate(packagePointer{Pointer: &p}, validation.By(validatePackage))
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

func validatePackage(value interface{}) error {
	p := *(value.(packagePointer).Pointer)
	failCondition := p.Inits == nil && p.Files == nil && p.Binaries == nil && p.Services == nil
	if failCondition {
		return errors.New("Package must contain a name, and one of binaries, files, services, or inits")
	}
	return nil
}
