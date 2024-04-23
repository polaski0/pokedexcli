package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hEllO WoRlD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "fOo BAR",
			expected: []string{"foo", "bar"},
		},
	}

	for _, c := range cases {
		val := cleanInput(c.input)

		for i, v := range val {
			if c.expected[i] != v {
				t.Errorf("expected: %v \nactual: %v", c.expected[i], v)
			}
		}
	}
}
