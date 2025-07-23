package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't get response from '%s': %v", rawURL, err)
	}
	defer res.Body.Close()

	// Handle 400+ status code
	if res.StatusCode >= 400 {
		return "", fmt.Errorf("get request to '%s' yielded error code '%d'", rawURL, res.StatusCode)
	}

	// Handle content-type header
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("'%s' has Content-Type '%s' which is not supported", rawURL, contentType)
	}

	// Read body of response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}

	return string(body), nil
}
