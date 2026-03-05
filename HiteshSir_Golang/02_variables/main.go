package main

import "fmt"

// This is a public var because its first letter is Capital
const LoginToken string = "gakjdhfiensnnfiusuhfdsfsdfkjdsbfjdffjsfsdnfsbfsnfsd"

func main() {
	var username string = "Aditya"
	fmt.Println(username)
	fmt.Printf("Variable is of type:%T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type:%T\n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type:%T\n", smallVal)

	var smallFloatVal float64 = 255.455555555555555555
	fmt.Println(smallFloatVal)
	fmt.Printf("Variable is of type:%T\n", smallFloatVal)

	// default values and some aliases
	var anotherVar int = 256
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type:%T\n", anotherVar)

	// implicit type
	var website = "https://adityakkpk.com"
	fmt.Println(website)


	// no var style : only inside a method or function
	numbersOfUsers := 300000
	fmt.Println(numbersOfUsers)


	fmt.Println(LoginToken);

}
