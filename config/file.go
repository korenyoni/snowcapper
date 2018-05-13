package config

import (
	"os"
)

type File struct {
	Path    string
	Content string
	Mode    os.FileMode
}

func NewFile(path string, content string, Mode os.FileMode) (file File, err error) {
	return File{
		Path:    path,
		Content: content,
	}, nil
}
