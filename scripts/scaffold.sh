#!/usr/bin/env sh

set -e

# DESCRIPTION: performs the scaffolding of a directory structure for a component
# INPUTS:
#   CAPABILITY: the capability, stored in the parent directory in (e.g. certificates, identity)
#   COMPONENT:  the project, within the CAPABILITY, to download
#
# USAGE:
#   CAPABILITY=secrets COMPONENT=external-secrets scripts/scaffold.sh

# retreive the common functions
source "$(dirname "$0")/common.sh"

# ensure inputs are provided
#   NOTE: also sets vars needed below
set_capability_component_vars

# create the capability directory structure
echo "creating capability directory structure at ${CAPABILITY_DIR}..."
mkdir -p \
    ${CAPABILITY_DIR}/.assets

# create the component directory structure
echo "creating component directory structure at ${COMPONENT_DIR}..."
mkdir -p \
    ${COMPONENT_DIR}/config/overlays \
    ${COMPONENT_DIR}/vendor \
    ${COMPONENT_DIR}/manifests

# create required configuration files
echo "creating required configuration files at ${COMPONENT_DIR}..."
touch \
    ${COMPONENT_DIR}/config/values.yaml \
    ${COMPONENT_DIR}/config/vendor.yaml \
    ${COMPONENT_DIR}/workload-${COMPONENT}.yaml

# create the specific configuration files
if [ ! -f ${CAPABILITY_DIR_BASE}/values.yaml ]; then
    echo "creating common values file at ${CAPABILITY_DIR_BASE}/values.yaml..."
    cat <<EOF > ${CAPABILITY_DIR_BASE}/values.yaml
---
operatorVersion: unstable
projectVersion: unstable
EOF
fi

if [ ! -f ${CAPABILITY_DIR_SOURCE}/values.yaml ]; then
    echo "creating capabilities values file at ${CAPABILITY_DIR_SOURCE}/values.yaml..."
    cat <<EOF > ${CAPABILITY_DIR_SOURCE}/values.yaml
---
capability: ${CAPABILITY}
EOF
fi

if [ ! -f ${CAPABILITY_DIR}/.assets/assets ]; then
    echo "creating assets file at ${CAPABILITY_DIR}/.assets/assets..."
    cat <<EOF > ${CAPABILITY_DIR}/.assets/assets
export MANAGED_ASSETS="
	../go.mod
	../README.md
	../.gitignore
	../.dockerignore
	../.github/workflows/release.yml
	../.goreleaser.yml
	../Dockerfile
	../LICENSE
	../config/default/kustomization.yaml
	../config/default/manager_auth_proxy_patch.yaml
"
EOF
fi
