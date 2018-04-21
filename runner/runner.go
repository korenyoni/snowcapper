package runner

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/download"
)

func Run(c config.Config) error {
	for _, p := range c.Packages {
		download.Download(makePath(p.Name), p.Source)
	}

	return nil
}

func makePath(fileName string) string {
	return "/usr/bin/" + fileName
}
