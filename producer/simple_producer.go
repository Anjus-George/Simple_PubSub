package producer

import (
	"context"
	"fmt"
	"log"
	"github.com/Anjus-George/pubsubpb"
)

//Producer - is a type of Client
type Producer struct {
}

//RpcProcessClient - producer's own implementation
func (p *Producer) RpcProcessClient(client1 pubsubpb.PubSubServiceClient, req *pubsubpb.ProduceRequest) {
	fmt.Println("Starting rpc Produce Test...")

	res, err := client1.Produce(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling produce rpc: %v", err)
	}

	log.Printf("Response after Produce call: %v", res.TopicLength)

}

//CreateClientRequest - producer's own implementation
func (p *Producer) CreateClientRequest(client1 pubsubpb.PubSubServiceClient, topic string, item string) {
	req := pubsubpb.ProduceRequest{
		Item:  item,
		Topic: topic,
	}

	p.RpcProcessClient(client1, &req)
}
