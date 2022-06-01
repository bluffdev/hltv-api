package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/landoniwnl/hltv-api/routes"
)

func main() {
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "3000"
	}

    router := mux.NewRouter()
    
    routes.RegisterRoutes(router)
    http.Handle("/", router)

    fmt.Println("Listening on port " + port)
    log.Fatal(http.ListenAndServe(":" + port, router))
}