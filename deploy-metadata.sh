#!/bin/bash

# Work in shared folder
cd /vagrant

# Check out latest version of metadata
git clone https://github.com/metno/S-ENDA-metadata

# Restart catalog-service-api
docker-compose restart catalog-service-api
