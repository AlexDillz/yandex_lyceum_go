// Напишите функцию SumTwoIntegers(a, b string) (int, error), которая получает две строки и, если это целые числа, возвращает их сумму.
// Если пользователь ввёл данные не целого типа, функция должна вернуть сообщение об ошибке (invalid input, please provide two integers).

package main

import (
	"errors"
	"strconv"
)

func SumTwoIntegers(a, b string) (int, error) {

	A, errA := strconv.Atoi(a)
	if errA != nil {
		return 0, errors.New("invalid input, please provide two integers")
	}

	B, errB := strconv.Atoi(b)
	if errB != nil {
		return 0, errors.New("invalid input, please provide two integers")
	}
	sum := A + B
	return sum, nil
}
