package main

import (
	"testing"
)

func TestTableDriven(t *testing.T) {
	cases := []struct {
		actual   string
		expected string
	}{
		{
			actual:   "actual",
			expected: "expected",
		},
		{
			actual:   "actual_2",
			expected: "expected_2",
		},
	}

	for _, c := range cases {
		if c.actual != c.expected {
			t.Errorf("actual:%v expected:%v\n", c.actual, c.expected)
		}
	}
}
