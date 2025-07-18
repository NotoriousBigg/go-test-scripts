package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://naughtyafrica.net/"
	fmt.Printf("%s\n", url)
	scraper(url)
}

func scraper(url string) {
	cl := http.Client{Timeout: 30 * time.Second}

	resp, err := cl.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".container").Each(func(i int, selection *goquery.Selection) {
		header := selection.Find("h1").Text()
		fmt.Println(header)

	})
}
