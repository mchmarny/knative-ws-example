package handlers

import (
	"html/template"
	"log"

	"github.com/mchmarny/knative-ws-example/pkg/utils"
)

var (
	// Templates for handlers
	templates           *template.Template
	knownPublisherToken string
)

// InitHandlers initializes OAuth package
func InitHandlers() {

	// Templates
	tmpls, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error while parsing templates: %v", err)
	}
	templates = tmpls

	// know publisher
	knownPublisherToken = utils.MustGetEnv("KNOWN_PUBLISHER_TOKEN", "")

}
