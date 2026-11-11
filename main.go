package main

import (
	"fmt"
	"log"

	"CrzxaExe/infus-server/src/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	fmt.Println("Starting Server")

	client := config.CreateClient()

	token := client.Subscribe("medical/infusion/#", 1, func(c mqtt.Client, m mqtt.Message) {
		log.Printf("Infusion | %s, %s", m.Topic(), m.Payload())
	})

	go func() {
		_ = token.Wait()
		if token.Error() != nil {
			log.Fatal(token.Error())
		}
	}()

	select {}
}
