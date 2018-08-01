package config

import (
	"os"
	"errors"

	"github.com/go-ozzo/ozzo-validation"
)

const (
	Sha512     int = 128
	Sha384     int = 96
	Sha256     int = 64
	Sha224     int = 56
	Sha1       int = 40
	Md5        int = 32
	Empty	   int = 0
	Invalid    int = -1
)

type Downloadable struct {
	Src    	string 		`yaml:"src"`
	SrcHash string 		`yaml:"src_hash"`
}

type Binary struct {
	Downloadable 		`yaml:",inline"`
	Name   	string 		`yaml:"name"`
	Format 	string 		`yaml:"format"`
	Mode   	os.FileMode 	`yaml:"mode"`
}

type Extend struct {
	Downloadable 	`yaml:",inline"`
}

func (b Binary) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required),
		validation.Field(&b.Src, validation.Required),
		validation.Field(&b.SrcHash, validation.By(validateSrcHash)),
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

func (e Extend) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Src, validation.Required),
		validation.Field(&e.SrcHash, validation.By(validateSrcHash)),
	)
}

func (e *Extend) GetDownloadPath() string {
	return "/tmp/extend_" + e.Src 
}

func validateSrcHash(value interface{}) error {
	s, _ := value.(string)
	if GetHashType(s) == Invalid {
		return errors.New("Invalid hash length")
	}
	return nil
}

func GetHashType(hash string) int {
	switch l := len(hash); l {
	case 0:
		return Empty
	case Sha512:
		return l 
	case Sha384:
		return l 
	case Sha256:
		return l 
	case Sha224:
		return l 
	case Sha1:
		return l 
	case Md5:
		return l
	default:
		return Invalid 
	}
}
