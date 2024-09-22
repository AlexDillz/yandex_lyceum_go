// Дан неотсортированный слайс целых чисел. Напишите функцию UnderLimit(nums []int, limit int, n int) ([]int, error),
// которая будет возвращать первые n (либо меньше, если остальные не подходят) элементов, которые меньше limit.
// В случае ошибки функция должна вернуть nil и описание ошибки.

package main

import (
	"fmt"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n should be positive")
	}

	var result []int

	for _, num := range nums {
		if num < limit {
			result = append(result, num)
		}
		if len(result) == n {
			break
		}
	}

	return result, nil
}

func main() {
	nums := []int{3, 7, 1, 9, 4}
	limit := 5
	n := 2
	result, err := UnderLimit(nums, limit, n)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
