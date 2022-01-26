#!/bin/sh
echo $DATABASE_URL
echo $DB_SSL_MODE

goose -env production up 