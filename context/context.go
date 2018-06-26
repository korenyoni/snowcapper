package context

import (
	"errors"
)

const (
	DefaultDryrun    string = "default"
	CommandErrDryrun string = "commandErr"
)

type Context struct {
	IsDryRun   bool
	DryRunType string
}

func New(isDryRun bool) Context {
	return Context{
		DryRunType: DefaultDryrun,
		IsDryRun:   isDryRun,
	}
}

func CommandErr(c Context) (Context, error) {
	if !c.IsDryRun {
		return c, errors.New("Context must be dry-run to set dry-run type")
	}
	return Context{
		IsDryRun:   c.IsDryRun,
		DryRunType: CommandErrDryrun,
	}, nil
}
