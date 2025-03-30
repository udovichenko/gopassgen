package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	specChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

// Default configuration values for password generation.
var (
	DefaultPasswordLength  = 20
	DefaultMinNumeric      = 1
	DefaultMaxNumeric      = 6
	DefaultMinSpecialChars = 1
	DefaultMaxSpecialChars = 6
)

func main() {
	// Define command line flags
	length := flag.Int("length", DefaultPasswordLength, "Total password length")
	minNums := flag.Int("min-nums", DefaultMinNumeric, "Minimum number of numeric characters")
	maxNums := flag.Int("max-nums", DefaultMaxNumeric, "Maximum number of numeric characters")
	minSpec := flag.Int("min-spec", DefaultMinSpecialChars, "Minimum number of special characters")
	maxSpec := flag.Int("max-spec", DefaultMaxSpecialChars, "Maximum number of special characters")
	help := flag.Bool("help", false, "Display help information")

	// Parse command line arguments
	flag.Parse()

	// Show help information if requested
	if *help {
		printHelp()
		os.Exit(0)
	}

	// Validate inputs
	if *length < 1 {
		fmt.Println("Error: Password length must be at least 1")
		os.Exit(1)
	}

	if *minNums < 0 || *maxNums < 0 || *minSpec < 0 || *maxSpec < 0 {
		fmt.Println("Error: Minimum and maximum values cannot be negative")
		os.Exit(1)
	}

	if *minNums > *maxNums {
		fmt.Println("Error: Minimum number of numeric characters cannot exceed maximum")
		os.Exit(1)
	}

	if *minSpec > *maxSpec {
		fmt.Println("Error: Minimum number of special characters cannot exceed maximum")
		os.Exit(1)
	}

	if *minNums+*minSpec > *length {
		fmt.Println("Error: Sum of minimum requirements exceeds total password length")
		os.Exit(1)
	}

	if *maxNums+*maxSpec > *length {
		fmt.Printf("Warning: Sum of maximum allowed characters (%d) exceeds password length (%d)\n", *maxNums+*maxSpec, *length)
		fmt.Println("Adjusting maximum values to fit password length")

		// Adjust maximums proportionally
		ratio := float64(*length) / float64(*maxNums+*maxSpec)
		*maxNums = int(float64(*maxNums) * ratio)
		*maxSpec = *length - *maxNums
	}

	// Generate password
	password := generatePassword(*length, *minNums, *maxNums, *minSpec, *maxSpec)
	fmt.Println(password)
}

func generatePassword(length, minNums, maxNums, minSpec, maxSpec int) string {
	// Create a local random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Determine actual counts to use
	numCount := minNums
	if maxNums > minNums {
		numCount += r.Intn(maxNums - minNums + 1)
	}

	specCount := minSpec
	if maxSpec > minSpec {
		specCount += r.Intn(maxSpec - minSpec + 1)
	}

	// Ensure we don't exceed total length
	if numCount+specCount > length {
		excess := numCount + specCount - length
		// Reduce counts proportionally
		if numCount > minNums && excess > 0 {
			reduction := min(numCount-minNums, excess)
			numCount -= reduction
			excess -= reduction
		}
		if specCount > minSpec && excess > 0 {
			reduction := min(specCount-minSpec, excess)
			specCount -= reduction
		}
	}

	// Calculate letter count
	letterCount := length - numCount - specCount

	// Create slices for our character sets
	var passwordChars []string

	// Add the specified number of each character type
	for i := 0; i < letterCount; i++ {
		if r.Intn(2) == 0 {
			passwordChars = append(passwordChars, string(lowercase[r.Intn(len(lowercase))]))
		} else {
			passwordChars = append(passwordChars, string(uppercase[r.Intn(len(uppercase))]))
		}
	}

	for i := 0; i < numCount; i++ {
		passwordChars = append(passwordChars, string(numbers[r.Intn(len(numbers))]))
	}

	for i := 0; i < specCount; i++ {
		passwordChars = append(passwordChars, string(specChars[r.Intn(len(specChars))]))
	}

	// Shuffle the characters
	r.Shuffle(len(passwordChars), func(i, j int) {
		passwordChars[i], passwordChars[j] = passwordChars[j], passwordChars[i]
	})

	return strings.Join(passwordChars, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printHelp() {
	fmt.Println("Password Generator - Create secure random passwords")
	fmt.Println("\nUsage:")
	fmt.Println("  password-generator [options]")
	fmt.Println("\nOptions:")
	fmt.Printf("  -length int     Total password length (default %d)\n", DefaultPasswordLength)
	fmt.Printf("  -min-nums int   Minimum number of numeric characters (default %d)\n", DefaultMinNumeric)
	fmt.Printf("  -max-nums int   Maximum number of numeric characters (default %d)\n", DefaultMaxNumeric)
	fmt.Printf("  -min-spec int   Minimum number of special characters (default %d)\n", DefaultMinSpecialChars)
	fmt.Printf("  -max-spec int   Maximum number of special characters (default %d)\n", DefaultMaxSpecialChars)
	fmt.Println("  -help           Display this help information")
	fmt.Println("\nExample:")
	fmt.Println("  password-generator -length 16 -min-nums 2 -max-nums 4 -min-spec 2 -max-spec 3")
}
