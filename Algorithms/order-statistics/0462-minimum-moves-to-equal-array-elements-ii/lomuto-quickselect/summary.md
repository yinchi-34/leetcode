# 462. Minimum Moves to Equal Array Elements II — Two-Way (Lomuto) Quickselect

English | [简体中文](./summary.zh-CN.md)

Terminology follows the [shared glossary](../../../../docs/GLOSSARY.md), and math notation follows the repository's [GitHub math writing conventions](../../../../docs/GITHUB_MATH.md).

This problem has three solutions in three directories that share one set of test cases:

- [**Sorting**](../sorting/summary.md): sort and take the median. The problem statement and the proof that the median is optimal (Lemmas 1 and 2) are there.
- [**Three-way quickselect**](../three-way-quickselect/summary.md): randomized quickselect that partitions into less, equal, and greater blocks. The induction proof for quickselect and the expected-time analysis are there.
- **Two-way (Lomuto) quickselect** (this page): CLRS's PARTITION, the version written first.

This page covers only how it differs from the three-way version: the partition, and why it degrades with many duplicates.

## 1. Problem Information

- **Problem**: [462. Minimum Moves to Equal Array Elements II](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/)
- **Official Difficulty**: Medium
- **Primary Category**: Order Statistics
- **Secondary Category**: Randomized Algorithms
- **Tags**: `Order Statistics`, `Randomized Algorithms`, `Median`, `Quickselect`, `Lomuto Partition`
- **Code**: [solution.go](./solution.go) · [solution.py](./solution.py)
- **Tests**: [solution_test.go](../solution_test.go) · [solution_test.py](../solution_test.py) (in the parent directory, shared by all three solutions)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Randomized Quickselect with Two-Way (Lomuto) Partitioning**

- Evaluation Scope: $`1 \le n \le 10^5`$, $`-10^9 \le a_i \le 10^9`$, the answer fits in a 32-bit integer; expected time $`O(n)`$ when elements are distinct, degrading to $`\Theta(n^2)`$ with many duplicates; $`O(1)`$ extra space besides the recursion stack; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2 | One must turn "minimum number of moves" into "minimum sum of absolute deviations", and then into finding the $`\lfloor n/2 \rfloor`$-th smallest element |
| Knowledge Prerequisites | 3 | Requires CLRS's PARTITION, RANDOMIZED-PARTITION, and RANDOMIZED-SELECT |
| Correctness Reasoning Difficulty | 4 | Correctness needs the same median optimality, partition invariant, and induction on range length; the time analysis must also point out that expected linear time assumes distinct elements |
| Implementation Difficulty | 2 | Follows the CLRS code; the details are starting `i` at `l-1` and finally swapping the pivot into `i+1` |

