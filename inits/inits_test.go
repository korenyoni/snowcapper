package inits

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"testing"
)

func TestInitCommand(t *testing.T) {
	ctx := context.New(true)
	init := config.Init{
		Type:    "command",
		Content: "echo test test2",
	}
	args := [...]string{"echo", "test", "test2"}
	expectedOut := fmt.Sprintf("%s", args)
	out, err := initCommand(&ctx, init)
	if err != nil {
		t.Fatal(err)
	}
	if out != expectedOut {
		t.Fatalf("expected: %s, got %s", expectedOut, out)
	}
}

func TestInitOpenRC(t *testing.T) {
	ctx := context.New(true)
	init := config.Init{
		Type:    "openrc",
		Content: "vault",
	}
	args := [...]string{"rc-update", "add", "vault"}
	expectedOut := fmt.Sprintf("%s", args)
	out, err := initOpenRC(&ctx, init)
	if err != nil {
		t.Fatal(err)
	}
	if out != expectedOut {
		t.Fatalf("expected: %s, got %s", expectedOut, out)
	}
}

func TestStartOpenRC(t *testing.T) {
	ctx := context.New(true)
	init := config.Init{
		Type:    "openrc",
		Content: "vault",
	}
	args := [...]string{"openrc"}
	expectedOut := fmt.Sprintf("%s", args)
	out, err := startOpenRC(&ctx, init)
	if err != nil {
		t.Fatal(err)
	}
	if out != expectedOut {
		t.Fatalf("expected: %s, got %s", expectedOut, out)
	}
}
