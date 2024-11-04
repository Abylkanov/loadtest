package toLearn

func FindLongestWord(s string) string {
	maxWord := ""
	current := ""

	for i := 0; i <= len(s); i++ {
		// Проверяем конец строки или пробел
		if i == len(s) || s[i] == ' ' {
			if len(current) > len(maxWord) {
				maxWord = current // Обновляем самое длинное слово
			}
			current = "" // Сбрасываем текущее слово
		} else {
			current += string(s[i]) // Собираем текущее слово
		}
	}

	return maxWord
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0 // НОК не существует, если одно из чисел 0
	}
	return (a * b) / GCD(a, b)
}

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
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
