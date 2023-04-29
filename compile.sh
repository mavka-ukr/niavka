#!/usr/bin/sh

IN=$1
OUT=$2

go run main.go "$IN" >"$OUT".ll
llc -filetype=obj "$OUT".ll -o "$OUT".o
clang "$OUT".o -o "$OUT"

rm "$OUT".ll "$OUT".o
