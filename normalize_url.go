package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(urlString string) (string, error) {
	urlObj, err := url.Parse(urlString)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %s", err)
	}

	hostname := urlObj.Hostname()
	path := urlObj.EscapedPath()

	fullPath := hostname + path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}
