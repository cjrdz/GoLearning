// Declaring main package
package main

// Importing the fmt package
import (
	"fmt"
)

// Declaring addition function
func addition() {
	// Declaring variables
	var a int = 10
	var b int = 20
	var sum int = a + b
	// Declaring a new variable and assigning the sum of a and b
	var sumAll = 10 + 20

	// Gradually adding the numbers
	var(

		sum1 = 10 + 20 
		sum2 = sum1 + 30
		sumTotal = sum2 + sum2
	)

	// Printing the sum
	fmt.Println("Arithmetic Operators")
	fmt.Println("***************************************")
	fmt.Println("Addition method")
	fmt.Println("Sum of", + a, "and", b, "=", sum)
	fmt.Println("Sum: ", sum)
	fmt.Println("Sum All: ", sumAll)
	fmt.Println("Sum Total: ", sumTotal)
}

// Declaring subtraction function
func subtraction() {
	// Declaring variables
	var a int = 20
	var b int = 10
	var subtraction int = a - b

	// Printing the difference
	fmt.Println("***************************************")
	fmt.Println("Subtraction method")
	fmt.Println("Difference between", + a, "and", b, "=", subtraction)
}

// Declaring multiplication function
func multiplication() {
	// Declaring variables
	var a int = 10
	var b int = 20
	var multiplication int = a * b

	// Printing the product
	fmt.Println("***************************************")
	fmt.Println("Multiplication method")
	fmt.Println("Product of", + a, "by", b, "=", multiplication)
}

// Declaring division function
func division() {
	// Declaring variables
	var a int = 20
	var b int = 10
	var division int = a / b

	// Printing the quotient
	fmt.Println("***************************************")
	fmt.Println("Division method")
	fmt.Println("Quotient of", + a, "and", b, "=", division)
}

// Declaring modulus function
func modulus() {
	// Declaring variables
	var a int = 20
	var b int = 10
	var modulus int = a % b

	// Printing the remainder
	fmt.Println("***************************************")
	fmt.Println("Modulus method")
	fmt.Println("Remainder of", + a, "and", b, "=", modulus)
}

// Declaring increment function
func increment() {
	// Declaring variables
	var a int = 10
	var increment int = a

	// Incrementing the variable
	increment++

	// Printing the increment
	fmt.Println("***************************************")
	fmt.Println("Increment method")
	fmt.Println("Increment of", + a, "=", increment)
}

// Declaring decrement function
func decrement() {
	// Declaring variables
	var a int = 10
	var decrement int = a

	// Decrementing the variable
	decrement--

	// Printing the decrement
	fmt.Println("***************************************")
	fmt.Println("Decrement method")
	fmt.Println("Decrement of", + a, "=", decrement)
	fmt.Println("***************************************")
}

// Declaring main function
func main() {
	// Calling addition function
	addition()
	// Calling subtraction function
	subtraction()
	// Calling multiplication function
	multiplication()
	// Calling division function
	division()
	// Calling modulus function
	modulus()
	// Calling increment function
	increment()
	// Calling decrement function
	decrement()
}