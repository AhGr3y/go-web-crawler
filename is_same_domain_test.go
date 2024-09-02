package main

import (
	"strings"
	"testing"
)

func TestIsSameDomain(t *testing.T) {
	rawBaseURL := "https://blog.boot.dev"
	cases := map[string]struct {
		input string
		want  bool
		err   string
	}{
		"empty":            {input: "", want: false},
		"same domain":      {input: "https://blog.boot.dev/path", want: true},
		"different domain": {input: "https://example.com/path", want: false},
		"invalid url":      {input: ":\\notavalidurl", want: false, err: "unable to parse URL"},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := isSameDomain(rawBaseURL, c.input)
			if err != nil && !strings.Contains(err.Error(), c.err) {
				t.Errorf("FAIL: unexpected error: %v", err)
				return
			}
			if got != c.want {
				t.Errorf("FAIL: expected: %v, actual: %v", c.want, got)
				return
			}
		})
	}
}
