package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, exist := cfg.pages[normalizedURL]
	if !exist {
		cfg.pages[normalizedURL] = 1
		isFirst = true
		return
	}
	cfg.pages[normalizedURL]++
	isFirst = false
	return
}
