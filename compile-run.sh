#!/bin/bash

go install avaatechSpe
go install app
export GOOS=linux
go build -o xrf-filetools-linux app
./xrf-filetools-linux