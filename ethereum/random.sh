#!/bin/bash
while :
do
    SLEEP_TIME=$(($RANDOM%11+1))
    echo ${SLEEP_TIME}
    ./killnode.sh node3
    sleep ${SLEEP_TIME}
    ./runnode.sh node3
    sleep ${SLEEP_TIME}
done
