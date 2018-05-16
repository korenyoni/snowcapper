package main

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"github.com/yonkornilov/snowcapper/runner"
	"testing"
)

func TestConfigBinData(t *testing.T) {
	_, err := Asset("config.yaml")
	if err != nil {
		t.Fatal(fmt.Sprint("Expecting no error, got \n%s", err))
	}
}

func TestConfigNew(t *testing.T) {
	configYaml, _ := Asset("config.yaml")
	_, err := config.New(configYaml)
	if err != nil {
		t.Fatal(fmt.Sprint("Expecting no error, got \n%s", err))
	}
}

func TestRunnerNew(t *testing.T) {
	configYaml, _ := Asset("config.yaml")
	conf, _ := config.New(configYaml)
	ctx := context.New(true)
	_, err := runner.New(&ctx, conf)
	if err != nil {
		t.Fatal(fmt.Sprint("Expecting no error, got \n%s", err))
	}
}
