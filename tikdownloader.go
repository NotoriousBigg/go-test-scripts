package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type TikTokResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Creator string `json:"creator"`
	Result  struct {
		Title       string `json:"title"`
		Author      string `json:"author"`
		NoWaterMark string `json:"nowm"`
		Watermark   string `json:"watermark"`
		Audio       string `json:"audio"`
		Thumbnail   string `json:"thumbnail"`
	} `json:"result"`
}

func main() {
	fmt.Print("Enter your Tiktok Video Link: ")

	reader := bufio.NewReader(os.Stdin)
	rawInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}
	rawInput = strings.TrimSpace(rawInput)
	fullApiUrl := fmt.Sprintf("https://abhi-api.vercel.app/api/download/tiktok?url=%s", rawInput)

	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(fullApiUrl)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	var NewTikTok TikTokResponse
	if err := json.NewDecoder(resp.Body).Decode(&NewTikTok); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Video Title: %s\n\n", NewTikTok.Result.Title)
	fmt.Print("What would you like to do?\n")
	fmt.Print("1. Download video with watermark\n")
	fmt.Print("2. Download video without watermark\n")
	fmt.Print("3. Download Audio\n")
	fmt.Print("4. Download video thumbnail\n")
	fmt.Print("5. Exit Program\n")

	fmt.Print("Enter your Selection: ")
	choice, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("Please wait as I download your video")
		downloader(NewTikTok.Result.Watermark, "watermarked.mp4")
	case "2":
		fmt.Println("Please wait as I download your video without watermark")
		downloader(NewTikTok.Result.NoWaterMark, "no_watermark.mp4")
	case "3":
		fmt.Println("Please wait as I download audio")
		downloader(NewTikTok.Result.Audio, "audio.mp3")
	case "4":
		fmt.Println("Downloading your thumbnail")
		downloader(NewTikTok.Result.Thumbnail, "thumb.jpg")
	case "5":
		fmt.Println("Thanks for using this program")
		break
	default:
		fmt.Println("What you entered is out of my capabilities.")

	}

}

func downloader(url string, filename string) {
	cl := http.Client{Timeout: 10 * time.Second}
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	resp, err := cl.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(file, resp.Body)
}
