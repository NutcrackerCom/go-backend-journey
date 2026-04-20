package main

import "fmt"

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int {

	var mDiff map[int]int = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if val, ok := mDiff[diff]; ok {
			return []int{val, i}
		} else {
			mDiff[nums[i]] = i
		}
	}
	return []int{}
}

func main() {
	var nums []int = []int{2, 7, 11, 15}
	var target int = 9
	fmt.Println(twoSum(nums, target))
}
