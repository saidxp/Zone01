#!/bin/bash

# Initialize the counter
i=1

# Use a loop to read each line from `ls -l` output
# read line from ls -l and i do -r for read it's as string  becs there is special char '\' 
ls -l | while read -r line; do

  if (( i % 2 == 0 )); then
    echo "$line"
  fi
  # Increment the counter
  ((i++))
done


