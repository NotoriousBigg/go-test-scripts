package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
	"time"
)

type Response struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Creator string `json:"creator"`
	Result  struct {
		SearchResults []string `json:"searchResults"`
	} `json:"result"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your Search Query: ")
	query, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}
	query = strings.TrimSpace(query)
	escapedQuery := url2.QueryEscape(query)

	url := fmt.Sprintf("https://abhi-api.vercel.app/api/search/gimage?text=%s", escapedQuery)
	cl := http.Client{Timeout: 10 * time.Second}

	res, err := cl.Get(url)

	if err != nil {
		log.Fatal(err)
		return
	}
	if res.StatusCode != 200 {
		log.Fatalf("STATUS: %s", res.Status)
		return
	}

	defer res.Body.Close()

	var googleSearch Response
	if err := json.NewDecoder(res.Body).Decode(&googleSearch); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Status Code: %d\n", googleSearch.Code)
	fmt.Printf("Success: %t\n\n", googleSearch.Status)
	fmt.Println("Search Results")
	for i, image := range googleSearch.Result.SearchResults {
		fmt.Printf("%d) %s\n", i+1, image)
	}

}
