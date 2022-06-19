package matches

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Event struct {
    Name string `json:"name"`
    Logo string `json:"logo"`
}

type MatchTeam struct {
    Name string `json:"name"`
    Logo string `json:"logo"`
}

type Matches struct {
    Id    int         `json:"id"`
    Date  string      `json:"date"`
    Time  string      `json:"time"`
    Event Event       `json:"event"`
    Stars int         `json:"stars"`
    Maps  string      `json:"maps"`
    Teams []MatchTeam `json:"teams"`
}

func ExtractMatchId(link string) int {
    idString := strings.Split(link, "/")
    id, _ := strconv.Atoi(idString[2])
    return id 
}

func GetMatches(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    c := colly.NewCollector()

    var matches []Matches
    c.OnHTML("div.liveMatch-container", func(h *colly.HTMLElement) {
        idString := h.ChildAttr("a.match.a-reset", "href")
        id := ExtractMatchId(idString)
        starsString := h.Attr("stars")
        stars, _ := strconv.Atoi(starsString)
        maps := h.ChildText("div.matchMeta")
        eventName := h.ChildText("div.matchEventName.gtSmartphone-only")
        eventLogo := h.ChildAttr("img.matchEventLogo", "src")

        if string(eventLogo[0]) == "/" {
            eventLogo = "https://www.hltv.org/" + eventLogo
        }

        var Teams []MatchTeam

        h.DOM.Find("div.matchTeams.text-ellipsis").Each(func(i int, s *goquery.Selection) {
            firstTeamName := s.Find("div.matchTeamName.text-ellipsis").Eq(0).Text()
            firstTeamLogo, _ := s.Find("img.matchTeamLogo").Eq(0).Attr("src")
            
            secondTeamName := s.Find("div.matchTeamName.text-ellipsis").Eq(1).Text()
            secondTeamLogo, _ := s.Find("img.matchTeamLogo").Eq(1).Attr("src")

            Teams = append(Teams, MatchTeam{
                firstTeamName,
                firstTeamLogo,
            })

            Teams = append(Teams, MatchTeam{
                secondTeamName,
                secondTeamLogo,
            })
        })

        matches = append(matches, Matches{
            id,
            "Today",
            "Live",
            Event{
                eventName,
                eventLogo,
            },
            stars,
            maps,
            Teams,
        })
    })

    c.OnHTML("div.upcomingMatchesSection", func(h *colly.HTMLElement) {
        date := h.DOM.Children().Eq(0).Text()

        h.DOM.Find("a.match.a-reset").Each(func(i int, s *goquery.Selection) {
            idString, _ := s.Attr("href")
            id := ExtractMatchId(idString)
            time := s.Find("div.matchTime").Text()
            starsString, _ := s.Parent().Attr("stars")
            stars, _ := strconv.Atoi(starsString)
            maps := s.Find("div.matchMeta").Text()
            eventName, _ := s.Find("img.matchEventLogo").Attr("title")
            eventLogo, _ := s.Find("img.matchEventLogo").Attr("src")

            var Teams []MatchTeam
            s.Find("div.matchTeamLogoContainer").Each(func(i int, s *goquery.Selection) {
                teamName, _ := s.Find("img.matchTeamLogo").Eq(0).Attr("title")
                teamLogo, _ := s.Find("img.matchTeamLogo").Eq(0).Attr("src")

                Teams = append(Teams, MatchTeam{
                    teamName,
                    teamLogo,
                })
            })

            matches = append(matches, Matches{
                id,
                date,
                time,
                Event{
                    eventName,
                    eventLogo,
                },
                stars,
                maps,
                Teams,
            })
        })
    })

    c.OnRequest(func(r *colly.Request) {
        log.Println("Visiting", r.URL)
    })

    c.OnError(func(r *colly.Response, err error) {
        log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
    })

    c.Visit("https://www.hltv.org/matches")

    json.NewEncoder(w).Encode(matches)
}