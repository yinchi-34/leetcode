# 146. LRU Cache

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md). The hash table itself (its definition, Go `map` behavior, and common usage patterns) is covered in the [hash table overview](../README.md); this note covers only how this problem combines a hash table with a doubly linked list.

## 1. Problem Information

- **Problem**: [146. LRU Cache](https://leetcode.cn/problems/lru-cache/)
- **Official Difficulty**: Medium
- **Primary Category**: Hash Table
- **Secondary Category**: Linked List, Design
- **Tags**: `Hash Table`, `Linked List`, `Design`, `Doubly Linked List`, `Sentinel`, `Representation Invariant`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Hash Map with Sentinel Doubly Linked List**

- Evaluation Scope: $`1 \le \mathrm{capacity} \le 3000`$, $`0 \le \mathrm{key} \le 10^4`$, $`0 \le \mathrm{value} \le 10^5`$, at most $`2 \times 10^5`$ calls to `get` and `put` in total; expected $`O(1)`$ per operation, $`O(\mathrm{capacity})`$ extra space; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 3 | One must view the recency order as a sequence that supports "move to front" and "remove the last", and see that no single structure does both in $`O(1)`$: the hash table locates a node by key, the list keeps the order, and node pointers link the two |
| Knowledge Prerequisites | 3 | Combines two common structures: expected $`O(1)`$ hash-table lookup and deletion, $`O(1)`$ unlink and insert in a doubly linked list given the node, plus the sentinel trick for boundary cases |
| Correctness Reasoning Difficulty | 3 | One must first turn LRU into a precise abstract specification using last-use times, then state a representation invariant linking the list and the hash table, and prove by induction on operations that every branch preserves it |
| Implementation Difficulty | 3 | Several pieces of state must agree: the order of pointer updates, storing the key in each node so eviction can delete it from the map, and checking whether the key exists before checking capacity in `Put`; any slip breaks the list–map consistency |

Ratings follow [RATING.md](../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

Design a cache of size `capacity` that stores key-value pairs and supports two operations:

- `get(key)`: return the key's value if the key is in the cache, otherwise `-1`.
- `put(key, value)`: if the key exists, update its value; otherwise insert the pair. If the insertion makes the number of keys exceed the capacity, evict the **least recently used** key.

A `get` hit or any `put` counts as a "use" of that key. Both operations must run in $`O(1)`$ average time.

**Example**

- Input: `["LRUCache","put","put","get","put","get","put","get","get","get"]` with arguments `[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]`
- Output: `[null,null,null,1,null,-1,null,-1,3,4]`

Trace: the capacity is 2. After putting 1 and 2, `get(1)` makes 1 the most recent, so putting 3 evicts 2. Putting 4 later evicts 1, which is now the least recent, leaving 3 and 4.

**Key Requirements**

1. Constraints (from the LeetCode statement): $`1 \le \mathrm{capacity} \le 3000`$, $`0 \le \mathrm{key} \le 10^4`$, $`0 \le \mathrm{value} \le 10^5`$, at most $`2 \times 10^5`$ calls.
2. Both `get` and `put` must take $`O(1)`$ average time, so one cannot scan all keys to find the least recently used one.
3. Edge cases: with capacity 1, every new key evicts the old one; `put` on an existing key only updates its value and recency and never evicts; a `get` miss changes no key's recency; values may be 0, so 0 cannot mean "absent".

## 4. Algorithm Specification

**Algorithm Name:** Hash Map with Sentinel Doubly Linked List

### 4.1 Input

A capacity $`c \ge 1`$ and a sequence of operations in time order

```math
O = \langle o_1, o_2, \ldots, o_q \rangle
```

where:

- Each $`o_t`$ is $`\mathrm{GET}(k)`$ or $`\mathrm{PUT}(k, v)`$ with nonnegative integers $`k`$ and $`v`$.
- $`t`$ also serves as the time at which the $`t`$-th operation happens, so different operations have different times.

### 4.2 Output

First we write the LRU semantics of the problem as a precise abstract specification. After the first $`t`$ operations, let $`K_t`$ be the set of keys in the cache ($`K_0 = \varnothing`$). Each $`k \in K_t`$ has a current value $`\mathrm{val}_t(k)`$ and a **last-use time** $`\mathrm{last}_t(k)`$, the time of the most recent GET hit or PUT on $`k`$. The $`t`$-th operation acts as follows:

- $`\mathrm{GET}(k)`$ with $`k \notin K_{t-1}`$: return $`-1`$; the state does not change.
- $`\mathrm{GET}(k)`$ with $`k \in K_{t-1}`$: return $`\mathrm{val}_{t-1}(k)`$ and set $`\mathrm{last}_t(k) = t`$; nothing else changes.
- $`\mathrm{PUT}(k, v)`$ with $`k \in K_{t-1}`$: set $`\mathrm{val}_t(k) = v`$ and $`\mathrm{last}_t(k) = t`$; the key set does not change.
- $`\mathrm{PUT}(k, v)`$ with $`k \notin K_{t-1}`$: if $`\lvert K_{t-1} \rvert = c`$, first remove the key with the smallest $`\mathrm{last}`$ (the least recently used); then add $`k`$ with $`\mathrm{val}_t(k) = v`$ and $`\mathrm{last}_t(k) = t`$.

Each time step uses at most one key, so the keys in $`K_{t-1}`$ have pairwise distinct $`\mathrm{last}`$ values, and "the key with the smallest $`\mathrm{last}`$" is unique.

The output must satisfy:

- Every GET returns the value given by the specification above.
- By the last case of the specification, $`\lvert K_t \rvert \le c`$ always holds.

### 4.3 Pseudocode

The problem does not index into arrays, so the pseudocode has no indices. `L.nil` is the sentinel; following CLRS §10.2, `L.nil.next` is the head and `L.nil.prev` is the tail.

```text
CONSTRUCTOR(c)
    L.nil.next ← L.nil
    L.nil.prev ← L.nil
    M ← empty hash map             // key → node
    cap ← c

GET(k)
    x ← SEARCH(M, k)
    if x = NIL
        return -1
    MOVE-TO-FRONT(x)
    return x.value

PUT(k, v)
    x ← SEARCH(M, k)
    if x ≠ NIL                      // existing key: update only, no eviction
        x.value ← v
        MOVE-TO-FRONT(x)
        return
    if |M| = cap                    // cache full: evict the tail
        y ← L.nil.prev
        LIST-DELETE'(L, y)
        DELETE(M, y.key)
    x ← new node with key k, value v
    LIST-INSERT'(L, x)
    INSERT(M, k, x)

LIST-DELETE'(L, x)                  // unlink x
    x.prev.next ← x.next
    x.next.prev ← x.prev

LIST-INSERT'(L, x)                  // insert x at the head
    x.next ← L.nil.next
    L.nil.next.prev ← x
    L.nil.next ← x
    x.prev ← L.nil

MOVE-TO-FRONT(x)
    LIST-DELETE'(L, x)
    LIST-INSERT'(L, x)
```

**State Definitions**

- `L.nil`: `sentinel` in the code, a node that holds no data and closes the list into a cycle.
- `M`: `nodes` in the code, a hash map from keys to list nodes.
- `cap`: `capacity` in the code.
- Node `x`: `node` in the code, holding `key`, `value`, `prev`, and `next`. The `key` is stored so that evicting the tail node can delete its key from `M`.

In the code, `remove`, `pushFront`, and `moveToFront` correspond to `LIST-DELETE'`, `LIST-INSERT'`, and `MOVE-TO-FRONT`. `GET` returns expressions directly and has no result variable.

### 4.4 Correctness

**Proof Method: Representation Invariant + Induction on the Number of Operations**

#### 1. Key Conditions for Correctness

Let $`x_1, x_2, \ldots, x_s`$ be the nodes visited when walking along `next` from `L.nil.next` until reaching `L.nil` ($`s = 0`$ means the list is empty).

Correctness relies on:

1. **Dictionary semantics of the hash table** (CLRS §11.2): `SEARCH`, `INSERT`, and `DELETE` correctly implement the map from keys to nodes.
2. **Locality of list operations**: `LIST-DELETE'` changes only the pointers of the two neighbors of `x`, removing $`x`$ from the sequence and keeping the relative order of the other nodes; `LIST-INSERT'` puts $`x`$ at the front and keeps the order of the others. With the sentinel, an empty list, a one-node list, and `x` being the head or the tail need no special cases (CLRS §10.2).
3. **Distinct last-use times**: the $`\mathrm{last}`$ values in the abstract specification are pairwise distinct, so sorting by $`\mathrm{last}`$ gives a unique order and the least recently used key is unique.

Together these reduce the problem to one claim: as long as the list order always matches decreasing $`\mathrm{last}`$, the tail is the key to evict.

#### 2. Correctness Proof

**Representation Invariant**

After the first $`t`$ operations, the following four statements hold:

**Invariant I: The list is well formed**

$`x_1, \ldots, x_s`$ and `L.nil` form a circular doubly linked list: every node $`y`$ (including the sentinel) has `y.next.prev = y` and `y.prev.next = y`.

**Invariant II: The hash table and the list correspond one to one**

The keys of $`x_1, \ldots, x_s`$ are pairwise distinct, and $`\{x_1.\mathrm{key}, \ldots, x_s.\mathrm{key}\} = K_t`$; the key set of `M` is also exactly $`K_t`$, and for each $`k \in K_t`$, `M[k]` is the list node whose key is $`k`$. In particular, $`s = \lvert M \rvert = \lvert K_t \rvert`$.

**Invariant III: Values are correct**

For each $`i`$, $`x_i.\mathrm{value} = \mathrm{val}_t(x_i.\mathrm{key})`$.

**Invariant IV: The list is in decreasing order of last-use time**

```math
\mathrm{last}_t(x_1.\mathrm{key}) > \mathrm{last}_t(x_2.\mathrm{key}) > \cdots > \mathrm{last}_t(x_s.\mathrm{key})
```

That is, the head holds the most recently used key and the tail $`x_s`$ holds the least recently used key.

**Initialization**

At $`t = 0`$, `CONSTRUCTOR` points both `next` and `prev` of `L.nil` to itself, so the list is empty; `M` is empty; $`K_0 = \varnothing`$. All four invariants hold (II, III, and IV hold trivially for an empty list).

**Maintenance**

Assume the invariants hold after the first $`t-1`$ operations and consider operation $`t`$. By Invariant II, `SEARCH(M, k)` returns a non-NIL node if and only if $`k \in K_{t-1}`$, and the node it returns is the list node $`x = x_j`$ whose key is $`k`$.

- **Case 1: GET miss ($`k \notin K_{t-1}`$)**. The algorithm returns $`-1`$ and modifies nothing; the specification also returns $`-1`$ and leaves the state unchanged. The invariants are preserved.
- **Case 2: GET hit ($`k = x_j.\mathrm{key}`$)**. By Invariant III, the returned `x.value` equals $`\mathrm{val}_{t-1}(k)`$, matching the specification. `MOVE-TO-FRONT` turns the sequence into $`x_j, x_1, \ldots, x_{j-1}, x_{j+1}, \ldots, x_s`$; the set of nodes is unchanged, so I and II are preserved, and no value changes, so III is preserved. The specification sets $`\mathrm{last}_t(k) = t`$, larger than every other key's $`\mathrm{last}`$, while the other keys keep their $`\mathrm{last}`$ and relative order, so the new sequence is still in decreasing order of $`\mathrm{last}_t`$, and IV is preserved.
- **Case 3: PUT on an existing key ($`k = x_j.\mathrm{key}`$)**. The algorithm sets $`x_j.\mathrm{value}`$ to $`v`$ and moves it to the front without evicting. In the specification the key set is unchanged, $`\mathrm{val}_t(k) = v`$, and $`\mathrm{last}_t(k) = t`$. III holds for $`x_j`$ and is unaffected for other nodes; I, II, and IV follow as in Case 2.
- **Case 4: PUT on a new key ($`k \notin K_{t-1}`$)**. Two steps.
  - *Eviction*: if $`\lvert M \rvert = c`$ ($`c \ge 1`$, so the list is nonempty), Invariant II gives $`\lvert K_{t-1} \rvert = c`$, and the specification removes the least recently used key. By Invariant IV this is the tail $`x_s`$, and by Invariant I `L.nil.prev` is exactly $`x_s`$. The algorithm unlinks $`x_s`$ and then deletes from `M` using the key stored in $`x_s`$, so the list and `M` both lose the same key, the one the specification removes. The remaining sequence $`x_1, \ldots, x_{s-1}`$ still satisfies I–IV. If $`\lvert M \rvert < c`$, neither the specification nor the algorithm evicts.
  - *Insertion*: a new node holding $`k`$ and $`v`$ is inserted at the head, and `M[k]` is set to it. Since $`k`$ was not in $`K_{t-1}`$, the keys in the list stay pairwise distinct and the key set equals the specification's $`K_t`$, so II is preserved; the new node's value is $`v = \mathrm{val}_t(k)`$, so III is preserved; $`\mathrm{last}_t(k) = t`$ exceeds every other key and the node is at the head, so IV is preserved; I is preserved by Key Condition 2.

Hence the invariants still hold after operation $`t`$.

**Termination**

Each operation runs a constant number of statements with no loops, so each operation terminates. After the induction reaches $`t = q`$, Cases 1 and 2 show that every GET returns the value the abstract specification prescribes; Invariant II and the eviction condition in Case 4 give $`\lvert M \rvert \le c`$ at all times.

#### 3. Conclusion

Invariant IV makes "the least recently used key" always equal to the tail, and Invariant II lets any key reach its node directly through the hash table. The three things the LRU specification needs (read a value by key, mark a key as most recent, evict the least recent key) therefore become one hash lookup, one `MOVE-TO-FRONT`, and one tail unlink, and the algorithm produces results matching the specification for every valid operation sequence.

Two orders in the code must not be swapped: `Put` checks whether the key exists before checking capacity, otherwise updating an existing key in a full cache would wrongly evict a key (violating Case 3); and eviction reads the key from the tail node before deleting the map entry, which is why the code calls `delete(c.nodes, lru.key)` with the key stored in the node just unlinked.

### 4.5 Complexity Analysis

Let $`c`$ be the capacity and $`q`$ the number of operations. "Expected" below refers to the expected $`O(1)`$ time of a single hash-table operation (see [overview §2.1](../README.md)).

**Time Complexity: expected O(1) per operation, expected O(q) in total**

- `GET`: one hash lookup, plus one `MOVE-TO-FRONT` on a hit. `MOVE-TO-FRONT` is `LIST-DELETE'` plus `LIST-INSERT'`, rewriting a constant number of pointers, worst case $`O(1)`$.
- `PUT`: one lookup; for an existing key, the same as `GET`; for a new key, at most one tail unlink, one map deletion, one head insertion, and one map insertion.

The list part of each operation is worst case $`O(1)`$ and the hash-table part is expected $`O(1)`$, so each operation takes expected $`O(1)`$ and $`q`$ operations take expected $`O(q)`$. Under heavy collisions a single hash-table operation can degrade to $`\Theta(c)`$; the Go runtime uses a random hash seed per `map`, so inputs cannot force this.

**Space Complexity: O(c)**

By Invariant II, at any time the list has at most $`c`$ nodes and `M` has at most $`c`$ entries, plus one sentinel node, so the extra space is $`O(c)`$. This counts live nodes and entries; a Go `map` does not return memory to the system after deletions, and the exact memory layout is up to the runtime and not discussed here.

## 5. Learning Notes

### Challenge 1: Why two structures? Wouldn't one do?

An LRU cache must answer two kinds of questions, "where is the node for key $`k`$" and "which key is least recently used", and must also be able to move any key to the most-recent position.

**Understanding Process**

Rule out single structures one by one:

- Hash table only: $`O(1)`$ lookup by key, but a hash table keeps no order, so finding the least recently used key means scanning all keys, $`O(c)`$.
- Linked list only: the order is there, but lookup by key walks from the head, $`O(c)`$.
- A slice of keys in recency order: moving a key from the middle to the end shifts the elements after it, $`O(c)`$.

When the two are combined, the hash table stores not the value but a **pointer to the list node**. One hash lookup then lands directly on a position in the list, and the list can unlink or insert a known node in $`O(1)`$. Each structure handles one kind of query, and the pointers tie them together.

**Transferable Insight:** When a structure must support both "locate by key" and "locate by order" and no single structure is good at both, use two structures, one for each, and let one store the "position" (pointer or index) of elements in the other. The price is that every update must keep both sides consistent, which is exactly what the representation invariant proves.

### Challenge 2: Why must the list be doubly linked, and why does each node also store its key?

**Understanding Process**

- **Doubly linked**: unlinking node $`x`$ requires changing the `next` of its predecessor. A singly linked list has no `prev`, so finding the predecessor means walking from the head, $`O(c)`$. CLRS §11.2 makes the same point when discussing deletion with chaining: given the element itself (not its key), a doubly linked list can delete it in $`O(1)`$. In this problem the hash table is what turns a key into the element itself.
- **Key in the node**: eviction reaches the tail node from the list side and then must delete the map entry pointing to it, which requires its key. If a node stored only the value, there would be no way back from the list to the hash table.
- **Sentinel**: without one, an empty list, a one-node list, and unlinking the head or tail all need `nil` checks. With a circular list and a sentinel, the head is `L.nil.next`, the tail is `L.nil.prev`, every node has a predecessor and a successor, and `LIST-DELETE'` and `LIST-INSERT'` need no branches.

**Transferable Insight:** When two structures refer to each other, check that navigation in both directions takes $`O(1)`$: from the map to the list via stored pointers, and from the list back to the map via the key stored in each node. Missing either direction makes some operation degrade.

### Challenge 3: The order of the two checks in `Put`

**Understanding Process**

If `Put` first checked "the cache is full, so evict the tail" and only then checked whether the key exists, a `put` on an existing key in a full cache would first evict the tail key. Unless that key is the one at the tail, this evicts a key that should have stayed, and the cache loses a key. The specification says an existing key is only updated, never causes eviction, so existence must be checked first. The test case "Updating an existing key does not evict" covers this branch.

**Transferable Insight:** In design problems, "update an existing element" and "insert a new element" usually take different branches, and constraints such as capacity or counts change only on insertion. Before writing code, split the specification by "does the key already exist", then decide the order of the checks.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Hash Map to Node Pointer**: here `M[k]` gives the list node directly, reducing "find the node by key" to expected $`O(1)`$. It applies to design problems that need a key index on top of an ordered structure (list, heap, array), such as an LFU cache or $`O(1)`$ insert/delete/get-random. The limits are that every update must touch both sides and the bound is only expected $`O(1)`$.
2. **Doubly Linked List with Sentinel**: here it gives $`O(1)`$ unlink of a known node and insertion at the head, and `L.nil.prev` gives the tail directly. It applies whenever elements must be removed from arbitrary positions and inserted at an end. It requires that the caller already holds the node pointer; knowing only the value still requires a search.
3. **Recency Order as Move-to-Front**: moving an element to the head on every use keeps the list in decreasing order of last-use time (Invariant IV). It applies to LRU-style problems ordered by the time of last access; if the ordering is by access count (LFU), moving to the head is not enough and counts must be bucketed.
4. **Representation Invariant**: here "the list and the hash table agree, and the list is in decreasing order of last-use time" is stated as an invariant and proved by induction on the number of operations. It applies to every data-structure design problem: write the abstract specification, write the invariant the concrete structure satisfies, then prove for each operation and each branch that the invariant is preserved and the return value matches the specification.

Go's standard `container/list` is also a circular doubly linked list with a sentinel, and its `MoveToFront`, `Remove`, `PushFront`, and `Back` match the four list operations here. This solution writes the list by hand so that every pointer update lines up with the pseudocode and the proof.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. "Relation" is either "Direct" or "Related idea"; when no specific section matches, the "CLRS Section" column says "No matching section" and only the concept name is kept.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Doubly linked lists and sentinels | §10.2 Linked lists | Direct | `remove` and `pushFront` are LIST-DELETE' and LIST-INSERT' with a sentinel, and the tail is `L.nil.prev` |
| Hash tables with chaining | §11.2 Hash tables | Direct | Lookup, insertion, and deletion in `nodes` take expected $`O(1)`$, which the expected $`O(1)`$ per operation relies on; the section notes that a doubly linked list deletes a given element in $`O(1)`$, and the hash table here turns a key into that element |
| Loop invariants | §2.1 Insertion sort | Related idea | The representation invariant follows the initialization–maintenance–termination structure, with maintenance taken over each operation rather than each loop iteration |
| Representation invariant | No matching section | Related idea | Consistency of the list and the hash table (Invariants I–IV) is proved by induction on the number of operations |

## 7. Review Log

| Date | Independent Completion | Understanding Gained | Areas for Improvement |
|---|---|---|---|
| To be recorded | To be recorded | To be recorded | To be recorded |

### Next Review Checklist

- [ ] Can I implement the algorithm without looking at the solution?
- [ ] Can I accurately define the input and output?
- [ ] Can I write the pseudocode independently?
- [ ] Can I explain why the algorithm is correct?
- [ ] Can I derive its time and space complexity?
- [ ] Can I recognize the same algorithmic pattern in other problems?

## 8. Related Problems

- [460. LFU Cache](https://leetcode.cn/problems/lfu-cache/) — eviction is by use count instead of last-use time (ties broken by recency); a common approach buckets keys by count, and each bucket is this problem's "hash table + doubly linked list".
- [432. All O\`one Data Structure](https://leetcode.cn/problems/all-oone-data-structure/) — also stores list-node pointers in a hash table, with the list ordered by count, so incrementing or decrementing a count only moves a node to a neighboring position.
- [380. Insert Delete GetRandom O(1)](https://leetcode.cn/problems/insert-delete-getrandom-o1/) — the same "hash table stores a position in another structure" idea: the map stores each value's index in an array, and deletion swaps with the last element; compare with this problem's map storing list-node pointers.
- [707. Design Linked List](https://leetcode.cn/problems/design-linked-list/) — practice list insertion and deletion on their own; a sentinel node removes the head and tail special cases.

---

**One-Sentence Summary**

The hash table finds a node by key, the sentinel doubly linked list keeps nodes in recency order and moves them in $`O(1)`$, and storing the key in each node lets each side reach the other, so both `get` and `put` take expected $`O(1)`$ time.
