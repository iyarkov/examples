#!/bin/bash

VERSION="1.0"

TARGET=generated

rm ${TARGET}/build.go

if [ -z "$BUILD_NUMBER" ]; then
  BUILD_NUMBER="Unknown"
fi
BUILD_DATE="$(date)"
#GIT_REVISION="Unknown"
GIT_REVISION="$(git rev-parse HEAD)"


echo 'package generated' > ${TARGET}/build.go
echo ''  >> ${TARGET}/build.go
echo "var Version=\"${VERSION}\""  >> ${TARGET}/build.go
echo "var BuildNumber=\"${BUILD_NUMBER}\""  >> ${TARGET}/build.go
echo "var BuildDate=\"${BUILD_DATE}\""  >> ${TARGET}/build.go
echo "var GitRevision=\"${GIT_REVISION}\""  >> ${TARGET}/build.go

