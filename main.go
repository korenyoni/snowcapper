package main

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/runner"
)

func main() {
	c := config.New()
	r := runner.New(c)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
