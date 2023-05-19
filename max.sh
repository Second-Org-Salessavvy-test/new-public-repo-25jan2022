#!/bin/bash

max=$(echo "$1 $2 $3" | tr ' ' '\n' | sort -n | tail -n 1)
echo $max

