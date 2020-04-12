package main

import (
	"net/http"

	"github.com/akantsevoi/integrationtests/cmd/database"
	"github.com/akantsevoi/integrationtests/cmd/server"
)

func main() {
	// TODO: create database.Options{}
	// TODO: get all the variables from os.Getenv()
	port := "80"
	r, _ := server.NewServer(database.Options{})
	http.ListenAndServe(":"+port, r.Router)
}
