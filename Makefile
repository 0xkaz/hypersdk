# _NAME := $(notdir $(CURDIR))
_NAME := weavedbvm
_DATE := $(shell date +%Y-%m-%d-%H-%M-%S)


build: 
	docker build  --platform linux/amd64  -t $(_NAME) . 
