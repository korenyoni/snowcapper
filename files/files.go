package files

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"io/ioutil"
)

func Run(c *context.Context, f config.File) error {
	fmt.Printf("Writing to %s ... \n", f.Path)
	data := []byte(f.Content)
	err := ioutil.WriteFile(f.Path, data, f.Mode)
	if err != nil {
		return err
	}
	fmt.Printf("Done.\n")
	return nil
}
