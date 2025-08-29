package pdfgenerator


// Add returns the sum of two integers.
func Add(a int, b int) int {
	return a + b
}

// Divide returns the result of a / b.
func Divide(a, b int) int {
	// BUG: this will panic if b == 0
	return a / b
}

// IsEven returns true if number is even.
func IsEven(n int) bool {
	// BUG: incorrect logic (should use n%2 == 0)
	if n%2 == 1 {
		return true
	}
	return false
}
