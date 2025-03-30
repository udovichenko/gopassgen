package main

import (
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name      string
		length    int
		minNums   int
		maxNums   int
		minSpec   int
		maxSpec   int
		wantCheck func(string) bool
	}{
		{
			name:    "Basic password generation",
			length:  20,
			minNums: 1,
			maxNums: 5,
			minSpec: 1,
			maxSpec: 5,
			wantCheck: func(password string) bool {
				return len(password) == 20 &&
					containsMinChars(password, numbers, 1) &&
					containsMinChars(password, specChars, 1)
			},
		},
		{
			name:    "Minimum length password",
			length:  8,
			minNums: 2,
			maxNums: 2,
			minSpec: 2,
			maxSpec: 2,
			wantCheck: func(password string) bool {
				return len(password) == 8 &&
					containsMinChars(password, numbers, 2) &&
					containsMinChars(password, specChars, 2)
			},
		},
		{
			name:    "Numbers only password",
			length:  10,
			minNums: 10,
			maxNums: 10,
			minSpec: 0,
			maxSpec: 0,
			wantCheck: func(password string) bool {
				return len(password) == 10 &&
					containsOnlyChars(password, numbers)
			},
		},
		{
			name:    "Special characters only password",
			length:  10,
			minNums: 0,
			maxNums: 0,
			minSpec: 10,
			maxSpec: 10,
			wantCheck: func(password string) bool {
				return len(password) == 10 &&
					containsOnlyChars(password, specChars)
			},
		},
		{
			name:    "Letters only password",
			length:  10,
			minNums: 0,
			maxNums: 0,
			minSpec: 0,
			maxSpec: 0,
			wantCheck: func(password string) bool {
				return len(password) == 10 &&
					containsOnlyChars(password, lowercase+uppercase)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password := generatePassword(tt.length, tt.minNums, tt.maxNums, tt.minSpec, tt.maxSpec)
			if !tt.wantCheck(password) {
				t.Errorf("generatePassword() = %v, did not match expected characteristics", password)
			}
		})
	}
}

func TestMinFunction(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{5, 10, 5},
		{10, 5, 5},
		{0, 5, 0},
		{-5, 5, -5},
		{5, -5, -5},
		{0, 0, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := min(tt.a, tt.b); got != tt.want {
				t.Errorf("min(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDefaultValuesValidity(t *testing.T) {
	// Test that min values don't exceed max values
	if DefaultMinNumeric > DefaultMaxNumeric {
		t.Errorf("Default min number of numeric characters (%d) exceeds max (%d)",
			DefaultMinNumeric, DefaultMaxNumeric)
	}

	if DefaultMinSpecialChars > DefaultMaxSpecialChars {
		t.Errorf("Default min number of special characters (%d) exceeds max (%d)",
			DefaultMinSpecialChars, DefaultMaxSpecialChars)
	}

	// Test that sum of minimums doesn't exceed total length
	if DefaultMinNumeric+DefaultMinSpecialChars > DefaultPasswordLength {
		t.Errorf("Sum of minimum requirements (%d + %d = %d) exceeds total password length (%d)",
			DefaultMinNumeric, DefaultMinSpecialChars,
			DefaultMinNumeric+DefaultMinSpecialChars, DefaultPasswordLength)
	}
}

// Helper functions for testing
func containsMinChars(s, charSet string, minCount int) bool {
	count := 0
	for _, c := range s {
		if strings.ContainsRune(charSet, c) {
			count++
		}
	}
	return count >= minCount
}

func containsOnlyChars(s, charSet string) bool {
	for _, c := range s {
		if !strings.ContainsRune(charSet, c) {
			return false
		}
	}
	return true
}
