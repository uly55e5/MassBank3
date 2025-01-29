#!/bin/bash

echo "REACT_APP_MB3_API_URL=$1" > .env
echo "REACT_APP_MB3_FRONTEND_URL=$2" >> .env
echo "REACT_APP_MB3_BASE_URL=$3" >> .env
echo "REACT_APP_MB3_VERSION=$4" >> .env
echo "REACT_APP_EXPORT_SERVICE_URL=$5" >> .env
echo "REACT_APP_GOOGLE_SEARCH_CONSOLE_KEY=$6" >> .env

npm run start