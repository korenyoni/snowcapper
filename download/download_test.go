package download

import (
	"testing"

	"github.com/yonkornilov/snowcapper/testasset"
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
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
	err := Run(&ctx, binary, target)
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestCheckHashIfExistsSha512(t *testing.T) {
	testBinary, err := testasset.Asset("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = checkHashIfExists(testBinary, "0e3e75234abc68f4378a86b3f4b32a198ba301845b0cd6e50106e874345700cc6663a86c1ea125dc5e92be17c98f9a0f85ca9d5f595db2012f7cc3571945c123")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestCheckHashIfExistsSha512Bad(t *testing.T) {
	testBinary, err := testasset.Asset("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = checkHashIfExists(testBinary, "fe3e75234abc68f4378a86b3f4b32a198ba301845b0cd6e50106e874345700cc6663a86c1ea125dc5e92be17c98f9a0f85ca9d5f595db2012f7cc3571945c123")
	if err == nil {
		t.Fatalf("Expected error, got nothing")
	}
}

func TestCheckHashIfExistsSha384(t *testing.T) {
	testBinary, err := testasset.Asset("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = checkHashIfExists(testBinary, "109bb6b5b6d5547c1ce03c7a8bd7d8f80c1cb0957f50c4f7fda04692079917e4f9cad52b878f3d8234e1a170b154b72d")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestCheckHashIfExistsSha384Bad(t *testing.T) {
	testBinary, err := testasset.Asset("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = checkHashIfExists(testBinary, "f09bb6b5b6d5547c1ce03c7a8bd7d8f80c1cb0957f50c4f7fda04692079917e4f9cad52b878f3d8234e1a170b154b72d")
	if err == nil {
		t.Fatalf("Expected error, got nothing")
	}
}

func TestCheckHashIfExistsSha256(t *testing.T) {
	testBinary, err := testasset.Asset("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = checkHashIfExists(testBinary, "f2ca1bb6c7e907d06dafe4687e579fce76b37e4e93b7605022da52e6ccc26fd2")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}

func TestCheckHashIfExistsSha256Bad(t *testing.T) {
	testBinary, err := testasset.Asset("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = checkHashIfExists(testBinary, "a2ca1bb6c7e907d06dafe4687e579fce76b37e4e93b7605022da52e6ccc26fd2")
	if err == nil {
		t.Fatalf("Expected error, got nothing")
	}
}
