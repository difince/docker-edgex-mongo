#!/bin/bash

###
# Launches edgex-mongo binary (must be previously built).
#
# Expects that MongoDB are already installed and running.
#
###

# disable secret-store integration
export EDGEX_SECURITY_SECRET_STORE=false

cd cmd
# Add `edgex-` prefix on start, so we can find the process family
exec -a edgex-mongo ./edgex-mongo