package extract

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"testing"
)

func TestExtractDryRun(t *testing.T) {
	ctx := context.New(true)
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.tar.gz",
		Format: "tar.gz",
		Mode:   0700,
	}
	_, err := Run(&ctx, binary, "/tmp/binary.tar.gz")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestExtractGetExtractedPath(t *testing.T) {
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.tar.gz",
		Format: "tar.gz",
		Mode:   0700,
	}
	downloadPath := "/tmp/test.tar.gz"
	extractedPath := getExtractedPath(binary.Format, downloadPath)
	expectedExtractedPath := "/tmp/test"
	if extractedPath != expectedExtractedPath {
		t.Fatalf("Expected %s, got %s", expectedExtractedPath, extractedPath)
	}
}

func TestExtractGetExtractedBinaryPath(t *testing.T) {
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.tar.gz",
		Format: "tar.gz",
		Mode:   0700,
	}
	downloadPath := "/tmp/test.tar.gz"
	extractedPath := getExtractedPath(binary.Format, downloadPath)
	extractedBinaryPath := getExtractedBinaryPath(binary, extractedPath)
	expectedExtractedBinaryPath := "/tmp/test/test"
	if extractedBinaryPath != expectedExtractedBinaryPath {
		t.Fatalf("Expected %s, got %s", expectedExtractedBinaryPath, extractedBinaryPath)
	}
}
