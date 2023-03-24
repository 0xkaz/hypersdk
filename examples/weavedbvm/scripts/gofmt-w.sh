#!/bin/sh
test -z "$(gofmt -s -l -w  -d $(find . -type f -name '*.go' -not -path "*/vendor/*") | tee /dev/stderr)"
