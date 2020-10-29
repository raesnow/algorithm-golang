package main

import "fmt"

type StackInterface interface {
	push(int) error
	pop() (int, error)
	print()
}

type Stack struct {
	data   []int
	top    int
	length int
}

func createStack(length int) *Stack {
	return &Stack{
		data:   make([]int, length),
		top:    -1,
		length: length,
	}
}

func (s *Stack) push(value int) error {
	if s.top == s.length-1 {
		return fmt.Errorf("stack is full")
	}
	s.top++
	s.data[s.top] = value
	return nil
}

func (s *Stack) pop() (int, error) {
	if s.top == -1 {
		return -1, fmt.Errorf("stack is empty")
	}
	value := s.data[s.top]
	s.top--
	return value, nil
}

func (s *Stack) print() {
	fmt.Printf("stack[%d/%d]: data: %v\n", s.top, s.length, s.data)
}

type LinkNode struct {
	value int
	next  *LinkNode
}

type LinkedStack struct {
	top *LinkNode
}

func createLinkedStack() *LinkedStack {
	return &LinkedStack{}
}

func (s *LinkedStack) push(value int) error {
	node := &LinkNode{
		value: value,
		next:  s.top,
	}
	s.top = node
	return nil
}

func (s *LinkedStack) pop() (int, error) {
	if s.top == nil {
		return -1, fmt.Errorf("stack is empty")
	}
	value := s.top.value
	s.top = s.top.next
	return value, nil
}

func (s *LinkedStack) print() {}

func main() {
	//stack := createStack(3)
	stack := createLinkedStack()
	fmt.Println("init stack")
	stack.print()

	stack.push(5)
	fmt.Println("push 5")
	stack.print()
	stack.push(7)
	fmt.Println("push 7")
	stack.print()
	stack.push(20)
	fmt.Println("push 20")
	stack.print()
	err := stack.push(20)
	fmt.Println("push 20")
	if err != nil {
		fmt.Printf("push error: %s\n", err)
	}
	stack.print()

	v, err := stack.pop()
	fmt.Printf("pop: %d\n", v)
	stack.print()
	v, err = stack.pop()
	fmt.Printf("pop: %d\n", v)
	stack.print()
	v, err = stack.pop()
	fmt.Printf("pop: %d\n", v)
	stack.print()
	v, err = stack.pop()
	if err != nil {
		fmt.Printf("push error: %s\n", err)
	}
	stack.print()

	stack.push(13)
	fmt.Println("push 13")
	stack.print()
	v, err = stack.pop()
	fmt.Printf("pop: %d\n", v)
	stack.print()
}
