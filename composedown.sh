#!/usr/bin/env bash
docker-compose down && \
docker system prune
echo "Docker-compose down!"