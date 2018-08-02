package runner

import (
	"os"
	"regexp"
	"io/ioutil"
	"errors"

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
	for _, e := range c.Extends {
		extendDownloadPath, err := r.getExtend(e)
		if err != nil {
			return err
		}
		extendConfig, err := r.createConfigFromExtend(extendDownloadPath)
		if err != nil {
			return err
		}
		extendRunner, err := New(r.Context, extendConfig)
		if err != nil {
			return err
		}
		extendRunner.Run()
		if err != nil {
			return err
		}
	}
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

func (r *Runner) getBinary(b config.Binary) (downloadPath string, err error) {
	remoteExp, err := regexp.Compile(`(http|https)://.*`)
	if err != nil {
		return "", err
	}
	if remoteExp.MatchString(b.Src) {
		downloadPath, err = download.Run(r.Context, download.DownloadableHolder{
			BinaryPointer: &b,
			Downloadable: b,
		})
		if err != nil {
			return "", err
		}
	}
	return downloadPath, nil
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

func (r *Runner) getExtend(e config.Extend) (downloadPath string, err error) {
	remoteExp, err := regexp.Compile(`(http|https)://.*\.snc`)
	if err != nil {
		return "", err
	}
	localExp, err := regexp.Compile(`.*\.snc`)
	if err != nil {
		return "", err
	}
	if remoteExp.MatchString(e.Src) {
		downloadPath, err = download.Run(r.Context, download.DownloadableHolder{
			ExtendPointer: &e,
			Downloadable: e,
		})
		if err != nil {
			return "", err
		}
	} else if !localExp.MatchString(e.Src) {
		return "", errors.New("Extend source is neither a local or remote *.snc file")
	} else {
		downloadPath = e.Src
		_, err = os.Stat(downloadPath)
		if err != nil {
			return "", err
		}
	}
	return downloadPath, nil
}

func (r *Runner) createConfigFromExtend(downloadPath string) (c config.Config, err error) {
	if r.Context.IsDryRun {
		return c, nil
	}
	extendConfigBytes, err := ioutil.ReadFile(downloadPath)
	if err != nil {
		return c, err
	}
	c, err = config.New(extendConfigBytes)
	if err != nil {
		return c, err
	}
	return c, nil
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
