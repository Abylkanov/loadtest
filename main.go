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
    matrix := [][]rune{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 11, 12, 13},
		{14, 15, 16, 17},
    }

    RotateMatrix180(matrix)
	fmt.Println(matrix)
}