// Напишите функцию AreAnagrams(str1, str2 string) bool, которая проверяет,
// являются ли две заданные строки анаграммами (то есть состоят ли они из одних и тех же символов, хотя и в разном порядке).
// Не учитывайте регистр символов.

// Примечания
// Используйте встроенные функции сортировки и сравнения строк.

package main

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	rune1 := []rune(str1)
	rune2 := []rune(str2)

	sort.Slice(rune1, func(i, j int) bool {
		return rune1[i] < rune1[j]
	})
	sort.Slice(rune2, func(i, j int) bool {
		return rune2[i] < rune2[j]
	})

	return string(rune1) == string(rune2)
}
