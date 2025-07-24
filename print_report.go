package main

import (
	"fmt"
	"strings"
)

func printReport(pages map[string]int, baseURL string) {
	lnLen := 15 + len(baseURL)
	fmt.Println(strings.Repeat("=", lnLen))
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println(strings.Repeat("=", lnLen))
	sortedPages := sortPages(pages)
	for _, pg := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", pg.count, pg.name)
	}
}
