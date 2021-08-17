#!/bin/bash


case $1 in
"cmd")
  go build -o bin/$2 cmd/$2/main.go
  sudo ./bin/$2
;;
"push")
  git add .
  git commit -m "$2"
  git pull
  git push
;;
esac