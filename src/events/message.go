package events

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MessageEvent mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	log.Printf("%s> %s", m.Topic(), m.Payload())
}
