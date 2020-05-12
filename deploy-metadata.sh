#!/bin/bash
echo "Webhook triggered." | systemd-cat -t webhook-handler

# Work in shared folder
cd /vagrant

mkdir -p isostore

# Check out latest version of metadata
if [ -d S-ENDA-metadata ]; then
  echo "Repository exists locally, running git pull." | systemd-cat -t webhook-handler
  git pull
else
  echo "Cloning repository." | systemd-cat -t webhook-handler
  git clone https://github.com/metno/S-ENDA-metadata
fi

docker-compose run -v /vagrant/isostore:/isostore -v /vagrant/S-ENDA-metadata:/mcfdir iso-converter convert-all-mcfs.py --mcfdir /mcfdir --outdir /isostore

# Restart catalog-service-api
docker-compose restart catalog-service-api
