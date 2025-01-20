package add

import "golang.org/x/exp/constraints"

// Add adds two numbers
//
// It has two parameters: num1 and num2.
// It adds the two parameters and returns the result
//
// More information on addition see [https://www.mathsisfun.com/numbers/addition.html]
func Add[T Number](num1, num2 T) T {
	return num1 + num2
}

// Number is the allowed type for the Add function
// 
// It allows values that implement either the Integer or Float interface from the [golang.org/x/exp/constraints] package
type Number interface {
	constraints.Integer | constraints.Float
}