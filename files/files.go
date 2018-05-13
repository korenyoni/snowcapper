package files

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"io/ioutil"
)

func Run(f config.File) error {
	fmt.Printf("Writing to %s ... \n", f.Path)
	data := []byte(f.Content)
	err := ioutil.WriteFile(f.Path, data, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Done.\n")
	return nil
}
