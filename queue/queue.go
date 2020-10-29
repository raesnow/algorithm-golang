package main

import (
	"fmt"
)

type QueueInterface interface {
	enqueue(value int) error
	dequeue() (int, error)
	print()
}

type Queue struct {
	data   []int
	tail   int
	length int
}

func createQueue(length int) *Queue {
	return &Queue{
		data:   make([]int, length),
		tail:   0,
		length: length,
	}
}

func (q *Queue) enqueue(value int) error {
	if q.tail == q.length {
		return fmt.Errorf("queue is full")
	}
	q.data[q.tail] = value
	q.tail++
	return nil
}

func (q *Queue) dequeue() (int, error) {
	if q.tail == 0 {
		return -1, fmt.Errorf("queue is empty")
	}
	value := q.data[0]
	for i := 0; i < q.tail-1; i++ {
		q.data[i] = q.data[i+1]
	}
	q.tail--
	return value, nil
}

func (q *Queue) print() {
	fmt.Printf("queue[%d/%d]: %v", q.tail, q.length, q.data)
}

type LinkNode struct {
	value int
	next  *LinkNode
}

type LinkedQueue struct {
	head *LinkNode
	tail *LinkNode
}

func createLinkedQueue() *LinkedQueue {
	sentinel := &LinkNode{}
	return &LinkedQueue{
		head: sentinel,
	}
}

func (q *LinkedQueue) enqueue(value int) error {
	node := &LinkNode{
		value: value,
	}
	if q.head.next == nil {
		q.tail = q.head
	}
	q.tail.next = node
	q.tail = q.tail.next
	return nil
}

func (q *LinkedQueue) dequeue() (int, error) {
	if q.head.next == nil {
		return -1, fmt.Errorf("queue is empty")
	}
	value := q.head.next.value
	q.head.next = q.head.next.next
	return value, nil
}

func (q *LinkedQueue) print() {
	p := q.head.next
	for p != nil {
		fmt.Printf(" -> %d", p.value)
		p = p.next
	}
	fmt.Println()
}

type DequeInterface interface {
	headPush(value int) error
	headPop() (int, error)
	tailPush(value int) error
	tailPop() (int, error)
	print()
}

type Deque struct {
	data   []int
	head   int
	tail   int
	length int
	empty  bool
}

func createDeque(length int) *Deque {
	return &Deque{
		data:   make([]int, length),
		head:   0,
		tail:   0,
		length: length,
		empty:  true,
	}
}

func (q *Deque) headPush(value int) error {
	if q.tail == q.head && !q.empty {
		return fmt.Errorf("queue is full")
	}
	q.empty = false
	q.head = (q.head - 1 + q.length) % q.length
	q.data[q.head] = value
	return nil
}

func (q *Deque) headPop() (int, error) {
	if q.tail == q.head && q.empty {
		return -1, fmt.Errorf("queue is empty")
	}
	value := q.data[q.head]
	q.head = (q.head + 1) % q.length
	if q.tail == q.head {
		q.empty = true
	}
	return value, nil
}

func (q *Deque) tailPush(value int) error {
	if q.tail == q.head && !q.empty {
		return fmt.Errorf("queue is full")
	}
	q.empty = false
	q.data[q.tail] = value
	q.tail = (q.tail + 1) % q.length
	return nil
}

func (q *Deque) tailPop() (int, error) {
	if q.tail == q.head && q.empty {
		return -1, fmt.Errorf("queue is empty")
	}
	q.tail = (q.tail - 1 + q.length) % q.length
	value := q.data[q.tail]
	if q.tail == q.head {
		q.empty = true
	}
	return value, nil
}

func (q *Deque) print() {
	fmt.Printf("queue[%d/%d/%d]: %v", q.head, q.tail, q.length, q.data)
}

type DequeNode struct {
	value int
	next  *DequeNode
	pre   *DequeNode
}

type LinkedDeque struct {
	head *DequeNode
	tail *DequeNode
}

