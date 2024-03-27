package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Manager struct {
	name       string
	clientList ClientList
	sync.RWMutex
}

func NewManager(managerName string) *Manager {
	return &Manager{
		name:       managerName,
		clientList: make(ClientList),
	}
}

func (m *Manager) addClient(client *Client) {

	m.Lock()
	defer m.Unlock()
	m.clientList[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clientList[client]; ok {
		client.connection.Close()
		delete(m.clientList, client)
	}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	log.Println("New Connection")

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m)
	m.addClient(client)

	go client.readMessages()
	go client.writeMessages()
}
