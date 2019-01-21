#!/bin/bash -e
DIFF="$(git --no-pager diff --name-only HEAD HEAD~1 | awk 'BEGIN {FS="/"} {print $2}' | sort -u)"
if [ ${#DIFF[@]} == 0 ]; then
	echo "sama aja biji"
else
	echo "beda bijinya:"
	for i in "${DIFF[@]}"
	do
		echo "$i"
	done
fi
