#!/bin/sh

if [ -e ".token" ]; 
then
  export GITHUB_TOKEN=$(cat .token)
fi

if [ "$1" = "dev" ]; 
then
  while :
  do
    go run $(ls src/*.go | grep -v test ) &
    pid=$!
    sleep 10
    kill $pid
    kill $(ps | grep go-build | awk '{ print $1}' | head -n 1)
  done
else
  go run $(ls src/*.go | grep -v test)
fi

