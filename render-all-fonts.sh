#!/bin/bash -ex
for i in $(ls fonts | grep -v .go) ; do echo $i ; sed -e "s/latoregular/${i}/g" < cmd/render-fonts/main.go > main-${i}.go ; time go run main-${i}.go -all -out fonts/${i}/${i}.png ; rm main-${i}.go ; done
