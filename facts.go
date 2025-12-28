package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Fact struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Creator string `json:"creator"`
	Result  string `json:"result"`
}

func main() {
	apiUrl := "https://abhi-api.vercel.app/api/fun/facts"
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	response, err := client.Get(apiUrl)
	// check for errors
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}
	// check for status code
	if response.StatusCode != 200 {
		log.Fatalf("Status Code: %s", response.StatusCode)
	}

	defer response.Body.Close()

	var thisFact Fact
	if err := json.NewDecoder(response.Body).Decode(&thisFact); err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	fmt.Printf("Fact: %s\n", thisFact.Result)

}
