#!/bin/sh

export MIGRATION_DIR="./database/migrations"
export DB_DSN="host=localhost port=5432 user=root password=secret dbname=densys sslmode=disable"

if [ "$1" = "--dryrun" ]; then
    goose -v -dir ${MIGRATION_DIR} postgres "${DB_DSN}" status
else
    goose -v -dir ${MIGRATION_DIR} postgres "${DB_DSN}" up
fi
