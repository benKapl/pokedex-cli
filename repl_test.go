package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " , cudo kiddo !",
			expected: []string{",", "cudo", "kiddo", "!"},
		},
		{
			input:    "alpha'12  ",
			expected: []string{"alpha'12"},
		},
		{
			input:    "COUcou tOI",
			expected: []string{"coucou", "toi"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test failed: lenght of input slice doesn't match with expected values")
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Test failed: actual word doesn't match witch expected")
			}

		}

	}
}
