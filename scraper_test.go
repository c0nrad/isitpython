package main

import (
	"testing"
)

func TestScrape(t *testing.T) {
	links, _, err := Scrape(BASE_URL)
	if err != nil {
		t.Error(err)
	}

	if len(links) == 0 {
		t.Error("There should be at least one link")
	}
}

func TestCodeScrape(t *testing.T) {
	links, snippets, err := Scrape("https://stackoverflow.com/questions/15112125/how-to-test-multiple-variables-against-a-value")
	if err != nil {
		t.Error(err)
	}

	if len(links) == 0 {
		t.Error("There should be at least one link")
	}

	if len(snippets) == 0 {
		t.Error("There should be more than one snippet")
	}
}
