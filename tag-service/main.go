package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	pb "playground/tag-service/proto"
	"playground/tag-service/server"

	"github.com/soheilhy/cmux"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8003", "啟動通訊埠編號")
	flag.Parse()
}

func RunTCPServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func RunHttpServer(port string) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("pong"))
		},
	)
	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
}

func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func main() {
	l, err := RunTCPServer(port)
	if err != nil {
		log.Fatalf("Run TCP Server err: %v", err)
	}
	m := cmux.New(l)
	grpcL := m.MatchWithWriters(
		cmux.HTTP2MatchHeaderFieldPrefixSendSettings(
			"content-type",
			"application/grpc",
		),
	)
	httpL := m.Match(cmux.HTTP1Fast())
	grpcS := RunGrpcServer()
	httpS := RunHttpServer(port)
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	err = m.Serve()
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}
