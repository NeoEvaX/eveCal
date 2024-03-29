#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
  export $(cat .env | grep -v '^#')
fi

# Get the values of the specific variables you want to include
GOOSE_DRIVER=$GOOSE_DRIVER
GOOSE_DBSTRING=$GOOSE_DBSTRING
GOOSE_MIGRATION_DIR=$GOOSE_MIGRATION_DIR

# Do something with the variables
goose ${@: 1}