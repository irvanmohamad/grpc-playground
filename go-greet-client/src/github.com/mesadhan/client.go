package main

import (
	"context"
	"fmt"
	greetpb "github.com/mesadhan/go-greet-client/services/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)


func main()  {
	fmt.Println("Hello, From Client!")

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	// fmt.Println(c)



	doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	//doStreamBoth(c)
}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a Unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sadhan",
			LastName: "Sarker",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)

}



func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sadhan",
			LastName: "Sarker",
		},
	}

	resStream, err := c.GreetManyTimesMethod(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}


func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sadhan",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ripon",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Karim",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Rahim",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req.Greeting.GetFirstName())

		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v", res)
}


func doStreamBoth(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a Bi-directional Streaming RPC...")
	//we create a stream by invoking the client

	stream, err := c.GreetEveryoneMethod(context.Background())
	if(err != nil) {
		log.Fatalf("Error while creating stream %v", err)
		return
	}

	waitc := make(chan struct{})

	//we send a bunch of messages to the server (go routine)
	go func ()  {
		requests := []*greetpb.GreetEveryoneRequest{
			{
				Greeting: &greetpb.Greeting{
					FirstName: "Sadhan",
				},
			},
			{
				Greeting: &greetpb.Greeting{
					FirstName: "Ripon",
				},
			},
			{
				Greeting: &greetpb.Greeting{
					FirstName: "Karim",
				},
			},
			{
				Greeting: &greetpb.Greeting{
					FirstName: "Rahim",
				},
			},
			{
				Greeting: &greetpb.Greeting{
					FirstName: "Hannan",
				},
			},
		}

		for _, req := range requests {
			fmt.Printf("Sending req: %v\n", req.Greeting.GetFirstName())
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	//then we print the response as we receive it (go routine)
	go func ()  {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if(err != nil) {
				log.Fatalf("Error while receiving stream: %v", err)
				break
			}
			fmt.Printf("Response from GreetEveryone: %v\n", res.GetResult())
		}
		close(waitc)
	}()
	//block until the entire thing is done

	<-waitc

}