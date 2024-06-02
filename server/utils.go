package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/translate"
	"os"
	"bufio"
	"strings"
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

func isProfane(message string, language string, profaneWord *string) bool {
	// Check if the message contains any profanity, if it is store in profaneWord
	// return true if profanity is found, false otherwise

	// load the profanity list based on the language
	filename := fmt.Sprintf("data/%s.txt", language)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		profane := scanner.Text()
		if strings.Contains(message, profane) {
			*profaneWord = profane
			return true
		}
	}
	return false
}

// Filter the message based on the language
func filterLanguage(message string, language string) string {
	// filter each word in the message based on the language
	words := strings.Fields(message)
	for i, word := range words {
		// check for profanity, and if there is profanity, replace it with ***
		var profaneWord string
		isProfane := isProfane(word, language, &profaneWord)
		if isProfane {
			fmt.Print("Profane word found: ", profaneWord)
			wordLength := len(profaneWord)
			words[i] = strings.Repeat("*", wordLength)
		}
	}

	return strings.Join(words, " ")
}
