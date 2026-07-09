// Package adder provides functionality to add 2 numbers.
package adder

// Number is a constraint that permits any integer or floating-point type.
type Number interface {
	~int | ~float64
}

// Add adds two Numbers .
// It has two parameters: two Numbers. Add returns the result
// of the two numbers added.

//
// [Mathsisfun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
