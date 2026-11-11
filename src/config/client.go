package config

import (
	"CrzxaExe/infus-server/src/events"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func CreateClient(configs Configuration) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", configs.BrokerUri, configs.BrokerPort))
	opts.SetClientID(configs.ClientName)
	opts.SetDefaultPublishHandler(events.MessageEvent)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Println("MQTT connected successfully")

	return client
}
