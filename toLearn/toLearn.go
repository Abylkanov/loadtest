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

func UniquePermutations(s string) []string {
	// Преобразуем строку в срез рун (чтобы работать с символами Unicode)
	str := []rune(s)
	result := []string{}
	used := make([]bool, len(str))

	// Вспомогательная рекурсивная функция для генерации перестановок
	var backtrack func(path []rune)
	backtrack = func(path []rune) {
		// Если длина текущего пути равна длине строки, добавляем результат
		if len(path) == len(str) {
			result = append(result, string(path))
			return
		}

		// Перебираем все символы строки
		for i := 0; i < len(str); i++ {
			// Если символ уже использован в текущей перестановке, пропускаем
			if used[i] {
				continue
			}

			// Пропускаем одинаковые символы, чтобы избежать дублирования перестановок
			// Проверяем, если символ одинаков с предыдущим и предыдущий не был использован
			if i > 0 && str[i] == str[i-1] && !used[i-1] {
				continue
			}

			// Помечаем символ как использованный
			used[i] = true
			// Добавляем символ в текущую перестановку
			backtrack(append(path, str[i]))
			// Отменяем пометку использованности
			used[i] = false
		}
	}

	// Начинаем рекурсию с пустым путем
	backtrack([]rune{})

	return result
}
