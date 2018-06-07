package runner

import (
	"os"
	"regexp"

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
			sourcePath, err := r.getBinary(b)
			if err != nil {
				return err
			}
			binaryPath, err := r.extract(b, sourcePath)
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

func (r *Runner) getBinary(b config.Binary) (sourcePath string, err error) {
	sourcePath = b.Src
	remoteExp, err := regexp.Compile(`(http|https)://.*`)
	if err != nil {
		return "", err
	}
	if remoteExp.MatchString(sourcePath) {
		sourcePath = b.GetDownloadPath()
		err = download.Run(r.Context, b, sourcePath)
		if err != nil {
			return "", err
		}
	}
	return sourcePath, nil
}

func (r *Runner) extract(b config.Binary, sourcePath string) (binaryPath string, err error) {
	binaryPath, err = extract.Run(r.Context, b, sourcePath)
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
