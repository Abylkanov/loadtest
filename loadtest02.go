package main

import (
"github.com/alem-platform/ap"
"fmt"
)

func importS(){
	ap.PutRune('a')
	fmt.Println(`test`)
}

func CountA(slice []string) int {
	count:= 0
	for _, str:=range slice {
		for _,ch:=range str {
			if ch == 'a' {
				count++
			}
		}
	}
return count
}

func FindFirstA(slice []string) string {
	for _, str:=range slice {
		for _,ch:=range str {
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
	s= s[len(s)-len(suffix):]
	for i:=len(suffix)-1; i>=0;i-- {
		if suffix[i]!= s[i] {
			return false
		}
	}
	return true
}

func CommonElements(slice1, slice2, slice3 []int) []int {
	result:=make([]int,0)
	commonEl:= make(map[int]bool)

	for i:=0; i < len(slice1); i++ {

		for j:=0; j < len(slice2); j++ {
			for k:=0; k<len(slice3); k++ {
				if slice1[i] == slice2[j] && slice1[i] == slice3[k] {
					commonEl[slice1[i]]=true
				}
			}	
		}
	}
	for key,_:=range commonEl {
		result = append(result,key)
	}
	return result
}