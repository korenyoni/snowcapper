package init

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"os/exec"
)

func Run(c *context.Context, p config.Package) error {
	if c.IsDryRun {
		fmt.Printf("DRY-RUN: Initializing %s via \n", p.Name, "")
		fmt.Printf("DRY-RUN: Done.\n")
		return nil
	}
	for _, i := range p.Inits {
		fmt.Printf("Initializing %s with init type %s and content %s\n", p.Name, i.Type, i.Content)
		var out string
		var err error
		if i.Type == config.Command {
			out, err = initCommand(i)
		}
		out, err = initOpenRC(i)
		if err != nil {
			return err
		}
		fmt.Printf("Output: %s\n", out)
	}
	return nil
}

func initCommand(i config.Init) (string, error) {
	out, err := exec.Command(i.Content).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func initOpenRC(i config.Init) (string, error) {
	out, err := exec.Command("rc-update add " + i.Content).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
