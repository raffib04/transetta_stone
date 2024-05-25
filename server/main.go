package main

import (
	"fmt"
	"net/http"
)

func servePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/home.html")
}

func main() {
	roomManager := newRoomManager()

	room := newRoom()
	go room.runRoom()

	// http.HandleFunc("/", servePage)
	fs := http.FileServer(http.Dir("./templates/build"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		roomName := r.URL.Query().Get("room")
		if roomName == "" {
			http.Error(w, "Room name is required", http.StatusBadRequest)
			return
		}
		room := roomManager.getRoom(roomName)
		serveWs(room, w, r)
	})

	http.HandleFunc("/rooms", func(w http.ResponseWriter, r *http.Request) {
		rooms := roomManager.listActiveRooms()
		for _, room := range rooms {
			fmt.Print(room)
		}
	})

	http.ListenAndServe(":8080", nil)
}
