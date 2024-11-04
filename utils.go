package main

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
