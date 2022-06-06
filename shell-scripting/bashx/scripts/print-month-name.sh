#!/bin/bash

declare -i month_num=$1
month_name

if [[ ${month_num} -eq 0 ]]; then
        printf "No month number given\n"

elif [[ ${month_num} -le 12 ]]; then
        month_name="$(date -d "${month_num}/1/1" +%B)"
        printf "%s\n" "${month_name}"
elif [[ ${month_num} -lt 0 ]] || [[ ${month_num} -gt 12 ]]; then
        printf "Invalid month number given\n"
fi
