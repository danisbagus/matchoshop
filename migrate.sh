#!/bin/sh
echo $DB_HOST
echo $DB_USER
echo $DB_PASSWORD
echo $DB_HOST
echo $DB_PORT
echo $DB_NAME
echo $DB_SSL_MODE

goose -env production up 