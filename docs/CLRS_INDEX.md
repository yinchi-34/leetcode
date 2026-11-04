# CLRS Index

English | [简体中文](CLRS_INDEX.zh-CN.md)

Section numbers follow CLRS, 3rd edition. This index collects the entries with section numbers from each problem's "CLRS Connections" table, sorted by section. Only sections linked to at least one problem are listed; concepts marked "No matching section" are not included.

Each section gets one table:

```markdown
## Chapter 16 Greedy Algorithms

### §16.2 Elements of the greedy strategy

| Problem | Relation | How It Appears Here |
|---|---|---|
| [Number. Title](../Algorithms/<category>/<problem>/summary.md) | Direct / Related idea | Which step of this problem |
```

## Chapter 2 Getting Started

### §2.1 Insertion sort

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Direct | Proves correctness with the outer-loop invariant "the first $`i`$ rows are correct" in the initialization, maintenance, and termination steps |
| [3644. Maximum K to Sort a Permutation](../Algorithms/bit-manipulation/3644-maximum-k-to-sort-a-permutation/summary.md) | Direct | Proves the code computes $`M`$ with the invariant "`res` is the AND of the support elements in the prefix" in the initialization, maintenance, and termination steps |
| [49. Group Anagrams (sorted key)](../data-structures/hash-table/0049-group-anagrams/sorted-key/summary.md) | Direct | Lemma 2 uses the invariant "the key set and each group match exactly the first $`i`$ strings", proved by initialization, maintenance, and termination, to show each string lands in exactly one group |
| [49. Group Anagrams (count key)](../data-structures/hash-table/0049-group-anagrams/count-key/summary.md) | Direct | Lemma 3 uses the invariant "`cnt` holds the letter counts of the prefix", proved by initialization, maintenance, and termination, to show the counts are correct |
| [128. Longest Consecutive Sequence](../data-structures/hash-table/0128-longest-consecutive-sequence/summary.md) | Direct | Lemma 1's inner invariant and the outer invariant "`res` is the longest run among processed heads" are each proved by initialization, maintenance, and termination |
| [453. Minimum Moves to Equal Array Elements (batched simulation)](../Algorithms/simulation/0453-minimum-moves-to-equal-array-elements/batched-simulation/summary.md) | Direct | The batched simulation is proved by the invariants "equal prefix, untouched suffix", through initialization, maintenance, and termination |
| [94. Binary Tree Inorder Traversal (Explicit Stack)](../data-structures/binary-tree/0094-binary-tree-inorder-traversal/iterative/summary.md) | Direct | Invariant I (output decomposition) and Invariant II (the stack is a leftward path) are each proved by initialization, maintenance, and termination |

## Chapter 3 Growth of Functions

### §3.2 Standard notations and common functions

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Related idea | The Fibonacci recurrence is the same kind as this problem's, and the shallow-diagonal sums of Pascal's triangle are Fibonacci numbers |

## Chapter 7 Quicksort

### §7.1 Description of quicksort

| Problem | Relation | How It Appears Here |
|---|---|---|
| [462. Minimum Moves to Equal Array Elements II (two-way)](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/lomuto-quickselect/summary.md) | Direct | `partition` is PARTITION, proved correct by its four-region loop invariant |
| [462. Minimum Moves to Equal Array Elements II (three-way)](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/three-way-quickselect/summary.md) | Related idea | Three-way partitioning reuses PARTITION's region-based invariant proof, extending the regions from two classes to three |

### §7.3 A randomized version of quicksort

| Problem | Relation | How It Appears Here |
|---|---|---|
| [462. Minimum Moves to Equal Array Elements II (two-way)](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/lomuto-quickselect/summary.md) | Direct | `randomPartition` is RANDOMIZED-PARTITION |
| [462. Minimum Moves to Equal Array Elements II (three-way)](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/three-way-quickselect/summary.md) | Direct | `randomPartition3` matches RANDOMIZED-PARTITION: swap a random element into `a[r]`, then partition |

### Problem 7-2 Quicksort with equal element values

| Problem | Relation | How It Appears Here |
|---|---|---|
| [462. Minimum Moves to Equal Array Elements II (two-way)](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/lomuto-quickselect/summary.md) | Direct | An all-equal array makes this solution degrade to $`\Theta(n^2)`$, exactly the case that problem discusses |
| [462. Minimum Moves to Equal Array Elements II (three-way)](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/three-way-quickselect/summary.md) | Direct | Three-way partitioning puts elements equal to the pivot in their own block, avoiding $`\Theta(n^2)`$ on all-equal arrays |

