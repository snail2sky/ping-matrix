#!/bin/bash

target=dist
declare -a GO_ARCH=(386 amd64 arm arm64)
declare -a GO_OS=(windows darwin linux)

build(){
  for arch in ${GO_ARCH[@]}; do
      for os in ${GO_OS[@]}; do
        export CGO_ENABLED=0 GOOS=$os GOARCH=$arch
        echo "building $os $arch program."
        if [ "$os" == "windows" ]; then
          cd rpc_client
          go build  -o ../$target/ping-matrix-client-${os}-${arch}.exe
          cd ../rpc_server
          go build  -o ../$target/ping-matrix-server-${os}-${arch}.exe
          cd ..
        else
          cd rpc_client
          go build  -o ../$target/ping-matrix-client-${os}-${arch}
          cd ../rpc_server
          go build  -o ../$target/ping-matrix-server-${os}-${arch}
          cd ..
        fi
      done
  done

}

rm -rf $target/*
build
