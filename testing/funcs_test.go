package testing

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		name string
		input string
		expectOutput string
	}{
		{
			name: "empty string",
			input: "",
			expectOutput: "",
		},
		{
			name: "single character",
			input: "a",
			expectOutput: "a",
		},
		{
			name: "two characters",
			input: "ab",
			expectOutput: "ba",
		},
		{
			name: "palindrome",
			input: "stressed",
			expectOutput: "desserts",
		},
		{
			name: "high unicode",
			input: "Hello, 世界",
			expectOutput: "界世 ,olleH",
		},
	}

	for _, c := range cases {
		if output := Reverse(c.input); output != c.expectOutput {
			t.Errorf("%s: got %s but expected %s", c.name, output, c.expectOutput)
		}

	}
}

func TestGetGreeting(t *testing.T) {
	cases := []struct{
		name string
		input string
		expectOutput string
	}{
		{
			name: "empty input",
			input: "",
			expectOutput: "Hello, World!",
		},
		{
			name: "valid input",
			input: "fan",
			expectOutput: "Hello, fan!",
		},
	}

	for _, c := range cases {
		if output := GetGreeting(c.input); output != c.expectOutput {
			t.Errorf("%s: got %s but expected %s", c.name, output, c.expectOutput)
		}
	}
}

func TestParseSize(t *testing.T) {
	cases := []struct{
		name string
		input string
		expectOutput *Size
	}{
		{
			name: "empty string",
			input: "",
			expectOutput: &Size{},
		},
	}

	for _, c := range cases {
		if output := ParseSize(c.input); output.Height != c.expectOutput.Height || output.Width != c.expectOutput.Width {
			t.Errorf("%s: got %v but expected %v", c.name, output, c.expectOutput)
		}
	}
}

func TestLateDaysConsume(t *testing.T) {
	ld := NewLateDays()
	for i := 3; i > -10; i-- {
		expectedLateDays := i
		if expectedLateDays < 0 {
			expectedLateDays = 0
		}
		if output := ld.Consume("test"); output != expectedLateDays {
			t.Errorf("iteration %d: got %d but expected %d", i, output, expectedLateDays)
		}
	}
}
