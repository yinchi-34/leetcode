# 3736. Minimum Moves to Equal Array Elements III

English | [简体中文](./summary.zh-CN.md)

Terminology follows the [shared glossary](../../../docs/GLOSSARY.md), and math notation follows the repository's [GitHub math writing rules](../../../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [3736. Minimum Moves to Equal Array Elements III](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-iii/)
- **Official difficulty**: Easy
- **Primary category**: Simulation
- **Secondary category**: Math
- **Tags**: `math`, `lower bound and achievability`, `negation`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External ratings**

- zerotrac rating: 1251.75 ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode official difficulty: Easy

**Four-dimension rating — Formula**

- Scope: $`1 \le n \le 100`$, $`1 \le a_i \le 100`$; time $`\Theta(n)`$, extra space $`O(1)`$; Go implementation.

| Dimension | Score | Rationale |
|---|---|---|
| Modeling | 1 | Moves can only increase values, so monotonicity fixes the target at the maximum |
| Prerequisite knowledge | 1 | Only a sum and a maximum are needed |
| Correctness argument | 1 | One lower bound plus one construction; no invariant or transformation is needed |
| Implementation | 1 | One pass; the value range is tiny, so overflow cannot occur |

Ratings follow the repository's [RATING.md](../../../docs/RATING.md); each dimension is rated independently, with no combined total.

## 2. Problem Understanding

Given an integer array `nums` of length $`n`$, one move picks **one** element and increases it by 1. Return the minimum number of moves to make all elements equal.

**Examples**

- Input: `nums = [2,1,3]`
- Output: `3` (add 1 to the first element and 2 to the second; all become 3)

- Input: `nums = [4,4,5]`
- Output: `2`

**Key requirements**

1. Constraints: $`1 \le n \le 100`$, $`1 \le a_i \le 100`$.
2. Each move changes exactly one element, and only upward. The final value is not prescribed.
3. Edge cases: if $`n = 1`$ or all elements are already equal, the answer is 0.

## 3. Initial Approach

### First Idea: By Analogy with 453

The [formula solution of 453](../0453-minimum-moves-to-equal-array-elements/formula/summary.md) views "add 1 to the other $`n-1`$" in reverse as "subtract 1 from one", with answer $`\sum a_i - n \cdot \min a`$. This problem's move is "add 1 to one", the mirror image of that equivalent problem: the direction flips from subtracting to adding, so the reference point moves from the minimum to the maximum, and the answer should be $`n \cdot \max a - \sum a_i`$.

### Difficulties Encountered

The analogy only gives a conjecture. Two things still need checking:

1. Whether the mirror is exact, that is, whether swapping "add" and "subtract" turns "minimum" into exactly "maximum".
2. Whether this problem can be proved directly, without going through 453's reverse view.

### Thought Transformation

From:

> This is a variant of 453 and relies on its result.

To:

> This is 453's equivalent problem itself, up to a sign; and since each move touches only one number, it can be proved directly with no transformation at all.

453 needs the reverse view because each of its moves touches $`n-1`$ numbers. Here each move touches one number, so the elements are independent from the start.

## 4. Algorithm Specification

**Algorithm:** Closed-Form Formula

### 4.1 Input

$`A = \langle a_0, a_1, \ldots, a_{n-1} \rangle`$

such that:

- $`1 \le n \le 100`$
- $`a_i \in \mathbb{Z}`$, $`1 \le a_i \le 100`$

### 4.2 Output

The minimum number of moves $`k^* = n \cdot \max_i a_i - \sum_i a_i`$.

### 4.3 Pseudocode

Indices start at 0.

```text
MIN-MOVES(A)
    S ← 0
    M ← A[0]
    for i ← 0 to n - 1
        S ← S + A[i]
        M ← max(M, A[i])
    return n · M - S
```

**State**

- `S`: sum of the elements scanned so far.
- `M`: maximum of the elements scanned so far.

### 4.4 Correctness

**Method: lower bound and achievability; as a cross-check, the same result is derived from 453's lemma by negation.**

**Notation**: $`S = \sum_i a_i`$, $`M = \max_i a_i`$. Consider any plan with $`k`$ moves in which element $`i`$ is increased $`c_i \ge 0`$ times, and all elements finally equal $`T`$.

#### I. Key Properties Required for Correctness

1. **Counting**: each move increases exactly one $`c_i`$ by 1, so $`k = \sum_i c_i`$; and the final value is $`a_i + c_i = T`$, so $`c_i = T - a_i`$.
2. **Monotonicity**: $`c_i \ge 0`$, so $`T \ge a_i`$ for every $`i`$, that is, $`T \ge M`$.

By counting,

```math
k = \sum_{i=0}^{n-1} (T - a_i) = nT - S .
```

The number of moves depends only on the final value $`T`$, not on the order of moves.

#### II. Correctness Proof

**Lower bound**: every valid plan has $`T \ge M`$, so

```math
k = nT - S \ge nM - S .
```

**Achievability**: take $`T = M`$ and increase element $`i`$ exactly $`M - a_i \ge 0`$ times, in any order. All elements end at $`M`$ after $`nM - S`$ moves.

