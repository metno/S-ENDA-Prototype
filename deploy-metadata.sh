#!/bin/bash
if [ -f "/vagrant/.env" ]; then
  source /vagrant/.env
fi

echo "Webhook triggered." | systemd-cat -t webhook-handler

# This copies the new or changes MMD to $MMD_IN
# We may also have to add functionality to remove deleted MMD files
./get_latest_metadata.sh

# Remove old iso files
rm $ISOSTORE/*

# Work in shared folder
cd /vagrant

# Translate from MMD to ISO19139
docker-compose run --rm \
    -e XSLTPATH=/usr/local/share/xslt \
    -v $ISOSTORE:/isostore \
    -v $MMD_IN:/mmddir \
    iso-converter \
    xmlconverter.py -i /mmddir -o /isostore -t /usr/local/share/xslt/mmd-to-iso.xsl

# We may prefer to have a separate container for indexing in pycsw..
# Ingest metadata from ISO19139 xml files
docker-compose exec -T catalog-service-api bash -c 'python3 /usr/bin/pycsw-admin.py -c load_records -f /etc/pycsw/pycsw.cfg -p $ISO_STORE -r -y'

# Clean up
rm $ISOSTORE/*
rm $MMD_IN/*
