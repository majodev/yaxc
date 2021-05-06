#!/bin/bash
# $@ > error.log
while [ $? -eq 0 ]; do
    echo "$(date) Trying..."
    $@ &> error.log
done

cat error.log
