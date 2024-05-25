package main

import (
	"net/http"
)

func servePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/home.html")
}

func main() {
	room := newRoom()
	go room.runRoom()

	http.HandleFunc("/", servePage)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(room, w, r)
	})

	http.ListenAndServe("10.105.210.1:8080", nil)
}
