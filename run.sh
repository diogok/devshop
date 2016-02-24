#!/bin/sh

if [ "$1" = "dev" ]; 
then
  while :
  do
    go run $(ls src/*.go) &
    pid=$!
    sleep 10
    kill $pid
    kill $(ps | grep go-build | awk '{ print $1}' | head -n 1)
  done
else
  go run $(ls src/*.go)
fi
