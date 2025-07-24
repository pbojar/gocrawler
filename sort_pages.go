package main

import "sort"

type page struct {
	name  string
	count int
}

func sortPages(pages map[string]int) []page {
	// Sorts keys of pages map alphabetically and in descending order
	// by page count.
	pagesSlice := make([]page, 0, len(pages))
	for pageName, pageCount := range pages {
		pagesSlice = append(pagesSlice, page{pageName, pageCount})
	}
	sort.Slice(pagesSlice, func(i, j int) bool {
		if pagesSlice[i].count == pagesSlice[j].count {
			return pagesSlice[i].name < pagesSlice[j].name
		}
		return pagesSlice[i].count > pagesSlice[j].count
	})
	return pagesSlice
}
