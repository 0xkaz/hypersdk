_NAME := $(notdir $(CURDIR))
_DATE := $(shell date +%Y-%m-%d-%H-%M-%S)

buildsh:
	./scripts/build.sh
run:
	./scripts/run.sh
watch:
	./scripts/watch.sh

docker: 
	docker build  --platform linux/amd64  -t $(_NAME) . 
