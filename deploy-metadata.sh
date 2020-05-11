#!/bin/bash
echo "Webhook triggered." | systemd-cat -t webhook-handler

# Work in shared folder
cd /vagrant

# Check out latest version of metadata
if [ -d S-ENDA-metadata ]; then
  echo "Repository exists locally, running git pull." | systemd-cat -t webhook-handler
  git pull
fi
  echo "Cloning repository." | systemd-cat -t webhook-handler
  git clone https://github.com/metno/S-ENDA-metadata
fi


# Restart catalog-service-api
docker-compose restart catalog-service-api
