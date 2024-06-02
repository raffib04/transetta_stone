package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var translateSession = translate.New(session.Must(session.NewSession(&aws.Config{
	Region: aws.String("us-east-2"),
})))

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	hostname := os.Getenv("MYSQL_HOSTNAME")
	port := os.Getenv("MYSQL_PORT")
	database_name := os.Getenv("MYSQL_DATABASE_NAME")

	connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, database_name)
	db, err = sql.Open("mysql", connection_string)
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
		return
	}

	roomManager := newRoomManager()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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

	http.HandleFunc("/getMessages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		roomName := r.URL.Query().Get("room")
		language := r.URL.Query().Get("language")

		room := roomManager.getRoom(roomName)

		w.Header().Set("Content-Type", "application/json")

		messages := room.messages

		// translate to target lang
		translated_history := translator(messages, language)

		// send to user
		w.Write(translated_history)
	})

	http.ListenAndServe(":8080", nil)
}
