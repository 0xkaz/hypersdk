_NAME := $(notdir $(CURDIR))
_DATE := $(shell date +%Y-%m-%d-%H-%M-%S)

buildsh:
	./scripts/build.sh
run: 
	./scripts/run.sh
kill-run: kill
	./scripts/run.sh
watch:
	./scripts/watch.sh

kill:
	ps auxww | grep avalanche | grep -v grep | awk '{print $$2}'|xargs  kill -9 







