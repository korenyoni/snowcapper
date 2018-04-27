package main

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"ioutil"
)

func Run(p config.Package) {
	for packageConfigFile := range p.ConfigFiles {
		fmt.Printf("Writing to %s ...", packageConfigFile.Path)
		data := []byte(packageConfigFile.Content)
		file, err := ioutil.WriteFile(packageConfigFile.Path, 0644)
		defer file.Close()
		fmt.Printf("Done.")
	}
}
