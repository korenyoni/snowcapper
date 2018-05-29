package files

import (
	"testing"

	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
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
	err := Run(&ctx, file)
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}
