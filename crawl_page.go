package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBase, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("couldn't parse base URL: %v\n", err)
		return
	}
	parsedCurr, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't parse current URL: %v\n", err)
		return
	}
	if parsedBase.Hostname() != parsedCurr.Hostname() {
		return
	}

	normCurr, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't normalize current URL: %v\n", err)
		return
	}
	if _, exists := pages[normCurr]; exists {
		pages[normCurr]++
		return
	} else {
		pages[normCurr] = 1
	}

	fmt.Printf("Attempting to get HTML from '%s'...\n", rawCurrentURL)
	currHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't get HTML: %v\n", err)
		return
	}
	currURLs, err := getURLsFromHTML(currHTML, rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't get URLs from HTML: %v\n", err)
		return
	}
	for _, currURL := range currURLs {
		crawlPage(rawBaseURL, currURL, pages)
	}
}
