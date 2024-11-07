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
func CoinsCombinations(sum int) [][]int {
    // Доступные номиналы монет
    coins := []int{100, 50, 10, 5, 1}
    
    // Результирующий срез для всех комбинаций
    var result [][]int
    
    // Рекурсивная вспомогательная функция
    var findCombinations func(target int, current []int, start int)
    findCombinations = func(target int, current []int, start int) {
        // Если сумма достигла 0, то добавляем текущую комбинацию в результат
        if target == 0 {
            // Создаём копию текущей комбинации, чтобы избежать изменений при дальнейших рекурсиях
            combination := make([]int, len(current))
            copy(combination, current)
            result = append(result, combination)
            return
        }
        
        // Проходим по монетам начиная с индекса start, чтобы избежать дублирования
        for i := start; i < len(coins); i++ {
            coin := coins[i]
            
            // Если текущая монета не превышает остаток суммы
            if coin <= target {
                // Добавляем монету в текущую комбинацию
                current = append(current, coin)
                
                // Рекурсивно продолжаем с уменьшенной суммой
                findCombinations(target-coin, current, i)
                
                // Убираем монету из текущей комбинации, чтобы попробовать другие варианты
                current = current[:len(current)-1]
            }
        }
    }
    
    // Инициализируем поиск с полной суммой и пустой комбинацией
    findCombinations(sum, []int{}, 0)
    
    return result
}

func isUnique(s string) bool {
	chars := make(map[rune]bool)
	for _, c := range s {
		if chars[c] {
			return false
		}
		chars[c] = true
	}
	return true
}

func generateCombinations(s []rune, n int, current []rune, start int, result *[]string) {
	if len(current) == n {
		*result = append(*result, string(current))
		return
	}

	for i := start; i < len(s); i++ {
		current = append(current, s[i])
		generateCombinations(s, n, current, i+1, result)
		current = current[:len(current)-1] 
	}
}

// Основная функция, которая генерирует все уникальные комбинации длины n
func UniqCombN(s string, n int) []string {
	// Если строка имеет дубликаты или невозможно сформировать комбинации, возвращаем пустой срез
	if !isUnique(s) || n > len(s) {
		return []string{}
	}

	// Преобразуем строку в срез рун (для корректной работы с Unicode)
	runes := []rune(s)
	var result []string

	// Генерируем все комбинации длины n
	generateCombinations(runes, n, []rune{}, 0, &result)

	return result
}


