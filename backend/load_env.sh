#!/bin/bash
# load config.yaml into environment variables
if [ -z "$ENV" ]; then
  echo "ENV variable is not set. Please set it to 'debug' or 'release'."
  exit 1
fi

export DB_USER=$(yq eval ".${ENV}.db.user" config.yaml)
export DB_PASSWORD=$(yq eval ".${ENV}.db.password" config.yaml)
export DB_NAME=$(yq eval ".${ENV}.db.name" config.yaml)
export DB_HOST=$(yq eval ".${ENV}.db.host" config.yaml)
export DB_PORT=$(yq eval ".${ENV}.db.port" config.yaml)

# Optionally print the variables to verify
echo "DB_USER=$DB_USER"
echo "DB_PASSWORD=$DB_PASSWORD"
echo "DB_NAME=$DB_NAME"
echo "DB_HOST=$DB_HOST"
echo "DB_PORT=$DB_PORT"
