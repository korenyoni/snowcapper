package inits

import (
	"fmt"
	"testing"

	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
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
	err := startOpenRC(&ctx, init)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckDaemon(t *testing.T) {
	ctx := context.New(true)
	init := config.Init{
		Type:    "openrc",
		Content: "vault",
	}
	pid, err := checkDaemon(&ctx, init)
	if err != nil {
		t.Fatal(err)
	}
	if pid != -1 {
		t.Fatalf("Expected pid -1, got pid %d\n", pid)
	}
}
