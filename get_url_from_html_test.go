package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	rawBaseURL := "https://blog.boot.dev"
	cases := map[string]struct {
		input string
		want  []string
		err   string
	}{
		"empty": {input: "", want: []string{}},
		"single absolute path": {
			input: `<html>
    <body>
        <a href="https://blog.boot.dev"><span>Go to Boot.dev, you React Andy</span></a>
    </body>
</html>`,
			want: []string{"https://blog.boot.dev"},
		},
		"single relative path": {
			input: `<html>
    <body>
	    <p>Check out this <a href="/education/is-there-a-course-on">article</a></p>
	</body>
</html>`,
			want: []string{"https://blog.boot.dev/education/is-there-a-course-on"},
		},
		"multiple mixed paths": {
			input: `<html>
    <body>
	    <h1>This is a Strage Title</h1>
		<p>This sentence has two links, <a href="https://blog.boot.dev/path/to/content">this</a> and <a href="/relative/path/to/content">that</a>.</p>
	</body>
</html>`,
			want: []string{"https://blog.boot.dev/path/to/content", "https://blog.boot.dev/relative/path/to/content"},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := getURLsFromHTML(c.input, rawBaseURL)
			if err != nil && !strings.Contains(err.Error(), c.err) {
				t.Errorf("FAIL: unexpected error: %v", err)
				return
			}
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("FAIL: expected: %v, actual: %v", c.want, got)
				return
			}
		})
	}
}
