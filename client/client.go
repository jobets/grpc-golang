package main

import (
	"context"
	"fmt"
	api "github.com/grpc-golang/api/generated"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {

	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %v", err)
	}
	defer cc.Close()

	c := api.NewHelloServiceClient(cc)
	checkUnary(c)
	//checkServerStreaming(c)
	//checkClientStreaming(c)
	//checkBidirectionalStreaming(c)
}

func checkUnary(c api.HelloServiceClient) {
	fmt.Println("Starting a Unary RPC...")
	req := &api.HelloRequest{
		Hello: &api.Hello{FirstName: "Jobet",
			LastName: " Samuel",
		},
	}
	res, err := c.SayHelloUnary(context.Background(), req)

	if err != nil {
		log.Fatal("error while calling Hello RPC: %v", err)
	}

	log.Printf("Respone from Unary Server: %v", res.Result)
}

func checkServerStreaming(c api.HelloServiceClient) {
	fmt.Println("Starting a Server Streaming RPC...")
	req := &api.HelloRequest{
		Hello: &api.Hello{FirstName: "Jobet",
			LastName: " Samuel",
		},
	}
	res, err := c.SayHelloServerStreaming(context.Background(), req)

	if err != nil {
		log.Fatal("error while calling Hello RPC: %v", err)
	}

	for {
		msg, err := res.Recv()

		if err == io.EOF {
			// Reached end of stream
			break
		}
		if err != nil {
			log.Fatal("Error while reading froms stream: %v", err)
		}
		log.Printf(" Response from Server Streaming RPC: %v", msg)
	}
}

func checkClientStreaming(c api.HelloServiceClient) {
	fmt.Println("Starting a Client Streaming RPC...")

	requests := []*api.HelloRequestMultipleTimes{
		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "Jobet",
			},
		},
		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "Mark",
			},
		},

		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "John",
			},
		},

		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "Matthew",
			},
		},
	}

	stream, err := c.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatal("Error while calling  HelloClient: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error while receiving response from Server: %v", err)
	}
	fmt.Printf("Response received from Client Streaming Server: %v\n", res)

}

func checkBidirectionalStreaming(c api.HelloServiceClient) {
	fmt.Println("Starting a Bidirectional Streaming RPC...")

	stream, err := c.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatal("Error while creating stream: %v", err)
	}

	requests := []*api.HelloRequestMultipleTimes{
		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "Jobet",
			},
		},
		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "Mark",
			},
		},

		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "John",
			},
		},

		&api.HelloRequestMultipleTimes{
			Hello: &api.Hello{
				FirstName: "Matthew",
			},
		},
	}

	waitc := make(chan struct{})

	go func() {
		//Send the above bunch of requests created
		for _, req := range requests {
			fmt.Printf("Sending requests: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Response Received from Bidirectional Server: %v\n", res.GetResult())
		}

		stream.CloseSend()
	}()
	<-waitc //Block until both are done.
}
