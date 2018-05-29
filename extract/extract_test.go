package extract

import (
	"testing"

	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
)

func TestExtractDryRunTar(t *testing.T) {
	ctx := context.New(true)
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.tar",
		Format: "tar",
		Mode:   0700,
	}
	_, err := Run(&ctx, binary, "/tmp/binary.tar")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestExtractDryRunRar(t *testing.T) {
	ctx := context.New(true)
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.rar",
		Format: "rar",
		Mode:   0700,
	}
	_, err := Run(&ctx, binary, "/tmp/binary.rar")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestExtractDryRunZip(t *testing.T) {
	ctx := context.New(true)
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.zip",
		Format: "zip",
		Mode:   0700,
	}
	_, err := Run(&ctx, binary, "/tmp/binary.zip")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestExtractDryRunTarGz(t *testing.T) {
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

func TestExtractDryRunTarBz2(t *testing.T) {
	ctx := context.New(true)
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.tar.bz2",
		Format: "tar.bz2",
		Mode:   0700,
	}
	_, err := Run(&ctx, binary, "/tmp/binary.tar.bz2")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestExtractDryRunTarXz(t *testing.T) {
	ctx := context.New(true)
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.tar.xz",
		Format: "tar.xz",
		Mode:   0700,
	}
	_, err := Run(&ctx, binary, "/tmp/binary.tar.xz")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestExtractDryRunTarLz4(t *testing.T) {
	ctx := context.New(true)
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.tar.lz4",
		Format: "tar.lz4",
		Mode:   0700,
	}
	_, err := Run(&ctx, binary, "/tmp/binary.tar.lz4")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestExtractDryRunTarSz(t *testing.T) {
	ctx := context.New(true)
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.tar.sz",
		Format: "tar.sz",
		Mode:   0700,
	}
	_, err := Run(&ctx, binary, "/tmp/binary.tar.sz")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestExtractDryRunUnsupported(t *testing.T) {
	ctx := context.New(true)
	binary := config.Binary{
		Name:   "test",
		Src:    "https://test.com/test.tar.fake",
		Format: "tar.fake",
		Mode:   0700,
	}
	_, err := Run(&ctx, binary, "/tmp/binary.tar.fake")
	if err == nil {
		t.Fatalf("Expected error, got nothing")
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
