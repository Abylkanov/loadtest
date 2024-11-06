package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func printAlphaOdd() {
	for i := 0; i < 26; i++ {
		if i%2 != 0 {
			ap.PutRune(rune(i) + 'a')
		} else {
			ap.PutRune(rune(i) + 'A')
		}
	}
}

func printAlphaOddRev() {
	for i := 0; i < 26; i++ {
		if i%2 != 0 {
			ap.PutRune(-rune(i) + 'z')
		} else {
			ap.PutRune(-rune(i) + 'Z')
		}
	}
}

func PutDigit(n int) {
	if n >= 0 && n < 10 {
		ap.PutRune(rune(n + '0'))
	}
}

func Swap(a, b *int) {
	temp := *b
	*b = *a
	*a = temp
}

func SwapCase(r [20]rune) [20]rune {
	res := [20]rune{}
	for i := 0; i < 20; i++ {
		if r[i] >= 'A' && r[i] <= 'Z' {
			res[i] = r[i] + 32
		} else if r[i] >= 'a' && r[i] <= 'z' {
			res[i] = r[i] - 32
		} else if r[i] == '\000' {
			break
		} else {
			res[i] = r[i]
		}
	}
	return res
}

func SetMax(a, b, c *int) {
	if *a > *b {
		*c = *a
	} else {
		*c = *b
	}
}

func PrintRectangle(a, b int) {
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			ap.PutRune('0')
			if j != b-1 {
				ap.PutRune(' ')
			}

		}
		ap.PutRune('\n')
	}
}

func PrintAlphaTri(n int) {
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			ap.PutRune(rune(j%26 + 'A'))
			if j != i-1 {
				ap.PutRune(' ')
			}
		}
		ap.PutRune('\n')
	}
}

func PrintAlphaPyramid(n int) {
	line := 1
	pos := 0

	for i := 0; i < n; i++ {
		ap.PutRune(rune(i%26 + 'A'))
		pos++
		if pos == line {
			line++
			pos = 0
			ap.PutRune('\n')
		}

	}
}

func RunesReverse(arr [20]rune) [20]rune {
	res := [20]rune{}
	count := 0
	for i := 0; i < 20; i++ {
		if arr[i] == '\000' {
			break
		}
		count++
	}

	for i := 0; i < count; i++ {
		res[i] = arr[count-i-1]
	}
	return res
}

func PutNumber(n int) {
	if n < 0 {
		ap.PutRune('-')
		n = -n
	}
	if n < 10 {
		ap.PutRune(rune(n + '0'))
	} else {
		PutNumber(n / 10)
		PutNumber(n % 10)
	}
}

func AtoN() {
	var n int
	fmt.Scanf("%d", &n)
	var str string
	fmt.Scanf("%s\n", &str)

	for i := 0; i < len(str)-1; i++ {
		PutNumber(int(str[i]))
		ap.PutRune(' ')
	}

	PutNumber(int(str[len(str)-1]))
}

func budget_desc() {
	var n int
	fmt.Scanf("%d\n", &n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		var rating, budget int
		fmt.Scanf("%d %d\n", &rating, &budget)
		arr[i] = budget
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
