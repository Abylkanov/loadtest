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
	result := ConvertNbrBase(1465, "0123456789")
	fmt.Println(result) // 1465

	result = ConvertNbrBase(1465, "01")
	fmt.Println(result) // 10110111001

	result = ConvertNbrBase(1465, "01234567")
	fmt.Println(result) // 2671

	result = ConvertNbrBase(1465, "0123456789ABCDEF")
	fmt.Println(result) // 5B9

	result = ConvertNbrBase(1465, "00")
	fmt.Println(result) //
}
