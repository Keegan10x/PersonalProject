package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler is the handler for the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

// AboutHandler is the handler for the about route
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us!")
}
