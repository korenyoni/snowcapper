package init

import (
	"errors"
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
)

func Run(c *context.Context, p config.Package) error {
	if c.IsDryRun {
		fmt.Printf("DRY-RUN: Initializing %s via \n", p.Name, "")
		fmt.Printf("DRY-RUN: Done.\n")
	} else {
		fmt.Printf("Initializing %s via \n", p.Name, "")
		err := errors.New("Empty")
		if err != nil {
			return err
		}
		fmt.Printf("Done.\n")
	}
	return nil
}
