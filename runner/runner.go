package runner

import (
	"os"

	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"github.com/yonkornilov/snowcapper/download"
	"github.com/yonkornilov/snowcapper/extract"
	"github.com/yonkornilov/snowcapper/files"
	"github.com/yonkornilov/snowcapper/inits"
	"github.com/yonkornilov/snowcapper/services"
)

type Runner struct {
	Config  *config.Config
	Context *context.Context
}

func (r *Runner) Run() (err error) {
	c := r.Config
	for _, p := range c.Packages {
		for _, b := range p.Binaries {
			downloadPath, err := r.download(b)
			if err != nil {
				return err
			}
			binaryPath, err := r.extract(b, downloadPath)
			if err != nil {
				return err
			}
			err = r.chmodBinary(binaryPath, b.Mode)
			if err != nil {
				return err
			}
		}
		for _, f := range p.Files {
			err = r.copyConfigFiles(f)
			if err != nil {
				return err
			}
		}
		err = r.applyServices(p)
		if err != nil {
			return err
		}
		err = r.applyInits(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Runner) download(b config.Binary) (downloadPath string, err error) {
	downloadPath = b.GetDownloadPath()
	err = download.Run(r.Context, b, downloadPath)
	if err != nil {
		return "", err
	}
	return downloadPath, nil
}

func (r *Runner) extract(b config.Binary, downloadPath string) (binaryPath string, err error) {
	binaryPath, err = extract.Run(r.Context, b, downloadPath)
	if err != nil {
		return "", err
	}
	return binaryPath, nil
}

func (r *Runner) chmodBinary(binaryPath string, mode os.FileMode) (err error) {
	if r.Context.IsDryRun {
		// Do nothing
		return nil
	}
	err = os.Chmod(binaryPath, mode)
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) copyConfigFiles(f config.File) (err error) {
	err = files.Run(r.Context, f)
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) applyServices(p config.Package) (err error) {
	err = services.Run(r.Context, p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) applyInits(p config.Package) (err error) {
	err = inits.Run(r.Context, p)
	if err != nil {
		return err
	}
	return nil
}

func New(context *context.Context, config config.Config) (runner Runner, err error) {
	return Runner{
		Context: context,
		Config:  &config,
	}, nil
}
