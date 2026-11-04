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

**Four-Dimensional Rating — Hash Map with Head and Tail Sentinel Doubly Linked List**

- Evaluation Scope: $`1 \le \mathrm{capacity} \le 3000`$, $`0 \le \mathrm{key} \le 10^4`$, $`0 \le \mathrm{value} \le 10^5`$, at most $`2 \times 10^5`$ calls to `get` and `put` in total; expected $`O(1)`$ per operation, $`O(\mathrm{capacity})`$ extra space; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 3 | One must view the recency order as a sequence that supports "move to front" and "remove the last", and see that no single structure does both in $`O(1)`$: the hash table locates a node by key, the list keeps the order, and node pointers link the two |
| Knowledge Prerequisites | 3 | Combines two common structures: expected $`O(1)`$ hash-table lookup and deletion, $`O(1)`$ unlink and insert in a doubly linked list given the node, plus the sentinel trick for boundary cases |
| Correctness Reasoning Difficulty | 3 | One must first turn LRU into a precise abstract specification using last-use times, then state a representation invariant linking the list and the hash table, and prove by induction on operations that every branch preserves it |
| Implementation Difficulty | 3 | Several pieces of state must agree: inserting at the head must update the old first node's `prev` before `head.next`, each node must store its key so eviction can delete it from the map, and `Put` must check whether the key exists before inserting and evicting; any slip breaks the list–map consistency |

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

## 3. Initial Approach

This section follows the original Python implementation and its code comments. [solution.go](./solution.go) in the repository is a Go port of that implementation and keeps its structure, helper functions, and order of operations: `DLinkedNode` becomes `dLinkedNode`, and `_remove` and `_add_to_head` become `remove` and `addToHead`.

### First Idea: Hash Map + Doubly Linked List with Head and Tail Sentinels

- A dictionary `cache` finds a node by key, and a doubly linked list keeps the usage order, with the most recently used at the head and the least recently used at the tail.
- The comments say the sentinels are there to "reduce checks": the dummy nodes `head` and `tail` always sit at the two ends, so insertion and removal need no null checks.
- `get`: on a hit, "first remove the current node, then put it at the head", that is, `_remove` followed by `_add_to_head`.
- `put`: for an existing key, update `value` and move it to the head; otherwise create a node, add it to the dictionary and the head first, then check whether the capacity is exceeded and, if so, remove the tail node.

### Key Point: The Order of Pointer Updates When Inserting at the Head

The key point recorded in the comments: inserting involves three nodes, `head`, `A`, and `B`, where `A` is the node being inserted and `B` is the old first node. First set `A`'s `prev` and `next`; after that, "you must update node B first and only then change `head.next`". See [Challenge 3](#challenge-3-why-update-bprev-before-headnext-when-inserting-at-the-head) for why.

## 4. Algorithm Specification

**Algorithm Name:** Hash Map with Head and Tail Sentinel Doubly Linked List

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

The problem does not index into arrays, so the pseudocode has no indices. `head` and `tail` are two sentinels that hold no data: `head.next` is the most recently used node and `tail.prev` is the least recently used node.

```text
CONSTRUCTOR(c)
    head.next ← tail
    tail.prev ← head
    M ← empty hash map             // key → node
    cap ← c

GET(k)
    x ← SEARCH(M, k)
    if x = NIL
        return -1
    REMOVE(x)
    ADD-TO-HEAD(x)
    return x.value

PUT(k, v)
    x ← SEARCH(M, k)
    if x ≠ NIL                      // existing key: update only, no eviction
        x.value ← v
        REMOVE(x)
        ADD-TO-HEAD(x)
        return
    x ← new node with key k, value v
    INSERT(M, k, x)
    ADD-TO-HEAD(x)
    if |M| > cap                    // over capacity: evict the tail
        y ← tail.prev
        REMOVE(y)
        DELETE(M, y.key)

REMOVE(x)                           // unlink x
    x.prev.next ← x.next
    x.next.prev ← x.prev

ADD-TO-HEAD(x)                      // insert x right after head
    x.prev ← head
    x.next ← head.next
    head.next.prev ← x              // update the old first node B first
    head.next ← x                   // then update head.next
```

**State Definitions**

- `head`, `tail`: the fields of the same names in the code, two sentinel nodes that always sit at the ends of the list.
- `M`: `cache` in the code, a hash map from keys to list nodes.
- `cap`: `capacity` in the code.
- Node `x`: `dLinkedNode` in the code, holding `key`, `value`, `prev`, and `next`. The `key` is stored so that evicting the tail node can delete its key from `M`.

