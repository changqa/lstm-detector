#!/bin/bash
for ((i=0;i<33;i++));
do
    randnum=$(((RANDOM%3)+1))
    cat test$randnum >> test_data
done