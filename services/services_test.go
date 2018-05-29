package services

import (
	"testing"

	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
)

func TestServicesDryRun(t *testing.T) {
	args := []string{
		"server",
		"-config /etc/vault/config.hcl",
	}
	ctx := context.New(true)
	service := config.Service{
		Binary: "vault",
		Args:   args,
	}
	services := []config.Service{service}
	pkg := config.Package{
		Services: services,
	}
	err := Run(&ctx, pkg)
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
}
