#!/bin/sh

# check if the first argument is empty
if [ -z "$1" ]; then
  echo "❌ Error: Please provide a command: ./migrate.sh up | down | status | create <name>"
  exit 1
fi

export $(grep -v '^#' .env | xargs)
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="user=$POSTGRES_USERNAME password=$POSTGRES_PASSWORD host=$POSTGRES_HOST dbname=$POSTGRES_DATABASE sslmode=$POSTGRES_SSL_MODE"

# run the goose command with the first argument
# echo "🪡 DB migration start! Command: $1"
# goose -dir ./app/migration/migrations $1
# echo "🎉 DB migration finished!"
# Menjalankan perintah sesuai argumen pertama
case "$1" in
  up|down|status)
    echo "🪡 Running DB migration: $1"
    goose -dir ./app/migration/migrations $1
    echo "🎉 Migration completed!"
    ;;
  create)
    if [ -z "$2" ]; then
      echo "❌ Error: Please provide a migration name: ./migrate.sh create <name>"
      exit 1
    fi
    echo "📝 Creating new migration: $2"
    goose -dir ./app/migration/migrations create "$2" sql
    echo "✅ Migration '$2' created!"
    ;;
  *)
    echo "❌ Error: Invalid command. Use: ./migrate.sh up | down | status | create <name>"
    exit 1
    ;;
esac