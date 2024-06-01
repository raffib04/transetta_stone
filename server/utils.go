package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/translate"
)

// Define a struct for the input message
type Message struct {
	Language string `json:"language"`
	Message  string `json:"message"`
	Username string `json:"username"`
}

// Define a struct for the translated message
type TranslatedMessage struct {
	Language string `json:"language"`
	Message  string `json:"message"`
	Username string `json:"username"`
}

// Translator method to translate messages
func translator(messages []byte, targetLanguage string) []byte {
	// Unmarshal the input JSON
	var msgs []Message

	err := json.Unmarshal(messages, &msgs)

	if err != nil {
		fmt.Println("Error unmarshalling messages: ", err)
		return []byte{}
	}

	// Create a slice to hold translated messages
	var translatedMsgs []TranslatedMessage

	// Loop over the messages and translate each one
	for _, msg := range msgs {
		response, err := translateSession.Text(&translate.TextInput{
			SourceLanguageCode: aws.String(msg.Language),
			TargetLanguageCode: aws.String(targetLanguage),
			Text:               aws.String(msg.Message),
		})
		if err != nil {
			fmt.Println("Error translating message: ", err)
			continue
		}

		// Create a translated message
		translatedMsg := TranslatedMessage{
			Language: targetLanguage,
			Message:  *response.TranslatedText,
			Username: msg.Username,
		}

		// Add the translated message to the slice
		translatedMsgs = append(translatedMsgs, translatedMsg)
	}

	// Marshal the translated messages to JSON
	translatedJSON, err := json.Marshal(translatedMsgs)
	if err != nil {
		fmt.Println("Error marshalling translated messages: ", err)
		return []byte{}
	}

	return translatedJSON
}
