package main

import (
	"net/http"

	"github.com/libdolf/go-api/config"
	"github.com/libdolf/go-api/handler"
)

func main() {

	datastore := config.NewDatastore()

	// Criar uma instância de UserHandler com o ponteiro para config.Datastore

	userH := handler.NewUserHandler(datastore)

	mux := http.NewServeMux()
	mux.Handle("/users", userH)
	mux.Handle("/users/", userH)

	http.ListenAndServe(":8080", mux)

}
