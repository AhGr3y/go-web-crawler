package main

import (
	"fmt"
	"net/url"
)

func isSameDomain(rawBaseURL, rawCurrentURL string) (bool, error) {
	rawBaseURLObj, err := url.Parse(rawBaseURL)
	if err != nil {
		return false, fmt.Errorf("unable to parse URL: %s", err)
	}

	rawCurrentURLObj, err := url.Parse(rawCurrentURL)
	if err != nil {
		return false, fmt.Errorf("unable to parse URL: %s", err)
	}

	return rawBaseURLObj.Hostname() == rawCurrentURLObj.Hostname(), nil
}
