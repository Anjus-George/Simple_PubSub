package consumer

import (
	"context"
	"fmt"
	"log"
	"github.com/Anjus-George/simple_pubsub/pubsubpb"
)

//Consumer - is a type of Client
type Consumer struct {
}

//RpcProcessClient - consumer's own implementation
func (c Consumer) RpcProcessClient(client1 pubsubpb.PubSubServiceClient, req *pubsubpb.ConsumeRequest) {
	fmt.Println("Starting rpc Consume Test...")

	res, err := client1.Consume(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling consume rpc: %v", err)
	}

	log.Printf("Response after Consume call:\n Response message: %v, Item List: %v, Last Consumed Index: %v\n", res.ResponseMessage, res.ItemList, res.LastConsumedIndex)
	//log.Printf("Response after Consume call: %v\n", res)
}

//CreateClientRequest - producer's own implementation
func (c Consumer) CreateClientRequest(client1 pubsubpb.PubSubServiceClient, topic string, startIndex int64, fetchSize int64) {

	req := pubsubpb.ConsumeRequest{
		Topic:      topic,
		StartIndex: startIndex,
		FetchSize:  fetchSize,
	}
	c.RpcProcessClient(client1, &req)
}
