#!/bin/bash

source .env

DB_NAME=$DB_NAME

sleep 5

# Создание базы данных
create_database() {
    psql -c "CREATE DATABASE ${DB_NAME}"
    >&2 echo "База данных ${DB_NAME} создана"
}

create_database