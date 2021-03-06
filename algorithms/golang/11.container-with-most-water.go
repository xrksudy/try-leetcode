/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (50.24%)
 * Likes:    6372
 * Dislikes: 595
 * Total Accepted:    678.1K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * Given n non-negative integers a1, a2, ..., an , where each represents a
 * point at coordinate (i, ai). n vertical lines are drawn such that the two
 * endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together
 * with x-axis forms a container, such that the container contains the most
 * water.
 *
 * Note: You may not slant the container and n is at least 2.
 *
 *
 *
 *
 *
 * The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In
 * this case, the max area of water (blue section) the container can contain is
 * 49.
 *
 *
 *
 * Example:
 *
 *
 * Input: [1,8,6,2,5,4,8,3,7]
 * Output: 49
 */

// @lc code=start
func maxArea(height []int) int {
	return maxArea2(height)
}

func maxArea2(height []int) int {
	max, left, right := 0, 0, len(height)-1
	for left < right {
		minVal, interval := height[left], right-left
		if height[right] < minVal {
			minVal = height[right]
		}
		total := minVal * interval
		if total > max {
			max = total
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func maxArea1(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			minVal, interval := height[i], j-i
			if height[j] < minVal {
				minVal = height[j]
			}
			total := minVal * interval
			if total > max {
				max = total
			}
		}
	}
	return max
}

// @lc code=end