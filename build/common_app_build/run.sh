#!/bin/bash

sleep_time=4
echo "sleep for $sleep_time to let the db start"
sleep $sleep_time

mig_conn="postgres://$POSTGRES_USERNAME:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/university?sslmode=disable"
echo "starting migrations..."

migrate -source file://"$GOPATH"/src/University/migrations/"$SERVICE_NAME"/ -database "$mig_conn" up 1

echo "starting app..."
./cmd/"$SERVICE_NAME"/app