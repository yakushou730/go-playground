package main

import (
	"context"
	"fmt"
	"time"

	"nhooyr.io/websocket/wsjson"

	"nhooyr.io/websocket"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, _, err := websocket.Dial(ctx, "ws://localhost:2021/ws", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close(websocket.StatusInternalError, "內部錯誤")

	err = wsjson.Write(ctx, conn, "Hello Web socket Server")
	if err != nil {
		panic(err)
	}

	var v interface{}
	err = wsjson.Read(ctx, conn, &v)
	if err != nil {
		panic(err)
	}

	fmt.Printf("接收到服務端回應: %v\n", v)
	conn.Close(websocket.StatusNormalClosure, "")
}
