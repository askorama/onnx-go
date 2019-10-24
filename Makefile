GOFILES := $(wildcard *.go)
TESTFILES := $(wildcard *test.go)
get ?= ./...
ONNX_COVERAGE ?= /tmp/report.md
TEST_LOCATION ?= backend/x/gorgonnx
COVER_TARGET ?= ${TEST_LOCATION}

help:
	@echo "\t\t\tONNX MAKE TASKS\n"
	@echo "\tinstall\t\t Get all dependencies (go get ./..)\n"
	@echo "\tgo get\t\t Run go get with the -get- variable as param (get=./.. by default)"
	@echo "\t\t\t Example: make install get=github.com/foo/bar\n"
	@echo "\tdoc\t\t Execute embedmd command with (-w README.md) as parameter\n"
	@echo "\tcoverage\t Create a coverage report in the -COVER_TARGET- location. You can change it with COVER_TARGET variable."
	@echo "\t\t\t Example: make coverage COVER_TARGET=/tmp/report.md"

## install: Install missing dependencies. 
## Runs `go get` internally. e.g; 
install: go-get

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@go get -v ${get}

doc: 
	@go get -v github.com/campoy/embedmd
	$(GOPATH)/bin/embedmd ${TESTFILES}
	@$(GOPATH)/bin/embedmd -w README.md	

coverage:
	@cd ${TEST_LOCATION} && \
	ONNX_COVERAGE=${ONNX_COVERAGE} go test && \
	cp ${ONNX_COVERAGE} ONNX_COVERAGE.md

all: install go-get doc
