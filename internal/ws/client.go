package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	egress     chan []byte
}

type ClientList map[*Client]bool

func NewClient(c *websocket.Conn, m *Manager) *Client {
	return &Client{
		connection: c,
		manager:    m,
		egress:     make(chan []byte),
	}
}

func (c *Client) readMessages() {

	defer func() {
		c.manager.removeClient(c)
		log.Println("client got disconnected.")
	}()

	for {
		_, msg, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			return
		}
		fmt.Println("manager = ", c.manager)
		fmt.Println(string(msg))
		c.egress <- msg
	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	// consumer, _ := kafka.Consumer("deez")

	go func() {
		log.Println("uhm in go thread running kafka dawg")
	}()

	for {
		select {
		case msg, ok := <-c.egress:
			if !ok {

				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed :", err)
				}
				return
			}
			for xc := range c.manager.clientList {
				if err := xc.connection.WriteMessage(websocket.TextMessage, msg); err != nil {
					log.Println("error in writing mesage to client : ", err)
					return
				}
			}

			log.Println("sent message ; clientList length =", len(c.manager.clientList))
		}
	}
}
