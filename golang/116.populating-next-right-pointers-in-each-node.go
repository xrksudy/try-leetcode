/*
 * @lc app=leetcode id=116 lang=golang
 *
 * [116] Populating Next Right Pointers in Each Node
 *
 * https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/
 *
 * algorithms
 * Medium (44.13%)
 * Likes:    2037
 * Dislikes: 144
 * Total Accepted:    365.1K
 * Total Submissions: 814.4K
 * Testcase Example:  '[1,2,3,4,5,6,7]'
 *
 * You are given a perfect binary tree where all leaves are on the same level,
 * and every parent has two children. The binary tree has the following
 * definition:
 *
 *
 * struct Node {
 * ⁠ int val;
 * ⁠ Node *left;
 * ⁠ Node *right;
 * ⁠ Node *next;
 * }
 *
 *
 * Populate each next pointer to point to its next right node. If there is no
 * next right node, the next pointer should be set to NULL.
 *
 * Initially, all next pointers are set to NULL.
 *
 *
 *
 * Follow up:
 *
 *
 * You may only use constant extra space.
 * Recursive approach is fine, you may assume implicit stack space does not
 * count as extra space for this problem.
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: root = [1,2,3,4,5,6,7]
 * Output: [1,#,2,3,#,4,5,6,7,#]
 * Explanation: Given the above perfect binary tree (Figure A), your function
 * should populate each next pointer to point to its next right node, just like
 * in Figure B. The serialized output is in level order as connected by the
 * next pointers, with '#' signifying the end of each level.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the given tree is less than 4096.
 * -1000 <= node.val <= 1000
 *
 *
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	return connect2(root)
}

// using recursive
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	var nextItem *Node
	root.Next = nextItem
	helper(root.Right, nextItem)
	helper(root.Left, root.Right)
	return root
}

func helper(root *Node, next *Node) {
	if root != nil {
		root.Next = next
		if root.Next != nil {
			helper(root.Right, root.Next.Left)
		} else {
			var nextItem *Node
			helper(root.Right, nextItem)
		}
		helper(root.Left, root.Right)
	}
}

// solution with bfs, using stack
func connect1(root *Node) *Node {
	if root == nil {
		return root
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		nextStack := []*Node{}
		var nextItem *Node
		for i := 0; i < len(stack); i++ {
			curNode := stack[i]
			curNode.Next = nextItem
			if curNode.Right != nil {
				nextStack = append(nextStack, curNode.Right)
			}
			if curNode.Left != nil {
				nextStack = append(nextStack, curNode.Left)
			}
			nextItem = curNode
		}
		stack = nextStack
	}
	return root
}

// @lc code=end