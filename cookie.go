package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// NewRouter : Set new router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods("GET").
		Path("/set").
		Name("SetCookie").
		HandlerFunc(SetCookie)
	router.
		Methods("GET").
		Path("/get").
		Name("GetCookie").
		HandlerFunc(GetCookie)

	return router
}

// SetCookie will set cookies in the user browser
func SetCookie(w http.ResponseWriter, r *http.Request) {

	// Set expiration time
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)

	// Generate cookie
	cookie := http.Cookie{Name: "username", Value: "rain", Expires: expiration}

	// Set cookie
	http.SetCookie(w, &cookie)
}

// GetCookie : Get cookies in the user browser
func GetCookie(w http.ResponseWriter, r *http.Request) {

	// Specified
	if cookie, err := r.Cookie("username"); err == nil {

		fmt.Fprintf(w, cookie.Name)
	}

	fmt.Fprintf(w, "\n")

	// Loop all
	for _, cookie := range r.Cookies() {

		fmt.Fprintf(w, cookie.Name)
	}
}
