package main

import (
	"fmt"
	"github.com/alem-platform/ap"
)

func importSave(){
	ap.PutRune('a')
	fmt.Println(`test`)
}

func main() {
    fmt.Println(CommonElements([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})) // [3]
    fmt.Println(CommonElements([]int{1, 2, 3}, []int{1, 1, 1}, []int{1, 2}))    // [1]
}