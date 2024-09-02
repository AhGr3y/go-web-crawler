package main

import "fmt"

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Return pages if rawCurrentURL does not have
	// same domain as rawBaseURL.
	isSameDomain, err := isSameDomain(rawBaseURL, rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - isSameDomain: %s\n", err)
	}
	if !isSameDomain {
		return
	}

	// Normalize rawCurrentURL
	normCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizeURL: %s\n", err)
	}

	// Increment pages if rawCurrentURL alraedy visited,
	// or set count to 1 if not visited.
	currentURLCount := pages[normCurrentURL]
	if currentURLCount > 0 {
		pages[normCurrentURL]++
		return
	}
	pages[normCurrentURL] = 1

	fmt.Printf("Now crawling thru: %s...\n", rawCurrentURL)

	// Get all internal links in rawCurrentURL
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %s\n", err)
	}
	rawInternalURLs, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("Error - getURLSFromHTML: %s\n", err)
	}

	// Recursively crawl all the internal links
	for _, rawNextURL := range rawInternalURLs {
		crawlPage(rawBaseURL, rawNextURL, pages)
	}
}
