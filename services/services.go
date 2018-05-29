package services

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"io/ioutil"
	"strings"
)

const (
	binaryTemplate  string = "$$BINARY"
	ArgsTemplate    string = "$$ARGS"
	serviceTemplate string = `#!/sbin/openrc-run

NAME=$$BINARY

supervisor=supervise-daemon
command=$NAME
command_args="$$ARGS"
command_background="Yes"
pidfile=/var/run/$NAME
`
)

func Run(c *context.Context, p config.Package) error {
	for _, s := range p.Services {
		err := createService(c, s)
		if err != nil {
			return err
		}
	}
	return nil
}

func createService(c *context.Context, s config.Service) error {
	var argsReplace string
	for _, arg := range s.Args {
		argsReplace = argsReplace + arg + " "
	}
	argsReplace = strings.TrimRight(argsReplace, " ")
	serviceContent := serviceTemplate
	serviceContent = strings.Replace(serviceContent, binaryTemplate, s.Binary, 1)
	serviceContent = strings.Replace(serviceContent, ArgsTemplate, argsReplace, 1)
	if c.IsDryRun {
		fmt.Printf("DRY-RUN: Writing %s to /etc/init.d/\n", s.Binary)
		fmt.Printf("DRY-RUN: %s\n", serviceContent)
		return nil
	}
	fmt.Printf("Writing %s to /etc/init.d/\n", s.Binary)
	fmt.Printf("%s\n", serviceContent)
	err := ioutil.WriteFile("/etc/init.d/"+s.Binary, []byte(serviceContent), 0700)
	if err != nil {
		return err
	}
	return nil
}
