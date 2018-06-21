.PHONY : binary vagrant clean

LATEST_TAG := $(shell git describe $(shell git rev-list --tags --max-count=1))

default: get binary vagrant

get:
	go get -u github.com/tmthrgd/go-bindata/...
	go-bindata config.yaml
	go get ./...

test:
	go test -v ./...

binary:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .

vagrant:
	vagrant up --provision

release: binary
	zip snowcapper_$(LATEST_TAG)_amd64.zip snowcapper && sha512sum snowcapper_$(LATEST_TAG)_amd64.zip > snowcapper_$(LATEST_TAG)_amd64.sha512

clean:
	vagrant destroy --force && rm -f snowcapper *.zip *.sha512
