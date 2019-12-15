#!/bin/bash
VER=$1
if [ "$VER" = "" ]; then
    echo 'please input pack version!'
    exit 1
fi
RELEASE="release-${VER}"
rm -rf release-*
mkdir ${RELEASE}

# windows amd64
echo 'Start pack windows amd64...'
GOOS=windows GOARCH=amd64 ./build.sh
tar -czvf "${RELEASE}/wmqx-${VER}-windows-amd64.tar.gz" -C release .

echo 'Start pack windows X386...'
GOOS=windows GOARCH=386 ./build.sh
tar -czvf "${RELEASE}/wmqx-${VER}-windows-386.tar.gz" -C release .

echo 'Start pack linux amd64'
GOOS=linux GOARCH=amd64 ./build.sh
tar -czvf "${RELEASE}/wmqx-${VER}-linux-amd64.tar.gz" -C release .

echo 'Start pack linux 386'
GOOS=linux GOARCH=386 ./build.sh
tar -czvf "${RELEASE}/wmqx-${VER}-linux-386.tar.gz" -C release .

echo 'Start pack mac amd64'
GOOS=darwin GOARCH=amd64 ./build.sh
tar -czvf "${RELEASE}/wmqx-${VER}-mac-amd64.tar.gz" -C release .

echo 'END'
