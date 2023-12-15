#!/usr/bin/env bash

# This script checks that the TAG arg does not exist, already.

TAG="$1"

if [ $(git tag -l $TAG) ]; then
	echo "Error; tag $TAG already exists"
	exit 1
else
	echo "tag $TAG does not exist"
fi
