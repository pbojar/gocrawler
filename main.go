package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawURL := args[1]
	fmt.Printf("starting crawl of: %s\n", rawURL)
	html, err := getHTML(rawURL)
	if err != nil {
		fmt.Printf("couldn't get html: %v\n", err)
	}
	fmt.Print(html)
}
