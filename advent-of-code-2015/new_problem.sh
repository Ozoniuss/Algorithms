#!/bin/bash

DAY=$1
URL="https://adventofcode.com/2015/day/$DAY"
DIR=$(python 'get_problem_details.py' "$(curl "$URL")")

mkdir "$DIR"

touch "$DIR/main.go"
touch "$DIR/statement.txt"
touch "$DIR/input.txt"

cp ../aocommon/template.go "$DIR/main.go"
echo $URL > "$DIR/statement.txt"