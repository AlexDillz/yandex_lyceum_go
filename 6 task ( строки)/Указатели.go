// Напишите функцию isLatin(input string) bool,
// которая принимает строку и выводит true, если все символы в строке латинские, false, если нет.

package main

import "unicode"

func isLatin(input string) bool {

	for _, r := range input {
		if !unicode.Is(unicode.Latin, r) {
			return false // Если хотя бы один символ не латинский, возвращаем false
		}
	}
	return true // Все символы латинские, возвращаем true
}
