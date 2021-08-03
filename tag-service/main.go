package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	pb "playground/tag-service/proto"
	"playground/tag-service/server"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

var grpcPort string
var httpPort string

func init() {
	flag.StringVar(&grpcPort, "grpc_port", "8001", "gRPC啟動埠編號")
	flag.StringVar(&httpPort, "http_port", "9001", "HTTP啟動埠編號")
	flag.Parse()
}

func RunHttpServer(port string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("pong"))
		},
	)
	return http.ListenAndServe(":"+port, serveMux)
}
func RunGrpcServer(port string) error {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

func main() {
	errs := make(chan error)
	go func() {
		err := RunHttpServer(httpPort)
		if err != nil {
			errs <- err
		}
	}()
	go func() {
		err := RunGrpcServer(grpcPort)
		if err != nil {
			errs <- err
		}
	}()

	select {
	case err := <-errs:
		log.Fatalf("Run server err: %v", err)
	}
}