In the code, `remove` and `addToHead` correspond to `REMOVE` and `ADD-TO-HEAD`. `GET` returns expressions directly and has no result variable.

### 4.4 Correctness

**Proof Method: Representation Invariant + Induction on the Number of Operations**

#### 1. Key Conditions for Correctness

Let $`x_1, x_2, \ldots, x_s`$ be the nodes visited when walking along `next` from `head.next` until reaching `tail` ($`s = 0`$ means the list is empty).

Correctness relies on:

1. **Dictionary semantics of the hash table** (CLRS §11.2): `SEARCH`, `INSERT`, and `DELETE` correctly implement the map from keys to nodes.
2. **Locality of list operations**: `REMOVE` changes only the pointers of the two neighbors of `x`, removing $`x`$ from the sequence and keeping the relative order of the other nodes; `ADD-TO-HEAD` puts $`x`$ at the front and keeps the order of the others. Because `head` and `tail` always sit at the ends, every real node has a non-null predecessor and successor, so an empty list, a one-node list, and `x` being the first or last node need no special cases (the same use of sentinels to remove boundary cases as in CLRS §10.2).
3. **Distinct last-use times**: the $`\mathrm{last}`$ values in the abstract specification are pairwise distinct, so sorting by $`\mathrm{last}`$ gives a unique order and the least recently used key is unique.

Together these reduce the problem to one claim: as long as the list order always matches decreasing $`\mathrm{last}`$, the tail is the key to evict.

#### 2. Correctness Proof

**Representation Invariant**

After the first $`t`$ operations, the following four statements hold:

**Invariant I: The list is well formed**

`head`, $`x_1, \ldots, x_s`$, `tail` form a doubly linked list in that order: `head.next.prev = head`, `tail.prev.next = tail`, and every $`x_i`$ has `x_i.next.prev = x_i` and `x_i.prev.next = x_i`.

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

At $`t = 0`$, `CONSTRUCTOR` sets `head.next = tail` and `tail.prev = head`, so the list is empty; `M` is empty; $`K_0 = \varnothing`$. All four invariants hold (II, III, and IV hold trivially for an empty list).

**Maintenance**

Assume the invariants hold after the first $`t-1`$ operations and consider operation $`t`$. By Invariant II, `SEARCH(M, k)` returns a non-NIL node if and only if $`k \in K_{t-1}`$, and the node it returns is the list node $`x = x_j`$ whose key is $`k`$.

- **Case 1: GET miss ($`k \notin K_{t-1}`$)**. The algorithm returns $`-1`$ and modifies nothing; the specification also returns $`-1`$ and leaves the state unchanged. The invariants are preserved.
- **Case 2: GET hit ($`k = x_j.\mathrm{key}`$)**. By Invariant III, the returned `x.value` equals $`\mathrm{val}_{t-1}(k)`$, matching the specification. `REMOVE` followed by `ADD-TO-HEAD` turns the sequence into $`x_j, x_1, \ldots, x_{j-1}, x_{j+1}, \ldots, x_s`$; the set of nodes is unchanged, so I and II are preserved, and no value changes, so III is preserved. The specification sets $`\mathrm{last}_t(k) = t`$, larger than every other key's $`\mathrm{last}`$, while the other keys keep their $`\mathrm{last}`$ and relative order, so the new sequence is still in decreasing order of $`\mathrm{last}_t`$, and IV is preserved.
- **Case 3: PUT on an existing key ($`k = x_j.\mathrm{key}`$)**. The algorithm sets $`x_j.\mathrm{value}`$ to $`v`$ and moves it to the front without evicting. In the specification the key set is unchanged, $`\mathrm{val}_t(k) = v`$, and $`\mathrm{last}_t(k) = t`$. III holds for $`x_j`$ and is unaffected for other nodes; I, II, and IV follow as in Case 2.
- **Case 4: PUT on a new key ($`k \notin K_{t-1}`$)**. The algorithm inserts first and evicts only if needed, the reverse of the specification's "evict, then insert"; we show both give the same result.
  - *Insertion*: a new node $`x_0`$ holding $`k`$ and $`v`$ is added to `M` and inserted at the head, so the sequence becomes $`x_0, x_1, \ldots, x_s`$. Since $`k`$ was not in $`K_{t-1}`$, the keys in the list stay pairwise distinct, and the key sets of the list and of `M` both become $`K_{t-1} \cup \{k\}`$, of size $`s + 1`$; the new node's value is $`v`$; the last-use time of $`k`$ is $`t`$, larger than every other key, and the node is at the head, so the sequence is still in decreasing order of last-use time; I is preserved by Key Condition 2.
  - *Eviction*: Invariant II gives $`s = \lvert K_{t-1} \rvert \le c`$, so $`\lvert M \rvert = s + 1 > c`$ if and only if $`\lvert K_{t-1} \rvert = c`$, exactly the condition under which the specification evicts. Then $`s = c \ge 1`$, so the list holds at least the two real nodes $`x_0`$ and $`x_s`$, and `tail.prev` is $`x_s`$, not the new node $`x_0`$. By Invariant IV (for $`K_{t-1}`$), $`x_s`$ holds the least recently used key of $`K_{t-1}`$, the very key the specification removes. The algorithm unlinks $`x_s`$ and then deletes from `M` using the key stored in $`x_s`$, so the list and `M` both lose that key, leaving $`x_0, x_1, \ldots, x_{s-1}`$ with key set exactly the specification's $`K_t`$. If $`s < c`$, neither the specification nor the algorithm evicts, and the key set is $`K_{t-1} \cup \{k\} = K_t`$.
  - In both situations I–IV hold for $`K_t`$: II follows from the equal key sets, III because the new node's value is $`v = \mathrm{val}_t(k)`$, and IV because removing the last element of a decreasing sequence leaves it decreasing.

