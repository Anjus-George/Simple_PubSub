package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"github.com/Anjus-George/pubsubpb"

	"google.golang.org/grpc"
)

var records = make(map[string][]string)

type server struct {
}

func (*server) Produce(ctx context.Context, req *pubsubpb.ProduceRequest) (*pubsubpb.ProduceResponse, error) {
	fmt.Printf("Produce rpc invoked with %v", req)
	//get topic
	topic := req.GetTopic()

	_, pres := records[topic]

	var res int64

	if pres != true { //if topic not in list create and add item in that
		records[topic] = []string{req.GetItem()}
		res = int64(len(records[topic]))
	} else { //if topic present append item
		records[topic] = append(records[topic], req.GetItem())
		res = int64(len(records[topic]))
	}

	topicLen := &pubsubpb.ProduceResponse{
		TopicLength: res,
	}

	fmt.Println(records)
	return topicLen, nil
}

func (*server) Consume(ctx context.Context, req *pubsubpb.ConsumeRequest) (*pubsubpb.ConsumeResponse, error) {
	fmt.Printf("Consume rpc invoked with %v\n", req)
	topic := req.GetTopic()
	startIndex := req.GetStartIndex()
	fetchSize := req.GetFetchSize()

	var resMessage string
	var s []string
	var lastIndex int64

	_, pres := records[topic]

	if pres != true {
		resMessage = "Topic does not exist!"
	} else {
		resMessage = "Topic exist!"
		s = records[topic][startIndex : startIndex+fetchSize]
		lastIndex = startIndex + fetchSize - 1

	}

	consResp := &pubsubpb.ConsumeResponse{
		ResponseMessage:   resMessage,
		ItemList:          s,
		LastConsumedIndex: lastIndex,
	}

	return consResp, nil
}

func main() {
	fmt.Println("hello I am the server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)

	}

	s := grpc.NewServer()
	pubsubpb.RegisterPubSubServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
