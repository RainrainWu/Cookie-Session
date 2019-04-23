package main

import (
	"encoding/gob"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// User define user log in data
type User struct {
	Username      string
	Authenticated bool
}

var store *sessions.FilesystemStore
var tpl *template.Template

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewFilesystemStore(
		"sessions/",
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(User{})

	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

// Deploy deploys a router and specify a target port to listen
func Deploy() {

	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/login", login)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/forbidden", forbidden)
	router.HandleFunc("/secret", secret)
	http.ListenAndServe(":8080", router)
}
