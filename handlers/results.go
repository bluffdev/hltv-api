package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func ExtractDate(date string) string {
    dateSlice := strings.Split(date, " ")
    newDate := strings.Join(dateSlice[2:], " ")
    return newDate
}

func ExtractMatchId2(link string) int {
    idString := strings.Split(link, "/")
    id, _ := strconv.Atoi(idString[2])
    return id 
}

func GetResults(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    c := colly.NewCollector()

    var Results []Result
    c.OnHTML("div.results-sublist", func(h *colly.HTMLElement) {
        resultsDate := h.DOM.Children().Eq(0).Text()
        date := ExtractDate(resultsDate)

        h.DOM.Find("div.result-con").Find("tr").Each(func(i int, s *goquery.Selection) {
            teamOneName, _ := s.Find("div.line-align.team1").Find("img").Eq(0).Attr("title")
            teamOneLogo, _ := s.Find("div.line-align.team1").Find("img").Eq(0).Attr("src")
            teamOneScoreString := s.Find("td.result-score").Find("span").Eq(0).Text()
            teamOneScore, _ := strconv.Atoi(teamOneScoreString)

            teamTwoName, _ := s.Find("div.line-align.team2").Find("img").Eq(0).Attr("title")
            teamTwoLogo, _ := s.Find("div.line-align.team2").Find("img").Eq(0).Attr("src")
            teamTwoScoreString := s.Find("td.result-score").Find("span").Eq(1).Text()
            teamTwoScore, _ := strconv.Atoi(teamTwoScoreString)

            eventName, _ := s.Find("img.event-logo").Attr("title")
            eventLogo, _ := s.Find("img.event-logo").Attr("src")
            maps := s.Find("div.map-text").Text()
            idString, _ := s.Parent().Parent().Parent().Parent().Attr("href")
            id := ExtractMatchId2(idString)

            Teams := []ResultTeam{
                {teamOneName, teamOneLogo, teamOneScore},
                {teamTwoName, teamTwoLogo, teamTwoScore},
            }

            Results = append(Results, Result{
                Event{
                    eventName,
                    eventLogo,
                },
                maps,
                date,
                Teams,
                id,
            })
        })
    })

    c.OnError(func(r *colly.Response, err error) {
        log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
    })

    c.Visit("https://www.hltv.org/results")

    json.NewEncoder(w).Encode(Results)
}