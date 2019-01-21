#!/bin/bash -e
echo git --no-pager diff --name-only HEAD HEAD~1 | sort -u | awk 'BEGIN {FS="/"} {print $2}' | uniq
