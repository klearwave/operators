#!/usr/bin/env sh

set -e

# DESCRIPTION: initialize the operator with operator-builder
#
# INPUTS:
#   CAPABILITY: the capability, stored in the parent directory in (e.g. certificates, identity)
#   COMPONENT:  the project, within the CAPABILITY, to download
#   VERSION:    the version of the operator to build
#
# USAGE:
#   VERSION=unstable CAPABILITY=secrets COMPONENT=external-secrets scripts/operator-init.sh

# retreive the common functions
source "$(dirname "$0")/common.sh"

# ensure inputs are provided
set_capability_component_vars
: ${VERSION?missing environment variable VERSION}

# ensure operator-builder is installed
if [ -z `which operator-builder` ]; then
    echo "operator-builder not installed...please install manually..."
    exit 1
fi

# ensure workload config exists
WORKLOAD_CONFIG=${COMPONENT_DIR}/config/workload.yaml
if [ ! -f ${WORKLOAD_CONFIG} ]; then
    echo "workload configuration file not found at ${WORKLOAD_CONFIG}"
    echo "please create the workload configuration file first..."
    exit 1
fi

# initialize the operator
pushd ${CAPABILITY_DIR}
operator-builder init \
    --workload-config .source/${COMPONENT}/config/workload.yaml \
    --repo ${GIT_REPO_URL}/capabilities/${CAPABILITY} \
    --controller-image ${GIT_REPO_URL}/${CAPABILITY}-operator:${VERSION} \
    --skip-go-version-check
popd
