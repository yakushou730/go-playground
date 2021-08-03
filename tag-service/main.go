package main

import (
	"log"
	"net"
	pb "playground/tag-service/proto"
	"playground/tag-service/server"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}
