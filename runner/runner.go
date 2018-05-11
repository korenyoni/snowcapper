package runner

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/download"
	"github.com/yonkornilov/snowcapper/extract"
	"os"
)

type Runner struct {
	Config     *config.Config
	BinaryMode os.FileMode
}

func (r *Runner) Run() error {
	c := r.Config
	for _, p := range c.Packages {
		downloadPath, err := r.Download(p)
		if err != nil {
			return err
		}
		binaryPath, err := r.Extract(p, downloadPath)
		if err != nil {
			return err
		}
		err = r.ChmodBinary(binaryPath, r.BinaryMode)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Runner) Download(p config.Package) (downloadPath string, err error) {
	downloadPath = p.GetDownloadPath()
	err = download.Run(p, downloadPath)
	if err != nil {
		return "", err
	}
	return downloadPath, nil
}

func (r *Runner) Extract(p config.Package, downloadPath string) (binaryPath string, err error) {
	binaryPath = p.GetBinaryPath()
	err = extract.Run(p, downloadPath, binaryPath)
	if err != nil {
		return "", err
	}
	return binaryPath, nil
}

func (r *Runner) ChmodBinary(binaryPath string, mode os.FileMode) (err error) {
	err = os.Chmod(binaryPath, mode)
	if err != nil {
		return err
	}
	return nil
}

func Make(config config.Config) Runner {
	return Runner{
		Config:     &config,
		BinaryMode: 0700,
	}
}
