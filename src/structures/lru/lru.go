package lru

import "container/list"


// LruCache 基于内存实现、不带过期时间
// 原理：map结构按照kv存储数据，双向链表保存数据新鲜度
// 扩展：支持过期时间可以增加一个双向链表按过期时间存储，

// LRUChainNode 链表节点
type LRUChainNode struct {
	pre   *LRUChainNode
	next  *LRUChainNode
	key   int
	value int
	ts    int32
}

// LRUCache 结构
type LRUCache struct {
	capacity int
	length   int
	store    map[int]*LRUChainNode
	head     *LRUChainNode
	tail     *LRUChainNode
}

// NewLRUCache constructor
func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		length:   0,
		store:    map[int]*LRUChainNode{},
	}
}

// Delete 删除key
func (c *LRUCache) Delete(key int) {
	node, exist := c.store[key]
	if !exist {
		return
	}
	delete(c.store, key)
	if c.length == 1 {
		c.head = nil
		c.tail = nil
		c.length--
		return
	}
	if node.pre == nil {
		c.head = c.head.next
		c.head.pre = nil
	} else {
		node.pre.next = node.next
	}
	if node.next == nil {
		c.tail = c.tail.pre
		c.tail.next = nil
	} else {
		node.next.pre = node.pre
	}
	c.length--
}

// Get 获取kv
func (c *LRUCache) Get(key int) int {
	node, exist := c.store[key]
	if !exist {
		return -1
	}
	c.Delete(key)
	c.Put(node.key, node.value)
	return node.value
}

// Put 插入
func (c *LRUCache) Put(key int, value int) {
	c.Delete(key)
	if c.length+1 > c.capacity {
		c.Delete(c.tail.key)
	}
	node := LRUChainNode{key: key, value: value}
	if c.length == 0 {
		c.head = &node
		c.tail = &node
		c.store[key] = &node
		c.length++
		return
	}
	// 头部处理
	c.head.pre = &node
	node.next = c.head
	c.head = &node
	c.store[key] = &node
	c.length++
}


// 第二种直接使用container/list 包
type LRUCache struct {
	Cap  int
	Keys map[int]*list.Element
	List *list.List
}

type pair struct {
	K,V int
}

func Constructor(capacity int) LRUCache{
	return LRUCache{
		Cap: capacity,
		Keys: make(map[int]*list.Element),
		List: list.New(),
	}
}

func (impl *LRUCache) Get(key int) int{
	if el,ok := impl.Keys[key];ok{
		impl.List.MoveToFront(el)
		return el.Value.(pair).V
	}
	return -1
}

func (impl *LRUCache) Put(key int,value int){
	if el,ok := impl.Keys[key];ok{
		el.Value = pair{K:key,V: value}
		impl.List.MoveToFront(el)
	}else{
		el := impl.List.PushFront(pair{K: key,V: value})
		impl.Keys[key] = el
	}
	if impl.List.Len() > impl.Cap{
		el := impl.List.Back()
		impl.List.Remove(el)
		delete(impl.Keys,el.Value.(pair).K)
	}
}


// 第三种直接实现


type LRUCache struct {
	head,tail *Node
	Keys map[int]*Node
	capacity int
}

type Node struct {
	Prev,Next *Node
	Key,Val int
}

func Constructor(capacity int) LRUCache{
	return LRUCache{
		Keys: make(map[int]*Node),
		capacity: capacity,
	}
}


func (impl *LRUCache) Get(key int) int{
	if node,ok := impl.Keys[key];ok {
		impl.Remove(node)
		impl.Add(node)
		return node.Val
	}
	return -1
}

func (impl *LRUCache) Put(key,value int){
	if node,ok := impl.Keys[key];ok{
		node.Val = value
		impl.Remove(node)
		impl.Add(node)
	}else{
		node = &Node{
			Key:key,
			Val: value,
		}
		impl.Keys[key] = node
		impl.Add(node)
	}
	if len(impl.Keys) > impl.capacity{
		delete(impl.Keys,impl.tail.Key)
		impl.Remove(impl.tail)
	}
	return
}

func (impl *LRUCache) Adds1(node *Node){
	if impl.head !=nil {
		impl.head.Prev = node
		node.Next = impl.head
	}
	 impl.head= node
	 if impl.tail == nil {
	 	impl.tail = node
	 	impl.tail.Prev = node
	 	impl.tail.Next = nil
	 }
}

func (impl *LRUCache)Adds(node *Node){
	if impl.head != nil{
		impl.head.Prev = node
		node.Next = impl.head
	}
	impl.head  = node
	if impl.tail == nil {
		impl.tail =node
		impl.tail.Prev = node
		impl.tail.Next = nil
	}
}

func (impl *LRUCache) Add(node *Node){
	if impl.head !=nil {
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

func (impl *LRUCache)Remove(node *Node){
	if impl.head == node {
		if node.Next != nil {
			node.Next.Prev = nil
		}
		impl.head= node.Next
		return
	}
	if impl.tail == node {
		impl.tail = impl.tail.Prev
		return
	}
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}



