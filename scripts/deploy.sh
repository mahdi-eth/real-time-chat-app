#!/bin/bash

# Stop and remove existing containers
docker-compose down

# Build and start containers
docker-compose up -d
