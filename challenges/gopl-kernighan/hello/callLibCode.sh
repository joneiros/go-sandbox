#!/bin/zsh

go mod init hello
go get gopl.io

echo Code:
# potential future error: path mismatch from version change
echo | cat $HOME/go/pkg/mod/gopl.io@v0.0.0-20211004154805-1ae3ec64947b/ch1/helloworld/main.go

echo ------------------------
echo Produces:
go run gopl.io/ch1/helloworld