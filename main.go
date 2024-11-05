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
	fmt.Println(UniquePermutations("aabb")) // ["aabb", "abab", "abba", "baab", "baba", "bbaa"]
	fmt.Println(UniquePermutations("abc"))  // ["abc", "acb", "bac", "bca", "cab", "cba"]
	fmt.Println(UniquePermutations("aaa"))  // ["aaa"]
}
