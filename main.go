package main

import (
	"github.com/gorilla/mux"
	"leanpub-app/app"
	"log"
	"net/http"
)

func main() {
	application := app.CreateApp()
	application.Router = mux.NewRouter()
	application.Setup()
	http.Handle("/", application.Router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
