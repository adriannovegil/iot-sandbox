package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// StartClient ...
func StartClient(broker string) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID("data-nerd")
	c := mqtt.NewClient(opts)
	return c
}
