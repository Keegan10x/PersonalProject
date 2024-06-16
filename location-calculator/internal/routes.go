package internal

import (
	viewers "main/dbl/locations"
	"main/location-calculator/internal/handlers"

	"github.com/gorilla/mux"
)

// PublicRoutes contains all public routes exposed by the service.
// is to be consumed by the main server package
func PublicRoutes(router *mux.Router, locationDBL viewers.LocationDBL) {
	router.Handle("/", handlers.HomeHandler()).Methods("GET")
	router.Handle("/location-calculator/{id}", handlers.LocationValues(locationDBL)).Methods("GET")
}
