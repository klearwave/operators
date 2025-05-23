#!/usr/bin/env sh

set -e

# DESCRIPTION: performs the download of a set of manifests, using vendir, from an upstream
# INPUTS:
#   CAPABILITY: the capability, stored in the parent directory in (e.g. certificates, identity)
#   COMPONENT:  the project, within the CAPABILITY, to download
#
# USAGE:
#   CAPABILITY=secrets COMPONENT=external-secrets scripts/download.sh

# retreive the common functions
source "$(dirname "$0")/common.sh"

# ensure inputs are provided
set_capability_component_vars

# ensure vendir is installed
if [ -z `which vendir` ]; then
    echo "vendir not installed...please run 'make vendir'..."
    exit 1
fi

# ensure helm is installed
if [ -z `which helm` ]; then
    echo "helm not installed...please install manually..."
    exit 1
fi

VENDIR_CONFIG=${COMPONENT_DIR}/config/vendor.yaml
VENDIR_CONFIG_LOCK=${COMPONENT_DIR}/config/vendor.yaml.lock

# ensure config file exists
if [ ! -f ${VENDIR_CONFIG} ]; then
    echo "vendir configuration file not found at ${VENDIR_CONFIG}"
    echo "please create the vendir configuration file first..."
    exit 1
fi

# sync the vendir configuration
vendir sync \
    --file $VENDIR_CONFIG \
    --lock-file $VENDIR_CONFIG_LOCK

# if we produced a vendor-helm directory in the project, we must template it
if [ -d ${COMPONENT_DIR}/vendor-helm ]; then
    # ensure the vendir directory exists
    mkdir -p ${COMPONENT_DIR}/vendor

    # ensure the vendir directory is clean
    rm -rf ${COMPONENT_DIR}/vendor/*

    helm template ${COMPONENT} ${COMPONENT_DIR}/vendor-helm --include-crds | tee ${COMPONENT_DIR}/vendor/${COMPONENT}.yaml
fi
