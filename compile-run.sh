#!/bin/bash

go install configureSpe
go install readAvaatechSpe
go install processAvaatechSpe
go install app
export GOOS=linux
go build -o xrf-filetools-linux app
./xrf-filetools-linux