#!/usr/bin/env bash

# run it inside the ffi rust project
cd ../..
cargo build --release -p ffi
cd rust/ffi || exit 1
mkdir -p build
cc main.c -o build/main
./build/main
