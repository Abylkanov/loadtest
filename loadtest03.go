package main

import (
	"fmt"

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
