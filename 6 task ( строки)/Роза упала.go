// Напишите функцию IsPalindrome(input string) bool,
// которая принимает строку и проверяет, является ли она палиндромом

package main

import (
	"strings"
)

func IsPalindrome(input string) bool {

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "") // убираем пробелы

	// Проверяем, является ли строка палиндромом
	length := len(input)
	for i := 0; i < length/2; i++ {
		if input[i] != input[length-1-i] {
			return false
		}
	}
	return true
}

// package main

// import (
// 	"fmt"
// )

// func isPalindrome(x int) bool {
//     // Отрицательные числа не могут быть палиндромами
// 	if x < 0 {
// 		return false
// 	}

// 	// Разворачиваем число
// 	original := x
// 	reversed := 0

// 	for x > 0 {
// 		reversed = reversed*10 + x%10
// 		x /= 10
// 	}

// 	// Сравниваем оригинальное число с перевернутым
// 	return original == reversed
// }
