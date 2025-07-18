package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
	"time"
)

type AptoideResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Creator string `json:"creator"`
	Result  struct {
		Name        string `json:"name"`
		Icon        string `json:"icon"`
		Developer   string `json:"developer"`
		Size        string `json:"size"`
		Version     string `json:"version"`
		Package     string `json:"package"`
		Downloads   int    `json:"downloads"`
		UpdatedOn   string `json:"updatedOn"`
		DownloadUrl string `json:"downloadUrl"`
	} `json:"result"`
}

func main() {
	fmt.Print("Please enter your search query: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}
	input = strings.TrimSpace(input)
	cleanInput := url2.QueryEscape(input)

	url := fmt.Sprintf("https://abhi-api.vercel.app/api/search/aptoide?text=%s", cleanInput)

	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()

	var APKForDownload AptoideResponse
	if err := json.NewDecoder(res.Body).Decode(&APKForDownload); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Search Successful")
	fmt.Printf("STATUS CODE: %d\n", APKForDownload.Code)
	fmt.Printf("SUCCESS: %t\n", APKForDownload.Status)
	fmt.Println("APK INFO")
	apkInfo := fmt.Sprintf(
		"Name: %s\nDeveloper: %s\nSize: %s\nVersion: %s\nPackage: %s\nTotal Downloads:%d\nLast Update: %s\nDownload Link: %s\n",
		APKForDownload.Result.Name, APKForDownload.Result.Developer, APKForDownload.Result.Size, APKForDownload.Result.Version, APKForDownload.Result.Package, APKForDownload.Result.Downloads, APKForDownload.Result.UpdatedOn, APKForDownload.Result.DownloadUrl,
	)
	fmt.Printf(apkInfo)

	fmt.Print("What would you like to do next?\n Use 1 to download app Icon, and 2 to download the apk\n")
	choice, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}
	choice = strings.TrimSpace(choice)
	switch choice {
	case "1":
		downloadApk(APKForDownload.Result.Icon, "image")
	case "2":
		downloadApk(APKForDownload.Result.DownloadUrl, "apk")

	}

}

func downloadApk(url string, fileType string) {
	var outputFile string
	if fileType == "apk" {
		outputFile = "output.apk"
	} else if fileType == "image" {
		outputFile = "output.jpg"
	}
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	cl := http.Client{Timeout: 10 * time.Second}
	resp, err := cl.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	io.Copy(file, resp.Body)
}
