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

	var translateSession *translate.Translate = translate.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})))

	var msg map[string]string
	json.Unmarshal(message, &msg)
	original_language := msg["language"]

	for client := range room.clients {

		receiver_language := client.language
		fmt.Println("Receiver Language: ", receiver_language)

		response, err := translateSession.Text(&translate.TextInput{
			SourceLanguageCode: aws.String(original_language),
			TargetLanguageCode: aws.String(receiver_language),
			Text:               aws.String(msg["message"]),
		})

		if err != nil {
			fmt.Println("Error translating message: ", err)
		}

		message = []byte(*response.TranslatedText)

		fmt.Println("Message: ", string(message))

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
