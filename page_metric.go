package main

import "sort"

type PageMetric struct {
	URL   string
	count int
}

func sortPageMetrics(pageMetrics []PageMetric) []PageMetric {
	// Sort by alphabet
	sort.Slice(pageMetrics, func(i, j int) bool {
		return pageMetrics[i].URL < pageMetrics[j].URL
	})
	// Sort by count
	sort.Slice(pageMetrics, func(i, j int) bool {
		return pageMetrics[i].count > pageMetrics[j].count
	})
	return pageMetrics
}
