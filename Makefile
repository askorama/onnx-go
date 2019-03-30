GOFILES := $(wildcard *.go)

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@go get $(get)

doc: $(GOPATH)/bin/embedmd $(wildcard *test.go)
	@$(GOPATH)/bin/embedmd -w README.md	

$(GOBIN)/embedmd:
	@go get github.com/campoy/embedmd

all:
