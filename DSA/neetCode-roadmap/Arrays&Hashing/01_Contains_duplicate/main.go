package main

import "fmt"

func main() {
	fmt.Println(hasDuplicate([]int{1, 2, 3, 3, 4}))
}

// solution
func hasDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = struct{}{}
	}

	return false
}

// Question: Why use struct for map value instead of bool/int
// Answer: Because int takes 8 bytes and Stores a whole number, bool takes 1 byte and Stores true/false and struct{} takes 0 bytes as It is an empty object with no fields.
