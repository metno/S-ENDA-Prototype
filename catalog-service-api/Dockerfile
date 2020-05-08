FROM metno/pycsw:2.4.2

ADD pycsw.cfg /etc/pycsw/pycsw.cfg
ADD pycsw_setup.sh /usr/local/bin/pycsw_setup.sh
ADD mmd2isofix.py /usr/local/bin/mmd2isofix.py
ADD mmd-to-iso.xsl /usr/local/share/mmd-to-iso.xsl

ENTRYPOINT ["/usr/local/bin/pycsw_setup.sh"]
