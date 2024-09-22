func romanToInt(s string) int {
	// Словарь римских символов и их значений
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	n := len(s)

	// Проходим по строке
	for i := 0; i < n; i++ {
		value := roman[s[i]]

		// Проверяем, есть ли следующий символ и больше ли он текущего
		if i+1 < n && roman[s[i+1]] > value {
			result -= value
		} else {
			result += value
		}
	}

	return result
}
