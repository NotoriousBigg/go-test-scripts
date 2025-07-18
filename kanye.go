package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Quote struct {
	Quote string `json:"quote"`
}

func main() {
	url := "https://api.kanye.rest/"
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	if res.StatusCode != 200 {
		log.Fatalf("response code: %s", res.Status)
		return
	}
	defer res.Body.Close()

	var kanyeQuote Quote
	if err := json.NewDecoder(res.Body).Decode(&kanyeQuote); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("QUOTE: %s\n", kanyeQuote.Quote)
}
