#!/bin/bash
set -e
set -v

VATZ_PATH="/root/vatz"
VATZ_PLUGIN_PATH="/root/.vatz"
API_PORT="1317"
VALOPER_ADDRESS="voloper_address"
VOTER_ADDRESS="voter_address"
SEI_HOME="$HOME/.sei"
WARN_CNT="20"
CRITICAL_CNT="95"
NODE_BLOCK_SYNC_PORT="10001"
NODE_IS_ALIVED_PORT="10002"
NODE_PEER_COUNT="10003"
NODE_ACTIVE_STATUS_PORT="10004"
NODE_GOVERNANCE_ALARM_PORT="10005"
PFD_STATUS_PORT="10006"

cd $VATZ_PATH
./vatz start --config default.yaml >> /var/log/vatz/vatz.log 2>&1 &
cd $VATZ_PLUGIN_PATH

./node_block_sync -port $NODE_BLOCK_SYNC_PORT >> /var/log/vatz/node_block_sync.log 2>&1 &
./node_is_alived -port $NODE_IS_ALIVED_PORT >> /var/log/vatz/node_is_alived.log 2>&1 &
./node_peer_count -port $NODE_PEER_COUNT >> /var/log/vatz/node_peer_count.log 2>&1 &
./node_active_status -port $NODE_ACTIVE_STATUS_PORT -valoperAddr $VALOPER_ADDRESS >> /var/log/vatz/node_active_status.log 2>&1 &
./node_governance_alarm -port $NODE_GOVERNANCE_ALARM_PORT -apiPort $API_PORT -voterAddr $VOTER_ADDRESS >> /var/log/vatz/node_governance_alarm.log 2>&1 &

./pfd_status -port $PFD_STATUS_PORT -warnCondition $WARN_CNT -criticalCondition $CRITICAL_CNT -valoperAddr $VALOPER_ADDRESS -seiHome $SEI_HOME  >> /var/log/vatz/pfd_status.log 2>&1 &

echo "true"
