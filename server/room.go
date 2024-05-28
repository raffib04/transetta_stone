package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

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
	// Initialize translation session once
	translateSession := translate.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})))

	var msg map[string]string

	err := json.Unmarshal(message, &msg)
	if err != nil {
		fmt.Println("Error unmarshalling message: ", err)
		return
	}

	originalLanguage := msg["language"]
	originalMessage := msg["message"]
	senderUsername := msg["username"]

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
