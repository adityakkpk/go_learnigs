package main

import "fmt"

/*
func main() {


	// Panic: the map has to be initialized before adding elements.
	// var currencyCode map[string]string
	// currencyCode["USD"] = "US Dollar"


	currencyCode := make(map[string]string)
	// fmt.Println(currencyCode)

	currencyCode["USD"] = "US Dollar"
	currencyCode["GBP"] = "Pound Sterling"
	currencyCode["EUR"] = "Euro"
	// currencyCode["INR"] = "Indian Rupee"
	// fmt.Println("currencyCode map contents:", currencyCode)


	// currency := "USD"
	// currencyName := currencyCode[currency]
	// fmt.Println("Currency name for currency code", currency, "is", currencyName)


	// cyCode := "INR"
	// if currencyName, ok := currencyCode[cyCode]; ok {
	// 	fmt.Println("Currency name for currency code", cyCode, " is", currencyName)
	// 	return
	// }
	// fmt.Println("Currency name for currency code", cyCode, "not found")

	// for code, name := range currencyCode {
	// 	fmt.Printf("Currency Name for currency code %s is %s\n", code, name)
	// }

	fmt.Println("map before deletion", currencyCode)
	delete(currencyCode, "EUR")
	fmt.Println("map after deletion", currencyCode)


}
*/

// Map of Structs
type currency struct {
	name   string
	symbol string
}

func main() {
	curUSD := currency{
		name:   "US Dollar",
		symbol: "$",
	}
	curGBP := currency{
		name:   "Pound Sterling",
		symbol: "£",
	}
	curEUR := currency{
		name:   "Euro",
		symbol: "€",
	}

	currencyCode := map[string]currency{
		"USD": curUSD,
		"GBP": curGBP,
		"EUR": curEUR,
	}

	for cyCode, cyInfo := range currencyCode {
		fmt.Printf("Currency Code: %s, Name: %s, Symbol: %s\n", cyCode, cyInfo.name, cyInfo.symbol)
	}

}

/*
func main() {
    map1 := map[string]int{
        "one": 1,
        "two": 2,
    }

    map2 := map1

    if map1 == map2 {  // invalid operation: map1 == map2 (map can only be compared to nil)

    }
}
*/
