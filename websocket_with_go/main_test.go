package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestWebSocket(t *testing.T) {
	serverAddr := "localhost:5051"
	server := &http.Server{
		Addr:    serverAddr,
		Handler: http.HandlerFunc(handleConnection),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Fatalf("Server failed: %v", err)
		}
	}()

	time.Sleep(1 * time.Second)

	url := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	originalMessage := "Hello, WebSocket!"
	err = conn.WriteMessage(websocket.TextMessage, []byte(originalMessage))
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	_, message, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	if string(message) != originalMessage {
		t.Errorf("Expected message %s, but got %s", originalMessage, message)
	}

	if err := server.Close(); err != nil {
		t.Fatalf("Failed to close server: %v", err)
	}
}
