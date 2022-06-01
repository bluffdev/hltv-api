package players

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type PlayersStats struct {
    Id         int    `json:"id"`
    Nickname   string `json:"nickname"`
    Team       string `json:"team"`
    Slug       string `json:"slug"`
    MapsPlayed string `json:"mapsPlayed"`
    Kd         string `json:"kd"`
    Rating     string `json:"rating"`
}

func ExtractIdAndSlug(link string) (int, string) {
    linkSlice := make([]string, 5)
    copy(linkSlice, strings.Split(link, "/"))
    id, _ := strconv.Atoi(linkSlice[3])
    slug := linkSlice[4]
    return id, slug
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    c := colly.NewCollector()

    var Players []PlayersStats

    c.OnHTML("table.stats-table.player-ratings-table", func(h *colly.HTMLElement) {
        h.DOM.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
            link, _ := s.Find("td.playerCol").Find("a").Attr("href")
            id, slug := ExtractIdAndSlug(link)
            nickname := s.Find("td.playerCol").Find("a").Text()
            team, _ := s.Find("td.teamCol").Find("a").Find("img").Attr("title")
            mapsPlayed := s.Find("td.statsDetail").First().Text()
            kd := s.Find("td.statsDetail").Eq(2).Text()
            rating := s.Find("td.ratingCol").Text()
            Players = append(Players, PlayersStats{
                id,
                nickname,
                team,
                slug,
                mapsPlayed,
                kd,
                rating,
            })
        })
    })

    c.OnRequest(func(r *colly.Request) {
        log.Println("Visiting", r.URL)
    })

    c.OnError(func(r *colly.Response, err error) {
        log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
    })

    c.Visit("https://www.hltv.org/stats/players")

    json.NewEncoder(w).Encode(Players)
}