package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(urlIn string) (string, error) {

	// Parse URL
	urlStruct, err := url.Parse(urlIn)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %v", err)
	}

	urlOut := strings.ToLower(urlStruct.Host + urlStruct.Path)
	urlOut = strings.TrimSuffix(urlOut, "/")

	return urlOut, nil
}
