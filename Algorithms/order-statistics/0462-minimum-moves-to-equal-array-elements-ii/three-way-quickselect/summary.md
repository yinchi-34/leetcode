# 462. Minimum Moves to Equal Array Elements II — Three-Way Quickselect

English | [简体中文](./summary.zh-CN.md)

Terminology follows the [shared glossary](../../../../docs/GLOSSARY.md), and math notation follows the repository's [GitHub math writing conventions](../../../../docs/GITHUB_MATH.md).

This problem has three solutions in three directories that share one set of test cases:

- [**Sorting**](../sorting/summary.md): sort and take the median. The problem statement and the proof that the median is optimal (Lemmas 1 and 2) are there.
- **Three-way quickselect** (this page): randomized quickselect that partitions into less, equal, and greater blocks.
- [**Two-way (Lomuto) quickselect**](../lomuto-quickselect/summary.md): CLRS's PARTITION, which degrades with many duplicates.

This page covers only how to find the median without sorting.

## 1. Problem Information

- **Problem**: [462. Minimum Moves to Equal Array Elements II](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/)
- **Official Difficulty**: Medium
- **Primary Category**: Order Statistics
- **Secondary Category**: Randomized Algorithms
- **Tags**: `Order Statistics`, `Randomized Algorithms`, `Median`, `Quickselect`, `Three-Way Partitioning`
- **Code**: [solution.go](./solution.go) · [solution.py](./solution.py)
- **Tests**: [solution_test.go](../solution_test.go) · [solution_test.py](../solution_test.py) (in the parent directory, shared by all three solutions)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Randomized Quickselect with Three-Way Partitioning**

- Evaluation Scope: $`1 \le n \le 10^5`$, $`-10^9 \le a_i \le 10^9`$, the answer fits in a 32-bit integer; expected time $`O(n)`$, worst case $`O(n^2)`$; $`O(1)`$ extra space besides the recursion stack; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2 | One must turn "minimum number of moves" into "minimum sum of absolute deviations", and then into finding the $`\lfloor n/2 \rfloor`$-th smallest element |
| Knowledge Prerequisites | 3 | Requires a working command of randomized quickselect: partitioning, random pivots, recursing into one side only, and knowing that duplicates call for three-way partitioning |
| Correctness Reasoning Difficulty | 4 | Needs several lemmas: optimality of the median, a loop invariant for three-way partitioning, strong induction on subarray length for quickselect, and a probabilistic argument for expected linear time |
| Implementation Difficulty | 3 | Three cooperating pointers; `i` must not advance when `a[i] > x`; the recursive ranges must use `lt-1` and `gt+1` to drop the whole equal block |

