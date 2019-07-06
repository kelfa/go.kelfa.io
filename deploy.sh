#!/bin/bash

GOOS=linux go build main.go
zip function.zip main
aws lambda update-function-code --function-name go-kelfa-io --zip-file fileb://function.zip
rm main
rm function.zip
