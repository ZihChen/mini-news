#!/bin/bash

WORK_PATH=$(dirname $(readlink -f $0))

# 專案名稱(取當前資料夾路徑最後一個資料夾名稱)
PROJECT_NAME=${WORK_PATH##*/}
# 環境變數
ENV="local"

echo "ENV=$ENV">.env
echo "PROJECT_NAME=$PROJECT_NAME">>.env

# 啟動容器服務
USER_PATH=$USER_PATH docker-compose up -d