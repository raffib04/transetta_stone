package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/translate"
)

type Room struct {
	clients   map[*Client]bool
	broadcast chan []byte
	register  chan *Client
	remove    chan *Client
	id        int
	messages  []byte
}

func newRoom(id int) *Room {
	return &Room{
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte),
		register:  make(chan *Client),
		remove:    make(chan *Client),
		id:        id,
		messages:  []byte{},
	}
}

func (room *Room) translateMessage(message []byte) {
	// Initialize translation session once

	var msg map[string]string

	err := json.Unmarshal(message, &msg)
	if err != nil {
		fmt.Println("Error unmarshalling message: ", err)
		return
	}

	originalLanguage := msg["language"]
	originalMessage := msg["message"]
	senderUsername := msg["username"]

	// store message to database
	_, err = db.Exec("INSERT INTO messages (room_id, text, original_language, creator_name, created_at) VALUES (?, ?, ?, ?, ?);", room.id, originalMessage, originalLanguage, senderUsername, time.Now())

	if err != nil {
		fmt.Println("Error storing into DB")
	}

	for client := range room.clients {
		receiverLanguage := client.language

		// Translate the message
		response, err := translateSession.Text(&translate.TextInput{
			SourceLanguageCode: aws.String(originalLanguage),
			TargetLanguageCode: aws.String(receiverLanguage),
			Text:               aws.String(originalMessage),
		})
		if err != nil {
			fmt.Println("Error translating message: ", err)
			continue
		}

		translatedMessage := *response.TranslatedText

		//send the translated message to the client as json
		jsonMsg := make(map[string]string)
		jsonMsg["username"] = senderUsername
		jsonMsg["message"] = translatedMessage

		data, err := json.Marshal(jsonMsg)
		if err != nil {
			break
		}
		select {
		case client.send <- data:
		default:
			// Properly handle client disconnection
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
			go room.translateMessage(message)
		}
	}
}

func (room *Room) getMessages() {
	roomid := room.id

	// query for messages
	rows, err := db.Query("SELECT * FROM messages WHERE room_id = ? ORDER BY created_at ASC;", roomid)
	if err != nil {
		fmt.Println("Error querying messages: ", err)
		return
	}
	defer rows.Close()

	var messages []map[string]string

	for rows.Next() {
		var id int
		var room_id int
		var text string
		var original_language string
		var creator_name string
		var created_at []uint8

		err = rows.Scan(&id, &room_id, &text, &original_language, &creator_name, &created_at)
		if err != nil {
			fmt.Println("Error scanning messages: ", err)
			return
		}

		msg := map[string]string{
			"username": creator_name,
			"language": original_language,
			"message":  text,
		}

		messages = append(messages, msg)
	}

	data, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("Error marshalling messages: ", err)
		return
	}

	room.messages = data
}

