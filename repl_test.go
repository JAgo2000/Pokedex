package main

import (
	"testing"
)

func test_clean_scantext(t *testing.T) {
	cases := []struct {
		input    string
		expectet []string
	}{
		{
			input: "Hello World",
			expectet: []string{
				"hello",
				"world",
			},
		},
	}
	for _, cs := range cases {
		actual := clean_scantext(cs.input)
		if len(actual) != len(cs.expectet) {
			t.Errorf("The lenghts are not equal %v vs %v", len(actual), len(cs.expectet))
			continue
		}
		for i := range actual {
			actualword := actual[i]
			expectetword := cs.expectet[i]
			if actualword != expectetword {
				t.Errorf("The words are not equal %v vs %v", actualword, expectetword)
			}
		}
	}
}
