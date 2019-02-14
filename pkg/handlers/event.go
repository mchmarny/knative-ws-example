package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cloudevents/sdk-go/v02"
)

const (
	knownPublisherTokenName = "token"
)

// CloudEventHandler submitted messages
func CloudEventHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// check method
	if r.Method != http.MethodPost {
		log.Printf("wring method: %s", r.Method)
		http.Error(w, "Invalid method. Only POST supported", http.StatusMethodNotAllowed)
		return
	}

	// parse form to update
	if err := r.ParseForm(); err != nil {
		log.Printf("error parsing form: %v", err)
		http.Error(w, fmt.Sprintf("Post content error (%s)", err),
			http.StatusBadRequest)
		return
	}

	// check for presense of publisher token
	srcToken := r.URL.Query().Get(knownPublisherTokenName)
	if srcToken == "" {
		log.Printf("nil token: %s", srcToken)
		http.Error(w, fmt.Sprintf("Invalid request (%s missing)", knownPublisherTokenName),
			http.StatusBadRequest)
		return
	}

	// check validity of poster token
	if knownPublisherToken != srcToken {
		log.Printf("invalid token: got(%s) expected(%s)", srcToken, knownPublisherToken)
		http.Error(w, fmt.Sprintf("Invalid publisher token value (%s)", knownPublisherTokenName),
			http.StatusBadRequest)
		return
	}

	converter := v02.NewDefaultHTTPMarshaller()
	event, err := converter.FromRequest(r)
	if err != nil {
		log.Printf("error parsing cloudevent: %v", err)
		http.Error(w, fmt.Sprintf("Invalid Cloud Event (%v)", err),
			http.StatusBadRequest)
		return
	}

	log.Printf("Event: %v", event)
	eventData, ok := event.Get("data")
	if !ok {
		http.Error(w, "Error, not a cloud event data", http.StatusBadRequest)
		return
	}

	// push event to the channel
	eventChannel <- eventData

	// response with the parsed payload data
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(eventData)

}
