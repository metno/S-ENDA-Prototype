# First stage
FROM alpine:3.11 AS builder

ENV MMD_REPO=https://github.com/metno/mmd
ENV MMD_VERSION=${MMD_VERSION:-master}


RUN apk add --no-cache \
  git

WORKDIR /

RUN git clone ${MMD_REPO}; \
  cd mmd; \
  git checkout ${MMD_VERSION}; \
  chmod +x /mmd/bin/convert_to_mmd

# Second stage
FROM alpine:3.11
LABEL maintainer="aheimsbakk@met.no"

ENV GEOMETA_VERSION=0.5.0

RUN apk add \
  libxml2 \
  libxslt \
  python3
RUN apk add --no-cache --virtual .build-deps \
  build-base \
  libxml2-dev \
  libxslt-dev \
  py3-pip \
  python3-dev; \
  pip3 install \
  pygeometa==${GEOMETA_VERSION} \
  lxml==4.5.0 \
  xmltodict==0.12.0 \
  parmap==1.5.2; \
  apk del .build-deps

RUN ln -s /usr/bin/python3 /usr/bin/python; \
    mkdir -p /usr/local/share/mmd-to-iso

COPY --from=builder /mmd/bin/convert_to_mmd /usr/local/bin
COPY --from=builder /mmd/mmd_utils/ /usr/lib/python3.8/site-packages/mmd_utils/
COPY convert-all-mcfs.py /usr/local/bin
COPY mmd2isofix.py /usr/local/bin
COPY mmd-to-iso.xsl /usr/local/share/mmd-to-iso

