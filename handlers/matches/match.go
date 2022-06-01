package matches

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

type Stats struct {
    Name     string  `json:"name"`
    Nickname string  `json:"nickname"`
    Id       int     `json:"id"`
    Kills    int     `json:"kills"`
    Deaths   int     `json:"deaths"`
    Adr      float64 `json:"adr"`
    Kast     float64 `json:"kast"`
    Rating   float64 `json:"rating"`
}

type TeamStats struct {
    Name    string  `json:"name"`
    Logo    string  `json:"logo"`
    Result  int     `json:"result"`
    Players []Stats `json:"players"`
}

type Match struct {
    Id    int         `json:"id"`
    Date  string      `json:"date"`
    Teams []TeamStats `json:"teams"`
}

func ExtractPlayerId(link string) int {
    linkSlice := strings.Split(link, "/")
    id, _ := strconv.Atoi(linkSlice[2])
    return id
}

func ExtractKillsAndDeaths(kd string) (int, int) {
    kdSlice := strings.Split(kd, "-")
    kills, _ := strconv.Atoi(kdSlice[0])
    deaths, _ := strconv.Atoi(kdSlice[1])
    return kills, deaths
}

func GetMatch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    id := params["id"]

    url := fmt.Sprintf("https://www.hltv.org/matches/%s/_", id)

    c := colly.NewCollector()

    var Teams []TeamStats
    var match Match

    c.OnHTML("div.contentCol", func(h *colly.HTMLElement) {
        teamOneName := h.DOM.Find("div.team1-gradient").Find("div.teamName").Text()
        teamOneLogo, _ := h.DOM.Find("div.team1-gradient").Find("img").Attr("src")
        teamOneScore, _ := strconv.Atoi(h.DOM.Find("div.team1-gradient").Children().Eq(1).Text())

        teamTwoName := h.DOM.Find("div.team2-gradient").Find("div.teamName").Text()
        teamTwoLogo, _ := h.DOM.Find("div.team2-gradient").Find("img").Attr("src")
        teamTwoScore, _ := strconv.Atoi(h.DOM.Find("div.team2-gradient").Children().Eq(1).Text())

        // time := h.DOM.Find("div.timeAndEvent").Find("div.time").Text()
        date := h.DOM.Find("div.timeAndEvent").Find("div.date").Text()

        var Stats1 []Stats
        var Stats2 []Stats

        h.DOM.Find("table.table.totalstats").Eq(0).Find("tr").Not("tr.header-row").Each(func(i int, s *goquery.Selection) {
            name := s.Find("div.statsPlayerName").Eq(0).Text()
            nickname := s.Find("span.player-nick").Text()
            idString, _ := s.Find("a.flagAlign.no-maps-indicator-offset").Attr("href")
            id := ExtractPlayerId(idString)
            kd := s.Find("td.kd.text-center").Text()
            kills, deaths := ExtractKillsAndDeaths(kd)
            adrString := s.Find("td.adr.text-center").Text()
            adr, _ := strconv.ParseFloat(adrString, 64)
            kastPercent := s.Find("td.kast.text-center").Text()
            kastSlice := strings.Split(kastPercent, "%")
            kast, _ := strconv.ParseFloat(kastSlice[0], 64)
            ratingString := s.Find("td.rating.text-center").Text()
            rating, _ := strconv.ParseFloat(ratingString, 64)

            Stats1 = append(Stats1, Stats{
                name,
                nickname,
                id,
                kills,
                deaths,
                adr,
                kast,
                rating,
            })
        })

        h.DOM.Find("table.table.totalstats").Eq(1).Find("tr").Not("tr.header-row").Each(func(i int, s *goquery.Selection) {
            name := s.Find("div.statsPlayerName").Eq(0).Text()
            nickname := s.Find("span.player-nick").Text()
            idString, _ := s.Find("a.flagAlign.no-maps-indicator-offset").Attr("href")
            id := ExtractPlayerId(idString)
            kd := s.Find("td.kd.text-center").Text()
            kills, deaths := ExtractKillsAndDeaths(kd)
            adrString := s.Find("td.adr.text-center").Text()
            adr, _ := strconv.ParseFloat(adrString, 64)
            kastPercent := s.Find("td.kast.text-center").Text()
            kastSlice := strings.Split(kastPercent, "%")
            kast, _ := strconv.ParseFloat(kastSlice[0], 64)
            ratingString := s.Find("td.rating.text-center").Text()
            rating, _ := strconv.ParseFloat(ratingString, 64)

            Stats2 = append(Stats2, Stats{
                name,
                nickname,
                id,
                kills,
                deaths,
                adr,
                kast,
                rating,
            })
        })

        Teams = append(Teams, TeamStats{
            teamOneName,
            teamOneLogo,
            teamOneScore,
            Stats1,
        })

        Teams = append(Teams, TeamStats{
            teamTwoName,
            teamTwoLogo,
            teamTwoScore,
            Stats2,
        })

        newId, _ := strconv.Atoi(id)

        match = Match{
            newId,
            date,
            Teams,
        }
    })

    c.OnError(func(r *colly.Response, err error) {
        log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
    })

    c.OnRequest(func(r *colly.Request) {
        log.Println("Visiting", r.URL)
    })

    c.Visit(url)

    json.NewEncoder(w).Encode(match)
}