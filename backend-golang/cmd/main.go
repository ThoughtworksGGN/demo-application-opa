package main

import (
	"demo-application-opa-golang/api"
	"demo-application-opa-golang/data"
	"log"
	"net/http"
)

func main() {
	store := data.Store{}
	store.Initialise()

	handler := api.InitialiseWebService(&store)

	log.Fatal(http.ListenAndServe(":8124", handler))
}
