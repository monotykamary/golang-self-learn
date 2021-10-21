package main

import "fmt"

var nums1 = []int{2, 7, 11, 15}
var target1 int = 9

// expected output = [0, 1]

var nums2 = []int{3, 2, 4}
var target2 int = 6

// expected output = [1, 2]

var nums3 = []int{3, 3}
var target3 int = 6

// expected output = [0, 1]

func twoSum(nums []int, target int) []int {
	output := []int{}       // output of indexes
	lookup := map[int]int{} // map of complement => index

	for index, num := range nums {
		complement := target - num
		value, exists := lookup[num]

		if exists == true {
			output = append(output, value, index)
		} else {
			lookup[complement] = index
		}
	}

	return output
}

func main() {
	answer1 := twoSum(nums1, target1)
	fmt.Printf("Input: nums = %v\n", nums1)
	fmt.Printf("Output: %v\n", answer1)

	answer2 := twoSum(nums2, target2)
	fmt.Printf("Input: nums = %v\n", nums2)
	fmt.Printf("Output: %v\n", answer2)

	answer3 := twoSum(nums3, target3)
	fmt.Printf("Input: nums = %v\n", nums3)
	fmt.Printf("Output: %v\n", answer3)
}