func createLinkedDeque() *LinkedDeque {
	tailSentinel := &DequeNode{}
	headSentinel := &DequeNode{}
	headSentinel.next = tailSentinel
	tailSentinel.pre = headSentinel
	return &LinkedDeque{
		head: headSentinel,
		tail: tailSentinel,
	}
}

func (q *LinkedDeque) headPush(value int) error {
	node := &DequeNode{
		value: value,
		next:  q.head.next,
		pre:   q.head,
	}
	q.head.next.pre = node
	q.head.next = node
	return nil
}

func (q *LinkedDeque) headPop() (int, error) {
	if q.head.next == q.tail {
		return -1, fmt.Errorf("queue is empty")
	}
	value := q.head.next.value
	q.head.next.next.pre = q.head
	q.head.next = q.head.next.next
	return value, nil
}

func (q *LinkedDeque) tailPush(value int) error {
	node := &DequeNode{
		value: value,
		pre:   q.tail.pre,
		next:  q.tail,
	}
	q.tail.pre.next = node
	q.tail.pre = node
	return nil
}

func (q *LinkedDeque) tailPop() (int, error) {
	if q.tail.pre == q.head {
		return -1, fmt.Errorf("queue is empty")
	}
	value := q.tail.pre.value
	q.tail.pre.pre.next = q.tail
	q.tail.pre = q.tail.pre.pre
	return value, nil
}

func (q *LinkedDeque) print() {
	p := q.head.next
	for p.next != nil {
		fmt.Printf(" -> %d", p.value)
		p = p.next
	}
	fmt.Println()
}

func main() {
	queue := createQueue(3)
	//queue := createLinkedQueue()
	fmt.Println("init queue")
	queue.print()

	queue.enqueue(5)
	fmt.Println("enqueue 5")
	queue.print()
	queue.enqueue(7)
	fmt.Println("enqueue 7")
	queue.print()
	queue.enqueue(20)
	fmt.Println("enqueue 20")
	queue.print()
	err := queue.enqueue(20)
	fmt.Println("enqueue 20")
	if err != nil {
		fmt.Printf("enqueue error: %s\n", err)
	}
	queue.print()

	v, err := queue.dequeue()
	fmt.Printf("dequeue: %d\n", v)
	queue.print()
	v, err = queue.dequeue()
	fmt.Printf("dequeue: %d\n", v)
	queue.print()
	v, err = queue.dequeue()
	fmt.Printf("dequeue: %d\n", v)
	queue.print()
	v, err = queue.dequeue()
	if err != nil {
		fmt.Printf("enqueue error: %s\n", err)
	} else {
		fmt.Printf("dequeue: %d\n", v)
	}
	queue.print()

	queue.enqueue(13)
	fmt.Println("enqueue 13")
	queue.print()
	v, err = queue.dequeue()
	fmt.Printf("dequeue: %d\n", v)
	queue.print()

	fmt.Println("****  DEQUE  ****")
	//deque := createDeque(3)
	deque := createLinkedDeque()
	fmt.Println("init deque")
	deque.print()

	deque.headPush(5)
	fmt.Println("headPush 5")
	deque.print()
	deque.headPush(7)
	fmt.Println("headPush 7")
	deque.print()
	deque.tailPush(20)
	fmt.Println("tailPush 20")
	deque.print()
	err = deque.tailPush(20)
	fmt.Println("tailPush 20")
	if err != nil {
		fmt.Printf("tailPush error: %s\n", err)
	}
	deque.print()

	v, err = deque.tailPop()
	fmt.Printf("tailPop: %d\n", v)
	deque.print()
	v, err = deque.tailPop()
	fmt.Printf("tailPop: %d\n", v)
	deque.print()
	v, err = deque.headPop()
	fmt.Printf("headPop: %d\n", v)
	deque.print()
	v, err = deque.headPop()
	if err != nil {
		fmt.Printf("headPop error: %s\n", err)
	} else {
		fmt.Printf("headPop: %d\n", v)
	}
	deque.print()

	deque.tailPush(13)
	fmt.Println("tailPush 13")
	deque.print()
	v, err = deque.headPop()
	fmt.Printf("headPop: %d\n", v)
	deque.print()
}
