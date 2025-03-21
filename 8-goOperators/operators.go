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
	fmt.Println("Sum: ", sum)
	fmt.Println("Sum All: ", sumAll)
	fmt.Println("Sum Total: ", sumTotal)
}

// Declaring main function
func main() {
	// Calling addition function
	addition()
}