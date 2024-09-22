// Напишите функцию IsPalindrome(input string) bool, 
// которая принимает строку и проверяет, является ли она палиндромом

package main
pachage string

import (
	"strings"
)

func  IsPalindrome(input string) bool {

firstLetter := input[0]


	for _, r := range input {
		if firstLetter == input[i] {
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