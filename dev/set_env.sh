#!/bin/bash

# Пример установки переменных окружения
set -a
source <(grep -v '^#' ../.env | sed 's/^/export /')
set +a
