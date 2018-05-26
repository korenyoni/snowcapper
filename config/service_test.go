package config

import (
	"fmt"
	"testing"
)

func TestConfigFileMissingBinary(t *testing.T) {
	config := `
packages:
  - name: vault
    binaries:
      - name: vault
        mode: 0755
        src: https://releases.hashicorp.com/vault/0.10.0/vault_0.10.0_linux_amd64.zip
        format: zip
    services:
      - args:
          - "server"
          - "-config /etc/vault/config.hcl"
`
	_, err := New([]byte(config))
	if err == nil {
		t.Fatal("Expected validation error, got nothing.")
	}
	t.Log(fmt.Sprintf("Got expected error: \n%s", err.Error()))
}
