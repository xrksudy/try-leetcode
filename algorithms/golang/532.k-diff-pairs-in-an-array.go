import "sort"

/*
 * @lc app=leetcode id=532 lang=golang
 *
 * [532] K-diff Pairs in an Array
 *
 * https://leetcode.com/problems/k-diff-pairs-in-an-array/description/
 *
 * algorithms
 * Easy (31.64%)
 * Likes:    645
 * Dislikes: 1357
 * Total Accepted:    108.2K
 * Total Submissions: 342K
 * Testcase Example:  '[3,1,4,1,5]\n2'
 *
 *
 * Given an array of integers and an integer k, you need to find the number of
 * unique k-diff pairs in the array. Here a k-diff pair is defined as an
 * integer pair (i, j), where i and j are both numbers in the array and their
 * absolute difference is k.
 *
 *
 *
 * Example 1:
 *
 * Input: [3, 1, 4, 1, 5], k = 2
 * Output: 2
 * Explanation: There are two 2-diff pairs in the array, (1, 3) and (3,
 * 5).Although we have two 1s in the input, we should only return the number of
 * unique pairs.
 *
 *
 *
 * Example 2:
 *
 * Input:[1, 2, 3, 4, 5], k = 1
 * Output: 4
 * Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3,
 * 4) and (4, 5).
 *
 *
 *
 * Example 3:
 *
 * Input: [1, 3, 1, 5, 4], k = 0
 * Output: 1
 * Explanation: There is one 0-diff pair in the array, (1, 1).
 *
 *
 *
 * Note:
 *
 * The pairs (i, j) and (j, i) count as the same pair.
 * The length of the array won't exceed 10,000.
 * All the integers in the given input belong to the range: [-1e7, 1e7].
 *
 *
 */

// @lc code=start
func findPairs(nums []int, k int) int {
	return findPairs3(nums, k)
}

func findPairs3(nums []int, k int) int {
	if k < 0 || len(nums) <= 1 {
		return 0
	}

	sort.Ints(nums)
	retVal := 0
	for begin, end := 0, 1; end < len(nums); {
		if begin >= end || nums[begin]+k > nums[end] {
			end++
		} else if begin > 0 && nums[begin] == nums[begin-1] || nums[begin]+k < nums[end] {
			// begin
			//  |
			// [1, 1, ...., 8, 8]
			//              |
			//             end
			begin++
		} else {
			begin++
			retVal++
		}
	}
	return retVal
}

// [1,2,3,4,5], -1
func findPairs2(nums []int, k int) int {
	if k < 0 || len(nums) <= 1 {
		return 0
	}

	hash := map[int]int{}
	for _, v := range nums {
		hash[v]++
	}

	retVal := 0
	for hKey, hValue := range hash { // iterator hash
		if k == 0 {
			if hValue >= 2 {
				retVal++
			}
		} else {
			if hash[hKey+k] > 0 {
				retVal++
			}
		}
	}
	return retVal
}

func findPairs1(nums []int, k int) int {
	hash := map[int]int{}
	for _, v := range nums {
		hash[v]++
	}

	visited, retVal := map[int]bool{}, 0
	for _, v := range nums {
		smallVal := v - k
		if !visited[v+smallVal] && ((k > 0 && hash[smallVal] > 0) || (k == 0 && hash[smallVal] > 1)) {
			retVal++
			visited[v+smallVal] = true
		}

		bigVal := v + k
		if !visited[v+bigVal] && ((k > 0 && hash[bigVal] > 0) || (k == 0 && hash[bigVal] > 1)) {
			retVal++
			visited[v+bigVal] = true
		}
	}
	return retVal
}

// @lc code=end