#!/usr/bin/env sh

set -e

# DESCRIPTION: restores assets that were manually modified.  to be used when regenerating operator
#              code in conjection with `make preseve` to preserve the assets.
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

# restore the assets
pushd ${CAPABILITY_DIR}/.assets
cp -R ./* ..
popd
