# 462. Minimum Moves to Equal Array Elements II — Sorting

English | [简体中文](./summary.zh-CN.md)

Terminology follows the [shared glossary](../../../../docs/GLOSSARY.md), and math notation follows the repository's [GitHub math writing conventions](../../../../docs/GITHUB_MATH.md).

This problem has three solutions in three directories that share one set of test cases:

- **Sorting** (this page): sort and take the median. This page gives the problem statement and the full proof that the median is optimal; the other two pages cite it.
- [**Three-way quickselect**](../three-way-quickselect/summary.md): randomized quickselect that partitions into less, equal, and greater blocks.
- [**Two-way (Lomuto) quickselect**](../lomuto-quickselect/summary.md): CLRS's PARTITION, which degrades with many duplicates.

## 1. Problem Information

- **Problem**: [462. Minimum Moves to Equal Array Elements II](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/)
- **Official Difficulty**: Medium
- **Primary Category**: Order Statistics
- **Secondary Category**: Sorting
- **Tags**: `Order Statistics`, `Sorting`, `Median`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](../solution_test.go) (in the parent directory, shared by all three solutions)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Median by Sorting**

- Evaluation Scope: $`1 \le n \le 10^5`$, $`-10^9 \le a_i \le 10^9`$, the answer fits in a 32-bit integer; time $`O(n \log n)`$, extra space $`\Theta(n)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2 | One must turn "minimum number of moves" into "minimum sum of absolute deviations" and realize the target is a median |
| Knowledge Prerequisites | 1 | Only sorting and absolute values |
| Correctness Reasoning Difficulty | 2 | Only median optimality is needed: one case split on "moving the target right by 1 changes the cost by $`2L - n`$" |
| Implementation Difficulty | 1 | Sort, take `a[n/2]`, and sum; the input is copied so the caller's slice is not modified |

Ratings follow [RATING.md](../../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

Given an integer array `nums` of length $`n`$, one move picks a single element and increments or decrements it by 1. Return the minimum number of moves needed to make all elements equal.

**Example**

- Input: `nums = [1,2,3]`
- Output: `2` (make everything 2: increment 1, decrement 3)
- Input: `nums = [1,10,2,9]`
- Output: `16`

**Key Requirements**

1. Input constraints: $`1 \le n \le 10^5`$, $`-10^9 \le a_i \le 10^9`$, and the answer is guaranteed to fit in a 32-bit integer.
2. Each move changes exactly one element by $`\pm 1`$; the final common value is not prescribed.
3. Edge cases: when $`n = 1`$ or all elements are already equal, the answer is 0; the array may contain many duplicates.

## 3. Initial Approach

### First Idea: Use the Median

Both versions had been written: a Python version that sorts and takes `nums[len(nums)//2]` as the target, and a Go version that finds the median with quickselect (see [two-way quickselect](../lomuto-quickselect/summary.md)). This page is the Python version written in Go.

### Difficulties Encountered

The open question was: if the median is far from the values on both sides (outliers), is it still the right target?

### Thought Transformation

From:

> The target should be "close" to every value, so outliers pull it toward them.

To:

> Moving the target by 1 changes the cost by an amount that depends only on how many values lie on each side, not on distances.

This observation is the proof of Lemma 2 below.

## 4. Algorithm Specification

**Algorithm Name:** Median by Sorting

### 4.1 Input

$`A = \langle a_0, a_1, \ldots, a_{n-1} \rangle`$

satisfying:

- $`1 \le n \le 10^5`$
- $`a_i \in \mathbb{Z}`$, $`-10^9 \le a_i \le 10^9`$

### 4.2 Output

The minimum number of moves

```math
\mathrm{res} = \min_{x \in \mathbb{Z}} \sum_{i=0}^{n-1} \lvert a_i - x \rvert .
```

### 4.3 Pseudocode

Indices are 0-based. Let $`m = \lfloor n/2 \rfloor`$.

```text
MIN-MOVES-2-SORT(A)
    B ← SORT(copy of A)              // ascending
    x ← B[m]
    res ← 0
    for i ← 0 to n - 1
        res ← res + |B[i] - x|
    return res
```

**State Definitions**

- `x`: the chosen target, the element at index $`m`$ after sorting.
- `res`: the total distance of all elements to `x`.

### 4.4 Correctness

**Proof Method:** reduce the problem to minimizing a sum of absolute deviations, then prove the median optimal by a case split on the cost difference.

Let $`b_0 \le b_1 \le \cdots \le b_{n-1}`$ be $`A`$ sorted in ascending order, and $`f(x) = \sum_i \lvert a_i - x \rvert`$.

#### I. Key Properties Required for Correctness

1. **Lemma 1 (Cost for a fixed target)**: making every element equal to $`x`$ takes at least, and at most, $`f(x)`$ moves.
2. **Lemma 2 (Median optimality)**: over the integers, $`f`$ attains its minimum at $`x = b_m`$.

Together they show the answer is $`f(b_m)`$, and after sorting `B[m]` is $`b_m`$. The other two solutions rely on the same two lemmas.

#### II. Correctness Proof

**Lemma 1**: each move changes one element by 1, so element $`i`$ needs at least $`\lvert a_i - x \rvert`$ moves, for a total of at least $`f(x)`$. Moving each element straight toward $`x`$ uses exactly $`f(x)`$ moves. Since every move changes a value by 1, the final value must be an integer, so

```math
\mathrm{res} = \min_{x \in \mathbb{Z}} f(x) .
```

**Lemma 2**: for an integer $`x`$, let $`L(x) = \#\{ i : a_i \le x \}`$. Moving the target from $`x`$ to $`x + 1`$ increases each term with $`a_i \le x`$ by 1 and decreases each term with $`a_i \ge x + 1`$ by 1, so

```math
f(x+1) - f(x) = L(x) - \bigl(n - L(x)\bigr) = 2L(x) - n .
```

- **When $`x < b_m`$**: since $`b_m > x`$, all of $`b_m, \ldots, b_{n-1}`$ exceed $`x`$, so $`L(x) \le m`$ and $`2L(x) - n \le 2m - n \le 0`$; $`f`$ does not increase.
- **When $`x \ge b_m`$**: all of $`b_0, \ldots, b_m`$ are at most $`x`$, so $`L(x) \ge m + 1`$ and $`2L(x) - n \ge 2m + 2 - n \ge 1`$; $`f`$ strictly increases.

So $`f`$ is non-increasing to the left of $`b_m`$ and strictly increasing from $`b_m`$ on, and its minimum is $`f(b_m)`$. $`\square`$

The argument uses only $`L(x)`$, how many values lie to the left, and not how far away outliers are. When $`n`$ is even, every integer in $`[b_{m-1}, b_m]`$ has the same cost, so $`b_m`$ suffices.

#### III. Conclusion

By Lemmas 1 and 2, the answer is $`f(b_m)`$. The algorithm takes $`B[m] = b_m`$ after sorting and sums over the sorted copy; sorting changes only the order, not the multiset, so the sum equals $`f(b_m)`$. The algorithm returns the correct answer on every valid input.

### 4.5 Complexity Analysis

Let the input size be $`n`$.

**Time Complexity: O(n log n)**

Sorting takes $`O(n \log n)`$ and the sum one $`\Theta(n)`$ pass.

**Space Complexity: Θ(n)**

$`\Theta(n)`$ for the copy, plus $`O(\log n)`$ stack space for `sort.Ints`. If modifying the input is allowed, sorting in place avoids the copy.

**Numeric range**: $`f(b_m) \le n \cdot 2 \cdot 10^9 = 2 \times 10^{14}`$, which fits in a 64-bit `int`; the problem also guarantees the answer fits in 32 bits.

## 5. Learning Notes

### Challenge 1: Is the median still right when it is far from both sides?

**Understanding Process**

The difference in Lemma 2 is $`2L(x) - n`$, which depends only on how many values lie to the left of $`x`$, not on distances. Moving the target 1 step toward an outlier lowers that outlier's cost by 1 but raises the cost of more elements by 1. For `[1,2,3,1000,1001]`, target 3 costs 1998 and target 1000 costs 2995.

**Transferable Insight:** minimizing a sum of absolute values depends only on counts, so the median is robust to outliers; minimizing a sum of squares depends on distances, so the mean is pulled by outliers.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Median Minimizes Absolute Deviation**: here the minimum number of moves becomes $`\min_x \sum \lvert a_i - x \rvert`$, attained at a median. Applies to "change one element by 1 per move until all are equal", one-dimensional facility location, and similar problems. Condition: the cost is a sum of absolute distances; if each move changes a value by a fixed $`d`$, first check that all values are congruent modulo $`d`$.
2. **Cost Difference Argument**: compare $`f(x+1)`$ with $`f(x)`$ and find where the difference turns from negative to positive. Applies when the objective is a sum of piecewise-linear convex functions.

Sorting is the most direct way to find a median; when only one order statistic is needed, quickselect can replace it, as the other two pages show.

### CLRS Connections

Chapter numbers follow the third edition of CLRS. "Relation" is either "Direct" or "Related idea"; when no specific section matches, the "CLRS Section" column says "No matching section" and only the concept is kept.

| Concept | CLRS Section | Relation | Where it appears in this problem |
|---|---|---|---|
| Median minimizes absolute deviation | No matching section | Related idea | The cost-difference argument in Lemma 2 |

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

- [453. Minimum Moves to Equal Array Elements](https://leetcode.com/problems/minimum-moves-to-equal-array-elements/) — also asks for all elements to become equal, but one move increments $`n-1`$ elements; after the equivalent "decrement one element" view, values can only go down and the target is pinned to the minimum rather than a median.
- [2033. Minimum Operations to Make a Uni-Value Grid](https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/) — each move adds or subtracts a fixed $`x`$; first check that all values are congruent modulo $`x`$, then choose a median likewise.

---

**One-Sentence Summary**

A sum of absolute deviations depends only on how many values lie on each side, so the target is a median; after sorting, take the element at index $`\lfloor n/2 \rfloor`$.
