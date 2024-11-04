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
	fmt.Println(LCM(4, 6))  // 12
	fmt.Println(LCM(21, 6)) // 42
	fmt.Println(LCM(5, 7))  // 35
}
