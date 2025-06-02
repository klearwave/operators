#!/usr/bin/env sh

set -e

# DESCRIPTION: clean the operator
#
# INPUTS:
#   CAPABILITY: the capability, stored in the parent directory in (e.g. certificates, identity)
#   COMPONENT:  the project, within the CAPABILITY, to download
#   VERSION:    the version of the operator to build
#
# USAGE:
#   VERSION=unstable CAPABILITY=secrets COMPONENT=external-secrets scripts/operator-create.sh

# retreive the common functions
source "$(dirname "$0")/common.sh"

# ensure inputs are provided
set_capability_vars
: ${VERSION?missing environment variable VERSION}

# ensure operator-builder is installed
if [ -z `which operator-builder` ]; then
    echo "operator-builder not installed...please install manually..."
    exit 1
fi

# for safety, double check that we have the CAPABILITY_DIR environment variable set
: ${CAPABILITY_DIR?unablet to find CAPABILITY_DIR}

# clean the capability operator
if [ ! -d ${CAPABILITY_DIR} ]; then
    echo "missing capability directory at ${CAPABILITY_DIR}..."
    exit 1
fi

rm -rf ${CAPABILITY_DIR}/bin/
rm -rf ${CAPABILITY_DIR}/*
