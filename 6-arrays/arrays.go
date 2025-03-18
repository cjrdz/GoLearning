// Declaring main package
package main
// Importing the fmt package
import ("fmt")

func arrays() {

	// Arrays in Go using var keyword and the array type declaration
	// Declaring and initializing an array
	// here length is defined
	var busPrices = [3]int{10, 20, 30}
	// here length is inferred
	var taxiPrices = [...]int{10, 20, 30}

	// Arrays in Go using := sign and the array type declaration
	// Declaring and initializing an array
	// here length is defined
	uberCarPrices := [3]int{10,20,30}
	// here length is inferred
	uberMotoPrices := [...]int{10,20,30}
	fmt.Println("***************************************")
	fmt.Println("Public Transport Prices")
	fmt.Println("Here are bus prices: $",busPrices)
	fmt.Println("Here are taxi prices: $",taxiPrices)
	fmt.Println("***************************************")
	fmt.Println("Uber Prices")
	fmt.Println("Here are uber prices: $",uberCarPrices)
	fmt.Println("Here are uber prices: $",uberMotoPrices)
	fmt.Println("***************************************")
}

// The mutiple arrays function
func multipleArrays() {
	// Declaring multiple arrays
	// Declaring and initializing an array
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	var cars = [3]string{"Toyota", "Honda", "Ford"}
	// Changing the value of the first element of the array
	arr2[0] = -4

	// Printing the arrays
	fmt.Println("Multiple Arrays")
	fmt.Println("Here are arr1: ",arr1)
	fmt.Println("Here are arr2: ",arr2)
	fmt.Println("My favorite brand cars are: ",cars)
	// Accessing the first element of the cars array
	fmt.Println("I have a car from: ",cars[0])
	fmt.Println("***************************************")
}

// The arrayIntialization function
func arrayIntialization() {
	// Array not initialized
	arr1 := [5]int{}
	// Array partially initialized
	arr2 := [5]int{1,2}
	// Array fully initialized
	arr3 := [5]int{1,2,3,4,5}

	// Printing the arrays
	fmt.Println("Array Initialization")
	fmt.Println("Here are arr1: ",arr1)
	fmt.Println("Here are arr2: ",arr2)
	fmt.Println("Here are arr3: ",arr3)
	fmt.Println("***************************************")

}

// The length of the array function
func arrayLength() {
	// Declaring arrays
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	// Declaring arrays using the ... operator
	arr2 := [...]int{1,2,3,4,5,6}

	// Using the arrays to avoid unused variable errors
	fmt.Println("Array arr1: ", arr1)
	fmt.Println("Array arr2: ", arr2)
  
	// Printing the length of the arrays
	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
	fmt.Println("***************************************")

}

// The main function
func main() {
	// Calling the arrays function
	arrays()
	// Calling the multipleArrays function
	multipleArrays()
	// Calling the arrayIntialization function
	arrayIntialization()
	// Calling the arrayLength function
	arrayLength()
}