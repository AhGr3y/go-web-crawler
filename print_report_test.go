package main

import (
	"reflect"
	"testing"
)

func TestSortReport(t *testing.T) {
	cases := map[string]struct {
		input []PageMetric
		want  []PageMetric
	}{
		"empty": {input: []PageMetric{}, want: []PageMetric{}},
		"one entry": {
			input: []PageMetric{
				{URL: "https://wagslane.dev", count: 1},
			},
			want: []PageMetric{
				{URL: "https://wagslane.dev", count: 1},
			},
		},
		"many entries": {
			input: []PageMetric{
				{URL: "https://wagslane.dev", count: 1},
				{URL: "https://wagslane.dev/a", count: 3},
				{URL: "https://wagslane.dev/b", count: 2},
			},
			want: []PageMetric{
				{URL: "https://wagslane.dev/a", count: 3},
				{URL: "https://wagslane.dev/b", count: 2},
				{URL: "https://wagslane.dev", count: 1},
			},
		},
		"many entries same count": {
			input: []PageMetric{
				{URL: "https://wagslane.dev", count: 1},
				{URL: "https://wagslane.dev/c", count: 3},
				{URL: "https://wagslane.dev/a", count: 3},
				{URL: "https://wagslane.dev/b", count: 3},
			},
			want: []PageMetric{
				{URL: "https://wagslane.dev/a", count: 3},
				{URL: "https://wagslane.dev/b", count: 3},
				{URL: "https://wagslane.dev/c", count: 3},
				{URL: "https://wagslane.dev", count: 1},
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got := sortPageMetrics(c.input)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("FAIL: expected: %v, actual: %v", c.want, got)
				return
			}
		})
	}
}
