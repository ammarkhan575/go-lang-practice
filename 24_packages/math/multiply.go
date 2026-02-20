package math

// Takes two input, and type can be int, float32, float64
// Returns the multiplication of the two inputs
func Multiply[T int | float32 | float64](a T, b T) T {
	return a * b
}