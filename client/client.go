package client

import (
	"fmt"
	"log"
	"github.com/Anjus-George/simple_pubsub/pubsubpb"

	"google.golang.org/grpc"
)

//Client - Base interface
type Client interface {
	CreateClientRequest()
	RpcProcessClient()
}

//ConnectToServer - common to all clients
func ConnectToServer() (*grpc.ClientConn, pubsubpb.PubSubServiceClient) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	client1 := pubsubpb.NewPubSubServiceClient(conn)
	fmt.Printf("created client, %f\n", client1)
	return conn, client1
}
