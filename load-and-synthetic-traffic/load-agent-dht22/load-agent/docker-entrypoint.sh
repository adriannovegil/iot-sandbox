#!/bin/ash
set -e

while true; do
  temperature="$((7 + RANDOM % 25))"
  humidity="0.$((70 + $((RANDOM % 20))))$((1 + RANDOM % 10))"
  mosquitto_pub -h $BROKER_URL -p 1883 -t v1/devices/me/test -m "{\"temperature\":${temperature},\"humidity\":${humidity}, \"computed\": {\"heat_index\":22.92} }"
  printf '.\n'
  sleep 10
done
