package main

import (
	"fmt"
	"sort"
)

func threeSumTarget(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var result [][]int

	for i := 0; i < n-2; i++ {

		// Skip duplicate fixed elements
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--

			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 9
	fmt.Println(threeSumTarget(nums, target))
}