The lower bound is achieved, so $`k^* = nM - S`$.

**Cross-check: from 453 by negation**

Let $`b_i = -a_i`$. Here "increase $`a_j`$ by 1" is "decrease $`b_j`$ by 1", and $`a`$ is all equal if and only if $`b`$ is. So the minimum number of moves equals that of "decrease one number by 1 per move" on $`b`$, which is exactly the equivalent problem obtained in 453's [Lemma 3](../0453-minimum-moves-to-equal-array-elements/formula/summary.md#44-correctness). Its answer is $`\sum_i b_i - n \cdot \min_i b_i`$. Since $`\min_i(-a_i) = -\max_i a_i`$,

```math
\sum_i b_i - n \cdot \min_i b_i = -S - n \cdot (-M) = nM - S ,
```

which agrees with the direct proof. Negation turns "only decrease, end at the minimum" into "only increase, end at the maximum"; this is the mirror described in Section 3.

#### III. Conclusion

For every valid input the minimum is $`nM - S`$. The algorithm computes $`S`$ and $`M`$ in one pass and returns $`nM - S`$, so it is correct.

### 4.5 Complexity Analysis

Let the input size be $`n`$.

**Time complexity: Θ(n)**

One pass, with one addition and one comparison per element.

**Space complexity: O(1)**

Two variables.

**Value range**: $`S \le 100 \cdot 100 = 10^4`$ and $`nM \le 10^4`$, so overflow cannot occur.

## 5. Learning Notes

### Challenge 1: why 453 needs the reverse view and this problem does not

453's formula first has to prove that "add 1 to the other $`n-1`$" is equivalent to "subtract 1 from one", whereas here $`c_i = T - a_i`$ can be written down immediately.

**How it was resolved**

The key is how many numbers a single move touches. In 453 each move changes $`n-1`$ numbers; how much element $`i`$ changes depends on the global step count and on how often it was skipped, so the elements are entangled, and the common rise has to be removed by a translation first. Here each move changes one number; how much element $`i`$ changes is just its own count $`c_i`$, so the elements are independent from the start.

**Transferable insight:** first ask how many quantities one move touches. If it touches one, the total can be counted per element; if it touches several, look for a transformation that reduces it to touching one.

### Challenge 2: why the reference point flips from minimum to maximum

**How it was resolved**

The reference point is the endpoint forced by the direction of monotonicity. If values only decrease, the endpoint cannot exceed any number, so the highest possible endpoint is the minimum. If values only increase, the endpoint cannot be below any number, so the lowest possible endpoint is the maximum. In both cases the closest endpoint is "the extreme value that never has to move". Negating $`b_i = -a_i`$ flips the direction and turns $`\min`$ into $`-\max`$, consistently.

**Transferable insight:** when moves can only change values in one direction, the optimal endpoint is the first feasible value in that direction, namely the corresponding extreme.

## 6. Key Takeaways

Reusable methods from this problem:

1. **Lower Bound and Achievability**: prove that every plan needs at least some number of moves, then construct a plan that uses exactly that many. Here the bound comes from $`T \ge M`$, and the construction raises every number to the maximum.
2. **Negation Symmetry**: when the goal only asks whether all values are equal, $`b_i = -a_i`$ swaps adding with subtracting and maximum with minimum at the same time, so an existing result translates directly. It applies when the problem imposes no directional constraint on the values themselves (for example, no non-negativity requirement).
3. **Per-element Counting**: when each move changes one element, the total equals the sum of per-element changes, regardless of order.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. "Relation" is either "Direct" or "Related idea"; when no specific section fits, the "CLRS section" column says "No matching section" and only the concept name is kept.

| Concept | CLRS section | Relation | Where it appears here |
|---|---|---|---|
| Lower bound and achievability | No matching section | Related idea | $`k = nT - S \ge nM - S`$, achieved by $`T = M`$ |

## 7. Review Log

| Date | Independent completion | Change in understanding | To improve |
|---|---|---|---|
| 2026-11-03 | Wrote the code independently | Found the mirror relation with 453 by analogy; added a direct proof and a negation cross-check | To be recorded |

### Next Review Checklist

- [ ] Can I implement it without looking at the code?
- [ ] Can I state the input and output precisely?
- [ ] Can I write the pseudocode independently?
- [ ] Can I explain why the algorithm is correct?
- [ ] Can I derive the time and space complexity?
- [ ] Can I recognize the same idea in other problems?

## 8. Related Problems

- [453. Minimum Moves to Equal Array Elements](https://leetcode.com/problems/minimum-moves-to-equal-array-elements/) — each move adds 1 to the other $`n-1`$; viewed in reverse it is "subtract 1 from one", and negating that gives this problem.
- [462. Minimum Moves to Equal Array Elements II](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/) — each move adds or subtracts 1 on one number; with both directions allowed, the endpoint is no longer pinned to an extreme and the median is chosen instead.

---

**One-sentence summary**

When you can only add 1, the endpoint is at least the maximum, and raising each number to it takes $`n \cdot \max a - \sum a_i`$ moves; negated, this is 453's "subtract 1 from one".
