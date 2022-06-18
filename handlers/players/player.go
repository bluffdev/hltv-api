package players

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

type PlayerTeam struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

type PlayerStats struct {
    Id           int        `json:"id"`
    Team         PlayerTeam `json:"team"`
    Image        string     `json:"image"`
    Nickname     string     `json:"nickname"`
    Age          int        `json:"age"`
    Rating       float64    `json:"rating"`
    Impact       float64    `json:"impact"`
    Dpr          float64    `json:"dpr"`
    Apr          float64    `json:"apr"`
    Kast         float64    `json:"kast"`
    Kpr          float64    `json:"kpr"`
    HsPercentage float64    `json:"hsPercentage"`
    MapsPlayed   int        `json:"mapsPlayed"`
}

// TODO: ADD THIS FUNCTION TO A UTILS FILE
func ExtractId(link string) int {
    stringId := strings.Split(link, "/")
    id, _ := strconv.Atoi(stringId[3])
    return id
}

func RemovePercent(percentage string) (float64, error) {
    split := strings.Split(percentage, "%")
    return strconv.ParseFloat(split[0], 64)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    params := mux.Vars(r)
    id := params["id"]

    url := fmt.Sprintf("https://www.hltv.org/stats/players/%s/_", id)

    c := colly.NewCollector()

    var hsPercentage float64 
    var mapsPlayed int
    var Player PlayerStats

    c.OnHTML("div.columns", func(h *colly.HTMLElement) {
        hsPercentageString := h.DOM.Find("div.col.stats-rows.standard-box").Eq(0).Find("span").Eq(3).Text()
        hsPercentage, _ = RemovePercent(hsPercentageString)
        mapsPlayedString := h.DOM.Find("div.col.stats-rows.standard-box").Eq(0).Find("span").Eq(13).Text()
        mapsPlayed, _ = strconv.Atoi(mapsPlayedString)
    })

    c.OnHTML("div.playerSummaryStatBox", func(h *colly.HTMLElement) {
        link := h.ChildAttr("a.a-reset.text-ellipsis", "href")
        teamId := ExtractId(link)
        teamName := h.ChildText("a.a-reset.text-ellipsis")
        image := h.ChildAttr("img.summaryBodyshot", "src")
        nickname := h.ChildText("h1.summaryNickname.text-ellipsis")
        age := h.ChildText("div.summaryPlayerAge")
        ageString := strings.Split(age, " ")
        ageInt, _ := strconv.Atoi(ageString[0])
        ratingString := h.DOM.Find("div.summaryStatBreakdownDataValue").Eq(0).Text()
        rating, _ := strconv.ParseFloat(ratingString, 64)
        impactString := h.DOM.Find("div.summaryStatBreakdownDataValue").Eq(3).Text()
        impact, _ := strconv.ParseFloat(impactString, 64)
        dprString := h.DOM.Find("div.summaryStatBreakdownDataValue").Eq(1).Text()
        dpr, _ := strconv.ParseFloat(dprString, 64)
        aprString := h.DOM.Find("div.summaryStatBreakdownDataValue").Eq(4).Text()
        apr, _ := strconv.ParseFloat(aprString, 64)
        kastString := h.DOM.Find("div.summaryStatBreakdownDataValue").Eq(2).Text()
        kprString := h.DOM.Find("div.summaryStatBreakdownDataValue").Eq(5).Text()
        kpr, _ := strconv.ParseFloat(kprString, 64)
        kast, _ := RemovePercent(kastString)

        idInt, _ := strconv.Atoi(id)

        Player = PlayerStats{
            idInt,
            PlayerTeam{
                teamId,
                teamName,
            },
            image,
            nickname,
            ageInt,
            rating,
            impact,
            dpr,
            apr,
            kast,
            kpr,
            hsPercentage,
            mapsPlayed,
        }
    })

    c.OnRequest(func(r *colly.Request) {
        log.Println("Visiting", r.URL)
    })

    c.OnError(func(r *colly.Response, err error) {
        log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
    })

    c.Visit(url)

    json.NewEncoder(w).Encode(Player)
}