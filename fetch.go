package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

const (
	target = "#008A2B"
)

func fetch() bool {
	url := "https://www.keishicho-gto.metro.tokyo.lg.jp/keishicho-u/reserve/offerList_detail?tempSeq=363"

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)

	}
	// Set the User-Agent header to mimic a browser
	headers := map[string]string{}
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	// Send the request
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// Check the response status code
	if res.StatusCode != http.StatusOK {
		panic("Failed to fetch the page: " + res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	found := false
	doc.Find("#height_auto_29の国･地域以外の方で、住民票のない方").Each(func(i int, s *goquery.Selection) {
		// Find the text of the element
		v := s.Text()
		// Print the text
		if strings.Contains(v, target) {
			found = true
		}
	})

	return found
}
