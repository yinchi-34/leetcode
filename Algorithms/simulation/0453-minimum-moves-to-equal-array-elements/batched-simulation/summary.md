# 453. Minimum Moves to Equal Array Elements — Batched Simulation

English | [简体中文](./summary.zh-CN.md)

Terminology follows the [shared glossary](../../../../docs/GLOSSARY.md), and math notation follows the repository's [GitHub math writing conventions](../../../../docs/GITHUB_MATH.md).

This problem has two solutions in two directories that share one set of test cases:

- **Batched simulation** (this page): after sorting, the group of equal elements chases the next element in batches. This page gives the problem statement, the potential-function lower bound, and the full proof of the batched simulation.
- [**Formula**](../formula/summary.md): one pass computing $`\sum a_i - n \cdot \min a`$, proved by the "reverse view" equivalence.

## 1. Problem Information

- **Problem**: [453. Minimum Moves to Equal Array Elements](https://leetcode.com/problems/minimum-moves-to-equal-array-elements/)
- **Official Difficulty**: Medium
- **Primary Category**: Simulation
- **Secondary Category**: Math
- **Tags**: `Simulation`, `Math`, `Potential Function`, `Lower Bound and Achievability`, `Sorting`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](../solution_test.go) (in the parent directory, shared by both solutions)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Batched Simulation**

- Evaluation Scope: $`1 \le n \le 10^5`$, $`-10^9 \le a_i \le 10^9`$, the answer fits in a 32-bit integer; time $`O(n \log n)`$, extra space $`O(n)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 3 | One must notice that consecutive steps skipping the same element can be merged, and that the element being chased and $`a_0`$ have never been skipped, so each batch has length equal to the original gap $`a_i - a_0`$ |
| Knowledge Prerequisites | 2 | Requires sorting and loop invariants |
| Correctness Reasoning Difficulty | 3 | Needs a loop invariant to show each batch is valid and ends in a caught-up prefix, plus a separate potential-function argument for the lower bound |
| Implementation Difficulty | 2 | The loop must compute `a[i] + res` before updating `res` and then overwrite `a[i]` with the caught-up value; the update order is easy to get wrong |

Ratings follow [RATING.md](../../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

Given an integer array `nums` of length $`n`$, one move increments $`n-1`$ of its elements by 1. Return the minimum number of moves needed to make all elements equal.

**Example**

- Input: `nums = [1,2,3]`
- Output: `3` (`[1,2,3] → [2,3,3] → [3,4,3] → [4,4,4]`)

**Key Requirements**

1. Input constraints: $`1 \le n \le 10^5`$, $`-10^9 \le a_i \le 10^9`$, and the answer is guaranteed to fit in a 32-bit integer.
2. Each move leaves exactly one element unchanged; below, this is called "skipping" that element. The final common value is not prescribed.
3. Edge cases: when $`n = 1`$ or all elements are already equal, the answer is 0.

## 3. Initial Approach

### First Idea: Look at It in Reverse

The code was written as "the answer is $`\sum a_i - n \cdot \min a`$" and passed the tests, but why the "reverse view" works was not understood. Two questions remained:

1. Why not simply simulate step by step? Nothing in the problem forbids it.
2. "Add 1 to the other $`n-1`$" and "subtract 1 from one" make no difference to the gaps between elements. But why does that alone justify switching to a different problem?

The second question is discussed in the [formula solution](../formula/summary.md); this page answers the first.

### Difficulties Encountered

Step-by-step simulation is valid; the issue is the number of steps. For `[1, 1000000000]`, each move shrinks the gap by at most 1, so about $`10^9`$ moves are needed; with $`n`$ up to $`10^5`$ and $`O(n)`$ work per move, the total can reach the order of $`10^{14}`$ operations.

So the real question is: can the number of moves be computed **without** making them one by one?

### Thought Transformation

From:

> Which element should be skipped at each step?

To:

> Many consecutive steps skip the same element. Can they be merged into one batch whose length is computed at once?

The process worked out along the way: after sorting, the group of elements that are already equal chases the next element. While chasing, the chased element is skipped and everything else increases by 1; the batch stops exactly when the group catches up, the chased element joins the group, and the group then chases the next one, until all are equal.

Two corrections to this process were needed:

- The length of a batch is not "current minimum versus global maximum", but the gap between the group's value and the **current** value of the chased element, which has been increasing in every earlier batch.
- Each batch skips the chased element, which is usually not the global maximum. So the basic step-by-step simulation must be stated as "skip any non-minimum element each step", not "skip the maximum"; only then is the batched simulation a special case of it, and the two can share one proof.

Pushing further shows that each batch has length exactly $`a_i - a_0`$ (Challenge 1 in Section 5).

## 4. Algorithm Specification

**Algorithm Name:** Batched Simulation

### 4.1 Input

$`A = \langle a_0, a_1, \ldots, a_{n-1} \rangle`$

satisfying:

- $`1 \le n \le 10^5`$
- $`a_i \in \mathbb{Z}`$, $`-10^9 \le a_i \le 10^9`$

### 4.2 Output

The minimum number of moves $`k^*`$: some sequence of $`k^*`$ moves makes all elements equal, and no shorter sequence does.

### 4.3 Pseudocode

Indices are 0-based.

```text
MIN-MOVES-BATCHED(A)
    B ← SORT(copy of A)              // ascending
    res ← 0
    for i ← 1 to n - 1
        cur ← B[i] + res             // B[i] has never been skipped, so it gained res
        res ← res + (cur - B[i-1])   // the prefix catches up: skip B[i] cur - B[i-1] times
        B[i] ← cur                   // the whole prefix now equals cur
    return res
