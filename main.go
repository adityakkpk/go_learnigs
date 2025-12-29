package main

// import (
// 	"fmt"
// )

// ################ Basics ##################

// func main() {
// 	fmt.Println("This is Me Aditya & my favorite number is: ", rand.Intn(10))

// 	// One of the example of Go Encapsulation
// 	fmt.Println(math.Pi); // This is a constant which is accessible without importing (Capital case letter defined variables can be accessible without importing like "math/rand", it can be directly used with the package name itself)
// }

// ################# Functions ###############

// func add (x int, y int) int {
// 	return x + y
// }

// func swap(x string, y string) (string, string) {
// 	return y, x
// }

// Named return values (not preferable)
// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return  // now here not have to explicitly write x and y with the return
// }

// func split(sum int) (int, int) {
// 	x := sum * 4 / 9
// 	y := sum - x

// 	return x, y
// }

// func main () {
// 	// fmt.Println(add(12, 13))

// 	// a, b := swap("Hello", "World")
// 	// fmt.Println(a, b)

// 	fmt.Println(split(17))
// }

// #################### Variables ##################

// In go every variable to initialized to its zero state

// Variable declared out side of the function should follow the "var" pattern

// ---------- Rules to use which patter when -------------
// 1. If there is no initial value required for a variable and we are fine with zero state than use : 'var i int'
// 2. If there is an initial value and it is important : i := 0
// 3. If we are outside of the function than: var c string / var c = "1"

// var c, python, java bool // All are of boolean type (zero state of bool is false)

// func main() {
// 	var i int // i is of integer type (zero state of int is 0)

// 	var f float64 // f is of integer Floating point (64) (zero state of float is 0)

// 	// Another way of defining variable inside a function is ":=" when you know the initial value of a variable

// 	x := 10 // It is just a convention (syntactic sugar) of writing var x int = 10

// 	fmt.Println(i, c, python, java, f, x)

// New convention of variable assignments
// var i, j int = 1, 2
// k := 3
// c, python, java := true, false, "no!"

// fmt.Println(i, j, k, c, python, java)
// }

// basic variables and their zero values
// func main() {
// 	var i int
// 	var f float64
// 	var b bool
// 	var s string
// 	fmt.Printf("%v %v %v %q\n", i, f, b, s)
// }

// ################## Datatype ##########################
// 1. Basics -> number, string, bool
// 2. Aggregate -> array, struct
// 3. Reference -> pointers, slices, functions, channels, maps
// 4. Interfaces

// ------------- Basic Types --------------------
/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

// var (
// 	ToBe	bool            = false
// 	MaxInt	uint64 	        = 1<<64 - 1
// 	z		complex128		= cmplx.Sqrt(-5 + 12i)
// )

// func main() {
// 	fmt.Printf("Type of value %v is %T\n", ToBe, ToBe)
// 	fmt.Printf("Type of value %v is %T\n", MaxInt, MaxInt)
// 	fmt.Printf("Type of value %v is %T\n", z, z)
// }

// %v is used to print the value of a variable
// %T is used to print the type of a variable
// %q is used to print the string in a quotes
// %s is used to print the value of a string

// ############### Type Conversion ################
/*
Type conversions
The expression T(v) converts the value v to the type T.

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

Or, put more simply:

i := 42
f := float64(i)
u := uint(f)
Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.
*/

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var x, y int = 3, 4
// 	var f float64 = math.Sqrt(float64(x*x + y*y))
// 	var z uint = uint(f)
// 	fmt.Println(x, y, z)
// }

// ################### Type Inference ######################
/*
When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

var i int
j := i // j is an int
But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
Try changing the initial value of v in the example code and observe how its type is affected.
*/

// import "fmt"

// func main() {
// 	v := 42.4554

// 	var x float64
// 	y := x

// 	fmt.Printf("v is of type %T\n", v)
// 	fmt.Printf("y is of type %T\n", y)
// }

