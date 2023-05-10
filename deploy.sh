#!/bin/sh
set -e

echo "Pulling The App from git"
# git pull

echo "Copying app.env file to place2connect-api"
cp app.env ./place2connect-api

echo "Copying .env file to place2connect-ui"
cp .env ./place2connect-ui

echo "Deploying The Entire app"

echo "Start Backend Deployment"

cd place2connect-api
make start-all
cd ..

echo "Backend Deployed"

echo "Start FrontEnd Deployment"

cd place2connect-ui
npm run dev
cd ..

echo "Front End Deployed"

echo "Building application For Caddy"
docker-compose up -d

# docker-compose up -d --build

# exec "$@"
