/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
 *
 * https://leetcode.com/problems/invert-binary-tree/description/
 *
 * algorithms
 * Easy (64.34%)
 * Likes:    3577
 * Dislikes: 58
 * Total Accepted:    555.4K
 * Total Submissions: 854K
 * Testcase Example:  '[4,2,7,1,3,6,9]'
 *
 * Invert a binary tree.
 *
 * Example:
 *
 * Input:
 *
 *
 * ⁠    4
 * ⁠  /   \
 * ⁠ 2     7
 * ⁠/ \   / \
 * 1   3 6   9
 *
 * Output:
 *
 *
 * ⁠    4
 * ⁠  /   \
 * ⁠ 7     2
 * ⁠/ \   / \
 * 9   6 3   1
 *
 * Trivia:
 * This problem was inspired by this original tweet by Max Howell:
 *
 * Google: 90% of our engineers use the software you wrote (Homebrew), but you
 * can’t invert a binary tree on a whiteboard so f*** off.
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	return invertTree2(root)
}

// iterative
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curNode.Left, curNode.Right = curNode.Right, curNode.Left
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
	}
	return root
}

// recursive
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftNode := invertTree(root.Left)
	rightNode := invertTree(root.Right)
	root.Left = rightNode
	root.Right = leftNode
	return root
}

// @lc code=end