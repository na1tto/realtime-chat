package main

import (
	"fmt"
	"net/http"

	"chat/pkg/websocket"
)

// This is the WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	// upgrading the http conn to a websocket conn
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	// listeing indefinitely for new messages comming
	// through our conn
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
