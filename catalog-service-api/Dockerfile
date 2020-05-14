FROM metno/pycsw:2.4.2

ADD pycsw.cfg /etc/pycsw/pycsw.cfg
ADD pycsw_setup.sh /usr/local/bin/pycsw_setup.sh

ENTRYPOINT ["/usr/local/bin/pycsw_setup.sh"]
