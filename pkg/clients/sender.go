package clients

import (
	"bytes"
	"log"
	"time"
	"net/url"
	"encoding/json"
    "net/http"
	"context"

	"github.com/mchmarny/myevents/pkg/utils"
	"github.com/cloudevents/sdk-go/v02"
)

// NewSender creates a already preconfigured Sender
func NewSender(targerURL string) (sender *Sender, err error) {

	tu, te := url.Parse(targerURL)
	if te != nil {
		return nil, te
	}

	su, se := url.Parse("https://github.com/mchmarny/myevents")
	if se != nil {
		return nil, se
	}

	s := &Sender{
		TargerURL: tu.String(),
		SourceURL: su,
	}

	return s, nil

}

// Sender sends messages
type Sender struct {
	TargerURL string
	SourceURL *url.URL
}

// SendMessages sends v02.Event based on the provided data
func (s *Sender) SendMessages(ctx context.Context, eventType, text string) error {

	now := time.Now().UTC()
	event := &v02.Event{
		SpecVersion: "0.2",
		Type:        eventType,
		Source:      *s.SourceURL,
		ID:          utils.MakeUUID(),
		Time: 		 &now,
		ContentType: "text/plain",
		Data: 		 text,
	}

	return s.SendEvent(ctx, event)

}


// SendEvent sends a v2 cloud event
func (s *Sender) SendEvent(ctx context.Context, event *v02.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error while marshaling event: %v", err)
	}
	return s.SendContent(ctx, data)
}


// SendContent sends the content
func (s *Sender) SendContent(ctx context.Context, content []byte) error {

	// request
	req, err := http.NewRequest("POST", s.TargerURL, bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	// update request
	req.WithContext(ctx)
	req.Header.Set("k-sender", "github.com/mchmarny/myevents")
	req.Header.Set("Content-Type", "application/json")

	// client
	client := &http.Client{}

	// send
    resp, err := client.Do(req)
    if err != nil {
        return err
	}

	// cleanup
    defer resp.Body.Close()
	log.Printf("Send status: %s", resp.Status)
	return nil
}





