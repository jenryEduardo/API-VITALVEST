package mosquitto

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Mosquitto() {
	// Definir opciones del cliente MQTT
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://localhost:1883").
		SetClientID("go-subscriber")

	// Callback cuando llega un mensaje
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("Conectado al broker MQTT")
		if token := c.Subscribe("sensores/temperatura", 0, messageHandler); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}
	}

	// Crear cliente
	client := mqtt.NewClient(opts)

	// Conectar al broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Mantener el programa corriendo
	fmt.Println("Esperando mensajes. Presiona Ctrl+C para salir")
	for {
		time.Sleep(1 * time.Second)
	}
}

// Callback cuando llega un mensaje
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Mensaje recibido en el tópico %s: %s\n", msg.Topic(), msg.Payload())
}
	