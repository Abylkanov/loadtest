package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func importS() {
	ap.PutRune('a')
	fmt.Println(`test`)
}

func CountA(slice []string) int {
	count := 0
	for _, str := range slice {
		for _, ch := range str {
			if ch == 'a' {
				count++
			}
		}
	}
	return count
}

func FindFirstA(slice []string) string {
	for _, str := range slice {
		for _, ch := range str {
			if ch == 'a' {
				return "a"
			}
		}
	}
	return ""
}

func EndsWith(s string, suffix string) bool {
	if len(suffix) > len(s) {
		return false
	}
	s = s[len(s)-len(suffix):]
	for i := len(suffix) - 1; i >= 0; i-- {
		if suffix[i] != s[i] {
			return false
		}
	}
	return true
}

func CommonElements(slice1, slice2, slice3 []int) []int {
	result := make([]int, 0)
	commonEl := make(map[int]bool)

	for i := 0; i < len(slice1); i++ {
		for j := 0; j < len(slice2); j++ {
			for k := 0; k < len(slice3); k++ {
				if slice1[i] == slice2[j] && slice1[i] == slice3[k] {
					commonEl[slice1[i]] = true
				}
			}
		}
	}
	for key, _ := range commonEl {
		result = append(result, key)
	}
	return result
}

func CapitalizeWords(s string) string {
	arr := []rune(s)

	if arr[0] >= 'a' && arr[0] <= 'z' {
		arr[0] -= 32
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == ' ' && arr[i+1] >= 'a' && arr[i-1] <= 'z' {
			arr[i+1] -= 32
		}
	}

	return string(arr)
}

func FactorialIterative(n int) int {
	res := 1
	switch {
	case n < 0:
		return -1
	case n == 0:
		return 1
	default:
		for i := 1; i <= n; i++ {
			res *= i
		}
		return res
	}
}

func FactorialRecursion(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}

	return FactorialRecursion(n-1) * n
}

func CountSubstring(s string, sub string) int {
	count := 0

	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			count++
		}
	}
	return count
}

func FindMaxMatrixElement(matrix [][]int) int {
	// if len(matrix) == 0 || len(matrix[0]) == 0 {
	// 	return -1
	// }
	max := matrix[0][0]
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if max < matrix[i][j] {
				max = matrix[i][j]
			}
		}
	}
	return max
}

func FindMaxValue(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

func FlattenMatrix(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	res := make([]int, 0)

	for i := 0; i < len(matrix); i++ {
		res = append(res, matrix[i]...)
	}
	return res
}

func CompressString(s string) string {
	count := 1
	res := ""
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] && i != len(s)-1 {
			count++
		} else {
			res += string(s[i-1])
			res += itoa(count)
			count = 1
		}
	}
	return string(res)
}

/********************************************       watch to Learn 		********************************************/
func FindLongestWord(s string) string {
	inword := false
	maxWord := ""
	cursor := 0

	arr := make([]string, 0)

	for i := 0; i <= len(s); i++ {
		if s[i] == ' ' || i == len(s) {
			inword = false
			arr = append(arr, s[cursor:i])
		}
		if !inword && (s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z') {
			inword = true
			cursor = i
		}
	}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if len(maxWord) < len(arr[i]) {
			maxWord = arr[i]
		}
	}
	return maxWord
}

func GCD(a, b int) int {
	res := 1
	for i := 1; i < a; i++ {
		if a%i == 0 && b%i == 0 {
			res = i
		}
		if i == b {
			break
		}
	}
	return res
}

func LCM(a, b int) int {
	res := 0
	for i := 1; i <= b; i++ {
		if (a*i)%b == 0 {
			res = a * i
			break
		}
	}
	return res
}

/****************************************************************************************************************/

func NthPrime(n int) int {
	if n <= 0 {
		return -1
	}
	count := 0
	num := 1
	for count < n {
		num++
		if isPrime(num) {
			count++
		}
	}

	return num
}

func isPrime(n int) bool {
	if n == 1 || n == 0 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func GeneratePrimes(n int) []int {
	if n < 2 {
		return nil
	}
	primes := make([]int, 0)

	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j*j < i; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

func PrimeFactors(n int) []int {
	if n < 0 {
		return nil
	}
	res := make([]int, 0)
	for n != 1 {
		for i := 0; i <= n; i++ {
			if isPrime(i) && n%i == 0 {
				res = append(res, i)
				n /= i
				break
			}
		}
	}
	return res
}

func MagicGrowth3() []string {
	res := make([]string, 0)

	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			for k := j + 1; k < 10; k++ {

				str := string(rune(i+'0')) + string(rune(j+'0')) + string(rune(k+'0'))
				res = append(res, str)
			}
		}
	}

	return res
}

func atoi(str string) int {
	isNegative := false
	var res int
	i := 0
	if str[0] == '-' {
		isNegative = true
		i++
	}

	for i < len(str) {
		if str[i] < '0' || str[i] > '9' {
			return 0
		} else {
			res += int(str[i] - '0')
			if i != len(str)-1 {
				res *= 10
			}
			i++
		}
	}
	if isNegative {
		res = -res
	}
	return res
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	arr := []rune{}
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	for n > 0 {
		arr = append(arr, rune(n%10+'0'))
		n /= 10
	}

	res := make([]rune, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = arr[len(arr)-i-1]
	}
	if isNegative {
		return "-" + string(res)
	}
	return string(res)
}

func UniqComb3(characters string) []string {
	uniq := make(map[rune]bool)

	for _, char := range characters {
		if !uniq[char] {
			uniq[char] = true
		} else {
			return []string{}
		}
	}
	if len(characters) < 3 {
		return []string{}
	}

	res := make([]string, 0)
	for i := 0; i < len(characters); i++ {
		for j := 0; j < len(characters); j++ {
			for k := 0; k < len(characters); k++ {
				if i != j && j != k && k != i {
					res = append(res, string([]byte{characters[i], characters[j], characters[k]}))
				}
			}
		}
	}
	return res
}

func TitleCase(s string) string {
	arr := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		// arr=append(arr, char)
		if s[i] >= 'A' && s[i] <= 'Z' {
			arr[i] = rune(s[i] + 32)
		} else {
			arr[i] = rune(s[i])
		}
	}
	inWord := false
	for i := 0; i < len(arr); i++ {

		if !inWord && toUpper(&arr[i]) {
			inWord = true
		}
		if arr[i] == ' ' {
			inWord = false
		}

	}
	return string(arr)
}

func toUpper(r *rune) bool {
	if *r >= 'a' && *r <= 'z' {
		*r = *r - 32
		return true
	}
	return false
}

func SumOfDigits(n int) int {
	if n < 0 {
		n = -n
	}

	if n == 0 {
		return 0
	}
	sum := n % 10
	return sum + SumOfDigits(n/10)
}

/********************************************       watch to Learn 		********************************************/

func UniquePermutations(s string) []string {
	str := []rune(s)
	result := []string{}
	used := make([]bool, len(str))

	var backtrack func(path []rune)
	backtrack = func(path []rune) {
		if len(path) == len(str) {
			result = append(result, string(path))
			return
		}

		for i := 0; i < len(str); i++ {
			if used[i] {
				continue
			}
			if i > 0 && str[i] == str[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			backtrack(append(path, str[i]))
			used[i] = false
		}
	}

	backtrack([]rune{})

	return result
}
