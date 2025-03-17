// Declaring main package
package main
// Importing the fmt package
import ("fmt")

// The students function
func students() {
	// Declaring strying variable
	var student1 string = "Carlos"
	// Declaring strying variable
	var student2 string = "Jonathan"
	// Declaring integer variable
	x := 2

	// Printing the values of the variables
	fmt.Println("************************************")
	fmt.Println("The students are", student1, "and", student2)
	fmt.Println("The value of x is", x)
	fmt.Println("************************************")
}

// The alphabet function
func alphabet() {
	// Declaring variables with no values
	var a string
	var b int
	var c bool

	// Printing the values of the variables
	fmt.Println("The value of a is", a)
	fmt.Println("The value of b is", b)
	fmt.Println("The value of c is", c)
	fmt.Println("************************************")
}

// The people function
func name() {
	// Declaring a variable with a value
	var name1 string = "Carlos"
	// Assigning a value to the variable
	name1 = "Carlos"

	// This is the best practice for declaring and assigning a value to a variable
	name2 := "Jonathan"
	age := 25
	
	// Printing the value of the variable
	fmt.Println("My first name is", name1)
	fmt.Println("My middle name is", name2)
	fmt.Println("My age is", age)
	// Printing the values of the variables
	fmt.Println("My name is", name1, name2 + ",", "and I am", age, "years old.")
	fmt.Println("************************************")
}

// The multiple function
func multiple() {
	// Declaring multiple variables
	var a, b, c, d, f int = 1, 2, 3, 4, 5

	// Declaring multiple variables with different data types
	// The variables are declared in a single line
	var g, h = 6, "Estas para vosotros"
	i, j := 7, "Siuuuuuuu!"


	// Printing the values of the variables
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)

	// Printing the values of the variables
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println("************************************")
}

// The main function
func main() {
	// Calling the students function
	students()
	// Calling the alphabet function
	alphabet()
	// Calling the people function
	name()
	// Calling the multiple function
	multiple()
}