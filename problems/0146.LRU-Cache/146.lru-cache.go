package problems

type DoubleLinkedList struct {
	head *Node
	tail *Node
	size int
}
type Node struct {
	key, val int
	next     *Node
	before   *Node
}

func (dll *DoubleLinkedList) remove(n *Node) {
	n.before.next = n.next
	n.next.before = n.before
	dll.size -= 1
}
func (dll *DoubleLinkedList) removeLast() int {
	key := dll.tail.before.key
	dll.remove(dll.tail.before)
	return key
}

func (dll *DoubleLinkedList) addFirst(key, val int) *Node {
	tmp := dll.head.next
	curr := &Node{key, val, tmp, dll.head}
	dll.head.next = curr
	tmp.before = curr
	dll.size += 1
	return curr
}

type LRUCache struct {
	capacity int
	dll      *DoubleLinkedList
	hashMap  map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{-1, -1, nil, nil}
	tail := &Node{-1, -1, nil, head}
	head.next = tail
	dll := &DoubleLinkedList{head, tail, 0}
	hashMap := make(map[int]*Node)
	return LRUCache{capacity, dll, hashMap}
}

func (this *LRUCache) Get(key int) int {
	if this.hashMap[key] != nil {
		this.use(key)
		return this.hashMap[key].val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.hashMap[key] != nil {
		this.update(key, value)
		return
	}
	if this.dll.size == this.capacity {
		removedKey := this.dll.removeLast()
		delete(this.hashMap, removedKey)
	}
	this.hashMap[key] = this.dll.addFirst(key, value)
}
func (this *LRUCache) use(key int) {
	if this.hashMap[key] != nil {
		val := this.hashMap[key].val
		this.dll.remove(this.hashMap[key])
		this.hashMap[key] = this.dll.addFirst(key, val)
	}
}
func (this *LRUCache) update(key, val int) {
	if this.hashMap[key] != nil {
		this.dll.remove(this.hashMap[key])
		this.hashMap[key] = this.dll.addFirst(key, val)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
