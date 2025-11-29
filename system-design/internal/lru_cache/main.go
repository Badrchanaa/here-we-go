package lru_cache

type LRUCache struct {
	capacity int
	size     int
	data     map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	prev  *Node
	next  *Node
	key   int
	value int
}

func insertFirst(head **Node, key int, value int) *Node {
	newNode := Node{
		prev:  nil,
		key:   key,
		next:  *head,
		value: value}
	if *head != nil {
		(*head).prev = &newNode
	}
	*head = &newNode
	return &newNode
}

func moveToFirst(node *Node, head **Node) {
	if node == nil || node.prev == nil || *head == nil {
		return
	}

	// Detach node
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev.next = node.next

	// Attach to head
	node.next = (*head)
	(*head).prev = node
	node.prev = nil
	*head = node
}

func removeLast(tail **Node) {
	if *tail == nil {
		return
	}
	prev := (*tail).prev
	if prev != nil {
		prev.next = nil
		*tail = prev
	} else {
		*tail = nil
	}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{data: make(map[int]*Node), capacity: capacity, size: 0, head: nil, tail: nil}
}

func (this *LRUCache) Get(key int) int {
	node, exists := this.data[key]
	if !exists {
		return -1
	}
	if node == this.tail && node.prev != nil {
		this.tail = node.prev
	}
	moveToFirst(node, &this.head)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, exists := this.data[key]
	if exists {
		node.value = value
		if node == this.tail && node.prev != nil {
			this.tail = node.prev
		}
		moveToFirst(node, &this.head)
		return
	}
	newNode := insertFirst(&this.head, key, value)
	this.data[key] = newNode
	if this.tail == nil {
		this.tail = newNode
	}
	if this.size == this.capacity {
		delete(this.data, this.tail.key)
		removeLast(&this.tail)
	} else {
		this.size += 1
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * saram_1 := obj.Get(key);
 * obj.Put(key,value);
 */
