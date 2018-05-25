package init

import (
	"errors"
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"os/exec"
	"strings"
)

func Run(c *context.Context, p config.Package) error {
	if c.IsDryRun {
		fmt.Printf("DRY-RUN: Initializing %s via \n", p.Name)
		fmt.Printf("DRY-RUN: Done.\n")
		return nil
	}
	for _, i := range p.Inits {
		fmt.Printf("Initializing %s with init type %s and content %s\n", p.Name, i.Type, i.Content)
		var out string
		var err error
		if i.Type == config.Command {
			out, err = initCommand(c, i)
			if err != nil {
				return err
			}
		} else if i.Type == config.OpenRC {
			out, err = initOpenRC(c, i)
			if err != nil {
				return err
			}
		} else {
			return errors.New(fmt.Sprint("Error: invalid init type: %s", i.Type))
		}
		fmt.Printf("Output: %s\n", out)
	}
	return nil
}

func initCommand(c *context.Context, i config.Init) (string, error) {
	splitContent := strings.Split(i.Content, " ")
	if c.IsDryRun {
		return fmt.Sprintf("%s", splitContent), nil
	}
	out, err := exec.Command(splitContent[0], splitContent[1:]...).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func initOpenRC(c *context.Context, i config.Init) (string, error) {
	args := [...]string{"/sbin/rc-update", "add", i.Content}
	if c.IsDryRun {
		return fmt.Sprintf("%s", args), nil
	}
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
