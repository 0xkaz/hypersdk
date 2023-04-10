#!/usr/bin/env bash


if ! [[ "$0" =~ scripts/watch.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

while [ true ]
do
    ./scripts/_watch.sh
    sleep 1
done



