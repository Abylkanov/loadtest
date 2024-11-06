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