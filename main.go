package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	// The program should accept 1 command-line argument,
	// not including the first command-line argument, which
	// is the program itself.
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	rawBaseURLObj, err := url.Parse(args[0])
	if err != nil {
		fmt.Printf("unable to parse url: %s", err)
		return
	}

	// Set up configuration
	const maxConcurrency = 10
	cfg := config{
		pages:              map[string]int{},
		baseURL:            rawBaseURLObj,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}
	fmt.Printf("Starting crawl of: %s...\n", cfg.baseURL.String())

	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL.String())
	cfg.wg.Wait()

	for page, count := range cfg.pages {
		fmt.Printf("%s: %d\n", page, count)
	}
}
