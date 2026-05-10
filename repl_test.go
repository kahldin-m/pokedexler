//go:build ignore

package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input      string
		expected   []string
	}{
		{
			input:     "   hello  world    ",
			expected:  []string{"hello", "world"},
		},
		{
			input:     " DREXler and  Dexter    ",
			expected:  []string{"drexler", "and", "dexter"},
		},
		{
			input:     "    ",
			expected:  []string{},
		},
		// add more cases here
	}

	// run testing on all cases
	for _, c := range cases {
		actual := cleanInput(c.input)
		actLength := len(actual)
		expLength := len(c.expected)
		if actLength != expLength {
			t.Errorf("Length of actual: %v does not meet expected: %v", actLength, expLength)
		}
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test failed. actual: %v does not equal expected: %v", word, expectedWord)
			}
		}
	}
}