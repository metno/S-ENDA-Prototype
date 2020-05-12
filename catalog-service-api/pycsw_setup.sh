#!/bin/sh

# xsltproc -o /home/pycsw/sentinel-mmd-to-iso.xml /data/mmd-to-iso.xsl /data/mmd_xml/S2A_MSIL1C_20190209T103201_N0207_R108_T32VPM_20190209T110450.xml

# for testring I usually remove the sqlite db and the output ISO
#rm -rf /home/pycsw/tests.db
#rm -rf $OUTPUT_ISO

if [  "$INDEXDB" = true ]; then
  mkdir -p "$OUTPUT_ISO"
  python3 /usr/local/bin/mmd2isofix.py -i "$METADATA" -o "$OUTPUT_ISO"
  python3 /usr/bin/pycsw-admin.py -c setup_db -f /etc/pycsw/pycsw.cfg
  python3 /usr/bin/pycsw-admin.py -c load_records -f /etc/pycsw/pycsw.cfg -p "$OUTPUT_ISO" -r -y
fi

python3 /usr/local/bin/entrypoint.py
