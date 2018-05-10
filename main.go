package main

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/runner"
	"log"
)

func main() {
	c := config.Make()
	r := runner.Make(c)
	err := r.Run
	if err != nil {
		log.Fatal(err)
	}
}
