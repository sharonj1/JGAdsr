package main

import (
	"fmt"
)

func main() {
	
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")
	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")
	var availableFunds int = 120
	
}
