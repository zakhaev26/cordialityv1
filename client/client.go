package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

const (
	webSocketURL = "ws://127.0.0.1:8080/channel?channelName=lol&password=deezBalls" // Replace with your actual URL
)

func main() {
	// Upgrade connection to WebSocket

	fmt.Println("Your Name?")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	senderName := sc.Text()
	conn, _, err := websocket.DefaultDialer.Dial(webSocketURL, http.Header{})
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// Channel for messages received from server
	messageChannel := make(chan []byte)

	// Go routine to read messages from the server
	go func() {
		defer close(messageChannel)
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("read error:", err)
				return
			}
			messageChannel <- msg
		}
	}()

	// Read user input for messages
	fmt.Println("Connected to WebSocket server!")
	fmt.Println("Enter a message (blank line to exit):")
	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for msg := range messageChannel {
			fmt.Println(string(msg))
		}
	}()

	for scanner.Scan() {
		message := scanner.Text()
		if message == "" {
			break
		}

		// Send message to server
		err := conn.WriteMessage(websocket.TextMessage, []byte("> "+senderName+" "+message))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

	fmt.Println("Exiting...")
}
