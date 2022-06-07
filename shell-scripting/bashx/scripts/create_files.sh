#!/bin/bash

FILE01="file01"
FILE02="file02"
FILE03="file03"

cd /home/bob || exit

echo "Creating file called ${FILE01}"
touch "${FILE01}"

echo "Creating file called ${FILE02}"
touch "${FILE02}"

echo "Creating file called ${FILE03}"
touch "${FILE03}"
