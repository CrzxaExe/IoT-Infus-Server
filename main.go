package main

import (
	"log"

	"CrzxaExe/infus-server/src/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	log.Println("Starting Server")
	configs := config.LoadConfig()

	client := config.CreateClient(configs)

	token := client.Subscribe("medical/infusion/data/#", 1, func(c mqtt.Client, m mqtt.Message) {
		log.Printf("Infusion | %s, %s", m.Topic(), m.Payload())
	})

	go func() {
		for {
			if !client.IsConnected() {
				return
			}

			_ = token.Wait()
			if token.Error() != nil {
				log.Fatal(token.Error())
			}
		}
	}()

	config.ServeWebsocket(configs)

	select {}
}
