package runner

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/download"
	"github.com/yonkornilov/snowcapper/extract"
	"github.com/yonkornilov/snowcapper/files"
	"os"
)

type Runner struct {
	Config     *config.Config
	BinaryMode os.FileMode
}

func (r *Runner) Run() error {
	c := r.Config
	for _, p := range c.Packages {
		downloadPath, err := r.download(p)
		if err != nil {
			return err
		}
		binaryPath, err := r.extract(p, downloadPath)
		if err != nil {
			return err
		}
		err = r.chmodBinary(binaryPath, r.BinaryMode)
		if err != nil {
			return err
		}
		err = r.copyConfigFiles(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Runner) download(p config.Package) (downloadPath string, err error) {
	downloadPath = p.GetDownloadPath()
	err = download.Run(p, downloadPath)
	if err != nil {
		return "", err
	}
	return downloadPath, nil
}

func (r *Runner) extract(p config.Package, downloadPath string) (binaryPath string, err error) {
	binaryPath = p.GetBinaryPath()
	err = extract.Run(p, downloadPath, binaryPath)
	if err != nil {
		return "", err
	}
	return binaryPath, nil
}

func (r *Runner) chmodBinary(binaryPath string, mode os.FileMode) (err error) {
	err = os.Chmod(binaryPath, mode)
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) copyConfigFiles(p config.Package) (err error) {
	err = files.Run(p)
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
