#!/bin/sh
set -e

echo "Pulling The App from git"
git pull

echo "Enabling swag"
cd /usr/local/go/bin/
PATH=$(go env GOPATH)/bin:$PATH
swag -v
cd /var/www/place2connect

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
# npm install
npm run build
cd ..
cd ..
sudo mv place2connect/place2connect-ui/dist www.place2connect.com

echo "Front End Deployed"

echo "Building application For Caddy"
sudo service caddy reload
# docker-compose up -d

# docker-compose up -d --build

# exec "$@"
