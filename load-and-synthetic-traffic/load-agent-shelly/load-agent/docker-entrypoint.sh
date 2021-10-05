#!/bin/ash
set -e

while true; do
  temperature="$((7 + RANDOM % 25))"
  mosquitto_pub -h $BROKER_URL -p 1883 -t shellies/living-room/sensor/temperature -m "${temperature}"
  printf '.\n'
  sleep 10
done
