package files

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"io/ioutil"
)

func Run(c *context.Context, f config.File) error {
	if c.IsDryRun {
		fmt.Printf("DRY-RUN: Writing to %s ... \n", f.Path)
		fmt.Printf("DRY-RUN: Done.\n")

		return nil
	}
	fmt.Printf("Writing to %s ... \n", f.Path)
	data := []byte(f.Content)
	err := ioutil.WriteFile(f.Path, data, f.Mode)
	if err != nil {
		return err
	}
	fmt.Printf("Done.\n")
	return nil
}
