package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// We need to define a upgrader to transform
// our HTTP connection in a WebSocket based connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We need to check the origin of our connections
	// this will allow us to make requests from our React
	// development server to here.
	// for now let's just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}