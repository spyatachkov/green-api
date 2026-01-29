package strutil

import (
	"reflect"
	"testing"
)

func TestTrimSpace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no spaces",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "leading spaces",
			input:    "  hello",
			expected: "hello",
		},
		{
			name:     "trailing spaces",
			input:    "hello  ",
			expected: "hello",
		},
		{
			name:     "both sides spaces",
			input:    "  hello  ",
			expected: "hello",
		},
		{
			name:     "tabs",
			input:    "\t\thello\t\t",
			expected: "hello",
		},
		{
			name:     "mixed spaces and tabs",
			input:    " \t hello world \t ",
			expected: "hello world",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only spaces",
			input:    "   ",
			expected: "",
		},
		{
			name:     "only tabs",
			input:    "\t\t\t",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrimSpace(tt.input)
			if result != tt.expected {
				t.Errorf("TrimSpace(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSplitAndTrim(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sep      string
		expected []string
	}{
		{
			name:     "simple split",
			input:    "a,b,c",
			sep:      ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "split with spaces",
			input:    "a, b, c",
			sep:      ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "split with extra spaces",
			input:    "  a  ,  b  ,  c  ",
			sep:      ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "split with tabs",
			input:    "\ta\t,\tb\t,\tc\t",
			sep:      ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "empty parts",
			input:    "a,,c",
			sep:      ",",
			expected: []string{"a", "c"},
		},
		{
			name:     "single element",
			input:    "hello",
			sep:      ",",
			expected: []string{"hello"},
		},
		{
			name:     "empty string",
			input:    "",
			sep:      ",",
			expected: nil,
		},
		{
			name:     "only separator",
			input:    ",",
			sep:      ",",
			expected: nil,
		},
		{
			name:     "multiple separators",
			input:    ",,,",
			sep:      ",",
			expected: nil,
		},
		{
			name:     "trailing separator",
			input:    "a,b,c,",
			sep:      ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "leading separator",
			input:    ",a,b,c",
			sep:      ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "mixed spaces and empty parts",
			input:    " a , , b ",
			sep:      ",",
			expected: []string{"a", "b"},
		},
		{
			name:     "URLs with commas",
			input:    "http://localhost:3000, http://localhost:8080, https://greenapi.com",
			sep:      ",",
			expected: []string{"http://localhost:3000", "http://localhost:8080", "https://greenapi.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitAndTrim(tt.input, tt.sep)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SplitAndTrim(%q, %q) = %v, want %v", tt.input, tt.sep, result, tt.expected)
			}
		})
	}
}

func BenchmarkTrimSpace(b *testing.B) {
	input := "  hello world  "
	for i := 0; i < b.N; i++ {
		TrimSpace(input)
	}
}

func BenchmarkSplitAndTrim(b *testing.B) {
	input := "http://localhost:3000, http://localhost:8080, https://greenapi.com"
	sep := ","
	for i := 0; i < b.N; i++ {
		SplitAndTrim(input, sep)
	}
}
