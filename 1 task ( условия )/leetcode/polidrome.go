package main

func isPalindrome(x int) bool {
	// Отрицательные числа не могут быть палиндромами
	if x < 0 {
		return false
	}

	// Разворачиваем число
	original := x
	reversed := 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// Сравниваем оригинальное число с перевернутым
	return original == reversed
}
