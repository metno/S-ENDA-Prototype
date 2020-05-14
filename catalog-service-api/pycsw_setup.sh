#!/bin/sh

python3 /usr/bin/pycsw-admin.py -c setup_db -f /etc/pycsw/pycsw.cfg
python3 /usr/bin/pycsw-admin.py -c load_records -f /etc/pycsw/pycsw.cfg -p "$ISO_STORE" -r -y

python3 /usr/local/bin/entrypoint.py
