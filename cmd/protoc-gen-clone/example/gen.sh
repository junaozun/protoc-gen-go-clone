#!/bin/bash

CUR_DIR="$( cd "$( dirname "$0"  )" && pwd  )"

set -xe

go install ../../protoc-gen-clone

deps=(common common1)
for i in ${deps[@]}
do
  path=$CUR_DIR/$i
  protoc \
  -I=$CUR_DIR/common \
  -I=$path \
  --go_out=paths=source_relative:$path \
  $path/*.proto
done

protoc \
-I=$CUR_DIR/common \
-I=$CUR_DIR/common1 \
-I=$CUR_DIR \
--go_out=paths=source_relative:$CUR_DIR \
--clone_out=paths=source_relative:$CUR_DIR \
$CUR_DIR/*.proto
