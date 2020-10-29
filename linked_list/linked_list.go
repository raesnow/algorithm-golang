package main

import "fmt"

type LinkNode struct {
	value int
	next  *LinkNode
}

func makeLinkedList(length int) *LinkNode {
	return makeLinkedListWithStartValue(length, 0)
}

func makeLinkedListWithStartValue(length int, startValue int) *LinkNode {
	if length == 0 {
		return nil
	}

	sentinel := &LinkNode{}
	pre := sentinel
	for i := 0; i < length; i++ {
		newNode := &LinkNode{
			value: startValue + i,
		}
		pre.next = newNode
		pre = newNode
	}

	return sentinel.next
}

func printLinkedList(head *LinkNode) {
	for head != nil {
		fmt.Printf("-> %d ", head.value)
		head = head.next
	}
	fmt.Println()
}

func reverseLinkedList(head *LinkNode) *LinkNode {
	if head == nil || head.next == nil {
		return head
	}

	var sentinel *LinkNode
	pre := sentinel
	for node := head; node != nil; {
		next := node.next
		node.next = pre
		pre = node
		node = next
	}
	return pre
}

func makeCircleLinkedList(length int, circleStartPos int) *LinkNode {
	if length == 0 || circleStartPos <= 0 {
		return nil
	}
	if circleStartPos > length {
		circleStartPos %= length
	}

	sentinel := &LinkNode{}
	pre := sentinel
	var circleStartNode *LinkNode
	for i := 0; i < length; i++ {
		newNode := &LinkNode{
			value: i,
		}
		if i == circleStartPos-1 {
			circleStartNode = newNode
		}
		pre.next = newNode
		pre = newNode
	}
	pre.next = circleStartNode

	return sentinel.next
}

func hasCircle(head *LinkNode) bool {
	if head == nil || head.next == nil {
		return true
	}
	slow := head.next
	quick := head.next.next
	for quick != nil && quick.next != nil && quick != slow {
		slow = slow.next
		quick = quick.next.next
	}
	if quick == slow {
		return true
	}
	return false
}

func findCircleStartPos(head *LinkNode) int {
	if head == nil || head.next == nil {
		return -1
	}

	slow := head.next
	quick := head.next.next
	for quick != nil && quick.next != nil && quick != slow {
		slow = slow.next
		quick = quick.next.next
	}
	if quick != slow {
		return -1
	}

	count := 1
	slow = head
	for quick != slow {
		slow = slow.next
		quick = quick.next
		count++
	}
	return count
}

func mergeLists(list1, list2 *LinkNode) *LinkNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	sentinel := &LinkNode{}
	pre := sentinel
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.value <= p2.value {
			pre.next = p1
			p1 = p1.next
		} else {
			pre.next = p2
			p2 = p2.next
		}
		pre = pre.next
	}
	if p1 != nil {
		pre.next = p1
		p1 = p1.next
		pre = pre.next
	}
	if p2 != nil {
		pre.next = p2
		p2 = p2.next
		pre = pre.next
	}
	return sentinel.next
}

func findKPosNode(head *LinkNode, pos int) *LinkNode {
	if head == nil || pos <= 0 {
		return nil
	}
	pos -= 1
	p := head
	for i := 0; i < pos && p != nil; i++ {
		p = p.next
	}
	return p
}

func findReversedKPosNode(head *LinkNode, pos int) *LinkNode {
	if head == nil {
		return nil
	}
	length := 0
	for p := head; p != nil; p = p.next {
		length++
	}
	return findKPosNode(head, length-pos+1)
}

func delReversedKPosNode(head *LinkNode, pos int) *LinkNode {
	if head == nil || pos <= 0 {
		return nil
	}
	length := 0
	for p := head; p != nil; p = p.next {
		length++
	}

	index := length - pos
	pre := &LinkNode{}
	p := head
	for i := 0; i < index && p != nil; i++ {
		pre = p
		p = p.next
	}
	pre.next = p.next
	return head
}

