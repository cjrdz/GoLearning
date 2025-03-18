// Declaring main package
package main
// Importing the fmt package
import ("fmt")

// Declaring ij function
func ij() {
	// Declaring and initializing two strings
	// Declaring two strings in a single line
	var i,j string = "Hello", "World"

	fmt.Println("**************************************************")
	// Printing the strings
	fmt.Println(i)
	fmt.Println(j)
	// Printing the strings using Println function
	// Println function adds a new line character at the end
	fmt.Println(i,j)
	// Printing the strings using Print function
	// Print function does not add a new line character at the end
	fmt.Print(i, "\n",j,"\n")
	fmt.Println("**************************************************")
}

func ik(){
	// Declaring and initializing a string and an integer
	// Declaring a string and an integer in a single line
	var i string = "Hello"
	var k int = 10

	// Printing the string and the integer
	// %v is used to print the value of the variable
	// %T is used to print the type of the variable
	// \n is used to add a new line character
	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T\n", k, k) 
	fmt.Println("**************************************************")
}

func itxt(){
	// Declaring and initializing a float64 and a string
	// Declaring a float64 and a string in a single line
	var i = 15.5
	var txt = "Hello World!"
	
	// Printing the float64 and the string
	// %v is used to print the value of the variable
	// %#v is used to print the value of the variable in a Go-syntax representation
	// %T is used to print the type of the variable
	// %v is used to print the value of the variable
	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)
  
	// Printing the string
	// %v is used to print the value of the variable
	// %#v is used to print the value of the variable in a Go-syntax representation
	// %T is used to print the type of the variable
	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt) 
	fmt.Println("**************************************************")
}

func ibool(){
	// Declaring and initializing two boolean variables
	// Declaring two boolean variables in a single line
	var i = true
	var j = false
	
	// Printing the boolean variables
	// %t is used to print the boolean value
	fmt.Printf("%t\n", i)
	fmt.Printf("%t\n", j)
	fmt.Println("**************************************************")
}

// main function
func main() {
	// Calling ij function
	ij()
	// Calling ik function
	ik()
	// Calling itxt function
	itxt()
	// Calling ibool function
	ibool()
}