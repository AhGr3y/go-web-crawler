package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Return an error if status code > 399
	if resp.StatusCode > 399 {
		return "", fmt.Errorf("response failed with status code: %d", resp.StatusCode)
	}

	// Return an error if response is not text/html
	respContentType := resp.Header.Get("Content-Type")
	if !strings.Contains(respContentType, "text/html") {
		return "", fmt.Errorf("got non-HTML content type: %s", respContentType)
	}

	htmlBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %s", err)
	}

	return string(htmlBody), nil
}
