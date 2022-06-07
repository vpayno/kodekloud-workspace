#!/bin/bash

declare input_file="/tmp/assets/mission-names.txt"

declare -a missions
declare mission

mapfile -t missions < "${input_file}"

for mission in "${missions[@]}"; do
        bash create-and-launch-rocket "${mission}"
        printf "\n"
done
