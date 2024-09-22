// Напишите функцию FindMaxKey(m map[int]int) int,
// которая принимает мапу и возвращает значение наибольшего ключа.

package main

func FindMaxKey(m map[int]int) int {
	// Инициализируем max первым ключом из мапы
	var max int
	for key := range m {
		max = key // Инициализируем max первым попавшимся ключом
		break     // Прерываем цикл после первой итерации
	}

	// Перебираем все остальные ключи
	for key := range m {
		if key > max {
			max = key
		}
	}
	return max
}
