// Declaring main package
package main
// Importing fmt package
import "fmt"

// Declaring typed constant
const number1Constat = 7
// Declaring untyped constant
const number2Constat = 3
// Declaring constant with float value
const PI = 3.14
// Declaring constant with multiple values
const (
	A = 1
	B = 3.14
	C = "Siuuuuu!"
)

// The number function
func fnNumber() {
	// // Unchanged constant value
	// const unchangable = 10
	// unchangable = 20
	// // Printing the constant value
	// fmt.Println("This is the constant value is: ", unchangable)

	// Printing the constant with multiple values
	fmt.Println("This is the constant values is: ", A)
	fmt.Println("This is the constant values is: ", B)
	fmt.Println("This is the constant values is: ", C)

	// Printing the constant values
	fmt.Println("This is the constant values is: ", number1Constat)
	fmt.Println("This is the constant values is: ", number2Constat)
	// Printing the addition of both constant values
	fmt.Println("The addition of both constant values is: ", number1Constat + number2Constat)
	// Printing the constant value of PI
	fmt.Println("The constant value of PI is: ", PI)
}

// Main function
func main() {
	fnNumber()
}