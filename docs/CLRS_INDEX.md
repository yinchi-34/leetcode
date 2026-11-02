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
| [128. Longest Consecutive Sequence](../data-structures/hash-table/0128-longest-consecutive-sequence/summary.md) | Direct | Lemma 1's inner invariant and the outer invariant "`res` is the longest run among processed heads" are each proved by initialization, maintenance, and termination |
| [453. Minimum Moves to Equal Array Elements](../Algorithms/simulation/0453-minimum-moves-to-equal-array-elements/summary.md) | Direct | Solution 1 is proved by the invariants "equal prefix, untouched suffix", through initialization, maintenance, and termination |

## Chapter 3 Growth of Functions

### §3.2 Standard notations and common functions

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Related idea | The Fibonacci recurrence is the same kind as this problem's, and the shallow-diagonal sums of Pascal's triangle are Fibonacci numbers |

## Chapter 11 Hash Tables

### §11.2 Hash tables

| Problem | Relation | How It Appears Here |
|---|---|---|
| [128. Longest Consecutive Sequence](../data-structures/hash-table/0128-longest-consecutive-sequence/summary.md) | Direct | Insertion and membership queries on `numSet` take expected $`O(1)`$, which the overall expected $`O(n)`$ relies on |

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

### §17.3 The potential method

| Problem | Relation | How It Appears Here |
|---|---|---|
| [453. Minimum Moves to Equal Array Elements](../Algorithms/simulation/0453-minimum-moves-to-equal-array-elements/summary.md) | Related idea | The state is measured by $`\Phi = \sum x_i - n \cdot \min x`$, and bounding $`\Delta\Phi \ge -1`$ per move yields the lower bound; here it proves a lower bound rather than performing amortized analysis |

## Appendix C Counting and Probability

### C.1 Counting

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Related idea | Proves $`P(i, j) = \binom{i}{j}`$ and uses it for the $`\binom{29}{14}`$ overflow bound |
| [3963. Create Grid With Exactly One Path](../Algorithms/constructive-algorithms/3963-create-grid-with-exactly-one-path/summary.md) | Related idea | The test uses $`\binom{4}{2} = 6`$ paths for the all-free $`3 \times 3`$ grid to check the validator `countPaths` |
