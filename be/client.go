package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

type Client struct {
	conn *websocket.Conn
}

func newClient(c *websocket.Conn) *Client {
	return &Client{conn: c}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
	client := newClient(conn)

	fmt.Println("New Client joined the hub !", client)
}
func main() {

}
