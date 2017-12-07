#!/bin/bash

go install timeUsage
go install app
export GOOS=linux
go build -o xrf-filetools-linux app
./xrf-filetools-linux