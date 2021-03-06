/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (35.81%)
 * Likes:    3473
 * Dislikes: 147
 * Total Accepted:    505K
 * Total Submissions: 1.4M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * If the target is not found in the array, return [-1, -1].
 *
 * Example 1:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 *
 * Example 2:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * nums is a non decreasing array.
 * -10^9 <= target <= 10^9
 *
 *
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	return searchRange2(nums, target)
}

// 摘抄leetcode-cn上一位同学的例子
func searchRange2(nums []int, target int) []int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
		return []int{-1, -1}
	}

	a := binarySearch(nums, target)
	b := binarySearch(nums, target+1) // 巧用target+1的逻辑

	if a == len(nums) || nums[a] != target { // 表明数组中不存在target
		return []int{-1, -1}
	}

	return []int{a, b - 1} // 数组中存在target, b-1是最后一个target的元素索引
}

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

// 直接常规遍历查询O(n)时间复杂度
func searchRange1(nums []int, target int) []int {
	ret := []int{-1, -1}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			ret[0] = i
			break
		}
	}

	if ret[0] == -1 {
		return ret
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			ret[1] = i
			break
		}
	}
	return ret
}

// @lc code=end