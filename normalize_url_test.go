package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	cases := map[string]struct {
		input string
		want  string
		err   string
	}{
		"empty":                         {input: "", want: ""},
		"invalid url":                   {input: ":\\invaldURL", want: "", err: "failed to parse URL"},
		"capitalize url":                {input: "https://BLOG.boot.dev/PATH", want: "blog.boot.dev/path"},
		"https trailing slash":          {input: "https://blog.boot.dev/path/", want: "blog.boot.dev/path"},
		"https no trailing slash":       {input: "https://blog.boot.dev/path", want: "blog.boot.dev/path"},
		"http trailing slash":           {input: "http://blog.boot.dev/path/", want: "blog.boot.dev/path"},
		"http no trailing slash":        {input: "http://blog.boot.dev/path", want: "blog.boot.dev/path"},
		"no protocol trailing slash":    {input: "blog.boot.dev/path/", want: "blog.boot.dev/path"},
		"no protocol no trailing slash": {input: "blog.boot.dev/path", want: "blog.boot.dev/path"},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := normalizeURL(c.input)
			if err != nil && !strings.Contains(err.Error(), c.err) {
				t.Errorf("Unexpected error: %s", err)
				return
			}
			if got != c.want {
				t.Errorf("%s != %s", got, c.want)
				return
			}
		})
	}
}
