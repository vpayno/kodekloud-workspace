#!/bin/bash

for f in /home/bob/images/*jpeg; do
        mv -v "${f}" "${f%.jpeg}".jpg
done
