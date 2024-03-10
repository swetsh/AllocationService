package main

import (
	"log"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	go main()

	conn, err := net.Dial("tcp", "localhost:8089")
	if err != nil {
		t.Fatalf("failed to connect to server: %s", err)
	}
	defer conn.Close()

	log.Println("Server started successfully")

}
