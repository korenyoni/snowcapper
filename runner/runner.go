package runner

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/download"
	"os"
)

func Run(c config.Config) error {
	for _, p := range c.Packages {
		filePath := makePath(p.Name)
		err := download.Download(filePath, p.Source)
		if err != nil {
			return err
		}
		err = os.Chmod(filePath, 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func makePath(fileName string) string {
	return "/usr/bin/" + fileName
}
