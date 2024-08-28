#!/bin/bash

# 运行构建脚本
sh ./script/build.sh

# 测试 ./xzdp 是否能正常运行，运行2秒后退出
timeout 2s ./xzdp &
sleep 2
kill $!

# 单元测试未编写，暂时只测试能否正确运行
# go test -v -cover ./...