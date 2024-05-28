package main

import (
	"fmt"
	"net/http"
)

func main() {
	roomManager := newRoomManager()

	room := newRoom()
	go room.runRoom()

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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		rooms := roomManager.listActiveRooms()
		w.Header().Set("Content-Type", "application/json")
		w.Write(rooms)

		fmt.Println("Active rooms:", string(rooms))
	})

	http.ListenAndServe(":8080", nil)
}
