// Напишите функцию GetCharacterAtPosition(str string, position int) (rune, error) для робота-помощника,
// которая получает на вход строку и целое число.
// Функция должна возвращать символ строки, который находится на позиции, указанной пользователем (и nil в качестве ошибки).
// Если пользователь ввёл число, которое выходит за пределы длины строки,
// функция должна возвращать в качестве ответа нулевую руну (0) сообщение об ошибке (position out of range).

// Примечания
// Как настоящий программист Гоша начинает нумерацию с 0.

// И напоминаем про руны

package main

import "errors"

func GetCharacterAtPosition(str string, position int) (rune, error) {

	if position >= len(str) || position < 0 {
		return 0, errors.New("position out of range")
	}

	// Преобразуем строку в руны для корректной работы с многобайтовыми символами
	runes := []rune(str)

	// Проверяем, что позиция не превышает количество рун
	if position >= len(runes) {
		return 0, errors.New("position out of range")
	}

	// Возвращаем руну по заданной позиции
	return runes[position], nil
}
