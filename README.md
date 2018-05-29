### Snowcapper - Immutable config management for Alpine Linux [![CircleCI](https://circleci.com/gh/yonkornilov/snowcapper.svg?style=svg)](https://circleci.com/gh/yonkornilov/snowcapper)

![snowcapper](_images/snowcapper.png)

Snowcapper is a single binary for bootstrapping services onto an Alpine Linux image.

### Example Config:

```
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
    inits:
      - type: openrc
        content: vault
```

### Usage:

```
make get
make binary
./snowcapper
```

### To test in an Alpine environment:

```
make
```

This builds the binary and provisions an Alpine VM using snowcapper and Vagrant.
