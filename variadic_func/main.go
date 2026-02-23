package main

import (
	"fmt"
)

// func find(num int, nums ...int) {
// 	fmt.Printf("type of nums is %T\n", nums)
// 	found := false
// 	for i, v := range nums {
// 		if v == num {
// 			fmt.Println(num, "found at index", i, "in", nums)
// 			found = true
// 		}
// 	}
// 	if !found {
// 		fmt.Println(num, "not found in ", nums)
// 	}
// 	fmt.Printf("\n")
// }
// func main() {
// 	find(89, 89, 90, 95, 89)
// 	find(45, 56, 67, 45, 90, 109)
// 	find(78, 38, 56, 98)
//     find(87)
// }


// using slice instead of variadic parameter

// func find(num int, nums []int) {  
//     fmt.Printf("type of nums is %T\n", nums)
//     found := false
//     for i, v := range nums {
//         if v == num {
//             fmt.Println(num, "found at index", i, "in", nums)
//             found = true
//         }
//     }
//     if !found {
//         fmt.Println(num, "not found in ", nums)
//     }
//     fmt.Printf("\n")
// }
// func main() {  
//     find(89, []int{89, 90, 95})
//     find(45, []int{56, 67, 45, 90, 109})
//     find(78, []int{38, 56, 98})
//     find(87, []int{})
// }

/*
The following of the advantages of using variadic arguments instead of slices.

1. There is no need to create a slice during each function call. If you look at the program above, we have created new slices during each function call. This additional slice creation can be avoided when using variadic functions
2. we are creating an empty slice just to satisfy the signature of the find function. This is totally not needed in the case of variadic functions. This line can just be find(87) when variadic function is used.
*/





// Passing a slice to a variadic function
func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {  
    nums := []int{89, 90, 95}
    // find(89, nums) // This will give error because find expects variadic arguments, not a slice
	find(89, nums...) // This will work because we are using ... to pass the slice as variadic arguments

	// According to the definition of a variadic function, nums ...int means that it will accept a variable number of arguments of type int. When we pass nums, which is a slice of int, it does not match the expected type of variadic arguments. However, when we use nums..., it tells the compiler to treat the slice as a series of individual arguments, which matches the expected type of variadic arguments.
}