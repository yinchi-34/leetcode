package solution

import (
	"math/rand"
	"slices"
	"testing"
)

// op is one call on the cache. For "get", want is the expected return value.
type op struct {
	name       string
	key, value int
	want       int
}

func get(key, want int) op  { return op{name: "get", key: key, want: want} }
func put(key, value int) op { return op{name: "put", key: key, value: value} }

// keysFromMRU walks the list from the most to the least recently used node.
func keysFromMRU(c *LRUCache) []int {
	var res []int
	for x := c.head.next; x != c.tail; x = x.next {
		res = append(res, x.key)
	}
	return res
}

// checkInvariant checks the representation invariant after every call: the
// list is a well-formed list from head to tail, its keys are exactly the map's keys,
// each map entry points at the node holding that key, and the size is within
// capacity.
func checkInvariant(t *testing.T, c *LRUCache) {
	t.Helper()
	if c.head.prev != nil || c.tail.next != nil {
		t.Fatalf("sentinels must not link outside the list")
	}
	if c.head.next.prev != c.head || c.tail.prev.next != c.tail {
		t.Fatalf("broken links at a sentinel")
	}
	count := 0
	for x := c.head.next; x != c.tail; x = x.next {
		if x.next.prev != x || x.prev.next != x {
			t.Fatalf("broken links at key %d", x.key)
		}
		if c.cache[x.key] != x {
			t.Fatalf("map entry for key %d does not point at its node", x.key)
		}
		count++
		if count > len(c.cache) {
			t.Fatalf("list longer than map (%d entries)", len(c.cache))
		}
	}
	if count != len(c.cache) {
		t.Fatalf("list has %d nodes, map has %d", count, len(c.cache))
	}
	if count > c.capacity {
		t.Fatalf("size %d exceeds capacity %d", count, c.capacity)
	}
}

func run(t *testing.T, capacity int, ops []op) *LRUCache {
	t.Helper()
	c := Constructor(capacity)
	for i, o := range ops {
		switch o.name {
		case "get":
			if got := c.Get(o.key); got != o.want {
				t.Fatalf("op %d: Get(%d) = %d, want %d", i, o.key, got, o.want)
			}
		case "put":
			c.Put(o.key, o.value)
		}
		checkInvariant(t, &c)
	}
	return &c
}

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		ops      []op
		order    []int // keys from most to least recently used at the end
	}{
		{
			name:     "Example 1",
			capacity: 2,
			ops: []op{
				put(1, 1), put(2, 2), get(1, 1), put(3, 3), get(2, -1),
				put(4, 4), get(1, -1), get(3, 3), get(4, 4),
			},
			order: []int{4, 3},
		},
		{
			name:     "Capacity 1: every new key evicts the old one",
			capacity: 1,
			ops:      []op{put(2, 1), get(2, 1), put(3, 2), get(2, -1), get(3, 2)},
			order:    []int{3},
		},
		{
			name:     "Get on empty cache",
			capacity: 2,
			ops:      []op{get(1, -1)},
			order:    nil,
		},
		{
			name:     "Updating an existing key does not evict",
			capacity: 2,
			ops:      []op{put(1, 1), put(2, 2), put(1, 10), get(1, 10), get(2, 2)},
			order:    []int{2, 1},
		},
		{
			name:     "Updating a key makes it most recent",
			capacity: 2,
			ops:      []op{put(2, 1), put(1, 1), put(2, 3), put(4, 1), get(1, -1), get(2, 3)},
			order:    []int{2, 4},
		},
		{
			name:     "A miss does not change recency",
			capacity: 2,
			ops:      []op{put(1, 1), put(2, 2), get(9, -1), put(3, 3), get(1, -1), get(2, 2)},
			order:    []int{2, 3},
		},
		{
			name:     "Get refreshes the least recent key",
			capacity: 3,
			ops:      []op{put(1, 1), put(2, 2), put(3, 3), get(1, 1), put(4, 4), get(2, -1), get(1, 1)},
			order:    []int{1, 4, 3},
		},
		{
			name:     "Value 0 and key 0 are stored like any other",
			capacity: 2,
			ops:      []op{put(0, 0), get(0, 0), put(1, 0), get(1, 0)},
			order:    []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := run(t, tt.capacity, tt.ops)
			if got := keysFromMRU(c); !slices.Equal(got, tt.order) {
				t.Errorf("recency order = %v, want %v", got, tt.order)
			}
		})
	}
}

// naiveLRU is a reference model: a slice of keys ordered from least to most
// recently used, plus a map of values. Every operation is O(capacity).
type naiveLRU struct {
	capacity int
	order    []int
	values   map[int]int
}

func (n *naiveLRU) touch(key int) {
	i := slices.Index(n.order, key)
	n.order = append(slices.Delete(n.order, i, i+1), key)
}

func (n *naiveLRU) get(key int) int {
	v, ok := n.values[key]
	if !ok {
		return -1
	}
	n.touch(key)
	return v
}

func (n *naiveLRU) put(key, value int) {
	if _, ok := n.values[key]; ok {
		n.values[key] = value
		n.touch(key)
		return
	}
	if len(n.order) == n.capacity {
		delete(n.values, n.order[0])
		n.order = n.order[1:]
	}
	n.order = append(n.order, key)
	n.values[key] = value
}

func TestLRUCacheRandomAgainstNaive(t *testing.T) {
	rng := rand.New(rand.NewSource(146))
	for trial := 0; trial < 300; trial++ {
		capacity := 1 + rng.Intn(6)
		keyRange := 1 + rng.Intn(10)
		c := Constructor(capacity)
		ref := &naiveLRU{capacity: capacity, values: map[int]int{}}
		for step := 0; step < 200; step++ {
			key := rng.Intn(keyRange)
			if rng.Intn(2) == 0 {
				if got, want := c.Get(key), ref.get(key); got != want {
					t.Fatalf("trial %d step %d: Get(%d) = %d, want %d", trial, step, key, got, want)
				}
			} else {
				value := rng.Intn(100)
				c.Put(key, value)
				ref.put(key, value)
			}
			checkInvariant(t, &c)
			want := slices.Clone(ref.order)
			slices.Reverse(want)
			if got := keysFromMRU(&c); !slices.Equal(got, want) {
				t.Fatalf("trial %d step %d: order = %v, want %v", trial, step, got, want)
			}
		}
	}
}
