#!/bin/bash

# 获取当前目录
current_dir=$(pwd)

# 定义执行命令的函数
execute_command() {
  local dir=$1
  echo "进入目录: $dir"
  cd "$dir" || exit 1
  if [ -f "go.mod" ]; then
    go mod tidy
    go install
  fi
  echo "退出目录: $dir"
  cd "$current_dir" || exit 1
}

# 递归遍历子目录并执行命令
traverse_directories() {
  local dir=$1
  for item in "$dir"/*; do
    if [ -d "$item" ]; then
      execute_command "$item"
      traverse_directories "$item"
    fi
  done
}

# 在当前目录执行命令
execute_command "$current_dir"

# 遍历子目录执行命令
traverse_directories "$current_dir"
