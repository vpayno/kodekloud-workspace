#!/bin/bash

color="$1"
red="$(tput setaf 1)"
green="$(tput setaf 2)"
reset="$(tput sgr0)"

case "${color}" in
	red)
		echo "${red}this is red${reset}"
		;;
	green)
		echo "${green}this is green${reset}"
		;;
	*)
		echo "${red}red${reset} and ${green}green${reset} are the only choices."
		;;
esac
