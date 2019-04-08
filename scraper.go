package main

import (
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var BASE_URL = "https://stackoverflow.com/questions/tagged/python?sort=frequent&pageSize=50"
var FutureQuestion = regexp.MustCompile(`^/questions/[0-9]+/`)

func Scrape(url string) ([]string, []Snippet, error) {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	request.Header.Set("User-Agent", "IsItPython? v0.1")
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, nil, err
	}

	links := ScrapeLinks(document)
	snippets := ScrapeSnippets(document)
	return links, snippets, nil
}

func ScrapeLinks(d *goquery.Document) []string {

	links := []string{}
	aTags := d.Find("a")

	for i := 0; i < aTags.Length(); i++ {
		element := aTags.Get(i)

		for _, a := range element.Attr {
			if a.Key == "href" && FutureQuestion.Match([]byte(a.Val)) {
				links = append(links, "https://stackoverflow.com/"+a.Val)
			}
		}
	}

	return links
}

func ScrapeSnippets(d *goquery.Document) []Snippet {
	snippets := []Snippet{}

	codeTags := d.Find("code")

	codeTags.Each(func(i int, s *goquery.Selection) {
		snippets = append(snippets, Snippet{Body: s.Text()})
	})

	return snippets
}
