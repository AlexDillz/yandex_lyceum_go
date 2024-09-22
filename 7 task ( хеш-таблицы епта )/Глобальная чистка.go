// Помогите Гоше почистить телефонную книгу. Создайте функцию DeleteLongKeys(m map[string]int) map[string]int,
// которая принимает мапу и возвращает новую мапу,
// где присутствуют только те ключи, у которых длинна не меньше 6 символов.

package main

func DeleteLongKeys(m map[string]int) map[string]int {
	newMap := make(map[string]int)

	for key, value := range m {
		if len(key) >= 6 {
			newMap[key] = value
		}
	}
	return newMap
}
