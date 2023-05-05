PREFIX?=$(shell pwd)

run:

build: 
	@echo "+ $@"
	@echo "Building..."
	go build -o ~/.terraform.d/providers/tiga -v

.PHONY: build