// Declaring main package
package main
// Importing the fmt package
import ("fmt")

func dataTypes() {
	// Declaring and initializing a string
	var name string = "Carlos"
	// Declaring and initializing an integer
	var age int = 25
	// Declaring and initializing a float64
	var weightthen float32 = 70.5
	var weightnow float64 = 80.5
	// Declaring and initializing a boolean
	var isStudent bool = false
	// Declaring and initializing a complex128
	var complexNumber complex128 = 1 + 2i
	// Declaring and initializing a byte
	var byteValue byte = 255
	// Declaring and initializing a rune
	var runeValue rune = 0x2318
	// Declaring and initializing an array
	var arrayValue [3]int = [3]int{1, 2, 3}
	// Declaring and initializing a slice
	var sliceValue []int = []int{1, 2, 3}

	// Printing the values of the variables
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("My weight was", weightthen, "KG", "and now is", weightnow, "KG")
	fmt.Println("Am I a student = ", isStudent)
	fmt.Println("A complexNumber is", complexNumber)
	fmt.Println("A byteValue is", byteValue)
	fmt.Println("A runeValue is", runeValue)
	fmt.Println("A arrayValue is", arrayValue)
	fmt.Println("A sliceValue is", sliceValue)
}

// The main function
func main() {

	// Calling the dataTypes function
	dataTypes()

}