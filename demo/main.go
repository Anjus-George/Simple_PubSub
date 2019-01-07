package main

import (
	"fmt"
	"vsc_workspace/simple_pubsub/client"
	"vsc_workspace/simple_pubsub/consumer"
	"vsc_workspace/simple_pubsub/producer"
)

func main() {
	fmt.Println("hello I am a client")
	conn, client1 := client.ConnectToServer()
	defer conn.Close()

	p1 := producer.Producer{}
	p1.CreateClientRequest(client1, "A", "0")
	p1.CreateClientRequest(client1, "A", "1")
	p1.CreateClientRequest(client1, "A", "2")
	p1.CreateClientRequest(client1, "A", "3")
	p1.CreateClientRequest(client1, "B", "0")
	p1.CreateClientRequest(client1, "B", "1")
	p1.CreateClientRequest(client1, "B", "2")

	c1 := consumer.Consumer{}
	c1.CreateClientRequest(client1, "A", 1, 2)
	c1.CreateClientRequest(client1, "B", 0, 1)
	c1.CreateClientRequest(client1, "C", 0, 1)

}