Hence the invariants still hold after operation $`t`$.

**Termination**

Each operation runs a constant number of statements with no loops, so each operation terminates. After the induction reaches $`t = q`$, Cases 1 and 2 show that every GET returns the value the abstract specification prescribes; Invariant II and the eviction condition in Case 4 give $`\lvert M \rvert \le c`$ at the end of every operation (it is briefly $`c + 1`$ between the insertion and the eviction of a PUT on a new key).

#### 3. Conclusion

Invariant IV makes "the least recently used key" always equal to the tail, and Invariant II lets any key reach its node directly through the hash table. The three things the LRU specification needs (read a value by key, mark a key as most recent, evict the least recent key) therefore become one hash lookup, one `REMOVE` plus `ADD-TO-HEAD`, and one tail unlink, and the algorithm produces results matching the specification for every valid operation sequence.

Three orders in the code must not be swapped: `Put` checks whether the key exists before inserting and evicting, otherwise an existing key would be treated as a new one (violating Case 3); `ADD-TO-HEAD` must update the old first node's `prev` before `head.next` (see [Challenge 3](#challenge-3-why-update-bprev-before-headnext-when-inserting-at-the-head)); and eviction deletes the map entry using the key stored in the tail node, which is why the code calls `delete(c.cache, lru.key)` with the key of the node just unlinked.

### 4.5 Complexity Analysis

Let $`c`$ be the capacity and $`q`$ the number of operations. "Expected" below refers to the expected $`O(1)`$ time of a single hash-table operation (see [overview §2.1](../README.md)).

**Time Complexity: expected O(1) per operation, expected O(q) in total**

- `GET`: one hash lookup, plus one `REMOVE` and one `ADD-TO-HEAD` on a hit, rewriting a constant number of pointers, worst case $`O(1)`$.
- `PUT`: one lookup; for an existing key, the same as `GET`; for a new key, one map insertion and one head insertion, plus at most one tail unlink and one map deletion.

The list part of each operation is worst case $`O(1)`$ and the hash-table part is expected $`O(1)`$, so each operation takes expected $`O(1)`$ and $`q`$ operations take expected $`O(q)`$. Under heavy collisions a single hash-table operation can degrade to $`\Theta(c)`$; the Go runtime uses a random hash seed per `map`, so inputs cannot force this.

**Space Complexity: O(c)**

By Invariant II, the list has at most $`c`$ real nodes and `M` has at most $`c`$ entries (briefly $`c + 1`$ during a PUT on a new key), plus two sentinel nodes, so the extra space is $`O(c)`$. This counts live nodes and entries; a Go `map` does not return memory to the system after deletions, and the exact memory layout is up to the runtime and not discussed here.

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
- **Sentinels**: without them, an empty list, a one-node list, and unlinking the first or last node all need `nil` checks. With the two sentinels `head` and `tail`, the most recently used node is `head.next`, the least recently used is `tail.prev`, every real node has a predecessor and a successor, and `REMOVE` and `ADD-TO-HEAD` need no branches. CLRS §10.2 uses a single circular sentinel `L.nil` as both ends, which works the same way with one node fewer.

**Transferable Insight:** When two structures refer to each other, check that navigation in both directions takes $`O(1)`$: from the map to the list via stored pointers, and from the list back to the map via the key stored in each node. Missing either direction makes some operation degrade.

### Challenge 3: Why update B.prev before head.next when inserting at the head?

This is the key point written down in the original implementation's comments: inserting `A` involves three nodes, `head`, `A`, and `B`, where `B` is the old first node, and "you must update node B first and only then change `head.next`".

