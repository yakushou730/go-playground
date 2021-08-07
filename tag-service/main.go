package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	pb "playground/tag-service/proto"
	"playground/tag-service/server"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"golang.org/x/net/http2"

	"golang.org/x/net/http2/h2c"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8004", "啟動通訊埠編號")
	flag.Parse()
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func RunServer(port string) error {
	httpMux := RunHttpServer()
	grpcS := RunGrpcServer()
	gatewayMux := RunGrpcGatewayServer()

	httpMux.Handle("/", gatewayMux)

	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

func RunHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("pong"))
		},
	)
	return serveMux
}

func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func RunGrpcGatewayServer() *runtime.ServeMux {
	endpoint := "0.0.0.0:" + port
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)

	return gwmux
}

func main() {
	err := RunServer(port)
	if err != nil {
		log.Fatalf("Run Server err: %v", err)
	}
}
