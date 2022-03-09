package array

import "fmt"

func TwoSumTest() {
	fmt.Println("\nTwo Sum:")
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(nums, target, TwoSum(nums, target))
	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(nums, target, TwoSum(nums, target))
	nums = []int{3, 3}
	target = 6
	fmt.Println(nums, target, TwoSum(nums, target))
}

func TwoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, val := range nums {
		if index, ok := m[target-val]; ok {
			return []int{index, i}
		}
		m[val] = i
	}

	return nums
}
