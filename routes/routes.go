package routes

import (
	"github.com/gorilla/mux"
	"github.com/landoniwnl/hltv-api/handlers"
)

func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/teams", handlers.GetTeams).Methods("GET")
    router.HandleFunc("/teams/{id}", handlers.GetTeam).Methods("GET")
	router.HandleFunc("/players", handlers.GetPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", handlers.GetPlayer).Methods("GET")
	router.HandleFunc("/matches", handlers.GetMatches).Methods("GET")
	router.HandleFunc("/results", handlers.GetResults).Methods("GET")
	router.HandleFunc("/matches/{id}", handlers.GetMatch).Methods("GET")
}