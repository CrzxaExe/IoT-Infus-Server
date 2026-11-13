package main

import (
	"fmt"
	"log"
	"net/http"

	"CrzxaExe/infus-server/src/config"
	"CrzxaExe/infus-server/src/constants"
	"CrzxaExe/infus-server/src/handler"
)

func main() {
	log.Println("Starting Server")
	configs := config.LoadConfig()

	hub := config.CreateHub()
	database := config.ConnectDatabase(configs)
	go hub.Run()

	mqttClient := config.CreateClient(configs, hub, database)
	mqttClient.Subscribe(fmt.Sprintf("%s/data/#", constants.TOPIC), 1, nil)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWS(hub, mqttClient, database, w, r)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configs.WebsocketPort), nil))
}
