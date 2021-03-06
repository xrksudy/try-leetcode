/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 *
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (42.50%)
 * Likes:    2434
 * Dislikes: 256
 * Total Accepted:    646.1K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * Given a sorted array and a target value, return the index if the target is
 * found. If not, return the index where it would be if it were inserted in
 * order.
 *
 * You may assume no duplicates in the array.
 *
 * Example 1:
 *
 *
 * Input: [1,3,5,6], 5
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [1,3,5,6], 2
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: [1,3,5,6], 7
 * Output: 4
 *
 *
 * Example 4:
 *
 *
 * Input: [1,3,5,6], 0
 * Output: 0
 *
 *
 */

// 不通过的测试case: [1,3]\n2; [1]\n1

// @lc code=start
func searchInsert(nums []int, target int) int {
	return searchInsert2(nums, target)
}

// use binary search, careful boundary, time space: O(log(n))
func searchInsert2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

// use for iterate, time space: O(n)
func searchInsert1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}

		if target < nums[i] {
			return i
		}
	}
	return len(nums)
}

// @lc code=end