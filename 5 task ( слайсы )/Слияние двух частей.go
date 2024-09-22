// Дан слайс nums, состоящий из 2n элементов в формате [x0,x1,...,xn,y0,y1,...,yn].
// Создайте функцию Mix(nums []int) []int, которая вернёт слайс,
// содержащий значения в следующем порядке: [x0,y0,x1,y1,...,xn,yn].

package main

func Mix(nums []int) []int {

	res := make([]int, len(nums))

	n := len(res) / 2

	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[n+i]

	}

	return res

}
