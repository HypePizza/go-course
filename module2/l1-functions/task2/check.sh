#!/usr/bin/env bash
# Прогон всех случаев из условия задачи
go build -o bin/stats . || exit 1   # собираем один раз, дальше запуск мгновенный

cases=(
    "3 4.5 10"
    "-2 7"
    "-7"
    ""
    "1 abc 3"
)

for c in "${cases[@]}"; do
  echo "\$ stats $c"
  ./bin/stats $c 2>&1 | perl -pe 's/\n/\$\n/'
  echo "exit=${PIPESTATUS[0]}"
done