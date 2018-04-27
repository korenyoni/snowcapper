package runner

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/download"
	"github.com/yonkornilov/snowcapper/extract"
	"os"
)

func Run(c config.Config) error {
	for _, p := range c.Packages {
		downloadPath := getDownloadPath(p)
		filePath := getPath(p)
		err := download.Run(p, downloadPath)
		if err != nil {
			return err
		}
		err = extract.Run(p, downloadPath, filePath)
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

func getPath(p config.Package) string {
	return "/usr/bin/" + p.Name
}

func getDownloadPath(p config.Package) string {
	return "/tmp/" + p.Name + "." + p.Type
}