Ratings follow [RATING.md](../../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

The statement, examples, and constraints are in the [sorting solution](../sorting/summary.md#2-problem-understanding). By Lemmas 1 and 2 there, the answer is $`\sum_i \lvert a_i - b_m \rvert`$, where $`b_m`$ is the element at index $`m = \lfloor n/2 \rfloor`$ in sorted order.

So the core subproblem is: **given an array and an index $`k`$, find the element at index $`k`$ in sorted order**, the $`(k+1)`$-th order statistic.

## 3. Initial Approach

### First Idea: Randomized Quickselect (Two-Way Partitioning)

The first Go version used CLRS's Lomuto partition with a random pivot; see [two-way quickselect](../lomuto-quickselect/summary.md).

### Difficulties Encountered

Three questions came up while discussing quickselect:

1. **Why recurse into one side only**: answered correctly; only the position of the $`k`$-th smallest is needed.
2. **Time complexity**: answered $`n^2`$. That is the worst case; the expected case is $`O(n)`$.
3. **What is Lomuto**: the name was unfamiliar, though it is exactly the `partition` already written.

Running the original code with $`n = 10^5`$: a random array took about 360,000 comparisons and 1 ms; **an all-equal array took about 3.7 billion comparisons and 7.7 seconds**.

### Thought Transformation

From:

> A random pivot avoids the worst case.

To:

> Randomization only protects against bad luck. Duplicates turn every pivot into a worst-case pivot, so elements equal to the pivot must be split off into their own block.

The Lomuto partition sends every element equal to the pivot to the left, so the pivot lands at the right end of that block; on an all-equal array each round removes only one element. Choosing a different one at random changes nothing, because they all have the same value. With three-way partitioning (less, equal, greater), one round removes the whole equal block, and the search returns immediately when the $`k`$-th smallest falls inside it. The same all-equal array dropped to about 0.26 ms.

## 4. Algorithm Specification

**Algorithm Name:** Randomized Quickselect with Three-Way Partitioning

### 4.1 Input

$`A = \langle a_0, a_1, \ldots, a_{n-1} \rangle`$

satisfying:

- $`1 \le n \le 10^5`$
- $`a_i \in \mathbb{Z}`$, $`-10^9 \le a_i \le 10^9`$

### 4.2 Output

```math
\mathrm{res} = \sum_{i=0}^{n-1} \lvert a_i - b_m \rvert, \qquad m = \lfloor n/2 \rfloor .
```

### 4.3 Pseudocode

Indices are 0-based.

```text
PARTITION-3(A, l, r)                 // pivot x = A[r]
    x ← A[r]
    lt ← l;  i ← l;  gt ← r
    while i ≤ gt
        if A[i] < x
            swap A[lt], A[i];  lt ← lt + 1;  i ← i + 1
        else if A[i] > x
            swap A[i], A[gt];  gt ← gt - 1    // i stays: the new A[i] is unexamined
        else
            i ← i + 1
    return (lt, gt)                  // A[l..lt-1] < x, A[lt..gt] = x, A[gt+1..r] > x

QUICK-SELECT(A, l, r, k)             // requires l ≤ k ≤ r
    p ← uniformly random index in [l, r]
    swap A[p], A[r]
    (lt, gt) ← PARTITION-3(A, l, r)
    if k < lt
        return QUICK-SELECT(A, l, lt - 1, k)
    if k > gt
        return QUICK-SELECT(A, gt + 1, r, k)
    return A[k]

MIN-MOVES-2(A)
    x ← QUICK-SELECT(A, 0, n - 1, ⌊n/2⌋)
    res ← 0
    for i ← 0 to n - 1
        res ← res + |A[i] - x|
    return res
```

**State Definitions**

- `lt`: left end of the equal block; $`A[l..lt-1]`$ are all less than the pivot.
- `i`: next position to examine; $`A[lt..i-1]`$ all equal the pivot.
- `gt`: right end of the unexamined block; $`A[gt+1..r]`$ are all greater than the pivot.

### 4.4 Correctness

**Proof Method:** a loop invariant for three-way partitioning; strong induction on subarray length for quickselect.

#### I. Key Properties Required for Correctness

1. **Lemmas 1 and 2** (see the [sorting solution](../sorting/summary.md#44-correctness)): the answer is $`f(b_m) = \sum_i \lvert a_i - b_m \rvert`$.
2. **Lemma 3 (Three-way partition)**: when `PARTITION-3` returns, $`A[l..r]`$ is a permutation of the original range and satisfies the three-block property stated in the return comment.
3. **Lemma 4 (Quickselect)**: `QUICK-SELECT(A, l, r, k)` returns the element at index $`k`$ in sorted order.

#### II. Correctness Proof

**Lemma 3: loop invariant for three-way partitioning**

Each time the `while` condition is tested:

**Invariant I: four regions**

- the elements of $`A[l..lt-1]`$ are all less than $`x`$;
- the elements of $`A[lt..i-1]`$ all equal $`x`$;
- $`A[i..gt]`$ has not been examined;
- the elements of $`A[gt+1..r]`$ are all greater than $`x`$.

**Invariant II: permutation**

$`A[l..r]`$ is always a permutation of the original range.

**Initialization**: $`lt = i = l`$ and $`gt = r`$. The first, second, and last regions are empty and the whole range is unexamined; no swaps have happened. The invariants hold.

**Maintenance**: suppose the invariants hold before a test with $`i \le gt`$.

- **$`A[i] < x`$**: swap $`A[lt]`$ and $`A[i]`$. If $`lt < i`$, the element moved to position $`i`$ is the first element of the equal block and equals $`x`$; if $`lt = i`$, the swap is a self-swap. Either way, after incrementing $`lt`$ and $`i`$, the less-than block has gained an element less than $`x`$ and the equal block has shifted right by one.
- **$`A[i] > x`$**: swap $`A[i]`$ and $`A[gt]`$ and decrement $`gt`$. The old $`A[i]`$ joins the greater-than block; the element now at position $`i`$ came from the unexamined block and is still unexamined, so $`i`$ stays.
- **$`A[i] = x`$**: the equal block ends at $`i`$, so incrementing $`i`$ adds it to the block.

Each case only swaps, so Invariant II is preserved, and Invariant I is preserved as well.

**Termination**: each iteration increments $`i`$ or decrements $`gt`$, so the unexamined length $`gt - i + 1`$ strictly decreases and the loop ends with $`i = gt + 1`$; the unexamined block is empty, leaving three blocks. The pivot itself equals $`x`$, so the equal block is nonempty: $`lt \le gt`$. $`\square`$

**Lemma 4: quickselect**

Call a range $`A[l..r]`$ a valid selection range if its elements occupy exactly the indices $`l..r`$ when the whole array is sorted. The top-level call with $`l = 0`$ and $`r = n - 1`$ is trivially valid.

Proceed by strong induction on the length $`r - l + 1`$, with $`l \le k \le r`$. By Lemma 3, after partitioning, the less-than, equal, and greater-than blocks occupy sorted indices $`l..lt-1`$, $`lt..gt`$, and $`gt+1..r`$, respectively.

- **$`k < lt`$**: the element at sorted index $`k`$ lies in the less-than block; the subrange $`[l, lt-1]`$ is still valid and $`l \le k \le lt - 1`$. Since the equal block is nonempty, the subrange is strictly shorter, and the induction hypothesis gives the correct result.
- **$`k > gt`$**: symmetric.
- **$`lt \le k \le gt`$**: the element at sorted index $`k`$ equals $`x`$, and so does $`A[k]`$, which is returned.

For length 1, $`lt = gt = l = k`$, which falls in the third case. Therefore `QUICK-SELECT(A, 0, n-1, m)` returns $`b_m`$. $`\square`$

#### III. Conclusion

By Lemma 4, `QUICK-SELECT` returns $`b_m`$. The sum is then taken over the rearranged array; rearranging does not change the multiset (Lemma 3), so the sum equals $`f(b_m)`$, which is the answer by Lemmas 1 and 2.

### 4.5 Complexity Analysis

Let the input size be $`n`$. `PARTITION-3` on a range of length $`s`$ makes $`\Theta(s)`$ comparisons.

**Time Complexity: expected O(n), worst case O(n²)**

**Expected $`O(n)`$ (phase argument)**: order the $`s`$ elements of the current range totally, by value and then by index. The random pivot's rank $`\rho`$ in this order is uniform on $`0..s-1`$. Every element of the less-than block ranks below $`\rho`$ and every element of the greater-than block ranks above it, so the next range has length at most $`\max(\rho, s - 1 - \rho)`$.

With probability at least $`1/2`$, $`\rho`$ falls in $`[\lfloor s/4 \rfloor, \lceil 3s/4 \rceil)`$, in which case the next range has length at most $`3s/4`$.

Say the algorithm is in phase $`j`$ while the current length lies in $`(n(3/4)^{j+1}, n(3/4)^j]`$. Each partition ends the current phase with probability at least $`1/2`$, so each phase has at most 2 partitions in expectation, each costing at most $`c \cdot n (3/4)^j`$. Hence

```math
\mathbb{E}[T(n)] \le \sum_{j \ge 0} 2c\,n\left(\tfrac{3}{4}\right)^{j} = 8c\,n = O(n) .
```

Three-way partitioning can only make the next range shorter, so duplicates do not affect this bound. The final sum is one $`\Theta(n)`$ pass.

**Worst case $`O(n^2)`$**: if every pivot is the unique minimum or maximum of its range, the length drops by only 1 each time, for about $`n^2/2`$ comparisons. Randomization makes this extremely unlikely but does not rule it out. The SELECT algorithm of CLRS §9.3 (median of medians) guarantees worst-case $`O(n)`$; this implementation does not use it.

**Space Complexity: O(1) besides the recursion stack**

Rearranges in place. The recursion depth is $`O(\log n)`$ in expectation and $`O(n)`$ in the worst case.

### 4.6 Implementation Notes: Go and Python

- **Recursion to a loop**: Go's `QuickSelect` is tail-recursive with expected depth $`O(\log n)`$. The Python version is a `while` loop, because CPython does not eliminate tail calls and stops at 1000 levels by default. See [recursion depth](./../../../../docs/LANGUAGE_NOTES.md#recursion-depth).
- **In-place reordering**: both versions reorder the caller's array, so the tests pass a copy (`append([]int(nil), ...)` in Go, `nums.copy()` in Python).
- **Random pivot**: Go's `rand.Intn(r-l+1) + l` is Python's `random.randint(l, r)`, which includes both ends.

## 5. Learning Notes

### Challenge 1: Why does a random pivot not protect against duplicates?

**Understanding Process**

The expected-time analysis needs the pivot's rank to be uniform. Lomuto partitioning sends every element equal to the pivot to the left, effectively treating each copy of a value as "the largest one". On an all-equal array, whichever element is picked, its rank is last. Three-way partitioning removes the whole equal block at once and returns immediately when the $`k`$-th smallest lies inside it.

**Transferable Insight:** randomization handles worst cases caused by input order or bad luck, not structural ones caused by equal values; when analyzing a randomized algorithm, check the assumptions behind the expectation, such as distinct elements.

### Challenge 2: When does `i` advance in three-way partitioning?

**Understanding Process**

Ask which region $`A[i]`$ belongs to. The equal block sits immediately left of $`i`$, so for an equal element $`i`$ simply advances. For a smaller element, swapping with the first element of the equal block brings back a value known to equal $`x`$ (or performs a self-swap), so $`i`$ advances. For a larger element, swapping with $`A[gt]`$ brings back an unexamined value, so $`i`$ must not advance.

**Transferable Insight:** when writing an in-place partition, first draw the regions and their boundary pointers; in each branch ask one question: "has the element now at the current position already been examined?"

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Quickselect**: here it finds the $`\lfloor n/2 \rfloor`$-th smallest element in expected $`O(n)`$. Applies whenever a single order statistic is needed rather than a full sort, such as the $`k`$-th largest or the $`k`$ smallest. Limitations: $`O(n^2)`$ worst case, and it reorders the array in place.
2. **Three-Way Partitioning**: here it prevents degradation on duplicates. Applies to quicksort and quickselect on inputs with many repeated values, and on its own to rearrange elements in place by comparison with a value into three classes.

### CLRS Connections

Chapter numbers follow the third edition of CLRS.

| Concept | CLRS Section | Relation | Where it appears in this problem |
|---|---|---|---|
| Partitioning and its loop invariant | §7.1 Description of quicksort | Related idea | Three-way partitioning reuses PARTITION's region-based invariant proof, extending the regions from two classes to three |
| Random pivot | §7.3 A randomized version of quicksort | Direct | `randomPartition3` matches RANDOMIZED-PARTITION: swap a random element into `a[r]`, then partition |
| Partitioning with equal elements | Problem 7-2 Quicksort with equal element values | Direct | Three-way partitioning puts elements equal to the pivot in their own block, avoiding $`\Theta(n^2)`$ on all-equal arrays |
| Selection in expected linear time | §9.2 Selection in expected linear time | Direct | `QuickSelect` is RANDOMIZED-SELECT: it recurses only into the side containing the $`k`$-th smallest, in expected $`O(n)`$ |
| Selection in worst-case linear time | §9.3 Selection in worst-case linear time | Related idea | This implementation is $`O(n^2)`$ in the worst case; choosing the pivot by median of medians would guarantee worst-case $`O(n)`$ |

## 7. Review Log

| Date | Independent Completion | Understanding Gained | To Improve |
|---|---|---|---|
| To be recorded | To be recorded | To be recorded | To be recorded |

### Next Review Checklist

- [ ] Can I implement it independently without looking at the code?
- [ ] Can I accurately describe the algorithm's input and output?
- [ ] Can I write the pseudocode independently?
- [ ] Can I explain why the algorithm is correct?
- [ ] Can I derive the time and space complexity?
- [ ] Can I recognize the same algorithmic idea in other problems?

## 8. Related Problems

- [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) — asks directly for the $`k`$-th largest element, a standalone exercise for the quickselect part of this page.

---

**One-Sentence Summary**

Finding the median needs no sort: randomized quickselect with three-way partitioning recurses into one side only, runs in expected $`O(n)`$, and removes a whole run of duplicates in one round.
