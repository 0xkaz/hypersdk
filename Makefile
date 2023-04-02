_NAME := weavedbvm
_DATE := $(shell date +%Y-%m-%d-%H-%M-%S)

build: 
	docker build  --platform linux/amd64  -t $(_NAME) . 

run: 
	docker run  -it $(_NAME) . /bin/sh 

