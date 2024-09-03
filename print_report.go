package main

import "fmt"

func printReport(pages map[string]int, rawBaseURL string) {
	pageMetrics := []PageMetric{}
	for page, count := range pages {
		pageMetrics = append(pageMetrics, PageMetric{
			URL:   "https://" + page,
			count: count,
		})
	}

	sortedPageMetrics := sortPageMetrics(pageMetrics)
	fmt.Println("=================================")
	fmt.Printf("REPORT for %s\n", rawBaseURL)
	fmt.Println("=================================")

	for _, pageMetric := range sortedPageMetrics {
		fmt.Printf("Found %d internal links to %s\n", pageMetric.count, pageMetric.URL)
	}
}
