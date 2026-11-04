package solution

// dLinkedNode is an element of the doubly linked list. It stores its key as
// well as its value so that evicting the tail can also remove the key from
// the map.
type dLinkedNode struct {
	key, value int
	prev, next *dLinkedNode
}

// LRUCache keeps at most capacity key-value pairs. The doubly linked list has
// two sentinels: head.next is the most recently used node and tail.prev is the
// least recently used one, so no operation needs a nil check. The map sends
// each key in the cache to its node, so both Get and Put run in expected O(1)
// time.
type LRUCache struct {
	capacity   int
	cache      map[int]*dLinkedNode
	head, tail *dLinkedNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &dLinkedNode{}, &dLinkedNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*dLinkedNode, capacity+1),
		head:     head,
		tail:     tail,
	}
}

// Get returns the value of key and marks it as most recently used, or returns
// -1 if key is not in the cache.
func (c *LRUCache) Get(key int) int {
	node, ok := c.cache[key]
	if !ok {
		return -1
	}
	c.remove(node)
	c.addToHead(node)
	return node.value
}

// Put sets the value of key and marks it as most recently used. A new key is
// inserted first; if that makes the cache exceed its capacity, the least
// recently used key is evicted.
func (c *LRUCache) Put(key, value int) {
	if node, ok := c.cache[key]; ok {
		node.value = value
		c.remove(node)
		c.addToHead(node)
		return
	}
	node := &dLinkedNode{key: key, value: value}
	c.cache[key] = node
	c.addToHead(node)
	if len(c.cache) > c.capacity {
		lru := c.tail.prev
		c.remove(lru)
		delete(c.cache, lru.key)
	}
}

// remove unlinks node from the list.
func (c *LRUCache) remove(node *dLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// addToHead links node right after head. The old first node's prev must be
// updated before head.next, because head.next is the only way to reach it.
func (c *LRUCache) addToHead(node *dLinkedNode) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}
