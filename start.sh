#!/bin/bash

# 第一引数に問題の名前を指定
if [ -z "$1" ]; then
  echo "Usage: ./create_problem.sh problem_name"
  exit 1
fi

PROBLEM_NAME=$1

# ディレクトリ作成
mkdir -p ./leetcode/$PROBLEM_NAME

# main.goのテンプレート作成
cat <<EOT >> ./leetcode/$PROBLEM_NAME/main.go
package main

import "fmt"

func main() {
    // Write your code here
    fmt.Println("$PROBLEM_NAME solution")
}
EOT


echo "$PROBLEM_NAME problem environment created."
