#!/bin/sh

docker-compose down -v

docker-compose down --rmi all

docker builder prune -f

docker volume prune -f

docker network prune -f

docker image prune -f

echo "Docker-compose cleanup completed."
