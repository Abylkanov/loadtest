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
	fmt.Println(NthPrime(1))  // 2
	fmt.Println(NthPrime(5))  // 11
	fmt.Println(NthPrime(10)) // 29
}
