package main

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/runner"
)

func main() {
	configYaml, err := Asset("config.yaml")
	if err != nil {
		panic(err)
	}
	c, err := config.New(configYaml)
	if err != nil {
		panic(err)
	}
	r, err := runner.New(c)
	if err != nil {
		panic(err)
	}
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
