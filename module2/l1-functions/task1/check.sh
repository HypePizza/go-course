#!/usr/bin/env bash
# Прогон всех случаев из условия задачи
go build -o bin/hsize . || exit 1   # собираем один раз, дальше запуск мгновенный

cases=(
  "512"
  "1536"
  "1.5 GB"
  "2.7 B"
  "-5"
  "abc GB"
  "abc"
  "10 xb"
  ""          # без аргументов
  "1 2 3"     # лишние аргументы
)

for c in "${cases[@]}"; do
  echo "\$ hsize $c"
  ./bin/hsize $c 2>&1 | perl -pe 's/\n/\$\n/'
  echo "exit=${PIPESTATUS[0]}"
done