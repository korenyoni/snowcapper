package config

import (
	"os"
)

type Binary struct {
	Name   string
	Src    string
	Format string
	Mode   os.FileMode
}

func (b *Binary) GetBinaryPath() string {
	return "/usr/bin/" + b.Name
}

func (b *Binary) GetDownloadPath() string {
	return "/tmp/" + b.Name + "." + b.Format
}
