package handlers

import (
	"fmt"
	viewers "main/dbl/locations"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
)

const customTimeLayout = "2006-01-02T15:04:05"

// HomeHandler is the handler for the home route
func HomeHandler() http.Handler {

	jsonUX := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "UX",
	}.Froze()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := jsonUX.NewEncoder(w)
		err := encoder.Encode("hello from the home handler")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			// log out error
			fmt.Printf("failed to hit home handler with error: %s", err.Error())
			return
		}
		fmt.Println("successfully returned from home handler") // log out success
	})
}

// LocationValues handles a request for a range of device locations
func LocationValues(locationDBL viewers.LocationDBL) http.HandlerFunc {

	jsonUX := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "UX",
	}.Froze()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		//log the incoming params
		params := mux.Vars(r)
		fmt.Println("logging params back from handler", params) // replace with logger

		//extract the id
		id, ok := params["id"]
		if !ok || len(id) == 0 {
			http.Error(w, "id not provided", http.StatusBadRequest)
			err := fmt.Errorf("id %s was mot found", id)

			//log out error
			fmt.Println(err) // replace with logger
			return
		}

		startTime, err := time.Parse(customTimeLayout, r.FormValue("startTime"))
		if err != nil {
			http.Error(w, "cannot parse startTime", http.StatusBadRequest)

			//log out error
			fmt.Printf("startTime cannot be parsed with err %s", err.Error()) //replace with logger
			return
		}

		stopTime, err := time.Parse(customTimeLayout, r.FormValue("stopTime"))
		if err != nil {
			http.Error(w, "cannot parse stopTime", http.StatusBadRequest)

			//log out error
			fmt.Printf("stopTime cannot be parsed with err %s", err.Error()) //replace with logger
			return
		}

		locations, _, err := locationDBL.GetLocationsByTimeRange(startTime, stopTime)
		if err != nil {
			http.Error(w, "cannot parse stopTime", http.StatusInternalServerError)

			//log out error
			fmt.Printf("failure occured in DeviceLocationValues(DBL) with err %s", err.Error()) //replace with logger
			return
		}

		encoder := jsonUX.NewEncoder(w)
		err = encoder.Encode(locations)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			//log out error
			fmt.Printf("failed to encode device location values with err %s", err.Error()) //replace with logger
			return
		}

		// log out success
		fmt.Println("successfully returned device location values") //replace with logger
	}
}
