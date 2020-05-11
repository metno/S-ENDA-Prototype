# First stage
FROM alpine:3.11 AS builder

ENV MMD_REPO=https://github.com/mortenwh/mmd
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

RUN apk add --no-cache \
  py3-pip; \
  pip3 install "pygeometa==${GEOMETA_VERSION}"; \
  ln -s /usr/bin/python3 /usr/bin/python

COPY --from=builder /mmd/bin/convert_to_mmd /usr/local/bin
COPY --from=builder /mmd/mmd_utils /usr/lib/python3.8/