// ################### Constants ######################
/*
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.
*/

// import "fmt"

// const Pi = 3.14

// func main() {
// 	const World = "世界"
// 	fmt.Println("Hello", World)
// 	fmt.Println("Happy", Pi, "Day")

// 	const Truth = true
// 	fmt.Println("Go rules?", Truth)
// }

// ################### Numeric Constants ######################
/*
Numeric constants are high-precision values.

An untyped constant takes the type needed by its context.

Try printing needInt(Big) too.

(An int can store at maximum a 64-bit integer, and sometimes less.)
*/

// import "fmt"

// const (
// 	// Create a huge number by shifting a 1 bit left 100 places.
// 	// In other words, the binary number that is 1 followed by 100 zeroes.
// 	Big = 1 << 100
// 	// Shift it right again 99 places, so we end up with 1<<1, or 2.
// 	Small = Big >> 99
// )

// func needInt(x int) int { return x*10 + 1 }
// func needFloat(x float64) float64 {
// 	return x * 0.1
// }

// func main() {
// 	fmt.Println(needInt(Small))
// 	fmt.Println(needFloat(Small))
// 	fmt.Println(needFloat(Big))
// }

// ########################## Loops & Conditionals ################################
// import (
// 	"fmt"
//  	"math"
// )

// ---------------- For loop ----------------------------
// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("The value of i is %v\n", i);
// 	}
// }

// --------------- If --------------------------
// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim {
// 		return v
// 	}

// 	return lim
// }

// func main() {
// 	fmt.Println(
// 		pow(3, 2, 10),
// 		pow(3, 3, 20),
// 	)
// }

// ######################### Switch Case Statement #########################
// import (
// 	"fmt"
// 	// "runtime"
// 	"time"
// )

// func main() {
// 	// fmt.Print("Go runs on ")
// 	// switch os := runtime.GOOS; os {
// 	// case "darwin":
// 	// 	fmt.Println("Os X.")
// 	// case "linux":
// 	// 	fmt.Println("Linux")
// 	// default:
// 	// 	fmt.Printf("%s.\n", os)
// 	// }

// 	// fmt.Println("When is Saturday?")
// 	// today := time.Now().Weekday()

// 	// switch time.Saturday {
// 	// case today + 0:
// 	// 	fmt.Printf("Today. %v\n", today + 0)
// 	// case today + 1:
// 	// 	fmt.Printf("Tomorrow. %v\n", today + 1)
// 	// case today + 2:
// 	// 	fmt.Printf("In two days. %v\n", today + 2)
// 	// default:
// 	// 	fmt.Printf("Too far away")
// 	// }

// 	t := time.Now()

// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good Morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("Good Afternoon!")
// 	default:
// 		fmt.Println("Good Evening!")
// 	}
// }

// ############################## Defer ###################################
// import (
// 	"fmt"
// )

// func main() {
// defer fmt.Println("World!")
// fmt.Println("Hello")

// fmt.Println("Counting")
// for i := 0; i < 10; i++ {
// 	defer fmt.Println(i)
// }
// fmt.Println("done")

// Why defer?
/*
	1. Closing the files, db close, etc to save from memory Leaks.
	ex: open(file)
		defer close(file)

	The defer statements runs even in the Panic situations
*/
// }

// #################### Reference Types #####################
// import (
// 	"fmt"
// )

// func main() {
// 	i, j := 42, 2345

// 	p := &i				// point to i
// 	fmt.Println(*p)		// read i through the pointed
// 	*p = 21				// Set i through the pointer
// 	fmt.Println(i)		// see the new value of i

// 	p = &j 				// point to j
// 	*p = *p / 37 		// divide j through the pointer
// 	fmt.Println(j)		// see the new value of j
// }

// ############# User Input ####################
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter Your Name: ")
	// var name string;
	// fmt.Scan(&name)
	// fmt.Println("My Name is", name)


	// Multi word Input
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello,", name)



}











// ###################### Struct ########################