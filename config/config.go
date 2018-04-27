package config

type PackageConfigFile struct {
	Path    string
	Content string
}

type Package struct {
	Name        string
	Source      string
	Type        string
	ConfigFiles []PackageConfigFile
}

type Config struct {
	Packages []Package
}

func Make() Config {
	packages := make([]Package, 0)
	packages = append(packages, Package{
		Name:   "vault",
		Source: "https://releases.hashicorp.com/vault/0.10.0/vault_0.10.0_linux_amd64.zip",
		Type:   "zip",
		ConfigFiles: []PackageConfigFile{PackageConfigFile{
			Path: "/etc/init.d/vault",
			Content: `
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
}`}}})
	return Config{
		Packages: packages,
	}
}
