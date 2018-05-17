package context

import (
	"testing"
)

func TestContext(t *testing.T) {
	ctx := New(false)
	isDryRun := ctx.IsDryRun
	if isDryRun {
		t.Fatalf("Expecting %s, got %s", false, isDryRun)
	}
}
