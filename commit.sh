#!/usr/bin/env bash

#自动提交master

msg=$1
if [ -n "$msg" ]; then
   git add .
   git commit -m"${msg}"
   git pull
   git status
   git push origin master
   echo "====== 提交完成 ====="
else
   echo "====== 遗漏注释 ====="
fi