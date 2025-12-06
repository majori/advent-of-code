package aoc

import (
	"reflect"
	"testing"
)

func TestAsRuneGrid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [][]rune
	}{
		{
			name:  "single line",
			input: "abc",
			expected: [][]rune{
				{'a', 'b', 'c'},
			},
		},
		{
			name:  "multiple lines",
			input: "abc\ndef\nghi",
			expected: [][]rune{
				{'a', 'b', 'c'},
				{'d', 'e', 'f'},
				{'g', 'h', 'i'},
			},
		},
		{
			name:  "empty input",
			input: "",
			expected: [][]rune{
				{},
			},
		},
		{
			name:  "input with spaces",
			input: "a b c\n d e f \ng h i",
			expected: [][]rune{
				{'a', ' ', 'b', ' ', 'c'},
				{' ', 'd', ' ', 'e', ' ', 'f', ' '},
				{'g', ' ', 'h', ' ', 'i'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aoc := AoC{input: tt.input}
			result := aoc.AsRuneGrid()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
