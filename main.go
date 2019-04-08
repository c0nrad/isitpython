package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	go RunScraper()
	go RunWebServer()
	select {}
}

func RunScraper() {
	futurelinks := []string{BASE_URL}

	for len(futurelinks) > 0 {
		link := futurelinks[0]
		futurelinks = futurelinks[1:]

		log.Println("Scraping ", link)
		links, snippets, err := Scrape(link)
		if err != nil {
			panic(err)
		}
		log.Println("Found", len(links), "links and", len(snippets), "snippets")

		futurelinks = append(futurelinks, links...)

		for _, snippet := range snippets {
			snippet.URL = link
			EvalSnippet(&snippet)
			SaveSnippet(snippet)
			fmt.Println("IsItPython?: ", snippet.IsValidPython, snippet.Error)

			time.Sleep(time.Second * 5)
		}

		time.Sleep(time.Second * 0)
	}
}
