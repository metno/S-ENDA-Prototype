#!/bin/bash
echo "Webhook triggered." | systemd-cat -t webhook-handler

# Work in shared folder
cd /vagrant

mkdir -p isostore

# Check out latest version of metadata
if [ -d S-ENDA-metadata ]; then
  echo "Repository exists locally, running git pull." | systemd-cat -t webhook-handler
  cd S-ENDA-metadata
  git pull
  cd ..
else
  echo "Cloning repository." | systemd-cat -t webhook-handler
  git clone https://github.com/metno/S-ENDA-metadata
fi

rm -rf /isostore/*
docker-compose run -v /vagrant/isostore:/isostore -v /vagrant/S-ENDA-metadata:/mcfdir iso-converter convert-all-mcfs.py --mcfdir /mcfdir --outdir /isostore
docker-compose run -v /vagrant/isostore:/isostore -v /vagrant/S-ENDA-metadata:/mmddir iso-converter mmd2isofix.py -i /mmddir -o /isostore

# Restart catalog-service-api
docker-compose restart catalog-service-api
