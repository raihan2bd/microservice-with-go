#!/bin/bash
# ./migrate.sh up
# ./migrate.sh down
# ./migrate.sh force 1

MIGRATE_PATH="./migrations"
DATABASE_URL="postgres://postgres:password@localhost:5432/product_service?sslmode=disable"

migrate -path $MIGRATE_PATH -database $DATABASE_URL "$@"