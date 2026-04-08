package repl

import "testing"

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Hello World",
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Pokemon names",
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			name:     "Empty",
			input:    "",
			expected: []string{},
		},
	}

	for _, tc := range tests {
		actual := CleanInut(tc.input)

		if len(actual) != len(tc.expected) {
			t.Errorf("Actual length: %d is not equal to expected length: %d", len(actual), len(tc.expected))
			return
		}

		for i := range tc.expected {
			word := actual[i]
			expectedWord := tc.expected[i]

			if word != expectedWord {
				t.Errorf("Test: %s. Expected: %s to be: %s", tc.name, word, expectedWord)
				return
			}
		}
	}
}
