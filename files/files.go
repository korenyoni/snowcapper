package files

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"io/ioutil"
)

func Run(p config.Package) error {
	for _, packageConfigFile := range p.ConfigFiles {
		fmt.Printf("Writing to %s ... \n", packageConfigFile.Path)
		data := []byte(packageConfigFile.Content)
		err := ioutil.WriteFile(packageConfigFile.Path, data, 0644)
		if err != nil {
			return err
		}
		fmt.Printf("Done.\n")
	}
	return nil
}
