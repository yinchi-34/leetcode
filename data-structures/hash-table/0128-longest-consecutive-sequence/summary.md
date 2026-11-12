# 128. Longest Consecutive Sequence

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md). The hash table itself (its definition, Go `map` behavior, and common usage patterns) is covered in the [hash table overview](../README.md); this note covers only how this problem uses a hash table.

## 1. Problem Information

- **Problem**: [128. Longest Consecutive Sequence](https://leetcode.cn/problems/longest-consecutive-sequence/)
- **Official Difficulty**: Medium
- **Primary Category**: Hash Table
- **Secondary Category**: Amortized Analysis
- **Tags**: `Hash Table`, `Hash Set`, `Amortized Analysis`, `Run-Head Enumeration`
- **Code**: [solution.go](./solution.go) · [solution.py](./solution.py)
- **Tests**: [solution_test.go](./solution_test.go) · [solution_test.py](./solution_test.py)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Hash Set with Run-Head Extension**

- Evaluation Scope: $`0 \le n \le 10^5`$, $`-10^9 \le \mathrm{nums}[i] \le 10^9`$; expected $`O(n)`$ time, $`O(n)`$ extra space; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 3 | One must view a consecutive sequence as a maximal run in the value set and see the key relation "extend only from run heads, where `x-1` is absent" to avoid an $`O(n \log n)`$ sort |
| Knowledge Prerequisites | 2 | Only hash-set insertion and membership queries, plus the Go `map` behavior that a missing key reads as the zero value |
| Correctness Reasoning Difficulty | 3 | One must prove every consecutive sequence lies inside some maximal run (completeness) and use an inner-loop invariant to show each head yields its full run length; the $`O(n)`$ bound needs a separate aggregate analysis |
| Implementation Difficulty | 2 | Few variables, but one must deduplicate through the set before iterating and get the head test `!numSet[num-1]` right; iterating the original array degrades when values repeat |

Ratings follow [RATING.md](../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

Given an unsorted integer array `nums`, find the length of the longest group of **numerically consecutive** values (such as 1, 2, 3, 4). The values need not be adjacent or ordered in the array. The required time complexity is $`O(n)`$.

**Example**

- Input: `nums = [100,4,200,1,3,2]`
- Output: `4` (the longest consecutive sequence is 1, 2, 3, 4)

**Key requirements**

1. $`0 \le \mathrm{nums.length} \le 10^5`$ and $`-10^9 \le \mathrm{nums}[i] \le 10^9`$.
2. The time complexity must be $`O(n)`$, so sorting followed by a scan, at $`O(n \log n)`$, does not meet the requirement.
3. The array may be empty (answer 0), may contain duplicates (a repeated value counts once, so `[1,0,1,2]` gives 3), and may contain negative values.

## 4. Algorithm Specification

**Algorithm name:** Hash Set Run-Head Extension

### 4.1 Input

$`A = \langle a_0, a_1, \ldots, a_{n-1} \rangle`$

where:

- $`0 \le n \le 10^5`$
- $`-10^9 \le a_i \le 10^9`$

Let $`S = \{a_0, a_1, \ldots, a_{n-1}\}`$ be the set of distinct values in $`A`$, and $`m = |S| \le n`$.

### 4.2 Output

Call an integer set $`\{x, x+1, \ldots, x+L-1\}`$ ($`L \ge 1`$) a **consecutive sequence** of length $`L`$. The output is

```math
\mathrm{res} = \max\bigl\{\, L \;:\; \exists x,\ \{x, x+1, \ldots, x+L-1\} \subseteq S \,\bigr\}
```

with $`\mathrm{res} = 0`$ when $`S = \varnothing`$.

### 4.3 Pseudocode

The algorithm never indexes into the array, so the pseudocode has no indices.

```text
LONGEST-CONSECUTIVE(A)
    S ← empty set
    for each a in A
        INSERT(S, a)

    res ← 0
    for each x in S
        if x - 1 ∉ S                     // x is a run head
            cur ← x
            len ← 1
            while cur + 1 ∈ S
                cur ← cur + 1
                len ← len + 1
            res ← max(res, len)

    return res
```

**State**

- `S`: `numSet` in the code, the hash set of all distinct values of `A`.
- `cur`: `currentNum`, the largest value the current run has reached.
- `len`: `currentStreak`, the length of the current run so far.
- `res`: the longest run length among the heads processed so far.

The code implements `max` as `if res <= currentStreak { res = currentStreak }`; on equality it reassigns the same value, so the result is unchanged.

### 4.4 Correctness

**Proof Method: Partition into Maximal Runs + Inner and Outer Loop Invariants**

#### 1. Key Conditions for Correctness

Two definitions first.

- **Maximal run**: a subset $`\{h, h+1, \ldots, h+L-1\}`$ of $`S`$ with $`h - 1 \notin S`$ and $`h + L \notin S`$, so it cannot be extended at either end.
- **Run head**: a value $`h`$ with $`h \in S`$ and $`h - 1 \notin S`$.

Correctness relies on:

1. **Finiteness of the set** (given): $`S`$ is finite, so walking up or down one step at a time from any element leaves $`S`$ after finitely many steps.
2. **Heads correspond one-to-one with maximal runs** (Lemma 1): each head $`h`$ determines the unique maximal run starting at it, of length $`L(h)`$; the smallest element of every maximal run is a head.
3. **Every consecutive sequence lies in some maximal run** (Lemma 2): this guarantees that looking only at heads misses no answer.

#### 2. Correctness Proof

**Lemma 1 (inner loop)**: Let $`x`$ be a run head. The inner `while` loop terminates, and on exit $`\mathrm{len} = L(x)`$.

*Loop invariant*: each time the `while` condition is tested, $`\mathrm{cur} = x + \mathrm{len} - 1`$ and $`\{x, x+1, \ldots, \mathrm{cur}\} \subseteq S`$.

- **Initialization**: $`\mathrm{cur} = x`$, $`\mathrm{len} = 1`$, and $`x \in S`$, so it holds.
- **Maintenance**: entering the body means $`\mathrm{cur} + 1 \in S`$. After `cur` and `len` each increase by 1, the equation still holds, and the set gains the element $`\mathrm{cur}+1 \in S`$, so it is still contained in $`S`$.
- **Termination**: the elements of $`\{x, \ldots, \mathrm{cur}\} \subseteq S`$ are distinct, so $`\mathrm{len} \le m`$ and the loop exits after at most $`m - 1`$ iterations. On exit $`\mathrm{cur} + 1 \notin S`$, and since $`x`$ is a head, $`x - 1 \notin S`$. So $`\{x, \ldots, \mathrm{cur}\}`$ cannot be extended at either end: it is the maximal run starting at $`x`$, and $`\mathrm{len} = L(x)`$.

**Lemma 2 (completeness)**: If $`\{a, a+1, \ldots, a+L-1\} \subseteq S`$, there is a run head $`h \le a`$ with $`L \le L(h)`$.

*Proof*: Walk down from $`a`$, continuing while the current value minus 1 is in $`S`$. By condition 1 this stops after finitely many steps at some $`h \le a`$ with $`h \in S`$ and $`h - 1 \notin S`$, so $`h`$ is a head, and $`\{h, \ldots, a\} \subseteq S`$. Then $`\{h, \ldots, a+L-1\} \subseteq S`$ is a consecutive sequence starting at $`h`$. The maximal run starting at $`h`$ is the longest consecutive sequence that starts at $`h`$, so $`L(h) \ge (a + L - 1) - h + 1 \ge L`$.

**Outer loop invariant**

Let $`T \subseteq S`$ be the set of elements the outer loop has processed. At the start of each outer iteration:

**Invariant: longest run among processed heads**

$`\mathrm{res} = \max\{L(h) : h \in T,\ h \text{ is a run head}\}`$, and $`\mathrm{res} = 0`$ when $`T`$ contains no head.

**Initialization**

$`T = \varnothing`$ and $`\mathrm{res} = 0`$, so it holds.

**Maintenance**

Assume the invariant holds before element $`x`$ is processed.

- **Case 1: $`x - 1 \in S`$**. $`x`$ is not a head and `res` is unchanged; $`T \cup \{x\}`$ has the same heads as $`T`$, so the invariant still holds.
- **Case 2: $`x - 1 \notin S`$**. $`x`$ is a head. By Lemma 1 the inner loop computes $`\mathrm{len} = L(x)`$, and then $`\mathrm{res} \leftarrow \max(\mathrm{res}, L(x))`$ brings $`x`$ into the maximum, so the invariant still holds.

**Termination**

The outer loop ends after visiting the $`m`$ elements of $`S`$ exactly once. Then $`T = S`$ and $`\mathrm{res} = \max\{L(h) : h \text{ is a run head}\}`$.

- **Achievable**: each head's maximal run is itself a consecutive sequence of length $`L(h)`$, so $`\mathrm{res}`$ does not exceed the required maximum.
- **Upper bound**: by Lemma 2, every consecutive sequence is no longer than some $`L(h)`$, so $`\mathrm{res}`$ is at least the required maximum.

Together, $`\mathrm{res}`$ equals the length of the longest consecutive sequence. When $`S = \varnothing`$ the outer loop does not run and the function returns 0, matching the definition.

#### 3. Conclusion

$`S`$ splits into disjoint maximal runs, each represented by a unique head. The algorithm computes the exact run length at each head (Lemma 1), and no consecutive sequence is longer than the maximal run containing it (Lemma 2), so the maximum over all heads is the answer. Go randomizes `map` iteration order, but the maximum does not depend on processing order.

Also, $`a_i \ge -10^9`$, so `num-1` and `currentNum+1` stay within Go's `int` (64 bits on LeetCode) and cannot overflow.

### 4.5 Complexity Analysis

Let $`n`$ be the array length and $`m = |S| \le n`$ the number of distinct values. "Expected" below refers to the expected $`O(1)`$ time of a single hash-table operation (see [overview §2.1](../README.md)).

**Time complexity: expected O(n)**

- Building the set: $`n`$ insertions, expected $`O(n)`$.
- Outer loop: $`m`$ iterations with one membership query each, expected $`O(m)`$.
- Inner loop: expected $`O(m)`$, by the aggregate analysis below.

Bounding each inner loop separately gives at most $`m - 1`$ iterations per head, and multiplying by $`m`$ outer iterations gives a loose $`O(m^2)`$. **Aggregate analysis** instead counts how many times the inner loop runs over the whole algorithm. It starts only at heads and runs $`L(h) - 1`$ times at head $`h`$, so the total is

```math
\sum_{h \text{ a run head}} \bigl(L(h) - 1\bigr) = m - (\text{number of heads}) \le m
```

This uses the fact that maximal runs are disjoint and cover $`S`$, so their lengths sum to $`m`$. Each inner iteration makes one membership query, for an expected $`O(m)`$ in total.

The total is expected $`O(n + m) = O(n)`$. If hashing collided badly, a single operation could degrade to $`\Theta(n)`$; the Go runtime uses a random hash seed per `map`, so inputs cannot deliberately trigger this in practice.

**Space complexity: O(n)**

The hash set stores $`m \le n`$ distinct values, for $`O(m) = O(n)`$ extra space; everything else is a constant number of variables.

### 4.6 Implementation Notes: Go and Python

- **Set**: Go uses `map[int]bool`, where a missing key reads as `false`, so `numSet[num-1]` alone tests membership. Python uses a `set` and writes `num - 1 not in num_set`. See [hash tables](../../../docs/LANGUAGE_NOTES.md#hash-tables).
- **Run length**: Go counts with `currentStreak`; Python computes `current - num + 1`. The two are equivalent.
- Everything else corresponds line by line, and integer range is not an issue.

## 5. Learning Notes

### Challenge 1: Why is the nested loop O(n) and not O(n²)?

The code nests a `while` inside a `for`, which by its shape looks like $`O(n^2)`$.

**Understanding Process**

The key is the head test `!numSet[num-1]`. Without it, the run `{1,2,3,4}` would be extended from each of 1, 2, 3 and 4, taking $`4 + 3 + 2 + 1`$ steps, so a run of length $`m`$ costs $`\Theta(m^2)`$. With the test, only 1 starts an extension and 2, 3, 4 are skipped, so the run is scanned once. The right question is not "how long is one inner loop" but "how many inner iterations happen in total": each element is reached at most once, by the head of its own run, so the total is at most $`m`$.

**Transferable Insight:** The cost of nested loops is not always the product of the two loop counts. If each inner step consumes an object that can be consumed only once (here, an element of the set), aggregate analysis bounds the total cost by the number of objects. "Each element is pushed and popped at most once" in a monotonic stack is the same argument.

### Challenge 2: Why iterate over the set rather than the original array?

Both give the same answer, but not the same complexity.

**Understanding Process**

Iterating the original array processes duplicate values repeatedly. If `nums` is $`k`$ copies of 1 followed by the run $`1, 2, \ldots, k`$, the head 1 appears $`k+1`$ times and each occurrence scans the whole run, so the inner loop runs about $`k^2`$ times in total, degrading to $`\Theta(n^2)`$. Iterating the set visits each value once, so each head starts one extension and the $`O(n)`$ bound above holds.

**Transferable Insight:** When a complexity argument says "each object is processed once", check that the objects being iterated have no duplicates. If the input can repeat values, deduplicate first (the set does this for free here), then do the amortized analysis.

### Challenge 3: Why not just sort?

Sorting and then scanning also finds the answer and is more direct, but it takes $`O(n \log n)`$ time and misses the required $`O(n)`$.

**Understanding Process**

Sorting provides total-order information, while this problem needs only the local membership fact "is $`x+1`$ present". A hash set answers that in expected $`O(1)`$, trading $`O(n)`$ space for the logarithmic factor of sorting.

**Transferable Insight:** When an algorithm depends only on whether a value exists, not on the relative order of elements, consider a hash set instead of sorting. Ordered queries such as "the smallest element greater than $`x`$" are beyond a hash table.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Hash Set Membership**: here it answers "is $`x+1`$ present" in expected $`O(1)`$. It applies whenever an algorithm repeatedly asks whether a value has appeared. Limits: it answers membership only, not ordered queries, and its $`O(1)`$ is expected rather than worst case.
2. **Run-Head Enumeration**: here it extends only from `x` with `x-1` absent, so each run is scanned once. It applies when the answer is built from disjoint chains whose predecessor can be tested in $`O(1)`$, for example replacing the chain $`x \to x+1`$ with $`x \to x^2`$. It requires the chains to be disjoint, with each element on exactly one chain.
3. **Aggregate Analysis**: here it bounds the total inner-loop count by $`m - (\text{number of heads})`$. It applies to nested loops where each inner step consumes an object that can be consumed only once, as in monotonic stacks and two pointers.
4. **Completeness via Maximality**: Lemma 2 shows every feasible solution lies inside some maximal run, so only maximal runs need checking. It applies to "find the longest structure with property P" when P is closed under taking subsets.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. "Relation" is either "Direct" or "Related idea"; when no specific section matches, the "CLRS Section" column says "No matching section" and only the concept name is kept.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Loop invariants | §2.1 Insertion sort | Direct | Lemma 1's inner invariant and the outer invariant "`res` is the longest run among processed heads" are each proved by initialization, maintenance, and termination |
| Hash tables with chaining | §11.2 Hash tables | Direct | Insertion and membership queries on `numSet` take expected $`O(1)`$, which the overall expected $`O(n)`$ relies on |
| Aggregate analysis | §17.1 Aggregate analysis | Direct | The inner `while` runs $`m - (\text{number of heads}) \le m`$ times in total, counted over the whole algorithm rather than per loop |

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

- [1. Two Sum](https://leetcode.cn/problems/two-sum/) — Also uses a hash table to answer "is this value present" in expected $`O(1)`$, but stores a value-to-index map and inserts while scanning; it contrasts with this problem's "build the full set, then query".
- [674. Longest Continuous Increasing Subsequence](https://leetcode.cn/problems/longest-continuous-increasing-subsequence/) — A similar name, but the elements must be **adjacent in the array**, so one linear scan suffices and no hash table is needed; the contrast makes clear that "consecutive" here refers to values, not positions.
- [2501. Longest Square Streak in an Array](https://leetcode.cn/problems/longest-square-streak-in-an-array/) — Replaces the chain $`x \to x+1`$ with $`x \to x^2`$ and works with the same hash set plus run-head enumeration: extend only from elements whose square root is not in the set.

---

**One-Sentence Summary**

Turn the array into a hash set and count forward only from run heads where `x-1` is absent; each value is counted once, so the nested loops still take $`O(n)`$ total time.
