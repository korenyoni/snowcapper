package main

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/runner"
)

func main() {
	c := config.Make()
	err := runner.Run(c)
	if err != nil {
		panic(err)
	}
}
