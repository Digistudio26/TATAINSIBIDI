#!/bin/bash

# Get the number of loops from the first argument
n=$1

# Limit to 100
if [ "$n" -gt 100 ]; then
    n=100
fi

# Loop from 1 to n
for ((i=1; i<=n; i++)); do
    echo "This is loop number $i"
done
