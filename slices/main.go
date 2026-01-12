package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices in Go lang")

	a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[1:4] // Creating a slice from a[1] to a[3]
	fmt.Println(b)

	// Another way of creating slice
	c := []int{3, 4, 5}
	fmt.Println("New Slice c:", c)

	// Modifying a Slice: A slice does not own any data of its own. It is just a representation of the underlying array. Any modifications done to the slice will be reflected in the underlying array and vice versa.
	darr := [...]int{34, 55, 64, 34, 53, 44, 56}
	dslice := darr[2:5]
	fmt.Println("Array before", darr)

	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("Array after", darr)

	// When a number of slices share the same underlying array, the changes that each one makes will be reflected in the array.
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	// length and capacity of a slice :
	// The length of the slice is the number of elements in the slice. The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "banana"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("Length of slide %d capacity %d\n", len(fruitslice), cap(fruitslice))

	//A slice can be re-sliced upto its capacity. Anything beyond that will cause the program to throw a run time error.
	fruitslice = fruitslice[:cap(fruitslice)]//re-slicing fruitslice till its capacity
	fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is", cap(fruitslice))




















	fmt.Println("Creating Slices using make function.")
	// Creating a slice using make
	// func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity. The capacity parameter is optional and defaults to the length. The make function creates an array and returns a slice reference to it.
	i := make([]int, 5, 5)
	fmt.Println(i)



	// Appending to a slice
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6



	// The zero value of a slice is nil
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:",names)
	}


	// It is also possible to append one slice to another using the ... operator.
	veggies := []string{"potato", "tomato", "Carrot"}
	fruits := []string{"apple", "banana"}
	food := append(veggies, fruits...)
	fmt.Println("food:",food)














	fmt.Println("Passing a slice to a function:")
	// A slice contains the length, capacity and a pointer to the zeroth element of the array. When a slice is passed to a function, even though it’s passed by value, the pointer variable will refer to the same underlying array. Hence when a slice is passed to a function as parameter, changes made inside the function are visible outside the function too. Lets write a program to check this out.
	nos := []int{3,4,5}
	fmt.Println("Slice befor :",nos)
	subtactOne(nos)
	fmt.Println("Slice after :",nos)







	fmt.Println("Multidimensional Slices:")
	pls := [][]string {
		{"C", "C++"},
		{"JS"},
		{"Go", "Rust"},
	}

	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Println("\n")
	}











	// Memory management : Slices hold a reference to the underlying array. As long as the slice is in memory, the array cannot be garbage collected. This might be of concern when it comes to memory management. Let’s assume that we have a very large array and we are interested in processing only a small part of it. Henceforth we create a slice from that array and start processing the slice. The important thing to be noted here is that the array will still be in memory since the slice references it.

	// One way to solve this problem is to use the copy function func copy(dst, src []T) int to make a copy of that slice. This way we can use the new slice and the original array can be garbage collected.
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}

func countries() []string { 
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}