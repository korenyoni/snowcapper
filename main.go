package main

import (
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"github.com/yonkornilov/snowcapper/runner"

	"log"
	"os/exec"
)

const (
	runtimeErrFormat string = "Runtime error: %s\n"
)

func main() {
	configYaml, err := Asset("config.yaml")
	if err != nil {
		log.Fatalf(runtimeErrFormat, err)
	}
	conf, err := config.New(configYaml)
	if err != nil {
		log.Fatalf(runtimeErrFormat, err)
	}
	ctx := context.New(false)
	r, err := runner.New(&ctx, conf)
	if err != nil {
		log.Fatalf(runtimeErrFormat, err)
	}
	err = r.Run()
	if err != nil {
		switch err.(type) {
		default:
			log.Fatalf(runtimeErrFormat, err)
		case *exec.ExitError:
			log.Fatalf(runtimeErrFormat, err.(*exec.ExitError).Stderr)
		}
	}
}
