package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// The program should accept 1 command-line argument,
	// not including the first command-line argument, which
	// is the program itself.
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		return
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error - Atoi: %s", err)
		return
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Error - Atoi: %s", err)
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %s\n", err)
		return
	}

	fmt.Printf("Starting crawl of: %s...\n", cfg.baseURL.String())

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for page, count := range cfg.pages {
		fmt.Printf("%s: %d\n", page, count)
	}

}
