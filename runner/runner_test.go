package runner

import (
	"testing"

	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
)

func TestRunnerDryRun(t *testing.T) {
	var packages []config.Package
	var binaries []config.Binary
	var files []config.File
	var inits []config.Init
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
	init := config.Init{
		Type:    "openrc",
		Content: "vault",
	}
	binaries = append(binaries, binary)
	files = append(files, file)
	inits = append(inits, init)
	packages = append(packages, config.Package{
		Name:     "test",
		Binaries: binaries,
		Files:    files,
		Inits:    inits,
	})
	conf := config.Config{
		Packages: packages,
	}
	runner := Runner{
		Config:  &conf,
		Context: &ctx,
	}
	err := runner.Run()
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}
