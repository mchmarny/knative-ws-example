package handlers

import (
	"log"

	"golang.org/x/net/websocket"
)

var (
	eventChannel = make(chan interface{}, 100)
	connections []*websocket.Conn
)

// WSHandler provides backing service for the UI
func WSHandler(ws *websocket.Conn) {
	log.Println("WS connection...")

	connections = append(connections, ws)

	for {
		select {
		case m := <-eventChannel:
			for _, w := range connections {
				if err := websocket.JSON.Send(w, m); err != nil {
					log.Printf("Error on write message: %v", err)
				}
			}
		}
	}
}
