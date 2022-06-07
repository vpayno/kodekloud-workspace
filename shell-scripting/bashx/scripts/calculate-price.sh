#!/bin/bash

quantity="$1"
price="$2"

printf "The total price for items is %d dollars\n" "$(( quantity * price ))"
