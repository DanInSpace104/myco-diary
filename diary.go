package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

var client http.Client

func init() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: diary <URL> <cookie> <password> <diary hypha> <text to append>")
		os.Exit(1)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Error creating cookie jar: %v", err)
		os.Exit(1)
	}
	client = http.Client{Jar: jar}
}

func main() {
	wikiUrl := os.Args[1]
	cookie := os.Args[2]
	diaryHypha := os.Args[3]
	appendText := strings.Join(os.Args[4:], " ")

	if strings.HasSuffix(wikiUrl, "/") {
		wikiUrl = wikiUrl[:len(wikiUrl)-1]
	}

	_url, err := url.Parse(wikiUrl)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}
	client.Jar.SetCookies(_url, []*http.Cookie{{Name: "mycorrhiza_token", Value: cookie}})

	date := time.Now().Format("2006-01-02")
	textUrl := fmt.Sprintf("%s/text/%s/%s/", wikiUrl, diaryHypha, date)

	resp, err := client.Get(textUrl)
	if err != nil {
		log.Fatalf("Error getting diary hypha: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	currentTime := time.Now().Format("15:04:05")
	newText := fmt.Sprintf("%s\n//%s// %s", body, currentTime, appendText)
	uploadUrl := fmt.Sprintf("%s/upload-text/%s/%s", wikiUrl, diaryHypha, date)

	resp, err = client.PostForm(uploadUrl, url.Values{"text": {newText}})
	if err != nil {
		log.Fatalf("Error uploading diary hypha: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error uploading diary hypha: %v", resp.Status)
	}

	fmt.Println("View the updated diary hypha at:", fmt.Sprintf("%s/hypha/%s/%s", wikiUrl, diaryHypha, date))
	os.Exit(0)
}