## Chapter 8 Sorting in Linear Time

### §8.2 Counting sort

| Problem | Relation | How It Appears Here |
|---|---|---|
| [49. Group Anagrams (count key)](../data-structures/hash-table/0049-group-anagrams/count-key/summary.md) | Related idea | Computing `cnt` is the same as the step of COUNTING-SORT that counts occurrences of each value; this solution stops there instead of expanding the counts into a sorted sequence, because the count vector itself can serve as the key |

## Chapter 9 Medians and Order Statistics

### §9.2 Selection in expected linear time

| Problem | Relation | How It Appears Here |
|---|---|---|
| [462. Minimum Moves to Equal Array Elements II (two-way)](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/lomuto-quickselect/summary.md) | Direct | `QuickSelect` is RANDOMIZED-SELECT; the expected $`O(n)`$ analysis assumes distinct elements |
| [462. Minimum Moves to Equal Array Elements II (three-way)](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/three-way-quickselect/summary.md) | Direct | `QuickSelect` is RANDOMIZED-SELECT: it recurses only into the side containing the $`k`$-th smallest, in expected $`O(n)`$ |

### §9.3 Selection in worst-case linear time

| Problem | Relation | How It Appears Here |
|---|---|---|
| [462. Minimum Moves to Equal Array Elements II (three-way)](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/three-way-quickselect/summary.md) | Related idea | This implementation is $`O(n^2)`$ in the worst case; choosing the pivot by median of medians would guarantee worst-case $`O(n)`$ |

## Chapter 10 Elementary Data Structures

### §10.1 Stacks and queues

| Problem | Relation | How It Appears Here |
|---|---|---|
| [94. Binary Tree Inorder Traversal (Explicit Stack)](../data-structures/binary-tree/0094-binary-tree-inorder-traversal/iterative/summary.md) | Direct | `stack` uses the end of a slice as its top, with `append` as PUSH and truncation as POP; last-in first-out makes the innermost ancestor resume first |

### §10.4 Representing rooted trees

| Problem | Relation | How It Appears Here |
|---|---|---|
| [104. Maximum Depth of Binary Tree](../data-structures/binary-tree/0104-maximum-depth-of-binary-tree/summary.md) | Direct | The `Left` and `Right` pointers of `TreeNode` are the left and right attributes of CLRS nodes, and an empty child `nil` corresponds to NIL |
| [94. Binary Tree Inorder Traversal (Recursive)](../data-structures/binary-tree/0094-binary-tree-inorder-traversal/recursive/summary.md) | Direct | `TreeNode`'s `Left` and `Right` pointers are the left and right attributes of CLRS nodes, and a `nil` child is NIL |
| [94. Binary Tree Inorder Traversal (Explicit Stack)](../data-structures/binary-tree/0094-binary-tree-inorder-traversal/iterative/summary.md) | Direct | Exercise 10.4-3 asks for an $`O(n)`$-time nonrecursive procedure that prints every key of a binary tree using a stack; this solution is one such procedure, in inorder |

## Chapter 11 Hash Tables

### §11.1 Direct-address tables

| Problem | Relation | How It Appears Here |
|---|---|---|
| [49. Group Anagrams (count key)](../data-structures/hash-table/0049-group-anagrams/count-key/summary.md) | Direct | `cnt` is a direct-address table keyed by the letter numbers $`0 \ldots 25`$, with worst-case $`O(1)`$ per count |

### §11.2 Hash tables

| Problem | Relation | How It Appears Here |
|---|---|---|
| [49. Group Anagrams (sorted key)](../data-structures/hash-table/0049-group-anagrams/sorted-key/summary.md) | Direct | `groups` is keyed by the sorted string; each string costs one lookup and one append, expected $`O(\ell_i)`$ |
| [49. Group Anagrams (count key)](../data-structures/hash-table/0049-group-anagrams/count-key/summary.md) | Direct | `groups` is keyed by `[26]int`; each string costs one lookup and one append, expected $`O(\lvert\Sigma\rvert)`$ |
| [128. Longest Consecutive Sequence](../data-structures/hash-table/0128-longest-consecutive-sequence/summary.md) | Direct | Insertion and membership queries on `numSet` take expected $`O(1)`$, which the overall expected $`O(n)`$ relies on |

## Chapter 12 Binary Search Trees

### §12.1 What is a binary search tree?

