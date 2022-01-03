package main

import (
	"context"
	"fmt"
	"github.com/mesadhan/go-greet-server/services/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

// unary
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := "Hello " + firstName + " : " + lastName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

// GreetManyTimesMethod server streaming
func (*server) GreetManyTimesMethod(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesMethodServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		// result := "Hello " + firstName + " number " + string(i)
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		fmt.Println("Server sending: ", result)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request\n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func (*server) GreetEveryoneMethod(stream greetpb.GreetService_GreetEveryoneMethodServer) error {
	fmt.Printf("GreetEveryoneMethod function was invoked with a streaming request\n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "
		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}
}

func main() {

	fmt.Println("GRPC Server is running: 0.0.0.0:8080")

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func allowCors(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
	resp.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
}