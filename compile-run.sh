#!/bin/bash

echo "BASH: Started build at: $(date)"
go install configureSpe
go install readAvaatechSpe
go install processAvaatechSpe
go install stdsProcess
go install app
export GOOS=linux
go build -o xrf-filetools-linux app
echo "BASH: Started running program at: $(date)"
./xrf-filetools-linux