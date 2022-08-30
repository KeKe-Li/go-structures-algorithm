package lru2

import "container/list"

// 第二种直接使用container/list
type LRUCache struct {
	Cap  int
	Keys map[int]*list.Element
	List *list.List
}

type pair struct {
	K, V int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Keys: make(map[int]*list.Element),
		List: list.New(),
	}
}

func (impl *LRUCache) Get(key int) int {
	if el, ok := impl.Keys[key]; ok {
		impl.List.MoveToFront(el)
		return el.Value.(pair).V
	}
	return -1
}

func (impl *LRUCache) Put(key int, value int) {
	if el, ok := impl.Keys[key]; ok {
		el.Value = pair{K: key, V: value}
		impl.List.MoveToFront(el)
	} else {
		el := impl.List.PushFront(pair{K: key, V: value})
		impl.Keys[key] = el
	}
	if impl.List.Len() > impl.Cap {
		el := impl.List.Back()
		impl.List.Remove(el)
		delete(impl.Keys, el.Value.(pair).K)
	}
}

// 第三种LRU直接实现
type LRUCacheStruct struct {
	head, tail *Node
	Keys       map[int]*Node
	capacity   int
}

type Node struct {
	Prev, Next *Node
	Key, Val   int
}

func ConstructorLRU(capacity int) LRUCacheStruct {
	return LRUCacheStruct{
		Keys:     make(map[int]*Node),
		capacity: capacity,
	}
}

func (impl *LRUCacheStruct) Get(key int) int {
	if node, ok := impl.Keys[key]; ok {
		impl.Remove(node)
		impl.Add(node)
		return node.Val
	}
	return -1
}

func (impl *LRUCacheStruct) Put(key, value int) {
	if node, ok := impl.Keys[key]; ok {
		node.Val = value
		impl.Remove(node)
		impl.Add(node)
	} else {
		node = &Node{
			Key: key,
			Val: value,
		}
		impl.Keys[key] = node
		impl.Add(node)
	}
	if len(impl.Keys) > impl.capacity {
		delete(impl.Keys, impl.tail.Key)
		impl.Remove(impl.tail)
	}
	return
}

func (impl *LRUCacheStruct) Add(node *Node) {
	if impl.head != nil {
		impl.head.Prev = node
		node.Next = impl.head
	}
	impl.head = node
	if impl.tail == nil {
		impl.tail = node
		impl.tail.Prev = node
		impl.tail.Next = nil
	}
	return
}

func (impl *LRUCacheStruct) Remove(node *Node) {
	if impl.head == node {
		if node.Next != nil {
			node.Next.Prev = nil
		}
		impl.head = node.Next
		return
	}
	if impl.tail == node {
		impl.tail = impl.tail.Prev
		return
	}
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}
