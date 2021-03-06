/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 *
 * https://leetcode.com/problems/rotate-array/description/
 *
 * algorithms
 * Easy (34.19%)
 * Likes:    2775
 * Dislikes: 809
 * Total Accepted:    486.8K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * Given an array, rotate the array to the right by k steps, where k is
 * non-negative.
 *
 * Follow up:
 *
 *
 * Try to come up as many solutions as you can, there are at least 3 different
 * ways to solve this problem.
 * Could you do it in-place with O(1) extra space?
 *
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,4,5,6,7], k = 3
 * Output: [5,6,7,1,2,3,4]
 * Explanation:
 * rotate 1 steps to the right: [7,1,2,3,4,5,6]
 * rotate 2 steps to the right: [6,7,1,2,3,4,5]
 * rotate 3 steps to the right: [5,6,7,1,2,3,4]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [-1,-100,3,99], k = 2
 * Output: [3,99,-1,-100]
 * Explanation:
 * rotate 1 steps to the right: [99,-1,-100,3]
 * rotate 2 steps to the right: [3,99,-1,-100]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * It's guaranteed that nums[i] fits in a 32 bit-signed integer.
 * k >= 0
 *
 *
 */

// @lc code=start
func rotate(nums []int, k int) {
	rotate2(nums, k)
}

func rotate2(nums []int, k int) {
	if k != 0 {
		copy(nums, append(nums[len(nums)-k%len(nums):], nums[0:len(nums)-k%len(nums)]...))
	}
}

func rotate1(nums []int, k int) {
	if k == 0 {
		return
	}
	for k > 0 {
		tmp := nums[len(nums)-1]
		// nums[1:len(nums)] = nums[0 : len(nums)-1]
		copy(nums[1:len(nums)], nums[0:len(nums)-1])
		nums[0] = tmp
		k--
	}
}

// [1,2,3,4,5,6,7]\n3

// @lc code=end

