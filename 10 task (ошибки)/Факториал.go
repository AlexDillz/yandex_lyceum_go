// Напишите функцию Factorial(n int) (int, error) для нахождения факториала без ошибок.
// Функция получает на вход целое число и выводит его факториал.
// Если пользователь ввёл отрицательное число, программа должна вернуть сообщение об ошибке (factorial is not defined for negative numbers).

// Примечания
// Как и в предыдущих примерах – не забудьте вернуть nil в качестве ошибки, если всё хорошо, и 0 в качестве ответа, если произошла ошибка.

package main

import (
	"errors"
)

func Factorial(n int) (int, error) {
	var result = 1
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result, nil

}
