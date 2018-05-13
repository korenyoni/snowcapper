.PHONY : binary vagrant clean

default: binary vagrant

get:
	go get ./...
	go get -u github.com/tmthrgd/go-bindata/...

test:
	go test -v ./...

binary:
	go-bindata config.yaml
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .

vagrant:
	vagrant up --provision

clean:
	vagrant destroy --force && rm -f snowcapper
