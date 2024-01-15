package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow connections from any origin
	},
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to websocket:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected")

	// Listen for messages from the client
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading from websocket:", err)
			break
		}

		log.Printf("Received message: %s\n", message)

		// Echo the message back to the client
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Error writing to websocket:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/echo", echoHandler)

	addr := "localhost:8080"
	fmt.Printf("WebSocket server started on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
