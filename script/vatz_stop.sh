#!/bin/bash
set -e
set -v

pid0=`pidof node_active_status`
pid1=`pidof node_block_sync`
pid2=`pidof node_governance_alarm`
pid3=`pidof node_is_alived`
pid4=`pidof node_peer_count`
pid5=`pidof pfd_status`
pid6=`pidof vatz`

# Kill process
kill -15 $pid0
kill -15 $pid1
kill -15 $pid2
kill -15 $pid3
kill -15 $pid4
kill -15 $pid5
kill -15 $pid6

# wait Kill process
sleep 1

while (( `lsof -p $pid0 | wc -l` > 0 ))
do
        echo `lsof -p $pid0 | wc -l`
        sleep 1
done
while (( `lsof -p $pid1 | wc -l` > 0 ))
do
        echo `lsof -p $pid1 | wc -l`
        sleep 1
done
while (( `lsof -p $pid2 | wc -l` > 0 ))
do
        echo `lsof -p $pid2 | wc -l`
        sleep 1
done
while (( `lsof -p $pid3 | wc -l` > 0 ))
do
        echo `lsof -p $pid3 | wc -l`
        sleep 1
done
while (( `lsof -p $pid4 | wc -l` > 0 ))
do
        echo `lsof -p $pid4 | wc -l`
        sleep 1
done
while (( `lsof -p $pid5 | wc -l` > 0 ))
do
        echo `lsof -p $pid5 | wc -l`
        sleep 1
done
while (( `lsof -p $pid6 | wc -l` > 0 ))
do
        echo `lsof -p $pid6 | wc -l`
        sleep 1
done

echo "seid $pid0 is killed."
echo "seid $pid1 is killed."
echo "seid $pid2 is killed."
echo "seid $pid3 is killed."
echo "seid $pid4 is killed."
echo "seid $pid5 is killed."
echo "seid $pid6 is killed."
