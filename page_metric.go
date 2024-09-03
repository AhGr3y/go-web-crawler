package main

import "sort"

type PageMetric struct {
	URL   string
	count int
}

func sortPageMetrics(pageMetrics []PageMetric) []PageMetric {
	// Sort by count then alphabet
	sort.Slice(pageMetrics, func(i, j int) bool {
		if pageMetrics[i].count == pageMetrics[j].count {
			return pageMetrics[i].URL < pageMetrics[j].URL
		}
		return pageMetrics[i].count > pageMetrics[j].count
	})
	return pageMetrics
}
