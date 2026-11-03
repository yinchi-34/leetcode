# Hash Table

English | [简体中文](./README.zh-CN.md)

Use the [shared glossary](../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../docs/GITHUB_MATH.md).

This is the overview note for the hash table: its definition, implementation, costs, and the usage patterns that recur across problems. Each application problem's note covers only how that problem uses a hash table; knowledge about the hash table itself lives here.

## 1. Definition

A hash table implements a **dictionary**: a dynamic set in which each element has a `key`, supporting three operations:

| Operation | Meaning | Name in CLRS |
|---|---|---|
| insert(x) | Insert element `x` | HASH-INSERT / CHAINED-HASH-INSERT |
| search(k) | Find the element with key `k`, or return NIL | HASH-SEARCH / CHAINED-HASH-SEARCH |
| delete(x) | Delete element `x` | CHAINED-HASH-DELETE |

When we only care whether a key is present and carry no satellite data, the dictionary is a **set**.

## 2. Implementation and Cost

### 2.1 The two approaches in CLRS (§11.1, §11.2)

- **Direct-address table**: when keys come from a small universe $`U = \{0, 1, \dots, m-1\}`$, use an array `T[0..m-1]` where `T[k]` holds the element with key `k`. All three operations take worst-case $`O(1)`$ time and the table uses $`\Theta(|U|)`$ space.
- **Hash table**: when the universe is too large, a hash function $`h : U \to \{0, 1, \dots, m-1\}`$ maps keys to slots. Two keys landing in the same slot is a **collision**; CLRS resolves collisions by **chaining**, keeping a linked list per slot.

With $`n`$ elements in $`m`$ slots, the **load factor** is $`\alpha = n/m`$. Under simple uniform hashing a search takes expected $`\Theta(1 + \alpha)`$ time (Theorems 11.1 and 11.2). Growing $`m`$ with $`n`$ so that $`\alpha = O(1)`$ makes all three operations expected $`O(1)`$. In the worst case every key hashes to the same slot and search degrades to $`\Theta(n)`$.

### 2.2 In Go

Go's built-in `map[K]V` is a hash table. The runtime grows it automatically when the load factor gets too high (the same amortized argument as the dynamic tables of CLRS §17.4):

```go
m := map[int]int{}        // empty table
m[k] = v                  // insert / update
v, ok := m[k]             // search: ok reports presence
delete(m, k)              // delete
n := len(m)               // number of elements
```

Notes:

- **Reading a missing key returns the zero value** without an error. So with `map[int]bool` used as a set, `set[k]` alone answers "is `k` present". When the zero value is meaningful (a count may be 0), use `v, ok := m[k]` to tell "absent" from "present with zero".
- **Two ways to write a set**: `map[int]bool` reads naturally; `map[int]struct{}` stores no value bytes but needs `_, ok := set[k]` for membership. Either is fine for these problems.
- **Iteration order is randomized** and may differ on every run. An answer must not depend on it; collect and sort the keys when order matters.
- When the size is known in advance, `make(map[int]bool, n)` allocates once and avoids regrowth.

### 2.3 Space

A hash table uses $`O(n)`$ extra space, where $`n`$ is the number of **distinct** keys. Trading this space for time typically turns $`O(n^2)`$ or $`O(n \log n)`$ into $`O(n)`$.

## 3. Common Usage Patterns

Each pattern lists what the table stores and when it applies. New problems are added under the matching pattern.

### 3.1 Membership queries: turn a scan into $`O(1)`$

- **What the table stores**: elements seen so far (a set).
- **When it applies**: the algorithm repeatedly asks "is this value present", which would otherwise need a linear scan or sorting plus binary search.
- **Typical variant**: complement lookup, as in Two Sum asking whether `target - x` has appeared; the table maps value to index.

### 3.2 Counting

- **What the table stores**: element to number of occurrences.
- **When it applies**: checking rearrangements, finding elements whose counts meet a condition, character counts inside a sliding window.

### 3.3 Grouping by canonical form

