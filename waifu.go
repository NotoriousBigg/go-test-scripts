package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type WaifuResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Creator string `json:"creator"`
	Result  string `json:"result"`
}

func main() {
	fmt.Println("Fetching a random Waifu")
	apiUrl := "https://abhi-api.vercel.app/api/anime/waifu"
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
		return
	}
	if res.StatusCode != 200 {
		log.Fatalf("Response code: %s", res.StatusCode)
	}

	defer res.Body.Close()

	var waifuResponse WaifuResponse
	if err := json.NewDecoder(res.Body).Decode(&waifuResponse); err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	fmt.Printf("Waifu URL: %s", waifuResponse.Result)

	downloadWaifu(waifuResponse.Result, "waifu.jpg")

}

func downloadWaifu(url string, filename string) {
	client := http.Client{Timeout: 30 * time.Second}
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	io.Copy(file, resp.Body)
}
