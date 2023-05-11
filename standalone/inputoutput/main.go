package main

import (
	"fmt"

	"frontendmasters.co/go/io/data"
)

// global variables scope
// var url  = "https://frontendmaster.com"
var name = "Adrian"

func main() {
	fmt.Print(data.MaxSpeed)
	// function-scoped variables
	// var message string = "Hello from Go"
	// price := 34.4 // cannot use shortcut := in the global scope. You will have to use var or const
	// isReady := true
	// print(message, price, url, isReady)
	fmt.Println("dfdfdf")
	PrintData()
}