- **What the table stores**: canonical form to the list of elements in that group.
- **When it applies**: elements that are equivalent under some transformation must be grouped. The key step is designing a canonical form that maps equivalent elements to the same key, such as the sorted string of an anagram.
- **Requirement on the key**: the key must be a **complete invariant** of the equivalence relation: equivalent elements have equal keys, and elements with equal keys are equivalent. Satisfying only the first merges different groups; satisfying only the second splits a group.
- **In Go**: a slice cannot be a `map` key; convert it to a `string`, or, when the value range is small, use a fixed-length array (such as `[26]int`) as the key.

### 3.4 Start only from chain heads: avoid rescanning

- **What the table stores**: the set of all elements.
- **When it applies**: the answer is built from disjoint chains (such as runs of consecutive integers), and extending a chain needs only $`O(1)`$ membership queries.
- **Method**: extend forward from `x` only when `x` starts a chain (`x-1` is not in the set). Each chain is then scanned in full exactly once.
- **Cost**: the inner extension loop runs at most $`n`$ times in total, because each element belongs to exactly one chain and is reached only from that chain's head. This is the aggregate analysis of CLRS §17.1: instead of bounding each inner loop separately, count how many times it runs over the whole algorithm.

### 3.5 Prefix sums plus a hash table

- **What the table stores**: prefix-sum value to its count (or earliest index).
- **When it applies**: counting subarrays with sum $`k`$. A subarray sum is $`S_j - S_i`$, so the problem becomes a complement lookup for $`S_j - k`$.

### 3.6 A key index for another structure

- **What the table stores**: key to the position of an element in another structure, such as a list-node pointer or an array index.
- **When it applies**: another structure (list, array, heap) keeps the order, but elements must also be found by key in $`O(1)`$ and then deleted or moved.
- **Requirements**: every update must touch both sides; when the other structure needs to get back to the table (for example, when evicting the list tail), each element must store its key.

## 4. When a Hash Table Is Unnecessary

1. **Small universe**: with only 26 letters, or $`0 \le x \le 10^5`$, a direct-address table (an array) is faster and has no worst-case degradation.
2. **$`O(n \log n)`$ is acceptable**: sort, then scan linearly. This needs $`O(1)`$ extra space (apart from the sort itself) but loses the original indices, which need separate handling if required.
3. **Order is needed**: queries such as "the smallest element greater than `x`" are beyond a hash table; use a balanced binary search tree or a sorted array.

## 5. Application Problems

| Problem | Location | Pattern | Notes |
|---|---|---|---|
| 49. Group Anagrams: [sorted key](./0049-group-anagrams/sorted-key/summary.md), [count key](./0049-group-anagrams/count-key/summary.md) | This directory | 3.2 Counting; 3.3 Grouping by canonical form | The sorted string or the `[26]int` count vector is the key; `map[key][]string` groups in one pass |
| [128. Longest Consecutive Sequence](./0128-longest-consecutive-sequence/summary.md) | This directory | 3.1 Membership; 3.4 Chain heads | The set holds all elements; extend only from `x` with `x-1` absent, for $`O(n)`$ total time |
| [146. LRU Cache](./0146-lru-cache/summary.md) | This directory | 3.6 A key index for another structure | `map[int]*node` stores doubly-linked-list node pointers, the list is in recency order, and `get` and `put` take expected $`O(1)`$ |

## 6. CLRS Connections

Section numbers follow the third edition.

| Topic | Section | Relation | Where it appears here |
|---|---|---|---|
| Direct-address tables | §11.1 Direct-address tables | Direct | Section 2.1; Section 4, item 1 |
| Chaining and load factor | §11.2 Hash tables | Direct | The expected $`\Theta(1 + \alpha)`$ result in Section 2.1 |
| Amortized analysis of dynamic tables | §17.4 Dynamic tables | Direct | Automatic `map` growth in Section 2.2 |
| Aggregate analysis | §17.1 Aggregate analysis | Direct | Section 3.4: the inner loop runs at most $`n`$ times in total |
