package main

import "fmt"

func main() {
	fmt.Println(isAnagram("racecar", "carrace"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charFrequency := make(map[rune]int)

	for _, val := range s {
		charFrequency[val]++
	}

	for _, val := range t {
		charFrequency[val]--

		if charFrequency[val] < 0 {
			return false
		}
	}

	return true
}
