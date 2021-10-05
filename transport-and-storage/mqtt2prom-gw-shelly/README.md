
# MQTT2Prometheus for Shelly Sensor (Metric Per Topic)

A easy way to setup a MQTT2Prometheus server using Docker Compose.

## Fast run

Build and launch the complete infrastructure:

```
make up
```

## Publish a message

Required is a MQTT client. I use this: https://hivemq.github.io/mqtt-cli/docs/installation/packages.html

To publish a message run:

```
mqtt pub --host localhost -p 1883 -t shellies/living-room/sensor/temperature '15'
```

## Contributing

All contributions are welcome. Please check the [CONTRIBUTING.md](./CONTRIBUTING.md) file to know how contribute to the project.

## License

See the [LICENSE](./LICENSE) file
