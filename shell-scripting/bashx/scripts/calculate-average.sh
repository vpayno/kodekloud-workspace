#!/bin/bash

num1="$1"
num2="$2"
num3="$3"

printf "%s\n" "$(bc <<< "scale=4; (${num1} + ${num2} + ${num3}) / 3.0")"
