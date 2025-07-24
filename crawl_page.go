package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	if cfg.overMax() {
		return
	}

	parsedCurr, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse current URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	if cfg.baseURL.Hostname() != parsedCurr.Hostname() {
		return
	}

	normCurr, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't normalize current URL: %v\n", err)
		return
	}
	isFirst := cfg.addPageVisit(normCurr)
	if !isFirst {
		return
	}

	fmt.Printf("Attempting to get HTML from '%s'...\n", rawCurrentURL)
	currHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't get HTML: %v\n", err)
		return
	}
	currURLs, err := getURLsFromHTML(currHTML, cfg.baseURL)
	if err != nil {
		fmt.Printf("couldn't get URLs from HTML: %v\n", err)
		return
	}
	for _, currURL := range currURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(currURL)
	}
}
