// This is the main package
package main
// Importing the fmt and quote packages
import ("fmt")
import "rsc.io/quote"

// The main function
func main() {
	// Printing the message
	fmt.Println("Hello, World!")
	// Printing the message from the quote package
	fmt.Println(quote.Go())
}