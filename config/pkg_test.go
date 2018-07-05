package config

import (
	"fmt"
	"testing"
)

func TestConfigMissingName(t *testing.T) {
	config := `
packages:
  - binaries:
      - name: vault
        mode: 0755
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

func TestConfigMissingBinaries(t *testing.T) {
	config := `
packages:
  - name: vault
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
	if err != nil {
		t.Fatal("Cannot validate config: " + err.Error())
	}
}

func TestConfigMissingFiles(t *testing.T) {
	config := `
packages:
  - name: vault
    binaries:
      - name: vault
        mode: 0755
        src: https://releases.hashicorp.com/vault/0.10.0/vault_0.10.0_linux_amd64.zip
        format: zip
    services:
      - binary: vault
        args:
          - "server"
          - "-config /etc/vault/config.hcl"
`
	_, err := New([]byte(config))
	if err != nil {
		t.Fatal("Cannot validate config: " + err.Error())
	}
}

func TestConfigMissingServices(t *testing.T) {
	config := `
packages:
  - name: vault
    binaries:
      - name: vault
        mode: 0755
        src: https://releases.hashicorp.com/vault/0.10.0/vault_0.10.0_linux_amd64.zip
        format: zip
    files:
      - path: /etc/vault/config.hcl
        mode: 0700
        content: |
          storage "file" {
            path    = "/mnt/vault/data"
          }

          listener "tcp" {
            address     = "0.0.0.0:8200"
            tls_disable = 1
          }
`
	_, err := New([]byte(config))
	if err != nil {
		t.Fatal("Cannot validate config: " + err.Error())
	}
}

func TestConfigEmptyPackage(t *testing.T) {
	config := `
packages:
  - name: vault
`
	_, err := New([]byte(config))
	if err == nil {
		t.Fatal("Expected validation error, got nothing.")
	}
	t.Log(fmt.Sprintf("Got expected error: \n%s", err.Error()))
}

func TestConfigOnlyBinaries(t *testing.T) {
	config := `
packages:
  - name: vault
    binaries:
      - name: vault
        mode: 0755
        src: https://releases.hashicorp.com/vault/0.10.0/vault_0.10.0_linux_amd64.zip
        format: zip
`
	_, err := New([]byte(config))
	if err != nil {
		t.Fatal("Cannot validate config: " + err.Error())
	}
}

func TestConfigOnlyFiles(t *testing.T) {
	config := `
packages:
  - name: vault
    files:
      - path: /etc/vault/config.hcl
        mode: 0700
        content: |
          storage "file" {
            path    = "/mnt/vault/data"
          }

          listener "tcp" {
            address     = "0.0.0.0:8200"
            tls_disable = 1
          }
`
	_, err := New([]byte(config))
	if err != nil {
		t.Fatal("Cannot validate config: " + err.Error())
	}
}

func TestConfigOnlyServices(t *testing.T) {
	config := `
packages:
  - name: vault
    services:
      - binary: vault
        args:
          - "server"
          - "-config /etc/vault/config.hcl"
`
	_, err := New([]byte(config))
	if err != nil {
		t.Fatal("Cannot validate config: " + err.Error())
	}
}

func TestConfigOnlyInits(t *testing.T) {
	config := `
packages:
  - name: vault
    inits:
      - type: command
        content: echo test
`
	_, err := New([]byte(config))
	if err != nil {
		t.Fatal("Cannot validate config: " + err.Error())
	}
}
