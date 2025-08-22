package main

import "testing"

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
			input:    "i need pokemon.    NOW!! ",
			expected: []string{"i", "need", "pokemon.", "NOW!!"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "justoneword",
			expected: []string{"justoneword"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Testing error: lengths do not match")
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Testing error: words do not match: %v and %v", word, expectedWord)
				t.Fail()
			}
		}
	}
}
