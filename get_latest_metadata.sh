#!/bin/bash
LIB="/vagrant/lib"
MMD_IN="/vagrant/lib/input_mmd_files"
MMD_REPO_FOLDER_NAME="mmd-example-xml-repo"
MMD_XML_REPO="https://github.com/metno/mmd-example-xml-repo.git"
if [ -f "/vagrant/.env" ]; then
  source /vagrant/.env
fi

# Check out latest version of metadata
if [ -d $MMD_REPO_FOLDER_NAME ]; then
  echo "$MMD_REPO_FOLDER_NAME repository exists locally, running git pull."
  cd $MMD_REPO_FOLDER_NAME
  # Get commit hash
  PREV=$(git rev-parse HEAD)
  git pull || echo "Could not read from git repository - continuing script execution..."
  NEW=$(git rev-parse HEAD)
  FILES=$(git diff --name-only $PREV $NEW)
else
  echo "Cloning repository."
  cd $LIB
  git clone $MMD_XML_REPO $MMD_REPO_FOLDER_NAME || echo "Could not read from git repository - continuing script execution..."
  cd $MMD_REPO_FOLDER_NAME
  FILES=$(git ls-tree -r master --name-only)
fi
# Get list of recent changes and copy the changed files to MMD input folder
echo "Copying files to "$MMD_IN
for FILE in $FILES; do
  if [ -f "${FILE}" ] && [[ "${FILE}" == *.xml ]]; then
    cp $FILE $MMD_IN
  fi
done
