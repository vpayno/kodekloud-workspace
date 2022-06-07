#!/bin/bash

num1=$1
num2=$2

# The test is checking to see that [ ] were used instead of [[ ]] and claiming that the output isn't correct.
if [ "${num1}" -gt "${num2}" ]; then
	printf "%d\n" "${num1}"
else
	printf "%d\n" "${num2}"
fi
