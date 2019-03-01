package handlers

import (
	"fmt"
	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/datacodec"
	"log"
)

const (
	knownPublisherTokenName = "token"
)

func init() {
	datacodec.AddDecoder("text/plain", plainTextDecoder)
}

func plainTextDecoder(in, out interface{}) error {
	if in == nil {
		return nil
	}

	b, ok := in.([]byte)
	if !ok {
		return fmt.Errorf("unable to decode input for non-[]byte type")
	}

	s, ok := out.(*string)
	if !ok {
		return fmt.Errorf("unable to decode output for non-*string type")
	}
	*s = string(b)
	return nil
}

func CloudEventReceived(event cloudevents.Event) {
	//// check for presence of publisher token
	var srcToken string
	ctx := event.Context.AsV02()
	if ctx.Extensions != nil {

		if t, ok := ctx.Extensions[knownPublisherTokenName]; ok {
			if srcToken, ok = t.(string); !ok {
				log.Printf("Invalid request (%s missing)", knownPublisherTokenName)
				return
			}
		}
	}

	if srcToken == "" {
		log.Printf("nil token: %s", srcToken)
		return
	}

	// check validity of poster token
	if knownPublisherToken != srcToken {
		log.Printf("invalid token: %s", srcToken)
		return
	}

	log.Printf("Event: %v", event)

	//var data *string
	data := ""
	if err := event.DataAs(&data); err != nil {
		log.Printf("Failed to DataAs: %s", err.Error())
		return
	}

	eventChannel <- data
}
