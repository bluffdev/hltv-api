package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

func GetResults(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    c := colly.NewCollector()

    c.OnHTML("div.results-sublist", func(h *colly.HTMLElement) {
        date := h.DOM.Children().Eq(0).Text()
        // date := h.DOM.Find("div.standard-headline").Text()
        fmt.Println(date)   
    })

    c.OnError(func(r *colly.Response, err error) {
        log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
    })

    c.Visit("https://www.hltv.org/results")
}