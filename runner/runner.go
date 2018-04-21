package runner

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/download"
	"os"
)

func Run(c config.Config) error {
	for _, p := range c.Packages {
		filePath := makePath(p.Name)
		download.Download(filePath, p.Source)
		os.Chmod(filePath, 0700)
	}

	return nil
}

func makePath(fileName string) string {
	return "/usr/bin/" + fileName
}
