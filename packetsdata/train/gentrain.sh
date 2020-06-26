#!/bin/bash
for ((i=0;i<33;i++));
do
    randnum=$(((RANDOM%3)+1))
    cat train$randnum >> train_data
done
