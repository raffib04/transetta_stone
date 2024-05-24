package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	room *Room

	conn *websocket.Conn

	send chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// intakes messages from the client and sends them to the room's broadcast channel
func (c *Client) intakeMessages() {
	defer func() {
		c.room.remove <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		c.room.broadcast <- message
	}

}

// sends messages from the room's broadcast channel to the client
func (c *Client) distributeMessages() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}

			c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}

}

func serveWs(room *Room, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &Client{
		room: room,
		conn: conn,
		send: make(chan []byte),
	}

	client.room.register <- client

	go client.intakeMessages()
	go client.distributeMessages()
}
