package main

import (
	"context"
	"io"
	"log"
	"os"
	pb "playground/grpc-demo/proto"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	//_ = SayHello(c)
	//_ = SayList(c, &pb.HelloRequest{Name: "kiroto"})
	//_ = SayRecord(c, &pb.HelloRequest{Name: "Ahhhhh"})
	_ = SayRoute(c, &pb.HelloRequest{Name: "Yeeees"})
}

func SayHello(c pb.GreeterClient) error {
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	return nil
}

func SayList(c pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := c.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp: %v", resp)
	}
	return nil
}

func SayRecord(c pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := c.SayRecord(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
	}
	resp, _ := stream.CloseAndRecv()
	log.Printf("resp: %v", resp.GetMessage())
	return nil
}

func SayRoute(c pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := c.SayRoute(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp err: %v", resp)
	}
	_ = stream.CloseSend()

	return nil
}
