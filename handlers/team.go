package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

func GetTeam(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    id := params["id"]

    url := fmt.Sprintf("https://www.hltv.org/team/%s/_", id)

    c := colly.NewCollector()
    
    var Players []Player

    c.OnHTML("a.col-custom", func(e *colly.HTMLElement) {
        fullname := e.ChildAttr("img.bodyshot-team-img", "title")
        image := e.ChildAttr("img.bodyshot-team-img", "src")
        nickname := e.ChildText("span.text-ellipsis.bold")
        countryName := e.ChildAttr("img.flag", "title")
        flag := e.ChildAttr("img.flag", "src")

        Players = append(Players, Player{
            fullname,
            image,
            nickname,
            Country{
                countryName,
                flag,
            },
        })
    })

    var selectedTeam Team 

    c.OnHTML("div.standard-box.profileTopBox.clearfix", func(h *colly.HTMLElement) {
    //    fmt.Println(h.ChildText("div.team-country.text-ellipsis"))
        idInt, _ := strconv.Atoi(id)
        ranking := h.DOM.Find("span.right").First().Text()
        ranking = strings.Replace(ranking, "#", "", 1)
        rankingInt, _ := strconv.Atoi(ranking)
        name := h.ChildText("h1.profile-team-name.text-ellipsis")
        logo := h.ChildAttr("img.teamlogo", "src")

        selectedTeam = Team{
            idInt,
            rankingInt,
            name,
            logo,
            Players,
        }
    })

    c.OnRequest(func(r *colly.Request) {
        log.Println("Visiting", r.URL)
    })

    c.OnError(func(r *colly.Response, err error) {
        log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
    })

    c.Visit(url)

    json.NewEncoder(w).Encode(selectedTeam)
}