| Problem | Relation | How It Appears Here |
|---|---|---|
| [104. Maximum Depth of Binary Tree](../data-structures/binary-tree/0104-maximum-depth-of-binary-tree/summary.md) | Related idea | INORDER-TREE-WALK makes one call per node, and Theorem 12.1 proves it takes $`\Theta(n)`$ time; 4.5 counts calls the same way ("one non-empty call per node"), except that this problem combines subtree results in postorder |
| [94. Binary Tree Inorder Traversal (Recursive)](../data-structures/binary-tree/0094-binary-tree-inorder-traversal/recursive/summary.md) | Direct | This solution is INORDER-TREE-WALK, except that CLRS prints on visit while this solution returns and concatenates subsequences, which causes the copying cost in §4.5; the BST case matches the discussion around Theorem 12.1 that an inorder walk prints keys in sorted order |
| [94. Binary Tree Inorder Traversal (Explicit Stack)](../data-structures/binary-tree/0094-binary-tree-inorder-traversal/iterative/summary.md) | Direct | Exercise 12.1-3 asks for a nonrecursive inorder tree walk, with a hint that an easy solution uses a stack |

## Chapter 15 Dynamic Programming

### §15.3 Elements of dynamic programming

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Related idea | Each entry of a row is reused by two positions in the next row; filling rows in increasing order computes values before use |

## Chapter 17 Amortized Analysis

### §17.1 Aggregate analysis

| Problem | Relation | How It Appears Here |
|---|---|---|
| [128. Longest Consecutive Sequence](../data-structures/hash-table/0128-longest-consecutive-sequence/summary.md) | Direct | The inner `while` runs $`m - (\text{number of heads}) \le m`$ times in total, counted over the whole algorithm rather than per loop |
| [94. Binary Tree Inorder Traversal (Explicit Stack)](../data-structures/binary-tree/0094-binary-tree-inorder-traversal/iterative/summary.md) | Direct | Each node is pushed and popped once, for $`2n`$ iterations in total; the same argument as the MULTIPOP example's "pops never exceed pushes" |

### §17.3 The potential method

| Problem | Relation | How It Appears Here |
|---|---|---|
| [453. Minimum Moves to Equal Array Elements (batched simulation)](../Algorithms/simulation/0453-minimum-moves-to-equal-array-elements/batched-simulation/summary.md) | Related idea | The state is measured by $`\Phi = \sum x_i - n \cdot \min x`$, and bounding $`\Delta\Phi \ge -1`$ per move yields the lower bound; here it proves a lower bound rather than performing amortized analysis |

### §17.4 Dynamic tables

| Problem | Relation | How It Appears Here |
|---|---|---|
| [94. Binary Tree Inorder Traversal (Recursive)](../data-structures/binary-tree/0094-binary-tree-inorder-traversal/recursive/summary.md) | Related idea | In §4.5, the extra copying from `append` growth adds only a constant factor |

## Appendix B Sets, Etc.

### B.5 Trees

| Problem | Relation | How It Appears Here |
|---|---|---|
| [104. Maximum Depth of Binary Tree](../data-structures/binary-tree/0104-maximum-depth-of-binary-tree/summary.md) | Direct | Challenge 2: CLRS defines depth and height by edges, while this problem counts nodes, so on a non-empty tree the answer is the CLRS height plus 1 |
| [94. Binary Tree Inorder Traversal (Recursive)](../data-structures/binary-tree/0094-binary-tree-inorder-traversal/recursive/summary.md) | Direct | 4.1 writes the input as either the empty tree or $`(r, L, R)`$, and both the definition of $`\mathrm{in}`$ and the proof follow this recursive structure |

### B.2 Relations

| Problem | Relation | How It Appears Here |
|---|---|---|
| [49. Group Anagrams (sorted key)](../data-structures/hash-table/0049-group-anagrams/sorted-key/summary.md) | Direct | Lemma 0 proves "being anagrams" is an equivalence relation; by Theorem B.1 its classes form a partition, and the output groups are exactly those classes |
| [49. Group Anagrams (count key)](../data-structures/hash-table/0049-group-anagrams/count-key/summary.md) | Direct | Cites Lemma 0 of the sorted-key solution: "being anagrams" is an equivalence relation, so by Theorem B.1 its classes form a partition, and the output groups are exactly those classes |

## Appendix C Counting and Probability

### C.1 Counting

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Related idea | Proves $`P(i, j) = \binom{i}{j}`$ and uses it for the $`\binom{29}{14}`$ overflow bound |
| [3963. Create Grid With Exactly One Path](../Algorithms/constructive-algorithms/3963-create-grid-with-exactly-one-path/summary.md) | Related idea | The test uses $`\binom{4}{2} = 6`$ paths for the all-free $`3 \times 3`$ grid to check the validator `countPaths` |
