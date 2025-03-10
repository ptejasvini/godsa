/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3, 2, 4], target = 6
Output: [1, 2]

Example 3:

Input: nums = [3, 3], target = 6
Output: [0, 1]

Constraints:

2 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
Only one valid answer exists.
*/

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
	fmt.Println(twoSum(nums, target))

	nums2 := []int{3,2,4}
	target2 := 6
	twoSum(nums2, target2)
	fmt.Println(twoSum(nums2, target2))

	nums3 := []int{3,3}
	target3 := 6
	twoSum(nums3, target2)
	fmt.Println(twoSum(nums3, target3))

}

func twoSum(nums []int, target int) []int {
	var res = []int{}

	var m = make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if _, ok := m[diff]; ok {
			
			res = append(res, m[diff])
			res = append(res, i)
		} else {
			m[num] = i
		}
	}
	return res
}
