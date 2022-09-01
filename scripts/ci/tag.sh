#!/bin/bash

echo "Checking variables"
echo "GITHUB_ACTIONS: $GITHUB_ACTIONS"
echo "GITHUB_REF: $GITHUB_REF"

# Checks if the tag script is being run on Github Actions
if [ ! "$GITHUB_ACTIONS" ]; then
    echo "This script needs to be run on Github Actions platform"
    exit 1
fi

if [ ! "$GITHUB_REF" == "refs/heads/master" ]; then
    echo "This build is not in the master branch"
    exit 1
fi

# Checks if the VERSION file exists from the makefile
if [ ! -f BUILD_VERSION.txt ]; then
    echo "BUILD_VERSION.txt file not found!"
    exit 2
fi

# Checks if the Github OAuth Token was set.
if [ -z "$GITHUB_TOKEN" ]; then
    echo "The GITHUB_TOKEN has not been set"
    exit 2
fi

GIT_TAG=$(cat BUILD_VERSION.txt)
readonly GIT_TAG

echo "Tagging release version $GIT_TAG"
echo "$GITHUB_REF"

git config --global user.email "noreply@github.com"
git config --global user.name "Github Actions"

git tag "$GIT_TAG" -a -m "Generated tag from Github Actions build $GITHUB_RUN_NUMBER"
git push --tag > /dev/null 2>&1
echo "Pushed tag to repo"
exit 0
