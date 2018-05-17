package download

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"testing"
)

func TestDownloadDryRun(t *testing.T) {
	target := "/tmp/test.tar.gz"
	ctx := context.New(true)
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.tar.gz",
		Format: "tar.gz",
		Mode:   0700,
	}
	Run(&ctx, binary, target)
}
