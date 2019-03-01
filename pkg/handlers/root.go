package handlers

import (
	"log"
	"net/http"
)

// RootHandler handles view page
func RootHandler(w http.ResponseWriter, r *http.Request) {

	proto := r.Header.Get("x-forwarded-proto")
	if proto == "" {
		proto = "http"
	}

	data := make(map[string]interface{})

	data["host"] = r.Host
	data["proto"] = proto

	log.Printf("data: %v", data)

	// anonymous
	if err := templates.ExecuteTemplate(w, "index", data); err != nil {
		log.Printf("Error in home template: %s", err)
	}

}
