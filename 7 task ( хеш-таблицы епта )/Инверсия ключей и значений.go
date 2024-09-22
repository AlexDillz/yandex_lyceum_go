// Напишите функцию SwapKeysAndValues(m map[string]string) map[string]string,
// которая принимает мапу и возвращает новую мапу, где ключи и значения поменялись местами.

package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	// Создаем новую мапу для хранения результата
	newMap := make(map[string]string)

	// Перебираем оригинальную мапу
	for key, value := range m {
		// Меняем местами ключ и значение
		newMap[value] = key
	}

	// Возвращаем новую мапу с измененными ключами и значениями
	return newMap
}
