package main

type Room struct {
	clients   map[*Client]bool
	broadcast chan []byte
	register  chan *Client
	remove    chan *Client
}

func newRoom() *Room {
	return &Room{
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte),
		register:  make(chan *Client),
		remove:    make(chan *Client),
	}
}

func (room *Room) translateMessage(message []byte) {
	for client := range room.clients {
		select {
		case client.send <- message:
		default:
			delete(room.clients, client)
			close(client.send)
		}
	}
}

func (room *Room) runRoom() {
	for {
		select {
		case client := <-room.register:
			room.clients[client] = true
		case client := <-room.remove:
			if _, ok := room.clients[client]; ok {
				delete(room.clients, client)
				close(client.send)
			}
		case message := <-room.broadcast:
			room.translateMessage(message)
		}
	}
}
