package main

import (
	"fmt"
	"github.com/Anjus-George/client"
	"github.com/Anjus-George/consumer"
	"github.com/Anjus-George/producer"
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
