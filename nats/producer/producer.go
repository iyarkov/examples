package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer func() {
		if err := nc.Drain(); err != nil {
			fmt.Printf("Failed to drain nats connection, %v", err)
		}
	}()

	nc.Publish("greet.joe", []byte("hello"))
}
