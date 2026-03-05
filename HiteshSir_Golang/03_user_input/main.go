package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "This is all about user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter Something:")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered Value is: ", input)
	fmt.Printf("Type of input is: %T", input)
}