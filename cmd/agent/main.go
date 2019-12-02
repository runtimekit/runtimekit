package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/runtimekit/runtimekit/pkg/revdial"
)

func Revdial(ctx context.Context, path string) (*websocket.Conn, *http.Response, error) {
	return websocket.DefaultDialer.Dial("ws://localhost:1313/", nil)
}

func main() {
	fmt.Println("RuntimeKit agent is up and running")

	listener := revdial.NewListener(conn, func(ctx context.Context, path string) (*websocket.Conn, *http.Response, error) {
		return Revdial(ctx, path)
	})
	defer listener.Close()
}
