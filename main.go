package main

import "fmt"

func main() {
	fmt.Println("Swag overload")
	_, snippets, err := Scrape("https://stackoverflow.com/questions/15112125/how-to-test-multiple-variables-against-a-value")
	if err != nil {
		panic(err)
	}

	for _, snippet := range snippets[0:20] {
		EvalSnippet(&snippet)
		fmt.Println("Result: ", snippet.IsValidPython, snippet.Output, snippet.Error)
	}
}
