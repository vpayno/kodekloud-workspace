#!/bin/bash

new_shell="$2"
user_name="$1"

usermod -s "${new_shell}" "${user_name}"
