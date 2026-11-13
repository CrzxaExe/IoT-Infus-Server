package config

import (
	"CrzxaExe/infus-server/src/models"
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateClient(configs Configuration, hub *Hub, database *mongo.Client) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", configs.BrokerUri, configs.BrokerPort))
	opts.SetClientID(configs.ClientName)

	opts.SetDefaultPublishHandler(func(c mqtt.Client, m mqtt.Message) {
		var data models.InfusPayload
		if err := json.Unmarshal(m.Payload(), &data); err != nil {
			log.Println("Failed to parse json from MQTT:", err)
			return
		}
		data.Timestamp = time.Now().String()

		msg := models.WebsocketPayload[models.InfusPayload]{
			Type:    "data",
			Message: "test",
			Result:  data,
		}
		_, err := InsertOne(database, "data", data)
		if err != nil {
			log.Println("Failed to insert data:", err)
			return
		}

		log.Printf("%s> %s", m.Topic(), m.Payload())
		broadcastBytes, _ := json.Marshal(msg)
		hub.Broadcast <- broadcastBytes
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Println("MQTT connected successfully")

	return client
}
