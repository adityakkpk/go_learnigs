package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("PLease enter a number between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	val, _ := reader.ReadString('\n')

	fmt.Println("You have entered: ", val)


	numInput, err := strconv.ParseFloat(strings.TrimSpace(val), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numInput + 1)
	}
}