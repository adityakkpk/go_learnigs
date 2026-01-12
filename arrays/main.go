// package main

// import (
// 	"fmt"
// )

// func main () {
// 	fmt.Println("Arrays in Go Lang.")

// 	var a [3]int
// 	a[0] = 12;
// 	a[1] = 13;
// 	a[2] = 14;
// 	fmt.Println(a);

// 	// Short hand declaration to create array
// 	b := [3]int{12, 78, 43} 
// 	fmt.Println(b);



// 	// Let the compiler decide the length of the array
// 	c := [...]int{1,2,3,4,5,6,7,8,9}
// 	fmt.Println(len(c), c);

// 	//The size of the array is a part of the type. Hence [5]int and [25]int are distinct types. Because of this, arrays cannot be resized.
// 	// d:= [3]int{5, 78, 8}
// 	// var e [5]int
// 	// e = d//not possible since [3]int and [5]int are distinct types


// 	//Arrays in Go are value types and not reference types. This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable. If changes are made to the new array, it will not be reflected in the original array.
// 	f := [...]string{"USA", "China", "India"}
// 	g := f
// 	g[0] = "UK"
// 	fmt.Println("f is", f)
// 	fmt.Println("g is", g)

// 	// Similarly when arrays are passed to functions as parameters, they are pass by value and the original array is unchanged.









// 	//For convenience from the for loop, Go provides a better and concise way to iterate over an array by using the range form of the for loop. range returns both the index and the value at that index. Let’s rewrite the above code using range.
// 	h := [...]float64{44.5, 44.3, 23.2, 32}
// 	sum := float64(0)
// 	for i, v := range h {// range returns both the index and value
// 		fmt.Printf("%d the element of h is %.2f\n", i, v);
// 		sum += v;
// 	}
// 	fmt.Println("\nSum of all elements of h", sum);
// }





// Multidimensional Arrays
package main

import (
	"fmt"
)

func printarray(a [3][2]string) {
	for i, v1 := range a {
		for j, v2 := range v1 {
			fmt.Printf("The co-ords: [%d,%d], value: %s\t", i, j, v2)
		}
		fmt.Println()
	}
}

func main() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b)
}
