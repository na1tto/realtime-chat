package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


// We need to define a upgrader to transform
// our HTTP connection in a WebSocket based connection
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	
	// We need to check the origin of our connections
	// this will allow us to make requests from our React 
	// development server to here.
	// for now let's just allow any connection
	CheckOrigin: func(r *http.Request) bool {return true},
}

// This is a reader endpoint that will only
// listen for new messages being sent to our
// WebSocket endpoint
func reader(conn *websocket.Conn) {
	for {
		// read a message
		messageType, p, err := conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return
		}
		// printout that message
		fmt.Println(string(p))
		
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		
	}
}

// This is the WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Host)
	
	// upgrading the http conn to a websocket conn
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	
	// listeing indefinitely for new messages comming
	// through our conn
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	
	// mapped our websocket endpoint 
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
