/*
 * @lc app=leetcode id=1089 lang=golang
 *
 * [1089] Duplicate Zeros
 *
 * https://leetcode.com/problems/duplicate-zeros/description/
 *
 * algorithms
 * Easy (54.44%)
 * Likes:    468
 * Dislikes: 185
 * Total Accepted:    56.4K
 * Total Submissions: 103.8K
 * Testcase Example:  '[1,0,2,3,0,4,5,0]'
 *
 * Given a fixed length array arr of integers, duplicate each occurrence of
 * zero, shifting the remaining elements to the right.
 *
 * Note that elements beyond the length of the original array are not written.
 *
 * Do the above modifications to the input array in place, do not return
 * anything from your function.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,0,2,3,0,4,5,0]
 * Output: null
 * Explanation: After calling your function, the input array is modified to:
 * [1,0,0,2,3,0,0,4]
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3]
 * Output: null
 * Explanation: After calling your function, the input array is modified to:
 * [1,2,3]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= arr.length <= 10000
 * 0 <= arr[i] <= 9
 *
 */

// @lc code=start
func duplicateZeros(arr []int) {
	duplicateZeros1(arr)
}

func duplicateZeros1(arr []int) {
	zeros := 0

	for _, val := range arr {
		if val == 0 {
			zeros++
		}
	}

	dest := make([]int, len(arr)+zeros, len(arr)+zeros)
	dIndex := 0
	for _, val := range arr {
		if val == 0 {
			// copy zero twice
			dest[dIndex] = 0
			dIndex++
			dest[dIndex] = 0
		} else {
			dest[dIndex] = val
		}
		dIndex++
	}
	copy(arr[0:], dest[0:len(arr)])
}

// @lc code=end
