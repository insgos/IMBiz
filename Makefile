default: help

help:
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets: build / image / all / clean'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

all:
	make clean
	make build
	make image

clean:
	@echo 'clean ./bin/*'
	rm -f ./bin/*

build:	dummy
	go build cmd/imbizapi/*.go
	mv ./imbizapi ./bin/	

image:
	docker build -t imbizapi:`date +%Y%m%d` . -f ./build/docker/Dockerfile

dummy: