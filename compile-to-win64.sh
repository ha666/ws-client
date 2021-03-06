
#!/bin/bash

build_dir=`pwd`

echo $(date +"%H:%M:%S")  "start compile"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -mod=vendor -ldflags "-s -w" -o ws-client.exe
echo $(date +"%H:%M:%S")  "finish compile"
