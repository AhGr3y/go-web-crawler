package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return &config{}, err
	}

	return &config{
		pages:              make(map[string]int),
		maxPages:           maxPages,
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}

func (cfg *config) checkPageLength() (isCapped bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if len(cfg.pages) >= cfg.maxPages {
		isCapped = true
	} else {
		isCapped = false
	}

	return
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
