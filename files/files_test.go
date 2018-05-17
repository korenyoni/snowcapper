package files

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"testing"
)

func TestFilesDryRun(t *testing.T) {
	ctx := context.New(true)
	file := config.File{
		Path: "/tmp/test",
		Mode: 0600,
		Content: `
		echo test
		`,
	}
	Run(&ctx, file)
}
