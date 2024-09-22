// Гоша презирает гадания и пользуется генератором случайных чисел, чтобы создавать предсказуемые случайности.
// Воспользуйтесь им и напишите функцию FindMinMaxInArray(array [10]int) (int, int), которая получает массив
// из десяти случайных целочисленных значений.
// Далее функция должна находить максимальное и минимальное значения в массиве и возвращать их.

package main

func FindMinMaxInArray(array [10]int) (int, int) {
	// Инициализируем min и max первым элементом массива
	min := array[0]
	max := array[0]

	// Проходим по всему массиву, начиная со второго элемента
	for i := 1; i < len(array); i++ {
		if array[i] < min {
			min = array[i] // Обновляем min, если нашли меньшее значение
		}
		if array[i] > max {
			max = array[i] // Обновляем max, если нашли большее значение
		}
	}

	// Возвращаем найденные min и max
	return max, min
}

// package main

// func FindMinMaxInArray(array [10]int) (int, int) {
//     min := array[0]
//     max := array[0]

//     for _, value := range array {
//         if value < min {
//             min = value
//         }
//         if value > max {
//             max = value
//         }
//     }

//     return max, min
// }
