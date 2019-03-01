package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	"github.com/mchmarny/knative-ws-example/pkg/handlers"
	"github.com/mchmarny/knative-ws-example/pkg/utils"

	"golang.org/x/net/websocket"
)

func main() {
	ctx := context.Background()

	// Server configured
	port, err := strconv.Atoi(utils.MustGetEnv("PORT", "8080"))
	if err != nil {
		log.Fatalf("failed to parse port to int, %s", err.Error())
	}

	// init configs
	handlers.InitHandlers()

	// Static
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	// UI Handlers
	http.HandleFunc("/", handlers.RootHandler)
	http.Handle("/ws", websocket.Handler(handlers.WSHandler))

	// Ingres API Handler
	c, err := client.NewHTTPClient(client.WithHTTPPath("/v1/event"))
	if err != nil {
		log.Fatalf("failed to create cloudevents client, %s", err.Error())
	}

	// Health Handler
	http.HandleFunc("/_health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})

	if err := c.StartReceiver(ctx, handlers.CloudEventReceived); err != nil {
		log.Fatalf("failed to start cloudevents receiver, %s", err.Error())
	}
	log.Printf("Server starting on port %d \n", port)

	// Block until done.
	<-ctx.Done()
}
