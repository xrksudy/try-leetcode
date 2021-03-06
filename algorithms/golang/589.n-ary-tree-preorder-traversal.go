/*
 * @lc app=leetcode id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 *
 * https://leetcode.com/problems/n-ary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (71.53%)
 * Likes:    620
 * Dislikes: 61
 * Total Accepted:    109.2K
 * Total Submissions: 151.5K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * Given an n-ary tree, return the preorder traversal of its nodes' values.
 *
 * Nary-Tree input serialization is represented in their level order traversal,
 * each group of children is separated by the null value (See examples).
 *
 *
 *
 * Follow up:
 *
 * Recursive solution is trivial, could you do it iteratively?
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: root = [1,null,3,2,4,null,5,6]
 * Output: [1,3,5,6,2,4]
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: root =
 * [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
 * Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]
 *
 *
 *
 * Constraints:
 *
 *
 * The height of the n-ary tree is less than or equal to 1000
 * The total number of nodes is between [0, 10^4]
 *
 *
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	return preorder2(root)
}

// using iterator, stack
func preorder2(root *Node) []int {
	retVals := []int{}
	if root == nil {
		return retVals
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		retVals = append(retVals, cur.Val)
		stack = stack[:len(stack)-1]
		childs := cur.Children
		for i := len(childs) - 1; i >= 0; i-- {
			stack = append(stack, childs[i])
		}
	}
	return retVals
}

// using recursive
func preorder1(root *Node) []int {
	retVals := []int{}
	helper(root, &retVals)
	return retVals
}

func helper(root *Node, retVals *[]int) {
	if root == nil {
		return
	}

	*retVals = append(*retVals, root.Val)
	if len(root.Children) == 0 {
		return
	}

	children := root.Children
	for i := 0; i < len(children); i++ {
		helper(children[i], retVals)
	}
}

// @lc code=end