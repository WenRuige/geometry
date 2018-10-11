#!/usr/bin/env bash

#自动提交master

msg=$1
if [ -n "$msg" ]; then
   git add .
   git commit -m"${msg}"
   git pull
   git status
   git push origin master
   echo "完成add、commit、pull，别忘了push"
else
    echo "请添加注释再来一遍"
fi