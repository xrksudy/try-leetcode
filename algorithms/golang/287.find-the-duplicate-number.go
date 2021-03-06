import "sort"

/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 *
 * https://leetcode.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (53.94%)
 * Likes:    4767
 * Dislikes: 567
 * Total Accepted:    346.3K
 * Total Submissions: 628.1K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * Given an array nums containing n + 1 integers where each integer is between
 * 1 and n (inclusive), prove that at least one duplicate number must exist.
 * Assume that there is only one duplicate number, find the duplicate one.
 *
 * Example 1:
 *
 *
 * Input: [1,3,4,2,2]
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [3,1,3,4,2]
 * Output: 3
 *
 * Note:
 *
 *
 * You must not modify the array (assume the array is read only).
 * You must use only constant, O(1) extra space.
 * Your runtime complexity should be less than O(n^2).
 * There is only one duplicate number in the array, but it could be repeated
 * more than once.
 *
 *
 */

// @lc code=start
// 官方的二分查找，位查找方法，快慢环的查找方式还没理解
func findDuplicate(nums []int) int {
	return findDuplicate2(nums)
}

// 采用hashSet方式，判断是否重复(时间复杂度O(n), 空间复杂度O(n))
func findDuplicate2(nums []int) int {
	hash := make(map[int]bool)
	for _, v := range nums {
		if _, ok := hash[v]; ok {
			return v
		}
		hash[v] = true
	}
	return -1
}

// 先排序，再变量比较n与n+1的数(时间复杂度O(nlog(n)), 空间复杂度O(1)或者O(n))
func findDuplicate1(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

// @lc code=end