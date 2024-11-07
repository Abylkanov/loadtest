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
	// Пример использования
	fmt.Println(UniqCombN("abc", 1))  // ["a", "b", "c"]
	fmt.Println(UniqCombN("abc", 2))  // ["ab", "ac", "bc"]
	fmt.Println(UniqCombN("ab", 2))   // ["ab", "ba"]
	fmt.Println(UniqCombN("a", 1))    // ["a"]
	fmt.Println(UniqCombN("ab", 3))   // []
	fmt.Println(UniqCombN("aa", 1))   // []
}