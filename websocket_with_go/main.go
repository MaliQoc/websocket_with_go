package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error while reading message:", err)
			break
		}
		fmt.Printf("Received message: %s\n", msg)

		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			fmt.Println("Error while writing message:", err)
			break
		}
	}
}

func runTest() {
	serverAddr := "localhost:6061"
	server := &http.Server{
		Addr:    serverAddr,
		Handler: http.HandlerFunc(handleConnection),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server failed: %v", err)
		}
	}()

	time.Sleep(1 * time.Second)

	url := "ws://" + serverAddr + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	originalMessage := "Hello, WebSocket!"
	err = conn.WriteMessage(websocket.TextMessage, []byte(originalMessage))
	if err != nil {
		fmt.Printf("Failed to send message: %v\n", err)
		return
	}

	_, message, err := conn.ReadMessage()
	if err != nil {
		fmt.Printf("Failed to read message: %v\n", err)
		return
	}

	if string(message) != originalMessage {
		fmt.Printf("Test failed: expected message %s, but got %s\n", originalMessage, message)
	} else {
		fmt.Println("Test passed: received expected message")
	}

	if err := server.Close(); err != nil {
		fmt.Printf("Failed to close server: %v\n", err)
	}
}

func main() {
	runTest()
}
