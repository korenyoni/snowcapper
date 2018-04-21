package runner

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/download"
	"github.com/yonkornilov/snowcapper/extract"
	"os"
)

func Run(c config.Config) error {
	for _, p := range c.Packages {
		filePath := makePath(p.Name)
		err := download.Run(p.Name, filePath, p.Source)
		if err != nil {
			return err
		}
		err = extract.Run(p.Name, filePath)
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
