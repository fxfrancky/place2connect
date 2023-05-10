#!/bin/sh
set -e
echo "Stoping The Entire app"
cd place2connect-api
docker-compose down
cd ..

echo "App Successfully stopped"

# exec "$@"