```

**State Definitions**

- `res`: total number of moves made so far.
- `B[i-1]`: the common value of the prefix $`B[0..i-1]`$ at the start of batch $`i`$.

### 4.4 Correctness

**Proof Method:** a potential function gives the lower bound; a loop invariant shows the batched simulation achieves it.

**Notation**: $`S = \sum_i a_i`$, $`\mu = \min_i a_i`$. The move $`\mathrm{op}_j`$ adds 1 to every element except index $`j`$, called "skipping $`j`$".

**Theorem**: the minimum number of moves is $`k^* = S - n\mu`$.

#### I. Key Properties Required for Correctness

1. **Lemma 1 (Counting)**: after $`k`$ moves in which index $`j`$ is skipped $`c_j`$ times, $`\sum_j c_j = k`$ and the current value is $`x_i = a_i + k - c_i`$. In particular, elements never skipped have all gained $`k`$, so the gaps between them are unchanged.
2. **Potential Function**: $`\Phi(x) = \sum_i x_i - n \cdot \min x`$. We have $`\Phi(x) \ge 0`$, with $`\Phi(x) = 0`$ if and only if all entries are equal.

Lemma 1: each move skips exactly one index, so $`\sum_j c_j = k`$; element $`i`$ fails to increase only in the $`c_i`$ moves that skip it.

#### II. Correctness Proof

**(1) Direct simulation: lower bound and achievability**

Consider the effect of any single move on $`\Phi`$. The sum increases by $`n-1`$; the minimum increases by at most 1, and by exactly 1 when the skipped element is not a minimum, since then every minimum element increases. Hence

```math
\Delta\Phi = (n-1) - n \cdot \Delta\min x \ \ge\ -1,
```

with equality when the skipped element is not a minimum.

- **Lower bound**: each move decreases $`\Phi`$ by at most 1, and the goal requires $`\Phi = 0`$, so every valid sequence needs at least $`\Phi(a) = S - n\mu`$ moves.
- **Achievability**: as long as not all elements are equal, skip a non-minimum element; $`\Phi`$ then drops by exactly 1 per move, and the process ends after $`S - n\mu`$ moves.

Therefore $`k^* = S - n\mu`$. The only problem with direct simulation is the number of moves.

**(2) Loop invariant for the batched simulation**

Let `B` be sorted in ascending order, so $`B_0 = \mu`$. View the process as actually performing moves on the sorted array. At the start of iteration $`i`$ of the `for` loop ($`1 \le i \le n`$, where $`i = n`$ means the loop has ended):

**Invariant I: equal prefix**

After `res` moves, the prefix $`0..i-1`$ has all current values equal to `B[i-1]`.

**Invariant II: untouched suffix**

Index $`0`$ and every $`j \ge i`$ have never been skipped; for $`j \ge i`$, `B[j]` still holds the original value $`B_j`$.

**Initialization**: when $`i = 1`$, `res = 0`, the prefix has a single element, and nothing has been skipped, so both invariants hold.

**Maintenance**: suppose the invariants hold before iteration $`i`$. By Invariant II and Lemma 1, the current value of index $`i`$ is $`B_i + \mathrm{res}`$, which is `cur`. Index $`0`$ has never been skipped either, so its current value is $`B_0 + \mathrm{res}`$, which equals `B[i-1]` by Invariant I. Hence the batch length is

```math
d_i = \mathrm{cur} - B[i-1] = (B_i + \mathrm{res}) - (B_0 + \mathrm{res}) = B_i - B_0 \ge 0 .
```

Nonnegativity follows from sorting, so the batch is well defined. Skipping index $`i`$ exactly $`d_i`$ times leaves index $`i`$ unchanged while the prefix rises by 1 per move and catches up to `cur` exactly; elements $`j > i`$ also gain $`d_i`$ without being skipped. After `res += d_i` and `B[i] = cur`, the prefix $`0..i`$ all equal `B[i]`, and only index $`i`$ has newly been skipped, so the invariants hold for $`i+1`$.

**Termination**: the loop ends at $`i = n`$. By Invariant I, all $`n`$ elements are equal, and the total number of moves is

```math
\mathrm{res} = \sum_{i=1}^{n-1} (B_i - B_0) = S - n\mu .
```

In each batch the skipped index $`i`$ is strictly greater than the prefix until it is caught, so every move skips a non-minimum element; this is one of the achieving schedules in (1).

#### III. Conclusion

Part (1) gives the lower bound $`S - n\mu`$ for every valid input, and Part (2) shows the batched simulation performs a schedule meeting it, so the algorithm returns $`k^*`$.

### 4.5 Complexity Analysis

Let the input size be $`n`$.

**Time Complexity: O(n log n)**

Sorting takes $`O(n \log n)`$, followed by one $`\Theta(n)`$ pass.

**Space Complexity: O(n)**

Copying the array before sorting uses $`\Theta(n)`$ extra space. If modifying the input is allowed, it can sort in place, leaving the extra space to `sort.Ints` (Go's pdqsort uses $`O(\log n)`$ stack space).

**Numeric range**: $`\lvert S \rvert \le 10^5 \cdot 10^9 = 10^{14}`$, and $`n \cdot \lvert\mu\rvert`$ is bounded likewise, which fits in a 64-bit `int`.

## 5. Learning Notes

### Challenge 1: Why each batch has length exactly $`a_i - a_0`$

Deriving batch 2 symbolically gives $`d_2 = (a_2 + d_1) - a_1`$, and substituting $`d_1 = a_1 - a_0`$ makes $`d_1`$ cancel. The first question is: why is the group's value exactly $`a_1`$?

**Understanding Process**

- Every batch stops exactly when the group catches up. While chasing, the chased element stays put and the group rises by 1 per move; continuing to skip it after catching up would push the group past it and waste moves that must later be undone. So after batch 1 the group equals exactly $`a_1`$.
- Earlier step counts always cancel because $`a_0`$ and the chased $`a_i`$ have both never been skipped before batch $`i`$. By Lemma 1 they have gained the same amount, so their gap is still the original $`a_i - a_0`$.

With concrete numbers `[1,2,5,7]`, the three batches take 1, 4, and 6 moves, which are $`2-1`$, $`5-1`$, and $`7-1`$, totaling 11.

**Transferable Insight:** when merging repeated steps, look for elements that have not been affected so far; their gaps are unchanged and often give the batch length directly.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Potential Function**: here $`\Phi = \sum x_i - n \cdot \min x`$ measures how far the state is from the goal. Showing "each move decreases it by at most 1" gives the lower bound, and a schedule that decreases it by exactly 1 per move gives achievability. Applies to minimum-move problems with unit-cost moves, provided one can find a quantity whose per-move change is bounded and which is 0 exactly at the goal.
2. **Batching Repeated Steps**: here consecutive moves skipping the same element are merged and counted at once, turning a simulation with up to $`10^9`$ moves into $`n-1`$ batches. Applies to simulations with long runs of identical actions, provided each batch length can be computed directly from the current state.

The batch lengths $`a_i - a_0`$ do not depend on one another, so all batches can be merged once more into a formula; see the [formula solution](../formula/summary.md).

### CLRS Connections

Chapter numbers follow the third edition of CLRS.

| Concept | CLRS Section | Relation | Where it appears in this problem |
|---|---|---|---|
| Loop invariant | §2.1 Insertion sort | Direct | Proved by the invariants "equal prefix, untouched suffix", through initialization, maintenance, and termination |
| Potential method | §17.3 The potential method | Related idea | The state is measured by $`\Phi = \sum x_i - n \cdot \min x`$, and bounding $`\Delta\Phi \ge -1`$ per move yields the lower bound; here it proves a lower bound rather than performing amortized analysis |

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

- [462. Minimum Moves to Equal Array Elements II](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/) — also asks for all elements to become equal, but each move adds or subtracts 1 from a single element; the target is no longer pinned to the minimum and must be chosen as a median.

---

**One-Sentence Summary**

Step-by-step simulation is correct but too slow; after sorting, the equal group chases the next element in batches of exactly $`a_i - a_0`$ moves, for $`\sum a_i - n \cdot \min a`$ in total.
