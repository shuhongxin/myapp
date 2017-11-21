
package main

import (
	"log"
	"fmt"
	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "myapp/echo/echo"
)

const (
	address     = "localhost:50051"

)


func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)
	stream, err := c.Push(context.Background(),&pb.Autoreply{})
	
    go func()  {
		for {
			in, err := stream.Recv()
            if err == io.EOF {
                // read done.
                fmt.Println("read done ")
                return
            }
            if err != nil {
                fmt.Println("Failed to receive a time : %v", err)
			}
			fmt.Println(in.Message)
		}
	}()
	// Contact the server and print out its response.
	for {
		 var str string
		 fmt.Scan(&str)
		 r, err := c.Send(context.Background(), &pb.EchoRequest{Message: str})
		 if err != nil {
			log.Fatalf("could not greet: %v", err)
		 }
		 fmt.Printf(r.Message+"\n")
	}
	
}