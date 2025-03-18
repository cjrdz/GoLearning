// Declaring main package
package main

// Importing the fmt package
import (
	"fmt"
)

func slice(){
	fmt.Println("***************************************")
	// Declaring a slice
	mySlice := []int{1, 2, 3, 4, 5}

	// Printing the slice
	fmt.Println("My Slice: ", mySlice)
}

func lencap(){
	fmt.Println("***************************************")
	mine := []int{1, 2, 3, 4, 5}
	fmt.Println("Using slice with numbers elements")
	// Printing the length of the slice
	fmt.Println("Length of the slice: ", len(mine))
	// Printing the capacity of the slice
	fmt.Println("Capacity of the slice: ", cap(mine))
	// Appending a new element to the slice
	mine = append(mine, 6)
	// printing mine
	fmt.Print("Here, we go: ")
	// Iterating over the updated slice to print its elements
	for _, val := range mine {
		fmt.Print(val, " ")
	}
	fmt.Println()
	fmt.Println("***************************************")
	fmt.Println("Using slice with string elements")
	// Declaring a slice
	yours := []string{"Hello", "World", "My", "G"}
	// Printing the length of the slice
	fmt.Println("Length of the slice: ", len(yours))
	// Printing the capacity of the slice
	fmt.Println("Capacity of the slice: ", cap(yours))
	// Appending a new element to the slice
	yours = append(yours, "!")
	// Printing yours without square brackets
	fmt.Print("Hey: ")
	// Iterating over the updated slice to print its elements
	for _, val := range yours {
		fmt.Print(val, " ")
	}
	fmt.Println()
	fmt.Println("***************************************")
	fmt.Println("Itering over the slice and printing the index and value")
	// Printing a new line
	for i, v := range yours {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
	fmt.Println("***************************************")

}

// The arraySlice function
func arraySlice() {
	arra1 := [5]int{1, 2, 3, 4, 5}
	mySlice := arra1[1:4]

	// Printing the arrays
	fmt.Println("Array Slice")
	// Printing the array and the slice
	fmt.Println("Here are arra1: ", arra1)
	// Printing the length and the capacity of the slice
	fmt.Println("Here are mySlice: ", len(mySlice), cap(mySlice), mySlice)
	fmt.Println("***************************************")
	fmt.Println("Itering over the slice and printing the index and value")
	// Iterating over the slice to print its elements
	for _, val := range mySlice {
		fmt.Println(val, " ")
	}
	fmt.Println("***************************************")
}

// Main function
func main() {
	// Calling the slice function
	slice()
	// Calling the lencap function
	lencap()
	// Calling the arraySlice function
	arraySlice()
}