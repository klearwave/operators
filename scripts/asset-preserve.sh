#!/usr/bin/env sh

set -e

# DESCRIPTION: preserves assets that were manually modified.  to be used when regenerating operator
#              code in conjection with `make restore` to restore the assets.
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

# source the list of managed assets
if [ ! -f ${CAPABILITY_DIR}/.assets/assets ]; then
    echo "missing assets file at ${CAPABILITY_DIR}/.assets/assets..."
    exit 1
fi
source ${CAPABILITY_DIR}/.assets/assets

# preserve the assets
pushd ${CAPABILITY_DIR}/.assets
for ASSET in ${MANAGED_ASSETS}; do
    DEST="./"
    if [ -f $ASSET ]; then
        echo "preserving file: $ASSET"
        cp -R $ASSET $DEST
    elif [ -d $ASSET ]; then
        echo "preserving directory: $ASSET"
        cp -R $ASSET/* $DEST
    else
        echo "not preserving non-existing asset $ASSET"
    fi
done
popd
