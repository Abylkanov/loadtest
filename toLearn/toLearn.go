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
