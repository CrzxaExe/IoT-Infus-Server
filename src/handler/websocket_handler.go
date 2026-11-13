package handler

import (
	"CrzxaExe/infus-server/src/config"
	"CrzxaExe/infus-server/src/constants"
	"CrzxaExe/infus-server/src/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWS(hub *config.Hub, mqttClient mqtt.Client, database *mongo.Client, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error while upgrading http:", err)
		return
	}

	defer func() {
		hub.RemoveClient(conn)
		conn.Close()
	}()

	hub.AddClient(conn)

	allData, err := config.FetchData[models.InfusPayload](database, "data", bson.D{}, nil)
	if err != nil {
		log.Println("Failed to fetch data:", err)
		return
	}

	data, err := json.Marshal(allData)
	if err != nil {
		log.Println("Failed to parse websocket:", err)
		return
	}
	conn.WriteMessage(websocket.TextMessage, data)

	for {
		_, payload, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var msg models.SocketCommand[any]
		if err := json.Unmarshal(payload, &msg); err != nil {
			log.Println("Failed to parse Webscoket:", err)
			continue
		}

		// Add default value
		msg.Timestamp = time.Now().String()
		msg.IsSuccess = false

		if msg.Type == "set" {
			topic := fmt.Sprintf("%s/command/%s", constants.TOPIC, msg.DeviceID)

			result, err := config.InsertOne(database, "command", msg)
			if err != nil || result.InsertedID == nil {
				log.Fatal("Failed to save into database:", err)
				continue
			}

			mqttClient.Publish(topic, 1, false, payload)
			log.Printf("WS > %s\n", payload)
		}
	}
}
