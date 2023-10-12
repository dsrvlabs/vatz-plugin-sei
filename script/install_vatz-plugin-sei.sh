#!/bin/bash
set -e
set -v

#Create vatz log folder
mkdir /var/log/vatz

# Clone VATZ repository
cd /root
git clone git@github.com:dsrvlabs/vatz.git

# Compile VATZ
cd /root/vatz
make

## You will see binary named vatz

# Initialize VATZ
./vatz init

# Install vatz-plugin-sei
./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/node_block_sync node_block_sync
./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/node_is_alived node_is_alived
./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/node_peer_count node_peer_count
./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/node_active_status node_active_status
./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/node_governance_alarm node_governance_alarm
./vatz plugin install github.com/dsrvlabs/vatz-plugin-sei/plugins/pfd_status pfd_status

# Check plugin list
./vatz plugin list
