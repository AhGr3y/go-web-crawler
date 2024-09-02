package main

import (
	"fmt"
	"os"
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

	rawURL := args[0]
	fmt.Printf("Starting crawl of: %s...\n", rawURL)

	pages := map[string]int{}
	crawlPage(rawURL, rawURL, pages)

	for page, count := range pages {
		fmt.Printf("%s: %d\n", page, count)
	}
}
