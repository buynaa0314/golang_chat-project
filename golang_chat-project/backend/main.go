package main

import (
	"fmt"
	"net/http"
)

func serveWS(pool *webSocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("webSocket endpoint reached")

	conn, err := webSocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
		return
	}

	client := &webSocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := webSocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	fmt.Println("Buynaa's full stack chat project")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
