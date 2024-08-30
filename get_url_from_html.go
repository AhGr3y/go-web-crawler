package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	// Convert rawBaseURL into a URL Object
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, fmt.Errorf("unable to parse url: %s", err)
	}

	// Parse HTML string into a Node Tree
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("unable to parse html: %s", err)
	}

	// Recursively traverse Node Tree
	urls := []string{}
	var extractInternalLinks func(*html.Node)
	extractInternalLinks = func(n *html.Node) {
		// Look for anchor element
		if n.Type == html.ElementNode && n.Data == "a" {
			attributes := n.Attr
			// Loop thru anchor element's attributes
			for _, attr := range attributes {
				// Add href's value to 'urls' slice
				if attr.Key == "href" {
					urlObj, err := url.Parse(attr.Val)
					if err != nil {
						log.Printf("unable to parse url: %s", err)
					}
					resolvedURL := baseURL.ResolveReference(urlObj)
					urls = append(urls, resolvedURL.String())
				}
			}
		}
		// Recursively traverse Node Tree to extract anchor element's
		// href value.
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			extractInternalLinks(child)
		}
	}
	extractInternalLinks(doc)

	return urls, nil
}
