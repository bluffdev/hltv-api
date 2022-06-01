package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/landoniwnl/hltv-api/routes"
)

func main() {
    router := mux.NewRouter()
    
    routes.RegisterRoutes(router)
    http.Handle("/", router)

    fmt.Println("Listening on port 3000")
    log.Fatal(http.ListenAndServe("localhost:3000", router))
}