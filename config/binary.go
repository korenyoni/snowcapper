package config

type Binary struct {
	Name    string
	Src     string
	SrcType string
}

func (b *Binary) GetBinaryPath() string {
	return "/usr/bin/" + b.Name
}

func (b *Binary) GetDownloadPath() string {
	return "/tmp/" + b.Name + "." + b.SrcType
}
