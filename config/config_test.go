package config

import (
	"fmt"
	"testing"
)

func TestConfigGood(t *testing.T) {
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

func TestConfigMissingPackages(t *testing.T) {
	config := `
`
	_, err := New([]byte(config))
	if err == nil {
		t.Fatal("Expected validation error, got nothing.")
	}
	t.Log(fmt.Sprintf("Got expected error: \n%s", err.Error()))
}
