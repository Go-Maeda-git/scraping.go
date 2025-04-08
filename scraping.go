package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	targetURL := "https://example.com"

	res, err := http.Get(targetURL)
	if err != nil {
		log.Fatalf("HTTP GET request failed: %v", err)
	}
	defer res.Body.Close()
	fmt.Println("HTTP Status Code:", res.StatusCode)

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}
	fmt.Println("HTML Parsed Successfully")

	title := doc.Find("title").Text()
	fmt.Printf("Title: %s\n", title)

	count := 0
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			fmt.Printf("Link %d: %s\n", i+1, link)
			count++
		}
	})
	fmt.Printf("Found %d links\n", count)
}
