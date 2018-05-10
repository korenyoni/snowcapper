package runner

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/download"
	"github.com/yonkornilov/snowcapper/extract"
	"os"
)

type Runner struct {
	Config       *config.Config
	downloadPath string
	binaryPath   string
}

func (r *Runner) Run() error {
	c := r.Config
	for _, p := range c.Packages {
		downloadPath := p.GetDownloadPath()
		filePath := p.GetBinaryPath()
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

func Make(config config.Config) Runner {
	return Runner{
		Config: &config,
	}
}
