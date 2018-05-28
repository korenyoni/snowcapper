package services

import (
	"fmt"
	"github.com/yonkornilov/snowcapper/config"
	"github.com/yonkornilov/snowcapper/context"
	"io/ioutil"
	"strings"
)

const (
	argSpacing      string = "                        "
	binaryTemplate  string = "$$BINARY"
	ArgsTemplate    string = "$$ARGS"
	serviceTemplate string = `#!/sbin/openrc-run

NAME=$$BINARY
DAEMON=/usr/bin/$NAME

supervisor=s6

start() {
        ebegin "Starting ${NAME}"
                start-stop-daemon --start \
                        --background \
                        --make-pidfile --pidfile /var/run/$NAME.pid \
                        --stderr "/var/log/$NAME.log" \
                        --stdout "/var/log/$NAME.log" \
                        --exec $DAEMON \
                        -- \$$ARGS
        eend $?
}

stop () {
        ebegin "Stopping ${NAME}"
                start-stop-daemon --stop \
                        --pidfile /var/run/$NAME.pid \
                        --exec $DAEMON
        eend $?
}
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
		argsReplace = argsReplace + "\n" + argSpacing + arg + "\\"
	}
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
