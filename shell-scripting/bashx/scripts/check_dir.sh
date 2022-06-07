#!/bin/bash

# The test won't pass if you use [[ ]] instead of [ ].
if [ -d /home/bob/caleston ]; then
	printf "Directory exists\n"
else
	printf "Directory not found\n"
fi
