// Напишите функцию IntToBinary(num int) (string, error), которая принимает на вход целое число и возвращает его двоичное представление.
// Если пользователь ввёл отрицательное число, программа должна возвращать сообщение об ошибке (negative numbers are not allowed).

// Примечания
// Обратите внимание, что функция возвращает строку, а не число. В случае ошибки верните пустую строку.

// Много стандартных операций уже реализовано во внешних пакетах.

package main

import (
	"errors"
	"strconv"
)

func IntToBinary(num int) (string, error) {

	if num < 0 {
		return "", errors.New("negative numbers are not allowed")
	}
	doublevision := strconv.FormatInt(int64(num), 2)
	return doublevision, nil
}
