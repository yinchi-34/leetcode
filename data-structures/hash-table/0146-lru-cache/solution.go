package solution

// node is an element of the doubly linked list. It stores its key as well as
// its value so that evicting the tail can also remove the key from the map.
type node struct {
	key, value int
	prev, next *node
}

// LRUCache keeps at most capacity key-value pairs. The doubly linked list is
// circular with a sentinel, as in CLRS §10.2: sentinel.next is the most
// recently used node and sentinel.prev is the least recently used one. The
// map sends each key in the cache to its node, so both Get and Put run in
// expected O(1) time.
type LRUCache struct {
	capacity int
	nodes    map[int]*node
	sentinel *node
}

func Constructor(capacity int) LRUCache {
	s := &node{}
	s.prev, s.next = s, s
	return LRUCache{
		capacity: capacity,
		nodes:    make(map[int]*node, capacity),
		sentinel: s,
	}
}

// Get returns the value of key and marks it as most recently used, or returns
// -1 if key is not in the cache.
func (c *LRUCache) Get(key int) int {
	x, ok := c.nodes[key]
	if !ok {
		return -1
	}
	c.moveToFront(x)
	return x.value
}

// Put sets the value of key and marks it as most recently used. If the key is
// new and the cache is full, the least recently used key is evicted first.
func (c *LRUCache) Put(key, value int) {
	if x, ok := c.nodes[key]; ok {
		x.value = value
		c.moveToFront(x)
		return
	}
	if len(c.nodes) == c.capacity {
		lru := c.sentinel.prev
		c.remove(lru)
		delete(c.nodes, lru.key)
	}
	x := &node{key: key, value: value}
	c.pushFront(x)
	c.nodes[key] = x
}

// remove unlinks x from the list (LIST-DELETE' in CLRS §10.2).
func (c *LRUCache) remove(x *node) {
	x.prev.next = x.next
	x.next.prev = x.prev
}

// pushFront links x right after the sentinel (LIST-INSERT' in CLRS §10.2).
func (c *LRUCache) pushFront(x *node) {
	x.next = c.sentinel.next
	c.sentinel.next.prev = x
	c.sentinel.next = x
	x.prev = c.sentinel
}

func (c *LRUCache) moveToFront(x *node) {
	c.remove(x)
	c.pushFront(x)
}
