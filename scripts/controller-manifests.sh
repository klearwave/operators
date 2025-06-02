#!/usr/bin/env sh

set -e

# DESCRIPTION: generates the controller manifests for a specific capability
#
# INPUTS:
#   CAPABILITY: the capability, stored in the parent directory in (e.g. certificates, identity)
#
# USAGE:
#   CAPABILITY=secrets scripts/preserve-assets.sh

# retreive the common functions
source "$(dirname "$0")/common.sh"

# ensure inputs are provided
set_capability_vars
: ${VERSION?missing environment variable VERSION}

# controller manifests
# NOTE: we need to preserve the original kustomization file as we always want our generated manifest to 
#       be set to the latest release except when we tag a new version.
CONTROLLER_IMG=ghcr.io/klearwave/${CAPABILITY}-operator:latest

pushd ${CAPABILITY_DIR} > /dev/null
make kustomize
mkdir -p config/deploy
cd config/manager
cp -f kustomization.yaml kustomization.yaml.bak
export KUSTOMIZE_CMD=$(ls ../../bin/kustomize*)
${KUSTOMIZE_CMD} edit set image controller=${CONTROLLER_IMG}
${KUSTOMIZE_CMD} build ../../config/default > ../../config/deploy/manifests.yaml
mv kustomization.yaml.bak kustomization.yaml
popd > /dev/null
