package inits

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
)

func Run(c *context.Context, p config.Package) error {
	for _, i := range p.Inits {
		if c.IsDryRun {
			fmt.Printf("DRY-RUN: Initializing %s with init type %s and content %s\n", p.Name, i.Type, i.Content)
		} else {
			fmt.Printf("Initializing %s with init type %s and content %s\n", p.Name, i.Type, i.Content)
		}
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
			err = startOpenRC(c, i)
			if err != nil {
				return err
			}
			pid, err := checkDaemon(c, i)
			if err != nil {
				return err
			}
			fmt.Printf("Service %s is running with pid %d", i.Content, pid)
		} else {
			return errors.New(fmt.Sprint("Error: invalid init type: %s", i.Type))
		}
		if c.IsDryRun {
			fmt.Printf("DRY-RUN: Output: %s\n", out)
		} else {
			fmt.Printf("Output: %s\n", out)
		}
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
	args := [...]string{"rc-update", "add", i.Content}
	if c.IsDryRun {
		return fmt.Sprintf("%s", args), nil
	}
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func startOpenRC(c *context.Context, i config.Init) error {
	args := [...]string{"rc-service", i.Content, "start"}
	if c.IsDryRun {
		return nil
	}
	err := exec.Command(args[0], args[1:]...).Start()
	if err != nil {
		return err
	}
	return nil
}

func checkDaemon(c *context.Context, i config.Init) (int, error) {
	args := [...]string{"pidof", i.Content}
	if c.IsDryRun {
		return -1, nil
	}
	pidofOut, err := exec.Command(args[0], args[1:]...).Output()
	pidString := strings.Trim(string(pidofOut[:]), "\n")
	if err != nil {
		return -1, err
	}
	pid, err := strconv.Atoi(pidString)
	if err != nil {
		return -1, err
	}
	return pid, nil
}
