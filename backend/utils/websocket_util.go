package utils

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

// Message object to handle incoming/outgoing messages
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

// HandleWebSocketConnections sets up WebSocket connections
func HandleWebSocketConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to upgrade to WebSocket: %v\n", err)
		return
	}

	// Register new client
	clients[conn] = true

	// Listen indefinitely for new messages
	for {
		var msg Message
		// Read message from client
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("Error reading JSON: %v\n", err)
			delete(clients, conn)
			conn.Close()
			break
		}
		// Send received message to broadcast channel
		broadcast <- msg
	}
}

func init() {
	// Start broadcasting messages in a separate goroutine
	go handleMessages()
}

// handleMessages listens for incoming messages on the broadcast channel
// and sends them to all connected clients
func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Printf("Error writing JSON: %v\n", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
