package handlers

import (
	"log"

	"golang.org/x/net/websocket"
)

var (
	eventChannel = make(chan interface{}, 1)
)

// WSHandler provides backing service for the UI
func WSHandler(ws *websocket.Conn) {
	log.Println("WS connection...")
	for {
		select {
		case m := <-eventChannel:
			if err := websocket.JSON.Send(ws, m); err != nil {
				log.Printf("Error on write message: %v", err)
				break
			}
		}
	}

}
