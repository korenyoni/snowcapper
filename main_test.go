package main

import (
	"fmt"
	"testing"

	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"github.com/yonkornilov/snowcapper/runner"
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

func TestRunnerCommandErr(t *testing.T) {
	configYaml, _ := Asset("config.yaml")
	conf, _ := config.New(configYaml)
	ctx := context.New(true)
	ctx, err := context.CommandErr(ctx)
	if err != nil {
		t.Fatal(fmt.Sprint("Expecting no error, got \n%s", err))
	}
	r, err := runner.New(&ctx, conf)
	if err != nil {
		t.Fatal(fmt.Sprint("Expecting no error, got \n%s", err))
	}
	err = r.Run()
	if err == nil {
		t.Fatal("Expect error, got nothing")
	}
	t.Logf("Got expected command error: %s\n", err)

}
