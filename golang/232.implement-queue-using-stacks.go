/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 *
 * https://leetcode.com/problems/implement-queue-using-stacks/description/
 *
 * algorithms
 * Easy (48.76%)
 * Likes:    1085
 * Dislikes: 143
 * Total Accepted:    217.2K
 * Total Submissions: 441.3K
 * Testcase Example:  '["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a queue using stacks.
 *
 *
 * push(x) -- Push element x to the back of queue.
 * pop() -- Removes the element from in front of queue.
 * peek() -- Get the front element.
 * empty() -- Return whether the queue is empty.
 *
 *
 * Example:
 *
 *
 * MyQueue queue = new MyQueue();
 *
 * queue.push(1);
 * queue.push(2);
 * queue.peek();  // returns 1
 * queue.pop();   // returns 1
 * queue.empty(); // returns false
 *
 * Notes:
 *
 *
 * You must use only standard operations of a stack -- which means only push to
 * top, peek/pop from top, size, and is empty operations are valid.
 * Depending on your language, stack may not be supported natively. You may
 * simulate a stack by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a stack.
 * You may assume that all operations are valid (for example, no pop or peek
 * operations will be called on an empty queue).
 *
 *
 */

// @lc code=start
type MyQueue struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{data: []int{}}

}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.data) > 0 {
		retVal := this.data[0]
		this.data = this.data[1:]
		return retVal
	}
	panic("MyQueue is empty")
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.data) > 0 {
		return this.data[0]
	}
	panic("MyQueue is empty")
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end