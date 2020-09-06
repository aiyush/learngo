package main
import "fmt"

func main() {
	fmt.Println("Hello!")
	
}

// This main block contains all the declared functions in a package.
// We can either call them in func main part above or we can call them individually in the packages.
// Check out the packages(bye.go and hey.go for more info)

func bye() {
	fmt.Println("Bye!")
}

func hey() {
	fmt.Println("hey!")
}