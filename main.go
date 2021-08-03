package main

import (
	"flag"
	_ "playground/grpc-demo/proto"
)

var (
	port string
)

func main() {
	flag.StringVar(&port, "p", "8000", "啟動通訊埠號")
	flag.Parse()
}
