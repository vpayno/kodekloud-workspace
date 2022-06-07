#!/bin/bash

declare -A operators=(
        [add]="+"
        [subtract]="-"
        [multiply]="*"
        [divide]="/"
)

declare -A choices=(
        [1]="add"
        [2]="subtract"
        [3]="multiply"
        [4]="divide"
        [5]="quit"
)

declare -i num1
declare -i num2
declare input
declare result

menu()
{
        local choice

        printf "Available commands:\n"
        for choice in "${!choices[@]}"; do
                printf "\t%d. %s\n" "${choice}" "${choices[${choice}]}"
        done
        printf "\n"
        read -r -p "command: " input
        input="${input:-5}"

        selection="${choices[${input}]}"
}

menu

while [ "${selection}" != "quit" ]; do
        printf "\n"
        read -r -p "number1: " num1
        read -r -p "number2: " num2
        printf "\n"

        result="$(bc -l <<< "scale=4; ${num1} ${operators[${selection}]} ${num2}")"

        printf "Answer=%s\n" "${result}"
        printf "\n"

        menu
done
