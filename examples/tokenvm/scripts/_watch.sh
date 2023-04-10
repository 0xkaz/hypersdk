#!/usr/bin/env bash
#!/bin/sh


set -e
if ! [[ "$0" =~ scripts/_watch.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi


LAST=`find . | egrep ".sh|.go" | grep -v 'build/'| xargs ls -l -D '%s' | awk '{print $6}' | sort -n | tail -n 1 `
echo "LAST=$LAST"

LAST2=""

while [ true ]
do
  LAST2=`find . | egrep ".sh|.go" | grep -v 'build/'| xargs ls -l -D '%s' | awk '{print $6}' | sort -n | tail -n 1 `
  
  # echo "LAST2=$LAST2"

  if ! [ "${LAST}" -eq "${LAST2}" ]; then

  # then
    echo "LAST=$LAST"
    echo "LAST2=$LAST2"
    rm ./build/* || true
    LAST=$LAST2
    ./scripts/build.sh
    find ./build
    ./build/token-cli version
  # else
  #   echo "..."
  fi
  # sleep 1

done


echo "end3"