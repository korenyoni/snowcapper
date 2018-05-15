package config

import (
	"fmt"
	"testing"
)

func TestConfigBinaryMissingName(t *testing.T) {
	config := `
packages:
  - name: vault
    binaries:
      - mode: 0755
        src: https://releases.hashicorp.com/vault/0.10.0/vault_0.10.0_linux_amd64.zip
        format: zip
    files:
      - path: /etc/init.d/vault
        mode: 0700
        content: |
          #!/sbin/openrc-run

          NAME=vault
          DAEMON=/usr/bin/$NAME

          depend() {
                  need net
                  after firewall
          }

          start() {
                  ebegin "Starting ${NAME}"
                          start-stop-daemon --start \
                                  --background \
                                  --make-pidfile --pidfile /var/run/$NAME.pid \
                                  --stderr "/var/log/$NAME.log" \
                                  --stdout "/var/log/$NAME.log" \
                                  --user $USER \
                                  --exec $DAEMON \
                                  -- \
                                  -config /etc/vault/config.hcl
                  eend $?
          }

          stop () {
                  ebegin "Stopping ${NAME}"
                          start-stop-daemon --stop \
                                  --pidfile /var/run/$NAME.pid \
                                  --user $USER \
                                  --exec $DAEMON
                  eend $?
          }
`
	_, err := New([]byte(config))
	if err == nil {
		t.Fatal("Expected validation error, got nothing.")
	}
	t.Log(fmt.Sprintf("Got expected error: \n%s", err.Error()))
}

func TestConfigBinaryMissingMode(t *testing.T) {
	config := `
packages:
  - name: vault
    binaries:
      - name: vault
        src: https://releases.hashicorp.com/vault/0.10.0/vault_0.10.0_linux_amd64.zip
        format: zip
    files:
      - path: /etc/init.d/vault
        mode: 0700
        content: |
          #!/sbin/openrc-run

          NAME=vault
          DAEMON=/usr/bin/$NAME

          depend() {
                  need net
                  after firewall
          }

          start() {
                  ebegin "Starting ${NAME}"
                          start-stop-daemon --start \
                                  --background \
                                  --make-pidfile --pidfile /var/run/$NAME.pid \
                                  --stderr "/var/log/$NAME.log" \
                                  --stdout "/var/log/$NAME.log" \
                                  --user $USER \
                                  --exec $DAEMON \
                                  -- \
                                  -config /etc/vault/config.hcl
                  eend $?
          }

          stop () {
                  ebegin "Stopping ${NAME}"
                          start-stop-daemon --stop \
                                  --pidfile /var/run/$NAME.pid \
                                  --user $USER \
                                  --exec $DAEMON
                  eend $?
          }
`
	_, err := New([]byte(config))
	if err == nil {
		t.Fatal("Expected validation error, got nothing.")
	}
	t.Log(fmt.Sprintf("Got expected error: \n%s", err.Error()))
}

func TestConfigMissingSrc(t *testing.T) {
	config := `
packages:
  - name: vault
    binaries:
      - name: vault
        mode: 0755
        format: zip
    files:
      - path: /etc/init.d/vault
        mode: 0700
        content: |
          #!/sbin/openrc-run

          NAME=vault
          DAEMON=/usr/bin/$NAME

          depend() {
                  need net
                  after firewall
          }

          start() {
                  ebegin "Starting ${NAME}"
                          start-stop-daemon --start \
                                  --background \
                                  --make-pidfile --pidfile /var/run/$NAME.pid \
                                  --stderr "/var/log/$NAME.log" \
                                  --stdout "/var/log/$NAME.log" \
                                  --user $USER \
                                  --exec $DAEMON \
                                  -- \
                                  -config /etc/vault/config.hcl
                  eend $?
          }

          stop () {
                  ebegin "Stopping ${NAME}"
                          start-stop-daemon --stop \
                                  --pidfile /var/run/$NAME.pid \
                                  --user $USER \
                                  --exec $DAEMON
                  eend $?
          }
`
	_, err := New([]byte(config))
	if err == nil {
		t.Fatal("Expected validation error, got nothing.")
	}
	t.Log(fmt.Sprintf("Got expected error: \n%s", err.Error()))
}

func TestConfigMissingFormat(t *testing.T) {
	config := `
packages:
  - name: vault
    binaries:
      - name: vault
        mode: 0755
        src: https://releases.hashicorp.com/vault/0.10.0/vault_0.10.0_linux_amd64.zip
    files:
      - path: /etc/init.d/vault
        mode: 0700
        content: |
          #!/sbin/openrc-run

          NAME=vault
          DAEMON=/usr/bin/$NAME

          depend() {
                  need net
                  after firewall
          }

          start() {
                  ebegin "Starting ${NAME}"
                          start-stop-daemon --start \
                                  --background \
                                  --make-pidfile --pidfile /var/run/$NAME.pid \
                                  --stderr "/var/log/$NAME.log" \
                                  --stdout "/var/log/$NAME.log" \
                                  --user $USER \
                                  --exec $DAEMON \
                                  -- \
                                  -config /etc/vault/config.hcl
                  eend $?
          }

          stop () {
                  ebegin "Stopping ${NAME}"
                          start-stop-daemon --stop \
                                  --pidfile /var/run/$NAME.pid \
                                  --user $USER \
                                  --exec $DAEMON
                  eend $?
          }
`
	_, err := New([]byte(config))
	if err == nil {
		t.Fatal("Expected validation error, got nothing.")
	}
	t.Log(fmt.Sprintf("Got expected error: \n%s", err.Error()))
}
