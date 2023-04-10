#!/bin/sh
test -z "$(go vet -all ./...  | tee /dev/stderr)"
