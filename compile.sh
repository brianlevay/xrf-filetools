#!/bin/bash

OPTION=$1

if [ $1 -eq 1 ]
then
    export GOOS=linux
    go build -o xrf-filetools-linux app
elif [ $1 -eq 2 ]
then
    export GOOS=windows
    go build -o xrf-filetools-windows.exe app
else
    echo "No option selected"
fi