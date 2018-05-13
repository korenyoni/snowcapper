package config

type Package struct {
	Name     string
	Binaries []Binary
	Files    []File
}

func NewPackage(name string, binaries []Binary, files []File) (pkg Package, err error) {
	return Package{
		Name:     name,
		Binaries: binaries,
		Files:    files,
	}, nil
}
