#!/bin/bash

A="$1"
B="$2"

printf "Sum is %d\n" "$(( A + B ))"

printf "Difference is %d\n" "$(( A - B ))"

printf "Product is %d\n" "$(( A * B ))"

printf "Quotient is %d\n" "$(( A / B ))"