**Understanding Process**

In `ADD-TO-HEAD`, `B` is not held in its own variable; the code reaches it through `head.next` (the third line is `head.next.prev`). If `head.next ← A` ran first and `head.next.prev ← A` second, the second line's `head.next` would already be `A`, so it would become `A.prev ← A`: `A`'s predecessor points to itself and `B.prev` still points to `head`. The list is now broken, and any later unlink of `A` or `B` rewrites the wrong pointers.

The tests confirm this: swapping those two lines in the Go code makes `go test` fail at once, because the tests check Invariant I after every call, and `A.prev = A` violates `A.prev.next = A`.

**Transferable Insight:** When rewriting pointers, first ask "which node can only be reached through the pointer about to be overwritten", handle it first, and only then overwrite that pointer. Alternatively, save `B := head.next` in a temporary variable, after which the order of the assignments no longer matters.

### Challenge 4: In `Put`, check existence first, then insert, then evict

**Understanding Process**

`Put` first checks whether the key exists: an existing key only has its value and order updated, and capacity is not involved. If capacity were checked first, updating an existing key in a full cache would first evict the tail, throwing out a key that should have stayed. The test case "Updating an existing key does not evict" covers this branch.

For a new key, this implementation inserts first and evicts the tail when `len(cache) > capacity`, instead of evicting when `len(cache) == capacity` and then inserting. Both give the same result (the proof of Case 4), but inserting first requires checking one thing: the evicted node is never the one just inserted. The new node is at the head, and when the capacity is exceeded the list has $`c + 1 \ge 2`$ nodes, so the tail is always another node. The cost is that `cache` briefly holds one extra entry, which is why the Go code preallocates with `make(map[int]*dLinkedNode, capacity+1)`.

**Transferable Insight:** In design problems, "update an existing element" and "insert a new element" usually take different branches, and constraints such as capacity or counts change only on insertion. Before writing code, split the specification by "does the key already exist", then decide the order of the checks; if you swap the order of "insert" and "evict", check that at the boundary the evicted element is still the one the specification requires.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Hash Map to Node Pointer**: here `M[k]` gives the list node directly, reducing "find the node by key" to expected $`O(1)`$. It applies to design problems that need a key index on top of an ordered structure (list, heap, array), such as an LFU cache or $`O(1)`$ insert/delete/get-random. The limits are that every update must touch both sides and the bound is only expected $`O(1)`$.
2. **Doubly Linked List with Sentinel**: here it gives $`O(1)`$ unlink of a known node and insertion at the head, and `tail.prev` gives the least recently used node directly. It applies whenever elements must be removed from arbitrary positions and inserted at an end. It requires that the caller already holds the node pointer; knowing only the value still requires a search.
3. **Recency Order as Move-to-Front**: moving an element to the head on every use keeps the list in decreasing order of last-use time (Invariant IV). It applies to LRU-style problems ordered by the time of last access; if the ordering is by access count (LFU), moving to the head is not enough and counts must be bucketed.
4. **Representation Invariant**: here "the list and the hash table agree, and the list is in decreasing order of last-use time" is stated as an invariant and proved by induction on the number of operations. It applies to every data-structure design problem: write the abstract specification, write the invariant the concrete structure satisfies, then prove for each operation and each branch that the invariant is preserved and the return value matches the specification.

Go's standard `container/list` is a circular doubly linked list with a single sentinel, and its `MoveToFront`, `Remove`, `PushFront`, and `Back` match the list operations used here. This solution follows the original implementation and writes the list by hand, so every pointer update lines up with the pseudocode and the proof.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. "Relation" is either "Direct" or "Related idea"; when no specific section matches, the "CLRS Section" column says "No matching section" and only the concept name is kept.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Doubly linked lists and sentinels | §10.2 Linked lists | Direct | `remove` is LIST-DELETE' and `addToHead` does the same job as LIST-INSERT'; this solution uses two sentinels `head` and `tail` instead of CLRS's single circular sentinel `L.nil`, and the least recently used node is `tail.prev` |
| Hash tables with chaining | §11.2 Hash tables | Direct | Lookup, insertion, and deletion in `cache` take expected $`O(1)`$, which the expected $`O(1)`$ per operation relies on; the section notes that a doubly linked list deletes a given element in $`O(1)`$, and the hash table here turns a key into that element |
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

The hash table finds a node by key, the doubly linked list with head and tail sentinels keeps nodes in recency order and moves them in $`O(1)`$, and storing the key in each node lets each side reach the other, so both `get` and `put` take expected $`O(1)`$ time.
