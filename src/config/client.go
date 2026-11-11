package config

import (
	"CrzxaExe/infus-server/src/events"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func CreateClient() mqtt.Client {
	config := LoadConfig()

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", config.BrokerUri, config.Port))
	opts.SetClientID(config.ClientName)
	opts.SetDefaultPublishHandler(events.MessageEvent)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}
