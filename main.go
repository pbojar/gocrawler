package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 4 {
		fmt.Println("too few arguments provided")
		fmt.Println("usage: ./crawler <\"URL\"> (maxConcurrency) (maxPages)")
		os.Exit(1)
	}
	if len(args) > 4 {
		fmt.Println("too many arguments provided")
		fmt.Println("usage: ./crawler <\"URL\"> (maxConcurrency) (maxPages)")
		os.Exit(1)
	}
	rawURL := args[1]
	maxConcurrency, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Error - strconv.Atoi: %v\n", err)
		return
	}
	maxPages, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Printf("Error - strconv.Atoi: %v\n", err)
		return
	}

	cfg, err := configure(rawURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v\n", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawURL)
	cfg.wg.Wait()

	for page, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, page)
	}
}
