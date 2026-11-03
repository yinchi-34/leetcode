# 453. Minimum Moves to Equal Array Elements — Formula

English | [简体中文](./summary.zh-CN.md)

Terminology follows the [shared glossary](../../../../docs/GLOSSARY.md), and math notation follows the repository's [GitHub math writing conventions](../../../../docs/GITHUB_MATH.md).

This problem has two solutions in two directories that share one set of test cases:

- [**Batched simulation**](../batched-simulation/summary.md): after sorting, the group of equal elements chases the next element in batches. The problem statement, Lemma 1 (counting), the potential-function lower bound, and the proof of the batched simulation are there.
- **Formula** (this page): one pass computing $`\sum a_i - n \cdot \min a`$, proved by the "reverse view" equivalence.

This page covers only two points: where the formula "compresses" the simulation away, and why the reverse view is valid.

## 1. Problem Information

- **Problem**: [453. Minimum Moves to Equal Array Elements](https://leetcode.com/problems/minimum-moves-to-equal-array-elements/)
- **Official Difficulty**: Medium
- **Primary Category**: Simulation
- **Secondary Category**: Math
- **Tags**: `Simulation`, `Math`, `Equivalent Transformation`, `Translation Invariance`, `Lower Bound and Achievability`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](../solution_test.go) (in the parent directory, shared by both solutions)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Formula**

- Evaluation Scope: $`1 \le n \le 10^5`$, $`-10^9 \le a_i \le 10^9`$, the answer fits in a 32-bit integer; time $`\Theta(n)`$, extra space $`O(1)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 4 | One must split "add 1 to the other $`n-1`$ elements" into "add 1 to all, then subtract 1 from one", and use the fact that a uniform shift does not affect equality to transform the problem into "subtract 1 from one element per move" |
| Knowledge Prerequisites | 1 | Only a sum and a minimum are needed |
| Correctness Reasoning Difficulty | 3 | Requires induction on the step count to show the two problems' states differ by $`t \cdot \mathbf{1}`$, followed by separate lower-bound and achievability arguments |
| Implementation Difficulty | 1 | One pass; $`\sum a_i`$ can reach about $`10^{14}`$, which fits in Go's 64-bit `int` |

Ratings follow [RATING.md](../../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

The statement, example, and constraints are in the [batched simulation solution](../batched-simulation/summary.md#2-problem-understanding).

## 3. Initial Approach

### First Idea: Look at It in Reverse

The code was written as `sum - len(nums)*mn`, but why the reverse view works was not understood: "add 1 to the other $`n-1`$" and "subtract 1 from one" make no difference to the gaps, but why does that alone justify switching to a different problem?

### Difficulties Encountered

Several questions came up in turn during the discussion:

1. How is this line computed? Why are there $`n`$ copies of $`\min`$?
2. The batched simulation already compresses each batch into one subtraction; where does the formula do its compression?
3. In the reverse view, could some value be decreased below the minimum, changing the target?
4. Going forward, "which element to operate on" keeps changing; does it not change in reverse too? What is the point of reversing?

### Thought Transformation

From:

> The reverse view is a different process, and its reference point moves too.

To:

> The reverse view is the same process seen from another frame of reference: after removing the uniform rise at each step, each step touches only one element, the elements do not affect one another, and the minimum stays put.

The answers to these questions are in Section 5.

## 4. Algorithm Specification

**Algorithm Name:** Closed-Form Formula

### 4.1 Input

$`A = \langle a_0, a_1, \ldots, a_{n-1} \rangle`$

satisfying:

- $`1 \le n \le 10^5`$
- $`a_i \in \mathbb{Z}`$, $`-10^9 \le a_i \le 10^9`$

### 4.2 Output

The minimum number of moves $`k^* = \sum_i a_i - n \cdot \min_i a_i`$.

### 4.3 Pseudocode

Indices are 0-based.

```text
MIN-MOVES(A)
    S ← 0
    μ ← A[0]
    for i ← 0 to n - 1
        S ← S + A[i]
        μ ← min(μ, A[i])
    return S - n · μ
