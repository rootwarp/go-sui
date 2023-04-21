.PHONY: test cover

all: test cover

test:
	echo unittesting...
	@go test -v ./rpc

cover:
	echo calculating code coverage...
	@go test -coverprofile cover.out ./rpc
