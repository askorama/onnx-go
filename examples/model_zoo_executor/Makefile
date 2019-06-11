.PHONY: all

all: $(wildcard models/*)

models/%: .FORCE
	MODELDIR=$@ go test

.FORCE:
