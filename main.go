package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func importSave() {
	ap.PutRune('a')
	fmt.Println(`test`)
}

func main() {
    fmt.Println(PrimeFactors(28))  // [2 2 7]
    fmt.Println(PrimeFactors(100)) // [2 2 5 5]
    fmt.Println(PrimeFactors(17))  // [17]
}