package main

import "fmt"

func (cfg *config) crawlPage(rawCurrentURL string) {
	// Increment concurrencyControl each time crawlPage is called
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	// Return pages if rawCurrentURL does not have
	// same domain as rawBaseURL.
	isSameDomain, err := isSameDomain(cfg.baseURL.String(), rawCurrentURL)
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
	if isFirst := cfg.addPageVisit(normCurrentURL); !isFirst {
		return
	}

	// Log current URL
	fmt.Printf("Now crawling thru: %s...\n", rawCurrentURL)

	// Get all internal links in rawCurrentURL
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %s\n", err)
	}
	rawInternalURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("Error - getURLSFromHTML: %s\n", err)
	}

	// Recursively crawl all the internal links
	for _, rawNextURL := range rawInternalURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(rawNextURL)
	}
}
