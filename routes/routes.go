package routes

import (
	"github.com/gorilla/mux"
	"github.com/landoniwnl/hltv-api/handlers"
)

func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/teams", handlers.GetTeams).Methods("GET")
    router.HandleFunc("/teams/{id}", handlers.GetTeam).Methods("GET")
}