package main

import (
	"fmt"
	"os"
	"github.com/alem-platform/ap"
)

func importS3() {
	ap.PutRune('a')
	fmt.Println(`test`)
}

func MapContains(m map[string]int, key string) bool {
	if _, exist := m[key]; !exist {
		return false
	}
	return true
}

func MapSet(m map[string]int, key string, value int) {
	m[key] = value
}

func MapDelete(m map[string]int, key string) {
	delete(m, key)
}

func MapSum(m map[string]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}

func MapMax(m map[string]int) string {
	if len(m) == 0 {
		return ""
	}

	var res string
	var max int
	for key, value := range m {
		if res == "" {
			max = value
			res = key
		} else if value > max {
			max = value
			res = key
		}
	}
	return res
}

func MapIntersection(m1, m2 map[string]int) map[string]int {
	res:= make(map[string]int)

	for key, value := range m1 {
		exist:=false
		if value, exist = m2[key]; exist {
			res[key]=value
		}
	}
	return res
}

func MapDifference(m1, m2 map[string]int) map[string]int {
	res := make(map[string]int)
	
	for key, value := range m1 { 
		if _, exist := m2[key]; !exist {
			res[key] = value
		}
	}
	return res
}

func MapInvert(m map[string]int) map[int]string {

	res:=make(map[int]string)

	for key,value := range m {
		res[value]=key
	}
	return res
}

func printArgs(){
	args:=os.Args

	for i:=1; i<len(args); i++ {
		printString(itoa(i)+ ": " + args[i] +"\n")
	}
	
}
func printString(s string) {
	for _,char:= range s {
		ap.PutRune(char) 
	}
}

func countVowels() {
	args:=os.Args[1:]
	vowels:= []rune{'A', 'E','U', 'I', 'O'}
	for i:=0; i<len(args); i++ {
		count:=0
		for _, char:= range args[i] {
			for _, vowel:= range vowels {
				if char==vowel || char==vowel+32 {
					count++
				} 
			}
		}
		printString(itoa(count)+"\n")
	}

}

func uppercaseArguments(){
	args := os.Args[1:]

	for _, arg := range args {
		for i:=0; i<len(arg); i++ {
			if arg[i] >='a' && arg[i] <='z' {
				ap.PutRune(rune(arg[i]-32))
			} else {
				ap.PutRune(rune(arg[i]))
			}
		}
		ap.PutRune('\n')
	}

}

func reverseEachArgument() {
	args:=os.Args[1:]
	
	for _, arg:=range args {
		for i:=len(arg)-1; i>=0; i-- {
			ap.PutRune(rune(arg[i]))
		}
		ap.PutRune('\n')
	}
}

func sumNumbers() {
	args:= os.Args[1:]
	sum:=0

	for _, arg:= range args {
		if _, err:=atoiFull(arg); !err {
			return
		}
	}
	for _, arg:=range args {
		num,_ :=atoiFull(arg)
		sum = sum+num
		printString(itoa(sum)+"\n")
	}
	

}

func atoiFull(s string) (int, bool) {

	var res int
	isNegative:=false
	i:=0
	if len(s) == 0 {
		return 0, false
	}
	if s[0] == '-' {
		isNegative= true
		i++
	}

	for i<len(s) {
		if s[i] <'0' || s[i] > '9'{
			return 0, false
		}
		res = res*10 + int(s[i]-'0')
		i++
	}
	if isNegative {
		res=-res
	}
	return res, true

}

func product_numbers() {
	args:=os.Args[1:]
	product:=1
	for _, arg:=range args {
		if _,ok:= (atoiFull(arg)); !ok {
			return
		}
	}
	for _, arg:=range args {
		num,_:=atoiFull(arg)
		product=product*num
		printString(itoa(product)+"\n")
	}
}

func powerCalculation() {
	args:= os.Args[1:]
	if len(args) !=2 {
		return
	}
	num, ok1:=atoiFull(args[0])
	power,ok2:=atoiFull(args[1])
	if !ok1 || !ok2 {
		return
	}
	res:=1
	for i:=0; i<power; i++ {
		res*=num
	}
	printString(itoa(res)+"\n")
}

func GetMathOperation(op string) *func(int, int) int {
	var f func(int, int) int

	switch op {
	case "add":
		f = func(a,b int) int {
			return a+b
		}
	case "subtract":
		f = func(a,b int) int {
			return a-b
		}
	case "multiply":
		f = func(a,b int) int {
			return a*b
	}
	case "divide":
		f = func(a,b int) int {
			if b ==0 {
				return 0
			}
			return a/b
		}
	default: 
		return nil
	}
	return &f
}

func GetIncrementor(start int, step int) func() int {
	var f func() int
	f = func() int {
		start+=step
		return start
	}
	return f
}

func MapFromKeys(keys []string, f func(string) int) map[string]int {
	res:=make(map[string]int)
	for _, key:= range keys {
		res[key]=f(key)
	}
	return res
}

func MapUpdate(m map[string]int, key string, f func(int) int) {

	for k,value:= range m {
		if k == key {
			m[k]=f(value)
		}
	}
}

func Sort(arr []int, fn func(int, int) bool) {

	for i:=0;i<len(arr)-1; i++ {
		for j:=0; j<len(arr)-i-1; j++ {
			if !fn(arr[j],arr[j+1]) {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}

}

func printHex() {
	args:=os.Args[1:]
	if len(args) != 1 {
		return
	}
	for _, char := range args[0] {
		if char < '0' || char > '9' {
			return
		}
	}
	num,_ :=atoiFull(args[0])
	fmt.Println(num)
	base16:="0123456789abcdef"
	var runes []rune
	for num >0 {
		runes = append(runes, rune(base16[num%16]))
		num/=16
	}
	ap.PutRune('0')
	ap.PutRune('x')
	for i:=len(runes)-1; i>=0; i-- {
		ap.PutRune(runes[i])
	}
	ap.PutRune('\n')
}
/********************************************       watch to Learn 		********************************************/

func primeSum() {
	args:=os.Args[1:]
	if len(args) != 1 {
		return
	}
	for _, char := range args[0] {
		if char < '0' || char > '9' {
			return
		}
	}
	num,_ :=atoiFull(args[0])

	res:=make([]int,0)
	if num < 2 {
		return
	}

	
	// Получаем список всех простых чисел до num
	primes := []int{}
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	// Основной цикл: разлагаем число на простые числа
	for num > 0 {
		for _, prime := range primes {
			// Если простое число меньше или равно текущему числу, вычитаем его
			if prime <= num {
				res = append(res, prime)
				num -= prime
				break
			}
		}
	}

	// Выводим результат
	fmt.Println(res)
}

/***************************************************************************************************************************/


func RotateMatrix180(matrix [][]rune) {


	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return
	}

	arr:=make([]rune,0)

	for i:=0; i<len(matrix);i++ {
		for j:=0; j<len(matrix[0]); j++ {
			arr=append(arr,matrix[i][j])
		}
	}
	for i:=0; i<len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] =  arr[len(arr)-1-i], arr[i]
	}

	for i:=0; i<len(matrix);i++ {
		for j:=0; j<len(matrix[0]); j++ {
			matrix[i][j] = arr[i*len(matrix)+j]

		}
	}
}
