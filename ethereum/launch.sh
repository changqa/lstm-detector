#!/bin/bash
#./wipeall.sh
#./bootnode.sh
#./runnode.sh
#for ((i=2;i<300;i++));
#do
#    ./runnode.sh node$i
#    echo node$i >> ./nodes
#done
while :
do
    restartNodeCount=$(((RANDOM%5)+35))
    while(($restartNodeCount>=1))
    do
        node=$(((RANDOM%299)+2))
        ./killnode.sh node$node
        ./runnode.sh node$node
        let "restartNodeCount--"
    done
done
