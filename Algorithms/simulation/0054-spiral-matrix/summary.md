# 54. Spiral Matrix

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [54. Spiral Matrix](https://leetcode.cn/problems/spiral-matrix/)
- **Official Difficulty**: Medium
- **Primary Category**: Simulation
- **Secondary Category**: Matrix
- **Tags**: `Simulation`, `Matrix`, `Boundary Shrinking`, `Loop Invariant`
- **Code**: [solution.go](./solution.go) · [solution.py](./solution.py)
- **Tests**: [solution_test.go](./solution_test.go) · [solution_test.py](./solution_test.py)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Boundary Shrinking**

- Evaluation Scope: $`1 \le m, n \le 10`$; time $`\Theta(mn)`$, auxiliary space $`O(1)`$ excluding the output; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 1 | Walk the matrix ring by ring as the statement describes; the only step is to represent the unvisited region as the rectangle bounded by `t`, `b`, `l`, `r` |
| Knowledge Prerequisites | 1 | Only 2D slices, loops, and index arithmetic are needed |
| Correctness Reasoning Difficulty | 3 | Needs an explicit invariant "the unvisited region is exactly one rectangle" and a separate argument for four degenerate shapes: one row, one column, two rows, two columns |
| Implementation Difficulty | 3 | The bounds of the four loops, the order of the shrinks, and the four crossing checks must fit together; removing any one check makes some non-square matrix re-read cells |

Ratings follow [RATING.md](../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

Given an $`m \times n`$ matrix, start at the top-left corner and walk clockwise (right, down, left, up) ring by ring toward the center. Return the elements in the order they are visited.

**Example**

- Input: `matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]`
- Output: `[1,2,3,4,8,12,11,10,9,5,6,7]`

**Key Requirements**

1. Input constraints: $`m = \text{matrix.length}`$, $`n = \text{matrix[i].length}`$, $`1 \le m, n \le 10`$, $`-100 \le \text{matrix[i][j]} \le 100`$.
2. Every element appears exactly once, in clockwise spiral order.
3. Edge cases: the matrix may have a single row or a single column, or be non-square so that the innermost part is a single row or column.

## 3. Initial Approach

### First Idea: Shrink Four Boundaries

The boundary-shrinking solution was written from the start: `l`, `r`, `t`, `b` record the rectangle not yet walked; walk the top row, right column, bottom row, and left column in turn; after each side, move that boundary inward by one and stop if the boundaries have crossed. This code is correct for every shape.

### Difficulties Encountered

The code was correct, but the role of the four `break`s was not understood precisely:

1. The check `l > r` after the right column and `r -= 1` was thought to protect the right side, with a sense that failures had to do with the matrix being "asymmetric".
2. "Turning" in the direction-array approach was taken to be the same thing as "break" here.

### Thought Transformation

From:

> Each check protects the side that was just walked.

To:

> Each check asks whether the rectangle left after shrinking is empty. It protects the next side that walks **backward**: a forward loop naturally runs zero times once the boundaries cross, but a backward loop re-reads the row or column just walked.

Removing the checks one at a time gives the failing examples 1×3, 3×1, 2×3, and 4×2, all non-square, which confirms the "asymmetry" intuition. Going further, the four checks correspond to the four degenerate shapes of the remaining rectangle: one row, one column, two rows, and two columns (see Section 4.4). As for "turning", it corresponds to the shrink after each side and happens four times per ring; "break" means there is nowhere to go after turning, which in the direction-array approach is "all $`mn`$ steps taken".

## 4. Algorithm Specification

**Algorithm Name:** Boundary Shrinking

Rows and columns are 0-based throughout this section.

### 4.1 Input

An $`m \times n`$ integer matrix $`A`$ with $`1 \le m, n \le 10`$.

### 4.2 Output

For a nonempty rectangle $`R = [t, b] \times [l, r]`$ (row range $`[t, b]`$, column range $`[l, r]`$), define its **ring sequence** $`\mathrm{ring}(R)`$:

- If $`t = b`$ (one row): row $`t`$ from column $`l`$ to column $`r`$.
- Otherwise, if $`l = r`$ (one column): column $`l`$ from row $`t`$ to row $`b`$.
- Otherwise: top row $`(t, l), \ldots, (t, r)`$; right column $`(t+1, r), \ldots, (b, r)`$; bottom row $`(b, r-1), \ldots, (b, l)`$; left column $`(b-1, l), \ldots, (t+1, l)`$.

Removing the ring leaves $`R' = [t+1, b-1] \times [l+1, r-1]`$. The spiral order starts from the whole matrix, takes the ring, and repeats on what is left until it is empty:

```math
\mathrm{spiral}(R) =
\begin{cases}
\langle\rangle, & R \text{ is empty} \\
\mathrm{ring}(R) \cdot \mathrm{spiral}(R'), & \text{otherwise}
\end{cases}
```

Here $`\cdot`$ is sequence concatenation. The output must be the values of $`A`$ in the order $`\mathrm{spiral}([0, m-1] \times [0, n-1])`$.

### 4.3 Pseudocode

```text
SPIRAL-ORDER(A)
    t ← 0, b ← m - 1, l ← 0, r ← n - 1
    res ← empty sequence
    loop
        for i ← l to r                 ▷ top row, rightward
            append A[t][i] to res
        t ← t + 1
        if t > b then break            ▷ one row was left

        for i ← t to b                 ▷ right column, downward
            append A[i][r] to res
        r ← r - 1
        if l > r then break            ▷ one column was left

        for i ← r downto l             ▷ bottom row, leftward
            append A[b][i] to res
        b ← b - 1
        if b < t then break            ▷ two rows were left

        for i ← b downto t             ▷ left column, upward
            append A[i][l] to res
        l ← l + 1
        if l > r then break            ▷ two columns were left
    return res
```

**State Definitions**

- `t`, `b`: the top and bottom rows of the unvisited region (inclusive).
- `l`, `r`: the leftmost and rightmost columns of the unvisited region (inclusive).
- `res`: the elements already output in spiral order.

### 4.4 Correctness

**Proof Method: Outer Loop Invariant + Case Analysis on Degenerate Shapes**

#### 1. Key Conditions for Correctness

1. **Recursive structure of the spiral order (definition in 4.2)**: the spiral order is the ring followed by the spiral order of the remaining rectangle.
2. **The four checks match four degenerate shapes (lemma to prove)**: if an outer iteration starts with rectangle $`[t, b] \times [l, r]`$, then checks 1 to 4 fire exactly when it has one row, one column, exactly two rows, and exactly two columns, respectively.
3. **A forward loop over an empty range does nothing (language fact)**: `for i ← t to b` runs zero times when $`t > b`$.

#### 2. Correctness Proof

**Loop Invariant**

At the start of each outer iteration, let the boundaries be $`t, b, l, r`$ and $`R = [t, b] \times [l, r]`$:

**Invariant I: The remaining region is a nonempty rectangle**

$`t \le b`$ and $`l \le r`$, and the elements not yet output are exactly the elements of $`R`$.

**Invariant II: The output so far is correct**

`res` followed by $`\mathrm{spiral}(R)`$ is exactly the spiral order of the whole matrix.

**Initialization**

At the start, $`R = [0, m-1] \times [0, n-1]`$ is the whole matrix, nonempty because $`m, n \ge 1`$, and `res` is empty. Both invariants hold.

**Maintenance**

Assume the invariants hold at the start of an iteration. Split by the shape of the rectangle; below, $`t, b, l, r`$ are their values at the start of the iteration.

- **Case 1: $`t = b`$ (one row)**. The top row outputs row $`t`$, columns $`l..r`$, which is $`\mathrm{ring}(R)`$. Then `t` becomes $`t+1 > b`$, check 1 fires, and the loop ends. The remaining rectangle is empty, so by Invariant II, `res` is the full spiral order.
- **Case 2: $`t < b`$ and $`l = r`$ (one column)**. Check 1 does not fire. The top row outputs only $`(t, l)`$ and the right column outputs $`(t+1, r), \ldots, (b, r)`$; together they are column $`l`$ from $`t`$ to $`b`$, i.e. $`\mathrm{ring}(R)`$. Then `r` becomes $`l - 1`$, check 2 fires, the loop ends, and the remaining rectangle is empty.
- **Case 3: $`b = t + 1`$ and $`l < r`$ (exactly two rows)**. Neither of the first two checks fires. The top row, right column, and bottom row output the first three parts of $`\mathrm{ring}(R)`$; the left-column part $`(b-1, l), \ldots, (t+1, l)`$ is empty because $`b - 1 < t + 1`$. Then `b` becomes $`b - 1 = t`$ while `t` is already $`t + 1`$, so check 3 fires and the loop ends. The row range $`[t+1, b-1]`$ of the remaining rectangle is empty.
- **Case 4: $`b \ge t + 2`$ and $`l < r`$**. None of the first three checks fires, and the four loops output the four parts of $`\mathrm{ring}(R)`$. After the four shrinks the boundaries are $`t+1, b-1, l+1, r-1`$, which is $`R'`$.
  - If $`r = l + 1`$ (exactly two columns), then $`l + 1 > r - 1`$, check 4 fires, and the loop ends. The column range of $`R'`$ is empty.
  - Otherwise both ranges of $`R'`$ are nonempty, and the next iteration starts.

The four cases cover every nonempty rectangle. In each case, the iteration outputs exactly $`\mathrm{ring}(R)`$, and each element once (no cell appears twice in $`\mathrm{ring}(R)`$). By the recursive definition in 4.2, $`\mathrm{spiral}(R) = \mathrm{ring}(R) \cdot \mathrm{spiral}(R')`$, so Invariant II is maintained when the loop continues; when the loop ends, $`R'`$ is empty and `res` is already the full spiral order. The loop continues only in the second branch of Case 4, where $`R'`$ is nonempty, so Invariant I is maintained too.

**Termination**

Each time the loop continues, the rectangle loses 2 rows and 2 columns, and both counts start as finite positive integers, so the outer loop runs at most $`\lceil \min(m, n) / 2 \rceil`$ times. The last iteration always ends in one of the break branches.

#### 3. Conclusion

For every matrix with $`1 \le m, n \le 10`$, each iteration outputs exactly the ring of the current rectangle, and the algorithm stops once the remaining rectangle is empty. So it returns the spiral order with every element exactly once.

**Note: why none of the four checks can be removed.** The four cases of the lemma show that each check is the only exit for one degenerate shape. Without it, that shape keeps going, and the next backward loop re-reads the row or column just walked. Measured counterexamples: without check 1, 1×3 gives `[1,2,3,2,1]`; without check 2, 3×1 reads the middle element twice; without check 3, 2×3 re-reads in the next round; without check 4, 4×2 re-reads in the next round. These shapes are all in the tests.

### 4.5 Complexity Analysis

Let the matrix have $`m`$ rows and $`n`$ columns.

**Time Complexity: $`\Theta(mn)`$**

By the proof in 4.4, each element is appended to `res` exactly once, and each append takes constant time (Go preallocates the capacity; Python's `append` is amortized constant time). The checks and boundary updates run a constant number of times per outer iteration, and there are at most $`\lceil \min(m, n)/2 \rceil`$ outer iterations. The total time is $`\Theta(mn)`$. The output itself has $`mn`$ numbers, so this is also a lower bound.

**Space Complexity: $`\Theta(mn)`$ (including the output)**

`res` holds $`mn`$ integers. Excluding the output, only `t`, `b`, `l`, `r`, and the loop variable are used, so the auxiliary space is $`O(1)`$.

### 4.6 Implementation Notes: Go and Python

- **End of a backward loop**: Go writes `for i := r; i >= l; i--`, where the end `l` is inclusive; Python writes `range(r, l - 1, -1)`, where the end is exclusive, hence `l - 1`. Both run zero times when the start has passed the end, so the four checks sit in the same places in both languages. See [Go and Python Implementation Notes: Range Loops and Backward Iteration](../../../docs/LANGUAGE_NOTES.md#range-loops-and-backward-iteration).
- **Allocating the result**: Go preallocates with `make([]int, 0, m*n)`, so `append` never grows the slice; Python appends to a list directly.
- **Empty-matrix guard**: both versions keep a check for an empty matrix (Python `if not matrix`, Go `len(matrix) == 0`). The constraints guarantee $`m, n \ge 1`$, so this is only defensive.

## 5. Learning Notes

### Difficulty 1: What Does Each Check Actually Protect?

Intuitively, checking `l > r` after the right column and `r -= 1` seems to protect the right column itself.

**Understanding Process**

Removing the checks one at a time shows that the forward loops (top row, right column) naturally run zero times once the boundaries cross and never go wrong; the failures always come from the backward loops (bottom row, left column). For example, in a 3×1 matrix, after the top row and right column, column 0 has been fully output and `r` becomes $`-1`$. Without the check, the bottom row `range(-1, -1, -1)` is empty, but the left column `range(b, t-1, -1)` walks up column 0 from row 1 and reads the middle element again.

So the check decides whether the remaining rectangle is empty. It protects every side after it, especially the one that turns back.

**Transferable Thinking:** When a few variables describe the remaining region, check that the region is nonempty right after each shrink, instead of counting on later loops to happen to run zero times.

### Difficulty 2: Why Do Only Non-Square Matrices Fail?

In a square matrix every ring has four full sides, and the walk either ends exactly or leaves one cell (one row and one column at the same time), so nothing goes wrong. In a non-square matrix the shorter dimension runs out first, and the innermost part is one row, one column, two rows, or two columns, exactly the four degenerate cases in Section 4.4.

**Understanding Process**

For a rectangle $`[t, b] \times [l, r]`$, a full ring of four sides needs at least 3 rows and 3 columns; otherwise some sides overlap or are empty. The four checks rule out one row, one column, two rows, and two columns in turn. Only when at least 3×3 remains does a full ring get walked and the next iteration start.

**Transferable Thinking:** When designing a ring-by-ring algorithm, list every degenerate shape first (1 or 2 rows or columns) and confirm where each one exits.

### Difficulty 3: "Turning" in the Direction-Array Approach vs. "Break" Here

The direction-array approach walks with `dirs = [(0,1),(1,0),(0,-1),(-1,0)]` and a visited grid: turn right when the next cell is outside the matrix or already visited, and stop after $`mn`$ steps.

**Understanding Process**

The two approaches correspond as follows:

| Boundary shrinking | Direction array |
|---|---|
| The shrink after a side (`t += 1`, etc.) | Turning after hitting a wall or a visited cell |
| Boundaries cross after a shrink, so break | All $`mn`$ steps taken, so stop |
| Four integers `t`, `b`, `l`, `r` | An $`O(mn)`$ visited grid |

Turning happens four times per ring, and after most turns there is still a way forward, so "turning" and "stopping" are different things. Boundary shrinking has no step counter, so it can only detect the end through "the remaining rectangle is empty". One could replace `while True` with `while len(res) < m * n` to count steps, but the checks in the middle are still needed, otherwise the backward loops re-read cells.

**Transferable Thinking:** A visited grid can be replaced by a more compact structure when the visited region has a regular shape. Here the visited region is always "the rings outside a rectangle", so four boundaries describe it fully.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Boundary Shrinking**: here, four inclusive endpoints `t`, `b`, `l`, `r` describe the unvisited rectangle, and one endpoint shrinks after each side. It applies to problems that process a matrix layer by layer or ring by ring, such as spiral traversal or rotating a matrix ring by ring. The condition is that the unvisited region stays a regular rectangle; if the visiting order breaks that shape, use a visited grid instead.
2. **Degenerate Case Analysis**: here, the four checks match four degenerate shapes: one row, one column, two rows, and two columns. It applies to any algorithm with a fixed structure in general that collapses at small sizes, such as ring-by-ring processing, two pointers meeting, or interval splitting. The method is to list every degenerate shape and check where the algorithm exits and whether anything is processed twice.
3. **Loop Invariant**: here, the invariant is "the elements not yet output form exactly one nonempty rectangle, and the output so far followed by that rectangle's spiral order is the answer". It applies to simulations that build the result layer by layer. The key is to describe the remaining work as a subproblem of the same kind as the original.
4. **Direction Array Simulation**: walk step by step with a direction array and a visited grid, turning whenever the way ahead is blocked, and stop after the total number of steps. It needs no analysis of degenerate shapes, but needs $`O(mn)`$ extra space, or modifies the matrix in place as the mark. It suits grid simulations whose walking rules are complex or whose visited region is irregular, such as problem 885.

These methods fit simulations on grids and matrices. With boundary shrinking, the usual mistakes are in which endpoints are inclusive and where backward loops end; after writing the code, test it on a single row, a single column, and matrices whose sides differ by 1.

### CLRS Connections

Section numbers follow the third edition of CLRS. "Relation" is either "Direct" or "Related idea"; when no chapter matches, the "CLRS Section" column says "No matching section" and only the concept is listed.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Loop invariant | §2.1 Insertion sort | Direct | 4.4 proves correctness with the outer loop invariant "the elements not yet output form a nonempty rectangle, and the output so far followed by its spiral order is the answer", in three steps: initialization, maintenance, termination |
| Boundary shrinking | No matching section | Related idea | 4.3 replaces a visited grid with four inclusive endpoints, cutting the auxiliary space from $`O(mn)`$ to $`O(1)`$ |

## 7. Review Log

| Date | Independent Completion | Understanding Change | To Improve |
|---|---|---|---|
| 2026-11-23 | Wrote a correct Python boundary-shrinking solution | The four checks decide whether the remaining rectangle is empty and match one row, one column, two rows, and two columns; they protect the backward sides that come next, not the side just walked | Tell apart "turning" and "stopping" in the direction-array approach |

### Next Review Checklist

- [ ] Can I implement it independently without looking at the code?
- [ ] Can I describe the algorithm's input and output precisely?
- [ ] Can I write the pseudocode independently?
- [ ] Can I explain why the algorithm is correct?
- [ ] Can I derive the time and space complexity?
- [ ] Can I recognize the same algorithmic idea in other problems?

## 8. Related Problems

- [59. Spiral Matrix II](https://leetcode.cn/problems/spiral-matrix-ii/) — The reverse of this problem: fill $`1..n^2`$ into an $`n \times n`$ matrix in spiral order. Boundary shrinking carries over directly; because the matrix is square, the two-row and two-column cases never occur.
- [885. Spiral Matrix III](https://leetcode.cn/problems/spiral-matrix-iii/) — Spiral outward from a cell inside the matrix. The path leaves the matrix and the unvisited region is no longer a rectangle, so boundary shrinking does not apply; a direction array with step lengths 1, 1, 2, 2, 3, 3, … fits better.
- [2326. Spiral Matrix IV](https://leetcode.cn/problems/spiral-matrix-iv/) — Fill the values of a linked list into an $`m \times n`$ matrix in spiral order. The matrix may be non-square, so the same degenerate shapes as here must be handled.
- [48. Rotate Image](https://leetcode.cn/problems/rotate-image/) — Also processes the matrix ring by ring; `t`, `b`, `l`, `r` can describe the current ring, with a four-way swap inside each ring.

---

**One-Sentence Summary**

A spiral matrix is peeled one outer ring at a time; with four boundaries describing the remaining rectangle, check whether it is empty after every shrink, and the four checks match exactly the four degenerate shapes: one row, one column, two rows, and two columns.
