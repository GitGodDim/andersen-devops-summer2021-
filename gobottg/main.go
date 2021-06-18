package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Create a struct that mimics the webhook response body
	type webhookReqBody struct {
		Message struct {
			Text string `json:"text"`
			Chat struct {
				ID int64 `json:"id"`
			} `json:"chat"`
		} `json:"message"`
	}
	var i int=0
	// This handler is called everytime telegram sends us a webhook event
	func Handler(res http.ResponseWriter, req *http.Request) {
		// First, decode the JSON response body
		body := &webhookReqBody{}
		//test :=json.NewDecoder(req.Body).Decode(body)
		if !strings.Contains(strings.ToLower(body.Message.Text), "/git") {
			return
		}
		if err := json.NewDecoder(req.Body).Decode(body); err != nil {
			fmt.Println("could not decode request body", err)
			return
		} 
		
		// Check if the message contains the github repo

		// if not, return without doing anything
		// If the text contains , call the `sayPolo` function, which
		// is defined b1elow

		if err := sayPolo(body.Message.Chat.ID, ); err != nil {
			
			fmt.Println("error in sending reply:", err)
			return
		}
	
		// log a confirmation message if the message is sent successfully
		fmt.Println("reply sent")
	}
	
	//The below code deals with the process of sending a response message
	// to the user
	
	// Create a struct to conform to the JSON body
	// of the send message request
	// https://core.telegram.org/bots/api#sendmessage
	type sendMessageReqBody struct {
		ChatID int64  `json:"chat_id"`
		Text   string `json:"text"`
	}
	
	// sayPolo takes a chatID and sends  to them
	func sayPolo(chatID int64, ) error {
		reqBody := &sendMessageReqBody{
			ChatID: chatID,
			Text:   "https://github.com/GitGodDim/andersen-devops-summer2021-",
		}
		// Create the JSON body from the struct
		reqBytes, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
	
		// Send a post request with your token
		res, err := http.Post("https://api.telegram.org/bot1856538537:AAHoHK3mn6-YV_4S2mmK3zfqnqDy3_VX0Qo/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusOK {
			return errors.New("unexpected status" + res.Status)
		}
		return nil
	
	}
	
	// FInally, the main funtion starts our server on port 3000
	func main() {
		http.ListenAndServe(":3000", http.HandlerFunc(Handler))
	}
