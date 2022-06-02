package teams

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// Extracts id from string and converts it to an int
func ExtractId(link string) int {
    stringId := strings.Split(link, "/")
    id, _ := strconv.Atoi(stringId[2])
    return id
}

// Extracts ranking from string and converts it to an int
func ExtractRanking(ranking string) int {
    ranking = strings.Replace(ranking, "#", "", 1)
    rankingInt, _ := strconv.Atoi(ranking)
    return rankingInt
}

func GetTeams(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    
    c := colly.NewCollector()

    var Teams = []Team{}

    c.OnHTML("div.ranked-team", func(h *colly.HTMLElement) {
        teamProfileLink := h.ChildAttr("a.moreLink", "href")
        id := ExtractId(teamProfileLink)
        rankingString := h.ChildText("span.position")
        ranking := ExtractRanking(rankingString)
        name := h.ChildText("span.name")
        logo := h.ChildAttr("img", "src")

        var Players = []Player{}

        // Finds the HTML table that contains a team's player info
        h.DOM.Find("table.lineup").Find("td.player-holder").Each(func(i int, s *goquery.Selection) {
            fullname, _ := s.Find("img.playerPicture").Attr("title")
            playerImage, _ := s.Find("img.playerPicture").Attr("src")
            nickname := s.Find("div.nick").Text()
            name, _ := s.Find("img.gtSmartphone-only.flag").Attr("title")
            flag, _ := s.Find("img.gtSmartphone-only.flag").Attr("src")

            Players = append(Players, Player{
                fullname,
                playerImage,
                nickname,
                Country{
                    name,
                    flag,
                },
            })
        })

        Teams = append(Teams, Team{
            id,
            ranking,
            name,
            logo,
            Players,
        })
    })

    c.OnRequest(func(r *colly.Request) {
        log.Println("Visiting", r.URL)
    })

    c.OnError(func(r *colly.Response, err error) {
        log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
    })

    c.Visit("https://www.hltv.org/ranking/teams")

    json.NewEncoder(w).Encode(Teams)
}