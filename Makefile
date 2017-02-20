export GOPATH:=$(CURDIR)
all: install

fmt:
	gofmt -l -w -s src/

dep:fmt
	go get github.com/j-keck/arping
	go get github.com/franela/goreq

install:dep
	go install main