Ratings follow [RATING.md](../../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

The statement, examples, and constraints are in the [sorting solution](../sorting/summary.md#2-problem-understanding). As in the three-way version, the core subproblem is finding $`b_m`$, the element at index $`m = \lfloor n/2 \rfloor`$ in sorted order.

## 3. Initial Approach

### First Idea: Randomized Quickselect, Following CLRS

This is the Go version written first: three functions, `partition`, `randomPartition`, and `quickSelect`, matching CLRS's PARTITION, RANDOMIZED-PARTITION, and RANDOMIZED-SELECT. The code in the repository is the original with only `rand.Seed` removed: it has been deprecated since Go 1.20, and the repository uses Go 1.22, where the global random source is seeded automatically.

### Difficulties Encountered

At the time, the name Lomuto partition was unfamiliar, and only the $`n^2`$ worst case came to mind. Measured with $`n = 10^5`$: a random array took about 360,000 comparisons and 1 ms; **an all-equal array took about 3.7 billion comparisons and 7.7 seconds**.

### Thought Transformation

From:

> A random pivot avoids the worst case.

To:

> Two-way partitioning puts every element equal to the pivot on the left; with many repeated values, randomization does not help.

The corrected version is [three-way quickselect](../three-way-quickselect/summary.md). This solution stays in the repository for comparison.

## 4. Algorithm Specification

**Algorithm Name:** Randomized Quickselect with Lomuto Partitioning

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
PARTITION(A, l, r)                   // CLRS §7.1, pivot x = A[r]
    x ← A[r]
    i ← l - 1
    for j ← l to r - 1
        if A[j] ≤ x
            i ← i + 1
            swap A[i], A[j]
    swap A[i+1], A[r]
    return i + 1                     // A[l..q-1] ≤ x, A[q] = x, A[q+1..r] > x

QUICK-SELECT(A, l, r, k)             // requires l ≤ k ≤ r
    p ← uniformly random index in [l, r]
    swap A[p], A[r]
    q ← PARTITION(A, l, r)
    if q = k
        return A[q]
    if q < k
        return QUICK-SELECT(A, q + 1, r, k)
    return QUICK-SELECT(A, l, q - 1, k)

MIN-MOVES-2(A)
    x ← QUICK-SELECT(A, 0, n - 1, ⌊n/2⌋)
    res ← Σ |A[i] - x|
    return res
```

**State Definitions**

- `i`: last position of the "$`\le x`$" region.
- `j`: next position to examine.
- `q`: position of the pivot after partitioning.

### 4.4 Correctness

**Proof Method:** a loop invariant for PARTITION; quickselect reuses the strong induction of the three-way version.

#### I. Key Properties Required for Correctness

1. **Lemmas 1 and 2** (see the [sorting solution](../sorting/summary.md#44-correctness)): the answer is $`f(b_m) = \sum_i \lvert a_i - b_m \rvert`$.
2. **Lemma 3′ (Lomuto partition)**: when `PARTITION` returns $`q`$, $`A[l..r]`$ is a permutation of the original range with $`A[l..q-1] \le x`$, $`A[q] = x`$, and $`A[q+1..r] > x`$.
3. **Lemma 4 (Quickselect)**: `QUICK-SELECT(A, l, r, k)` returns the element at index $`k`$ in sorted order.

#### II. Correctness Proof

**Lemma 3′: loop invariant for PARTITION** (CLRS §7.1)

Each time the `for` condition is tested:

- the elements of $`A[l..i]`$ are all $`\le x`$;
- the elements of $`A[i+1..j-1]`$ are all $`> x`$;
- $`A[j..r-1]`$ is unexamined;
- $`A[r] = x`$.

**Initialization**: $`i = l - 1`$ and $`j = l`$, so the first two regions are empty.

**Maintenance**:

- If $`A[j] > x`$, incrementing $`j`$ extends the second region.
- If $`A[j] \le x`$, incrementing $`i`$ and swapping $`A[i]`$ with $`A[j]`$ adds an element $`\le x`$ to the first region; the element moved to $`j`$ is the first element of the old second region (or the swap is a self-swap), so the second region shifts right by one.

**Termination**: at $`j = r`$ nothing is unexamined. The final swap of $`A[i+1]`$ and $`A[r]`$ places the pivot between the two regions, and $`q = i + 1`$ is returned. Only swaps are performed, so the result is a permutation. $`\square`$

**Lemma 4**: the proof is that of [Lemma 4 in the three-way version](../three-way-quickselect/summary.md#44-correctness), with the three blocks replaced by $`A[l..q-1]`$, $`A[q]`$, and $`A[q+1..r]`$. By Lemma 3′ they occupy sorted indices $`l..q-1`$, $`q`$, and $`q+1..r`$; $`A[l..q-1]`$ may contain elements equal to $`x`$, which does not change their sorted positions. The pivot itself is excluded, so each subrange is strictly shorter. $`\square`$

#### III. Conclusion

By Lemma 4, `QUICK-SELECT` returns $`b_m`$; rearranging does not change the multiset (Lemma 3′), so the sum is $`f(b_m)`$, which is the answer by Lemmas 1 and 2. The result is always correct; only the running time is at issue.

### 4.5 Complexity Analysis

Let the input size be $`n`$.

**Time Complexity: expected O(n) for distinct elements; Θ(n²) with many duplicates**

**Distinct elements**: the phase argument of the three-way version needs "the next range has length at most $`\max(\rho, s - 1 - \rho)`$", where $`\rho`$ is the pivot's rank. With distinct elements, Lomuto partitioning satisfies this, so the expectation is still $`O(n)`$, in agreement with CLRS §9.2.

**With duplicates**: Lomuto partitioning puts every element equal to the pivot on the left, and the requirement above fails. On an all-equal array, whichever element is chosen, `A[j] ≤ x` always holds, $`q`$ always equals $`r`$, and each level shrinks by only 1. With $`k = \lfloor n/2 \rfloor`$ the range must shrink from $`n`$ to about $`n/2`$, for about

```math
\sum_{s = n/2}^{n} s \approx \frac{3n^2}{8}
```

comparisons. For $`n = 10^5`$ that is about $`3.75 \times 10^9`$; the measurement was about $`3.7 \times 10^9`$ comparisons and 7.7 seconds, matching the estimate.

The final sum is one $`\Theta(n)`$ pass.

**Space Complexity: O(1) besides the recursion stack**

Rearranges in place. The recursion depth is $`O(\log n)`$ in expectation and $`O(n)`$ in the worst case; on an all-equal array it is $`\Theta(n)`$.

### 4.6 Implementation Notes: Go and Python

- **Recursion to a loop**: in this solution the rewrite is required. With many duplicates, each partition removes only one element and the recursion is nearly $`n`$ levels deep. Go's stack grows on demand and copes; CPython raises `RecursionError` at 1000 levels by default. The Python version is therefore a `while` loop. See [recursion depth](./../../../../docs/LANGUAGE_NOTES.md#recursion-depth).
- **Cost of the degenerate case**: the same $`\Theta(n^2)`$ degradation is far slower in Python, so neither language's tests run this solution on a large all-equal array.
- **Random pivot**: Go's `rand.Intn(r-l+1) + l` is Python's `random.randint(l, r)`, which includes both ends.

## 5. Learning Notes

### Challenge 1: Why a random pivot does not help here

**Understanding Process**

Randomization guarantees that the pivot's rank in sorted order is uniform. But `A[j] ≤ x` sends every element equal to the pivot to the left, effectively treating the pivot as "the largest of the equal values". On an all-equal array, whichever element is chosen, its rank is last, so every partition is the worst one. See [Challenge 1 in the three-way version](../three-way-quickselect/summary.md#challenge-1-why-does-a-random-pivot-not-protect-against-duplicates) for more.

**Transferable Insight:** proofs of expected complexity often assume distinct elements; with many repeated values, check that assumption again.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Lomuto Partition**: CLRS's PARTITION scans with a single pointer and swaps elements $`\le x`$ to the front. It is simple and a common starting point for quicksort and quickselect. Limitation: with many repeated values the split becomes badly unbalanced; use three-way or Hoare partitioning instead.
2. **Measuring the Worst Case**: here a measurement on an all-equal array exposed the hidden assumption of distinct elements in the theoretical analysis. Applies to checking any algorithm that claims an "expected" complexity.

### CLRS Connections

Chapter numbers follow the third edition of CLRS.

| Concept | CLRS Section | Relation | Where it appears in this problem |
|---|---|---|---|
| Partitioning and its loop invariant | §7.1 Description of quicksort | Direct | `partition` is PARTITION, proved correct by its four-region loop invariant |
| Random pivot | §7.3 A randomized version of quicksort | Direct | `randomPartition` is RANDOMIZED-PARTITION |
| Partitioning with equal elements | Problem 7-2 Quicksort with equal element values | Direct | An all-equal array makes this solution degrade to $`\Theta(n^2)`$, exactly the case that problem discusses |
| Selection in expected linear time | §9.2 Selection in expected linear time | Direct | `QuickSelect` is RANDOMIZED-SELECT; the expected $`O(n)`$ analysis assumes distinct elements |

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

- [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) — asks directly for the $`k`$-th largest element; with many repeated values in the test data, two-way partitioning degrades there too.

---

**One-Sentence Summary**

Randomized quickselect written after CLRS always returns the right answer, but `A[j] ≤ x` crowds all repeated values onto one side, and on an all-equal array it degrades to $`\Theta(n^2)`$.
