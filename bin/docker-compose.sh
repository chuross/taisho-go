#!/bin/bash

set -e

docker-compose -f deployment/local/docker-compose.yml $@