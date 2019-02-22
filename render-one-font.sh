#!/bin/bash -ex

echo $@
sed -e "s/latoregular/$@/g" < cmd/render-fonts/main.go > main-$@.go
time go run main-$@.go -all -out fonts/$@/$@.png
# rm main-$@.go