func findHalfPosNode(head *LinkNode) *LinkNode {
	if head == nil || head.next == nil {
		return head
	}

	quick := head.next.next
	slow := head.next
	for quick != nil && quick.next != nil {
		quick = quick.next.next
		slow = slow.next
	}
	return slow
}

type KVLinkNode struct {
	key   string
	value string
	next  *KVLinkNode
	pre   *KVLinkNode
}

type LRUCache struct {
	head     *KVLinkNode
	tail     *KVLinkNode
	length   int
	capacity int
	cache    map[string]*KVLinkNode
}

func makeLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*KVLinkNode),
	}
}

func (c *LRUCache) GetLength() int {
	return c.length
}

func (c *LRUCache) GetCapacity() int {
	return c.capacity
}

func (c *LRUCache) Add(key, value string) {
	kvLinkNode := &KVLinkNode{
		key:   key,
		value: value,
	}

	if c.head == nil {
		c.head = kvLinkNode
		c.tail = kvLinkNode
	} else {
		kvLinkNode.next = c.head
		c.head.pre = kvLinkNode
		c.head = kvLinkNode
	}
	c.length++
	c.cache[key] = kvLinkNode

	if c.length > c.capacity {
		delete(c.cache, c.tail.key)
		c.tail = c.tail.pre
		if c.tail != nil {
			c.tail.next = nil
		}
		c.length--
	}
}

func (c *LRUCache) Get(key string) string {
	node, ok := c.cache[key]
	if ok {
		return node.value
	}
	return ""
}

func main() {
	list := makeLinkedList(10)
	printLinkedList(list)

	reversedHead := reverseLinkedList(list)
	fmt.Println("After reverse")
	printLinkedList(reversedHead)
	fmt.Println()

	fmt.Printf("Check has circle: %v\n", hasCircle(reversedHead))
	circleList := makeCircleLinkedList(10, 4)
	fmt.Printf("Check has circle: %v\n", hasCircle(circleList))
	fmt.Printf("Find circle start position: %d\n", findCircleStartPos(circleList))
	fmt.Println()

	list1 := makeLinkedList(15)
	printLinkedList(list1)
	list2 := makeLinkedListWithStartValue(10, 2)
	printLinkedList(list2)
	mergedList := mergeLists(list1, list2)
	fmt.Println("After merged")
	printLinkedList(mergedList)

	node := findKPosNode(mergedList, 7)
	fmt.Printf("Find list1 %d node: %d\n", 7, node.value)
	node = findKPosNode(mergedList, 1)
	fmt.Printf("Find list1 %d node: %d\n", 1, node.value)
	node = findReversedKPosNode(mergedList, 7)
	fmt.Printf("Find list1 %d node: %d\n", 7, node.value)
	delReversedKPosNode(mergedList, 7)
	fmt.Println("After delete reversed pos 7 node")
	printLinkedList(mergedList)

	list3 := makeLinkedList(10)
	printLinkedList(list3)
	node = findHalfPosNode(list3)
	fmt.Printf("Find list1 half node: %d\n", node.value)

	lru := makeLRUCache(3)
	fmt.Printf("Create LRU Cache, length: %d, capacity: %d\n", lru.GetLength(), lru.GetCapacity())
	lru.Add("1", "value1")
	lru.Add("2", "value2")
	lru.Add("3", "value3")
	fmt.Printf("1: %s, 2: %s, 3: %s\n", lru.Get("1"), lru.Get("2"), lru.Get("3"))
	lru.Add("4", "value4")
	fmt.Printf("LRU Cache, length: %d, capacity: %d\n", lru.GetLength(), lru.GetCapacity())
	fmt.Printf("1: %s, 2: %s, 3: %s, 4: %s\n",
		lru.Get("1"), lru.Get("2"), lru.Get("3"), lru.Get("4"))
}
