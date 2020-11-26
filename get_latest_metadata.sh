#!/bin/bash

cd /vagrant/lib

# Check out latest version of metadata
if [ -d s-enda-mmd-xml ]; then
  echo "s-enda-mmd-xml repository exists locally, running git pull."
  cd s-enda-mmd-xml
  # Get commit hash
  PREV=$(git rev-parse HEAD)
  git pull || echo "Could not read from git repository - continuing script execution..."
  NEW=$(git rev-parse HEAD)
  FILES=$(git diff --name-only $PREV $NEW)
else
  echo "Cloning repository."
  git clone git@gitlab.met.no:mmd/s-enda-mmd-xml.git || echo "Could not read from git repository - continuing script execution..."
  #git clone https://gitlab.met.no/mmd/s-enda-mmd-xml.git
  cd s-enda-mmd-xml
  FILES=$(git ls-tree -r master --name-only)
fi
# Get list of recent changes and copy the changed files to MMD input folder
for FILE in $FILES; do
  echo "Copying files..."
  if [ -f "${FILE}" ] && [[ "${FILE}" == *.xml ]]; then
    cp $FILE $MMD_IN
  fi
done
