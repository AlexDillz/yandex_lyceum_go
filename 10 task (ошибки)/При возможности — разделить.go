// Напишите функцию DivideIntegers(a, b int) (float64, error), которая возвращает результат деления первого числа на второе.
// Если второе число равно нулю, функция должна возвращать в качестве ответа 0 и сообщение об ошибке (division by zero is not allowed).
// Если второе число не равно нулю – верните в качестве ошибки nil.

package main

import "errors"

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return float64(a) / float64(b), nil //тк надо вернуть float

}
