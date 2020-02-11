cd sample_data

xsltproc -o /home/pycsw/sentinel-mmd-to-iso.xml /data/mmd-to-iso.xsl /data/mmd_xml/S2A_MSIL1C_20190209T103201_N0207_R108_T32VPM_20190209T110450.xml

python3 /usr/local/bin/pycsw-admin.py -c setup_db -f /etc/pycsw/pycsw.cfg
python3 /usr/local/bin/pycsw-admin.py -c load_records -f /etc/pycsw/pycsw.cfg -p /home/pycsw/sentinel-mmd-to-iso.xml

python3 /usr/local/bin/entrypoint.py
# with transaction == True, havesting of wms can be done with:
# python3 /usr/local/bin/pycsw-admin.py -c post_xml -u http://pycsw:8000/pycsw/csw.py -x sample_data/xml/post.xml