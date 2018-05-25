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
    inits:
      - type: openrc
        content: vault
```

### Usage:

```
make get
make binary
```

### To test in an Alpine environment:

```
make
```

This builds the binary and provisions an Alpine VM using snowcapper and Vagrant.
