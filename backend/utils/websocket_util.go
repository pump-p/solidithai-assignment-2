package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pump-p/solidithai-assignment-2/backend/config"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

// Message represents the structure for incoming/outgoing WebSocket messages
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

// HandleWebSocketConnections handles new WebSocket connections
func HandleWebSocketConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to upgrade to WebSocket: %v\n", err)
		return
	}

	// Register new client
	clients[conn] = true

	// Listen for new messages from this client
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("Error reading JSON: %v\n", err)
			delete(clients, conn)
			conn.Close()
			break
		}
		msg.Time = time.Now().Format(time.RFC3339) // Add timestamp to message
		broadcast <- msg

		// Store the message in Elasticsearch
		go storeLog(msg)
	}
}

func init() {
	// Start broadcasting messages
	go handleMessages()
}

// handleMessages listens for incoming messages on the broadcast channel and sends them to all clients
func handleMessages() {
	for {
		msg := <-broadcast
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

// storeLog stores a log message in Elasticsearch
func storeLog(msg Message) {
	esClient := config.ESClient
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("Failed to marshal message: %v\n", err)
		return
	}

	// Index the log into Elasticsearch
	_, err = esClient.Index(
		"streaming_logs",
		bytes.NewReader(data),
		esClient.Index.WithContext(context.Background()),
	)
	if err != nil {
		fmt.Printf("Failed to index message: %v\n", err)
	}
}
