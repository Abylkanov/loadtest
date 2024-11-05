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
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	maxKey := MapMax(m)
	fmt.Println(maxKey) // three
}
