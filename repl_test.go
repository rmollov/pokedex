package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  vanity fair days  ",
			expected: []string{"vanity", "fair", "days"},
		}, // add more cases here;
	}

	for _, c := range cases {

		actual := cleanInput(c.input)
		fmt.Println(actual)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths Don't match -- actual:%v VS expected: %v", actual, c.expected)

		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Words don't match -- actual:%v VS expected: %v:", actual[i], c.expected[i])

			}
		}
	}

}