```

**State Definitions**

- `S`: sum of the elements seen so far.
- `μ`: minimum of the elements seen so far.

### 4.4 Correctness

**Proof Method:** an equivalent transformation (translation invariance plus induction on the step count), followed by lower-bound and achievability arguments for the equivalent problem.

**Notation**: $`S = \sum_i a_i`$, $`\mu = \min_i a_i`$. $`\mathbf{1} = (1, \ldots, 1)`$, and $`e_j`$ is the vector with 1 in position $`j`$ and 0 elsewhere.

#### I. Key Properties Required for Correctness

1. **Lemma 2 (Translation Invariance)**: for any constant $`t`$, all entries of $`x`$ are equal if and only if all entries of $`x + t\mathbf{1}`$ are equal.
2. **Lemma 3 (Equivalent Problems)**: the original problem and the problem "subtract 1 from one element per move" become all-equal at the same move.

Lemma 2: $`x_i = x_j \iff x_i + t = x_j + t`$.

#### II. Correctness Proof

**Lemma 3**: start from $`a`$ and use the same index sequence $`j_0, j_1, \ldots`$. In the original problem, move $`t`$ skips $`j_t`$; in the new problem, move $`t`$ subtracts 1 from $`j_t`$. Let $`x^{(t)}`$ and $`y^{(t)}`$ be the respective states after $`t`$ moves. Then

```math
x^{(t)} = y^{(t)} + t\,\mathbf{1} .
```

*Proof*: by induction on $`t`$. For $`t = 0`$ both equal $`a`$. If it holds for $`t`$, then since skipping $`j_t`$ equals "add 1 to all, then subtract 1 from $`j_t`$",

```math
x^{(t+1)} = x^{(t)} + \mathbf{1} - e_{j_t} = \bigl(y^{(t)} - e_{j_t}\bigr) + (t+1)\mathbf{1} = y^{(t+1)} + (t+1)\mathbf{1} . \qquad \square
```

By Lemma 2, $`x^{(t)}`$ is all-equal if and only if $`y^{(t)}`$ is. Move sequences of the two problems correspond one to one with equal lengths, so their minimum move counts coincide. Note that the two problems **end at different values**; what coincides is the move at which all elements become equal.

**Solving the equivalent problem**: in the new problem, values only decrease.

- **Lower bound**: if all end at $`T`$, then $`T \le \mu`$, and the number of moves is $`\sum_i (a_i - T) = S - nT \ge S - n\mu`$.
- **Achievability**: take $`T = \mu`$ and decrease element $`i`$ exactly $`a_i - \mu`$ times, for $`S - n\mu`$ moves.

Hence $`k^* = S - n\mu`$.

#### III. Conclusion

By Lemma 3, the minimum for the original problem equals the minimum $`S - n\mu`$ for the equivalent problem, which the algorithm returns after one pass computing $`S`$ and $`\mu`$. The same value also follows independently from the potential-function lower bound in the [batched simulation solution](../batched-simulation/summary.md#44-correctness).

### 4.5 Complexity Analysis

Let the input size be $`n`$.

**Time Complexity: Θ(n)**

One pass with one addition and one comparison per element.

**Space Complexity: O(1)**

Two variables.

**Numeric range**: $`\lvert S \rvert \le 10^5 \cdot 10^9 = 10^{14}`$, and $`n \cdot \lvert\mu\rvert`$ is bounded likewise, which fits in a 64-bit `int`.

## 5. Learning Notes

### Challenge 1: Why $`n`$ copies of $`\min`$, and where the compression happens

**Understanding Process**

- **Splitting the sum**: $`\sum_{i=0}^{n-1} (a_i - \mu) = \sum a_i - \sum \mu`$. $`\mu`$ does not depend on $`i`$ and there are $`n`$ terms, so the second sum is $`n\mu`$. The minimum's own term is $`\mu - \mu = 0`$, but it still counts among the $`n`$ terms.
- **Two compressions**: the batched simulation compresses the $`a_i - a_0`$ moves of batch $`i`$ into one subtraction; that is the first. Each batch length depends only on $`a_i`$ and the minimum, not on the order or lengths of other batches, so there is no need to sort or chase in order; summing all batches directly is the second, and it is the line `sum - n*mn`.
- **As a bar chart**: draw each value as a column of cells. There are $`S`$ cells in total, $`n\mu`$ of them at or below the minimum's level, and each cell above that level is one move. The batched simulation counts those cells column by column; the formula counts all cells and subtracts the bottom rectangle. Both count the same cells.

**Transferable Insight:** if the batch lengths do not depend on one another, the order of batches does not matter, and the sum can be written in closed form.

### Challenge 2: Does the reference point move in the reverse view?

**Understanding Process**

There are two kinds of "reference point":

- **The endpoint**: in reverse, values only decrease, and going below the minimum only adds moves (each unit lower costs one extra move for each of the $`n`$ elements), so an optimal schedule never touches the minimum and the endpoint stays at $`\mu`$.
- **Which element to operate on next**: if the reverse view insisted on "always decrease the maximum", this choice would change too. But in reverse each move touches one element, the elements do not affect one another, and element $`i`$ needs exactly $`a_i - \mu`$ decrements in any order, so no choice is needed at all.

The forward view works the same way: as long as element $`i`$ is skipped $`a_i - \mu`$ times, any order makes all elements equal at the end (Lemma 1). It is hard to see going forward because every move touches $`n-1`$ elements.

**Frame of reference**: the reverse view watches the forward process from an elevator that rises by 1 per move, $`y = x - t\mathbf{1}`$. The minimum is never skipped and rises by exactly 1 per move, so from inside the elevator it stands still. For `[1,2,3]`:

| Move $`t`$ | Forward $`x`$ | Reverse $`y = x - t\mathbf{1}`$ |
|---|---|---|
| 0 | `[1,2,3]` | `[1,2,3]` |
| 1 (skip the 3rd) | `[2,3,3]` | `[1,2,2]` |
| 2 (skip the 2nd) | `[3,3,4]` | `[1,1,2]` |
| 3 (skip the 3rd) | `[4,4,4]` | `[1,1,1]` |

**Transferable Insight:** the reverse view is not a new algorithm but a point of view: it turns "each move touches $`n-1`$ elements" into "each move touches one element", which reveals that the elements are independent.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Translation Invariance**: when the goal depends only on differences, decompose a move into a uniform shift plus a local change and pass to an equivalent problem; here "add 1 to $`n-1`$ elements" becomes "subtract 1 from one element". Limitation: valid only when the goal truly ignores uniform shifts; if the final value were prescribed, the shift could not be dropped.
2. **Closed Form from Independent Batches**: when batch lengths do not depend on one another, skip the simulation and sum them directly. Applies to counting problems where one can first show that each part's contribution does not depend on the order.

### CLRS Connections

Chapter numbers follow the third edition of CLRS. "Relation" is either "Direct" or "Related idea"; when no specific section matches, the "CLRS Section" column says "No matching section" and only the concept is kept.

| Concept | CLRS Section | Relation | Where it appears in this problem |
|---|---|---|---|
| Translation invariance and equivalent transformation | No matching section | Related idea | Lemma 3: $`x^{(t)} = y^{(t)} + t\mathbf{1}`$ |

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

- [462. Minimum Moves to Equal Array Elements II](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/) — each move adds or subtracts 1 from a single element, so values can move in both directions; the target is no longer pinned to the minimum by monotonicity and must be chosen as a median.

---

**One-Sentence Summary**

Removing the uniform rise at each step turns "add 1 to the others" into "subtract 1 from one"; each element decreases to the minimum on its own, for $`\sum a_i - n \cdot \min a`$ moves in total.
