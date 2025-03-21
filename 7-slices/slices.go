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

// Create a slice with a function
func funcSlice(){
	// Declaring the function slice
	funcSlice1 := make([]int, 5, 10)
	fmt.Println("Function Slice created with make")
	fmt.Printf("Function Slice: %v\n", funcSlice1)
	// Printing the length of the slice
	fmt.Printf("Length of the slice: %d\n", len(funcSlice1))
	// Printing the capacity of the slice
	fmt.Printf("Capacity of the slice: %d\n", cap(funcSlice1))
	
	fmt.Println("***************************************")
	fmt.Println("Ommiting / Changing the capacity")
	// With omitted capacity
	funcSlice2 := make([]int, 5)
	fmt.Printf("Function Slice2: %v\n", funcSlice2)
	// Printing the length of the slice
	fmt.Printf("Length of the slice: %d\n", len(funcSlice2))
	// Printing the capacity of the slice
	fmt.Printf("Capacity of the slice: %d\n", cap(funcSlice2))
	fmt.Println("***************************************")
}

func elementSlice(){
	fmt.Println("Accessing elements of a slice")
	// Declaring a slice
	mySlice1:= []int{1, 2, 3, 4, 5}
	// Printing the slice
	fmt.Println("My Slice: ", mySlice1)
	// Printing the first element of the slice
	fmt.Println("First element of the slice: ", mySlice1[0])
	// Printing the last element of the slice
	fmt.Println("Last element of the slice: ", mySlice1[len(mySlice1)-1])
	fmt.Println("***************************************")
	// Changing the value of the first element of the slice
	fmt.Println("Changing the value of the first element of the slice")
	mySlice2 := []int{1, 2, 3, 4, 5}
	// Changing the value of the first element of the slice
	mySlice2[0] = -1

	// Printing the slice
	fmt.Println("My Slice: ", mySlice2)
	fmt.Println("***************************************")

}

func appendSlice(){
	print("Appending elements to a slice")
	// Declaring two slices
	mySlice1 := []int{1, 2, 3}
	mySlice2 := []int{4, 5, 6}

	// Appending the second slice to the first slice
	mySlice3 := append(mySlice1, mySlice2...)

	// Printing the slices
	fmt.Printf("myslice3=%v\n", mySlice3)
	// Printing the length of the slice
	fmt.Println("Length of the slice: ", len(mySlice3))
	// Printing the capacity of the slice
	fmt.Println("Capacity of the slice: ", cap(mySlice3))
	fmt.Println("***************************************")

}

// Memory efficient using copy() function
func memoryEf(){
	fmt.Println("Memory efficient using copy() function")
	// Declaring two slices
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// Original Slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length of numbers = %d\n", len(numbers))
	fmt.Printf("capacity of numbers = %d\n", cap(numbers))

	// Create copy of the slice only need numbers
	neededNumbers := numbers[:len(numbers)-10]
	// Copy of the slice
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	// Printing the slices
	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length of numbersCopy = %d\n", len(numbersCopy))
	fmt.Printf("capacity of numbersCopy = %d\n", cap(numbersCopy))
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
	// Calling the funcSlice function
	funcSlice()
	// Calling the elementSlice function
	elementSlice()
	// Calling the appendSlice function
	appendSlice()
	// Calling the memoryEf function
	memoryEf()
}