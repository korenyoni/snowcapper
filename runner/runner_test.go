package runner

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"testing"
)

func TestRunnerDryRun(t *testing.T) {
	var packages []config.Package
	var binaries []config.Binary
	var files []config.File
	ctx := context.New(true)
	file := config.File{
		Path: "/tmp/test",
		Mode: 0600,
		Content: `
		echo test
		`,
	}
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.tar.gz",
		Format: "tar.gz",
		Mode:   0700,
	}
	binaries = append(binaries, binary)
	files = append(files, file)
	packages = append(packages, config.Package{
		Name:     "test",
		Binaries: binaries,
		Files:    files,
	})
	conf := config.Config{
		Packages: packages,
	}
	runner := Runner{
		Config:     &conf,
		Context:    &ctx,
		BinaryMode: 0700,
	}
	runner.Run()
}
