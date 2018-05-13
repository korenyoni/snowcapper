package config

type File struct {
	Path    string
	Content string
}

func NewFile(path string, content string) (file File, err error) {
	return File{
		Path:    path,
		Content: content,
	}, nil
}
