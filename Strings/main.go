package main

import (
	"fmt"
	// "unicode/utf8"
)

// func printBytes(s string) {
// 	fmt.Printf("Bytes: ")
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x ", s[i])
// 	}
// }

// func printChars(s string) {
// 	fmt.Printf("Characters: ")
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%c ", s[i])
// 	}

// }

// func main() {
// 	name := "Go Programming"
// 	fmt.Printf("String: %s\n", name)
// 	printBytes(name) // 47 6f 20 50 72 6f 67 72 61 6d 6d 69 6e 67 (These are the Inicode UT8-encoding values of "Go Programming")
// 	fmt.Printf("\n")
// 	printChars(name)
// }

// func main() {
// 	name := "Hello World"
// 	fmt.Printf("String: %s\n", name)
// 	printChars(name)
// 	fmt.Printf("\n")
// 	printBytes(name)
// 	fmt.Printf("\n\n")

// 	name = "Señor"
// 	fmt.Printf("String: %s\n", name)
// 	printChars(name) // S e Ã ± o r
// 	fmt.Printf("\n")
// 	printBytes(name)
// }

// In line no. 39 of the program above, we are trying to print the characters of Señor and it outputs S e Ã ± o r which is wrong. Why does this program break for Señor when it works perfectly fine for Hello World . The reason is that the Unicode code point of ñ is U+00F1 and its UTF-8 encoding occupies 2 bytes c3 and b1. We are trying to print the characters assuming that each code point will be one byte long which is wrong. In UTF-8 encoding a code point can occupy more than 1 byte. So how do we solve this? This is where rune saves us.

//////////////////////// RUNE ////////////////////////////
// func printBytes(s string) {
// 	fmt.Printf("Bytes: ");
// 	runes := []rune(s)
// 	for i := 0; i < len(runes); i++ {
// 		fmt.Printf("%x ", runes[i])
// 	}
// }

// func printChars(s string) {
// 	fmt.Printf("Chars: ");
// 	runes := []rune(s)
// 	for i := 0; i < len(runes); i++ {
// 		fmt.Printf("%c ", runes[i])
// 	}
// }

// func main() {
// 	name := "Hello World"
// 	fmt.Printf("String: %s\n", name)
// 	printChars(name)
// 	fmt.Printf("\n")
// 	printBytes(name)
// 	fmt.Printf("\n\n")

// 	name = "Señor"
// 	fmt.Printf("String: %s\n", name)
// 	printChars(name)
// 	fmt.Printf("\n")
// 	printBytes(name)
// }

// ---------- Accessing individual runes using for range loop ---------
// func chatAndBytePosition(s string) {
// 	for index, rune := range(s) {
// 		fmt.Printf("%c starts at byte %d\n", rune, index)
// 	}
// }

// func main() {
// 	name:="Señor"
// 	chatAndBytePosition(name) // From the output, it’s clear that ñ occupies 2 bytes since the next character o starts at byte 4 instead of byte 3 😀
// }

// ******************* Creating a string from a slice of bytes *******************
// func main() {
// 	// byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
// 	byteSlice := []byte{67, 97, 102, 195, 169}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
// 	str := string(byteSlice)
// 	fmt.Println(str)
// }

// ******************* Creating a string from a slice of runes *******************
// func main() {
// 	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
// 	str := string(runeSlice)
// 	fmt.Println(str)
// }

// ************************ String Length **************************
/*
As discussed earlier, len(s) is used to find the number of bytes in the string and it doesn’t return the string length. As we already discussed, some Unicode characters have code points that occupy more than 1 byte. Using len to find out the length of those strings will return the incorrect string length.

The RuneCountInString(s string) (n int) function of the utf8 package can be used to find the length of the string. This method takes a string as an argument and returns the number of runes in it.

func main() {
	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Lenght: %d\n", utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d\n", len(word1))

	fmt.Printf("\n")
	word2 := "Pets"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d\n", len(word2))
}
// Oputput
String: Señor
Lenght: 5
Number of bytes: 6

String: Pets
Length: 4
Number of bytes: 4
*/


// ************************ String Comparision **************************
/*
The == operator is used to compare two strings for equality. If both the strings are equal, then the result is true else it’s false.

func compareStrings(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return false
}
func main() {
	str1:= "Go"
	str2:= "Go"
	fmt.Printf("%s is eqlal to %s: %t\n", str1, str2, compareStrings(str1, str2))

	str3:= "Hello"
	str4:= "Aditya"
	fmt.Printf("%s is eqlal to %s: %t\n", str3, str4, compareStrings(str3, str4))

}

*/


// ******************* Strings are Imutable *******************
/*
In Go, strings are immutable which means that once a string is created, it cannot be changed. If you want to modify a string, you need to create a new string with the desired changes.

func mutate(s string)string {
	s[0] = 'a'//any valid unicode character within single quote is a rune 
	return s
}
func main() {  
    h := "hello"
    fmt.Println(mutate(h))
}

// In above program, we try to change the first character of the string to 'a'. Any valid Unicode character within a single quote is a rune. We try to assign the rune a to the zeroth position of the slice. This is not allowed since the string is immutable and hence the program fails to compile with error ./prog.go:165:2: cannot assign to s[0] (neither addressable nor a map index expression)

// To workaround this string immutability, strings are converted to a slice of runes. Then that slice is mutated with whatever changes are needed and converted back to a new string.
func mutate(s []rune) string {
	s[0] = 'a' 
	return string(s)
}
func main() {  
    h := "hello"
    fmt.Println(mutate([]rune(h)))
}
*/