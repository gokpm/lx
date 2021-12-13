package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-s":
			synonyms()
			break
		default:
			definition()
			break
		}
	}
	return
}

func scrap(url string, selector string) (elements []string) {
	c := colly.NewCollector()
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		elements = append(elements, strings.TrimSpace(e.Text))
	})
	c.Visit(url)
	return
}

func synonyms() {
	query := strings.Join(os.Args[2:], "+")
	url := fmt.Sprintf("https://www.thesaurus.com/browse/%[1]s", query)
	content := scrap(url, "a.css-1n6g4vv.eh475bn0")
	for _, word := range content {
		fmt.Println(word)
	}
	return
}

func definition() {
	query := strings.Join(os.Args[1:], "+")
	url := fmt.Sprintf("https://gcide.gnu.org.ua/?q=%[1]s&db=gcide&define=1", query)
	content := scrap(url, "pre")
	for _, word := range content {
		fmt.Printf("\n%[1]s\n\n", word)
	}
	return
}
