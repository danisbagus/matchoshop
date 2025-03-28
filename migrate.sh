#!/bin/sh

# check if the first argument is empty
if [ -z "$1" ]; then
  echo "❌ Error: Please provide a command: ./migrate.sh up | down | status"
  exit 1
fi

export $(grep -v '^#' .env | xargs)
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="user=$POSTGRES_USERNAME password=$POSTGRES_PASSWORD host=$POSTGRES_HOST dbname=$POSTGRES_DATABASE sslmode=$POSTGRES_SSL_MODE"

# run the goose command with the first argument
echo "🪡 DB migration start! Command: $1"
goose -dir ./app/migration/migrations $1
echo "🎉 DB migration finished!"