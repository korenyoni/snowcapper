package main

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"github.com/yonkornilov/snowcapper/runner"
)

func main() {
	configYaml, err := Asset("config.yaml")
	if err != nil {
		panic(err)
	}
	conf, err := config.New(configYaml)
	if err != nil {
		panic(err)
	}
	ctx := context.New(false)
	r, err := runner.New(&ctx, conf)
	if err != nil {
		panic(err)
	}
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
