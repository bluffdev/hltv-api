package routes

import (
	"github.com/gorilla/mux"
	"github.com/landoniwnl/hltv-api/handlers/matches"
	"github.com/landoniwnl/hltv-api/handlers/players"
	"github.com/landoniwnl/hltv-api/handlers/results"
	"github.com/landoniwnl/hltv-api/handlers/teams"
)

func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/teams", teams.GetTeams).Methods("GET")
    router.HandleFunc("/teams/{id}", teams.GetTeam).Methods("GET")
    router.HandleFunc("/players", players.GetPlayers).Methods("GET")
    router.HandleFunc("/players/{id}", players.GetPlayer).Methods("GET")
    router.HandleFunc("/matches", matches.GetMatches).Methods("GET")
    router.HandleFunc("/matches/{id}", matches.GetMatch).Methods("GET")
    router.HandleFunc("/results", results.GetResults).Methods("GET")
}