package handlers

import (
	"log"
	"net/http"
)

// RootHandler handles view page
func RootHandler(w http.ResponseWriter, r *http.Request) {

	// if POST on root
	if r.Method == http.MethodPost {
		CloudEventHandler(w, r)
		return
	}

	data := make(map[string]interface{})

	// anonymous
	if err := templates.ExecuteTemplate(w, "index", data); err != nil {
		log.Printf("Error in home template: %s", err)
	}

}
