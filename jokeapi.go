package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type JokeStruct struct {
	Id            string `json:"id"`
	Question      string `json:"question"`
	Answer        string `json:"answer"`
	Permalink     string `json:"permalink"`
	PermalinkHtml string `json:"permalink_html"`
}

func main() {
	url := "https://teehee.dev/api/joke"

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	if res.StatusCode != 200 {
		log.Fatalf("RESPONSE CODE: %s", res.Status)
		return
	}
	defer res.Body.Close()

	var joke JokeStruct
	if err := json.NewDecoder(res.Body).Decode(&joke); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("New Joke Received")
	fmt.Printf("Joke ID: %s\n", joke.Id)
	fmt.Printf("Question: %s\n", joke.Question)
	fmt.Printf("Answer: %s\n", joke.Answer)
}
