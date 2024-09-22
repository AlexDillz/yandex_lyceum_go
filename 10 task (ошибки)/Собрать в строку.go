// Напишите функцию ConcatStringsAndInt(str1, str2 string, num int) string,
// которая принимает две строки и одно целое число, а затем выполняет конкатенацию строк и числа в одну строку.

package main

import "strconv"

func ConcatStringsAndInt(str1, str2 string, num int) string {

	sum := str1 + str2 + strconv.Itoa(num)

	return sum
}

// strconv пакет (строки <-> числа):

// strconv.Itoa(int) — преобразует целое число в строку.
// strconv.Atoi(string) — преобразует строку в целое число.
// strconv.FormatFloat(float64, fmt byte, prec int, bitSize int) — преобразует число с плавающей точкой в строку.
// strconv.ParseFloat(string, bitSize int) — преобразует строку в число с плавающей точкой.
// strconv.FormatBool(bool) — преобразует логическое значение в строку.
// strconv.ParseBool(string) — преобразует строку в логическое значение.
