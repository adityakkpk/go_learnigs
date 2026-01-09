package main

import (
	"fmt"
)

func main() {
	fmt.Println("If - Else Statement")

	// if Statement
	// num := 10
	// if num%2 == 0 {
	// 	fmt.Println("The number", num, "is even.")
	// 	return
	// }
	// fmt.Println("The number", num, "is odd.")

	// if-else statement
	// num := 10
	// if num%2 == 0 {
	// 	fmt.Println("The number", num, "is even.")
	// } else {
	// 	fmt.Println("The number", num, "is odd.")
	// }

	// if - else if - else statement
	// age := 10
	ticketPrice := 0
	// if age <= 0 {
	if age := 10; age <= 0 { // if with assignment
		ticketPrice = 0
	} else if age >= 5 && age <= 18 {
		ticketPrice = 10
	} else {
		ticketPrice = 15
	}
	fmt.Println("Ticket Price is", ticketPrice)
}

/* Gotcha
The else statement should start in the same line after the closing curly brace } of the if statement. If not the compiler will complain.


The reason is because of the way Go inserts semicolons automatically. You can read about the semicolon insertion rule here https://go.dev/ref/spec#Semicolons.

In the rules, it’s specified that a semicolon will be inserted after closing brace }, if that is the final token of the line. So a semicolon is automatically inserted after the if statement’s closing braces } in line no. 11 by the Go compiler.

Example:
...
if num%2 == 0 {
      fmt.Println("the number is even")
};  //semicolon inserted by Go Compiler
else {
      fmt.Println("the number is odd")
}

*/

/* Idiomatic Go
In Go’s philosophy, it is better to avoid unnecessary branches and indentation of code. It is also considered better to return as early as possible.


Example :
func main() {
	num := 10;
	if num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
		return
	}
	fmt.Println(num, "is odd")

}

In the above program, as soon as we find out the number is even, we return immediately. This avoids the unnecessary else code branch.

*/
