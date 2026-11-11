package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func ServeWebsocket(configs Configuration) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Failed to upgrade websocket, ", err)
		}

		defer conn.Close()

		log.Println("New Client connected")

		for {
			msgType, payload, err := conn.ReadMessage()
			if err != nil {
				log.Println("Client disconnected")
				return
			}

			log.Println(msgType, payload)
		}
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configs.WebsocketPort), nil))
}
