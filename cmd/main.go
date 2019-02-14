package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mchmarny/knative-ws-example/pkg/handlers"
	"github.com/mchmarny/knative-ws-example/pkg/utils"

	"golang.org/x/net/websocket"
)

func main() {

	// init configs
	handlers.InitHandlers()


	// Static
	http.Handle("/static/", http.StripPrefix("/static/",
		  http.FileServer(http.Dir("static"))))

	// UI Handlers
	http.HandleFunc("/", handlers.RootHandler)
	http.Handle("/ws", websocket.Handler(handlers.WSHandler))

	// Ingres API Handler
	http.HandleFunc("/v1/event", handlers.CloudEventHandler)

	// Health Handler
	http.HandleFunc("/_health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})

	// Server configured
	port := utils.MustGetEnv("PORT", "8080")

	log.Printf("Server starting on port %s \n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
