package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// ErrorHandler handles all app errors
func ErrorHandler(w http.ResponseWriter, r *http.Request, err error, code int) {

	log.Printf("Error: %v", err)
	errMsg := fmt.Sprintf("%+v", err)

	w.WriteHeader(code)
	if err := templates.ExecuteTemplate(w, "error", map[string]interface{}{
		"error":       errMsg,
		"status_code": code,
		"status":      http.StatusText(code),
	}); err != nil {
		log.Printf("Error in error template: %s", err)
	}

}
