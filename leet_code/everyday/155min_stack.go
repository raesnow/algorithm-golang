package main

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.

Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2


Constraints:

Methods pop, top and getMin operations will always be called on non-empty stacks.
*/

type MinStack struct {
	values    []int
	minValues []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		values:    make([]int, 0),
		minValues: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.values = append(this.values, x)
	if len(this.values) == 1 {
		this.minValues = []int{x}
	} else {
		min := this.minValues[len(this.minValues)-1]
		if x < min {
			this.minValues = append(this.minValues, x)
		} else {
			this.minValues = append(this.minValues, min)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.values) == 0 {
		return
	}
	length := len(this.values)
	this.values = this.values[:length-1]
	this.minValues = this.minValues[:length-1]
}

func (this *MinStack) Top() int {
	if len(this.values) == 0 {
		return 0
	}
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minValues) == 0 {
		return 0
	}
	return this.minValues[len(this.minValues)-1]
}
