# 3963. Create Grid With Exactly One Path

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [3963. Create Grid With Exactly One Path](https://leetcode.cn/problems/create-grid-with-exactly-one-path/)
- **Official Difficulty**: Easy
- **Primary Category**: Constructive Algorithms
- **Tags**: `Constructive Algorithms`, `Grid`, `Monotone Path`, `Validator Testing`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: 1310.81 ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Easy

**Four-Dimensional Rating — L-Shaped Corridor Construction**

- Evaluation Scope: $`1 \le m, n \le 25`$; time $`O(m + n)`$, space $`O(m + n)`$ including the output, since the rows share one string; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2 | One transformation: turn "exactly one path" into "leave a single corridor of width 1", then pick a corridor that is easy to write (first row + last column) |
| Knowledge Prerequisites | 1 | Only 2D grid indices, string concatenation, and loops |
| Correctness Reasoning Difficulty | 2 | Existence is checked directly; uniqueness only needs to look at which column the path's first downward move is in, a simple local argument |
| Implementation Difficulty | 1 | Only two kinds of rows: the all-`.` first row and the `n-1` copies of `#` followed by `.` for the rest, built once and reused; $`m = 1`$ and $`n = 1`$ need no special case |

Ratings follow the repository's [RATING.md](../../../docs/RATING.md). Each dimension is rated independently, and no total score is computed.

## 2. Problem Understanding

Given two integers $`m`$ and $`n`$, build an $`m \times n`$ grid in which every cell is `'.'` (free) or `'#'` (obstacle). Starting at the top-left cell $`(0, 0)`$, moving only right or down and stepping only on free cells, there must be **exactly one** path to the bottom-right cell $`(m-1, n-1)`$. The answer is not unique; any valid grid is accepted.

**Example**

- Input: `m = 2, n = 3`
- Output: `["..#","#.."]`, whose only path is $`(0,0) \to (0,1) \to (1,1) \to (1,2)`$. This solution returns a different valid answer, `["...","##."]`.

**Key Requirements**

1. $`1 \le m, n \le 25`$, so a $`1 \times 1`$ grid, a single row, and a single column are all valid inputs.
2. Paths move only right or down, and the start and end cells must be free.
3. "Exactly one" means two things: at least one path (existence) and at most one path (uniqueness).

## 4. Algorithm Specification

**Algorithm Name:** L-Shaped Corridor Construction

### 4.1 Input

Two integers $`m, n`$ with $`1 \le m, n \le 25`$.

### 4.2 Output

An array $`G`$ of $`m`$ strings, each of length $`n`$, with characters from $`\{\texttt{.}, \texttt{\#}\}`$.

Let the set of free cells be $`F = \{(i, j) : G[i][j] = \texttt{.}\}`$. A **monotone path** is a sequence of cells $`(0,0) = c_0, c_1, \ldots, c_k = (m-1, n-1)`$ where each step is $`c_{t+1} = c_t + (0, 1)`$ (right) or $`c_{t+1} = c_t + (1, 0)`$ (down). The output must satisfy:

- Exactly one monotone path has all of its cells in $`F`$.

### 4.3 Pseudocode

Indices are 0-based.

```text
CREATE-GRID(m, n)
    res ← array of m strings
    res[0] ← n copies of '.'

    row ← (n - 1) copies of '#' followed by one '.'
    for i ← 1 to m - 1
        res[i] ← row

    return res
```

The constructed free set is the union of the first row and the last column:

```math
F = \{(0, j) : 0 \le j < n\} \cup \{(i, n-1) : 0 \le i < m\}
```

**State Definitions**

- `res`: the result grid, `grid` in `solution.go`.
- `res[0]`: the first row, all free.
- `row`: the content of every row after the first; only the last column $`n-1`$ is free. It is built once.
- `res[i]` ($`i \ge 1`$): all equal to `row`.

### 4.4 Correctness

**Proof Method: direct construction for existence + the first downward move for uniqueness**

#### I. Key Properties Required for Correctness

The algorithm's correctness depends on:

1. **Fixed step counts** (follows from the move rules): every monotone path makes exactly $`n-1`$ right moves and $`m-1`$ down moves, because the column goes from 0 to $`n-1`$, the row goes from 0 to $`m-1`$, and each step increases exactly one of them by 1.
2. **Rows $`i \ge 1`$ are free only in the last column** (given by the construction): entering any row $`i \ge 1`$ can only land in column $`n-1`$.
3. **Start and end are free** (given by the construction): $`(0, 0)`$ is in the first row, and $`(m-1, n-1)`$ is in the last column. This also holds when $`m = 1`$ or $`n = 1`$.

Together, properties 1 and 2 say that once the path moves down, it is already in the last column and can only keep moving down.

#### II. Correctness Proof

Let $`P^*`$ be the path that goes along the first row to $`(0, n-1)`$ and then down the last column to $`(m-1, n-1)`$:

```math
P^* : (0,0) \to (0,1) \to \cdots \to (0,n-1) \to (1,n-1) \to \cdots \to (m-1,n-1)
```

**Existence**

The first $`n`$ cells of $`P^*`$ are in the first row and the rest are in the last column, so all of them are in $`F`$, and every step is right or down. So $`P^*`$ is a valid path. When $`m = 1`$ it has only the rightward part, and when $`n = 1`$ only the downward part; the conclusion is the same.

**Uniqueness**

Take any valid path $`P`$.

- **If $`m = 1`$**: the path has no down moves, so it moves right $`n-1`$ times in a row, which means $`P = P^*`$.
- **If $`m \ge 2`$**: the path moves down at least once. Let the first down move go from $`(0, j)`$ to $`(1, j)`$. Before it, the path only moves right along the first row, so it has made $`j`$ right moves. $`(1, j)`$ must be free, so by property 2, $`j = n-1`$; all $`n-1`$ right moves are used up. By property 1, the remaining $`m-1`$ steps are all down moves, and the path follows the last column to the end. So $`P = P^*`$.

In both cases every valid path equals $`P^*`$.

#### III. Conclusion

For all $`1 \le m, n \le 25`$, the constructed grid has a valid path, and $`P^*`$ is the only one, so the output meets the requirement. The proof does not use the upper bounds on $`m, n`$, so it holds for all positive integers $`m, n`$.

### 4.5 Complexity Analysis

Let the grid have $`m`$ rows and $`n`$ columns.

**Time Complexity: O(m + n)**

- `make([]string, m)` allocates and zeroes $`m`$ string headers, $`O(m)`$.
- The first row, `strings.Repeat(".", n)`, writes $`n`$ characters, $`O(n)`$.
- `row` is built by `strings.Repeat` producing $`n-1`$ copies of `#` and then concatenating a new string of length $`n`$, $`O(n)`$.
- The loop runs $`m-1`$ times; each iteration only assigns `row`'s string header (pointer and length) to `grid[i]` without copying characters, $`O(1)`$.

The total is $`O(m + n)`$. Note that the $`mn`$ characters are never written one by one: every row after the first shares the same character data. Calling `strings.Repeat` again for every row would make the time $`O(mn)`$.

**Space Complexity: O(m + n), including the output**

- `grid` itself is $`m`$ string headers, $`O(m)`$.
- There are only two blocks of character data: $`n`$ bytes for the first row and $`n`$ bytes for `row`; the temporary string of length $`n-1`$ made while building `row` can be reclaimed afterward.

So the total space is $`O(m + n)`$. This relies on Go strings being immutable: several elements sharing one string cannot affect each other. If a caller converts a row to `[]byte` and modifies it, the conversion copies the data, so that is still safe.

### 4.6 Test Design

The problem accepts any valid grid, so the tests do not compare against a fixed answer. Instead, a **validator** `checkGrid` checks three things:

1. There are $`m`$ rows, each of length $`n`$.
2. Every character is `.` or `#`.
3. The number of valid paths counted by `countPaths` is 1.

`countPaths` counts with dynamic programming: $`dp[i][j]`$ is the number of valid paths from the start to $`(i, j)`$; on a free cell $`dp[i][j] = dp[i-1][j] + dp[i][j-1]`$ (out-of-range terms count as 0), and on an obstacle it is 0.

| Test | Content | Purpose |
|---|---|---|
| `TestCreateGrid` | The three example sizes, plus $`1 \times 1`$, $`4 \times 1`$, and $`25 \times 25`$ | Covers a single cell, a single row, a single column, and the maximum size |
| `TestCreateGridAllSizes` | All 625 sizes with $`1 \le m, n \le 25`$ | The constraint range is small enough to enumerate fully |
| `TestCountPaths` | All-free $`2 \times 2`$ (2 paths), all-free $`3 \times 3`$ (6 paths), blocked end (0 paths), the answer from example 3 (1 path) | Confirms the validator itself is correct; otherwise a buggy validator could let tests pass wrongly |

The expected value 6 for the all-free $`3 \times 3`$ grid comes from counting: choose 2 of the 4 steps to be down moves, $`\binom{4}{2} = 6`$.

Running `go test ./Algorithms/constructive-algorithms/3963-create-grid-with-exactly-one-path/` locally passes all tests.

## 5. Learning Notes

### Challenge 1: "Exactly one" must be proved in two halves

Checking that the end is reachable is not enough. An all-free grid always reaches the end, but when $`m, n \ge 2`$ it has more than one path.

**Understanding Process**

Split "exactly one" into "at least one" and "at most one". The first only needs a concrete path. The second needs an argument that every valid path is that same path. The uniqueness argument looks at the first down move: below the first row only the last column is free, so the first down move can only happen in the last column.

**Transferable Insight:** To prove "exactly $`k`$", prove the lower bound and the upper bound separately. Uniqueness is often shown by taking an arbitrary solution and proving it must equal the constructed one.

### Challenge 2: How to test when the answer is not unique

The example gives `["..#","#.."]`, while this solution returns `["...","##."]`. Both are correct, so comparing strings directly would fail correct code.

**Understanding Process**

Tests should check whether the output satisfies the problem's condition, not whether it equals a particular answer. Here the condition "the path count is 1" can be computed directly with grid path counting DP. The validator itself also needs its own tests on small grids with known answers.

**Transferable Insight:** For construction problems and problems that accept any valid answer, write the test as a validator. The validator's rule should be simpler than the algorithm under test and easy to confirm.

### Challenge 3: Is it safe for the rows to share one string?

`solution.go` builds `row` once and assigns it to rows $`1`$ through $`m-1`$. One might worry: if one row is changed, do the others change too?

**Understanding Process**

Go's `string` is immutable. Assignment copies only the string header (a pointer to the character data and a length), not the data, and no operation can modify that data in place. So rows sharing the same data cannot affect each other, and $`m-1`$ copies are saved, bringing time and space from $`O(mn)`$ down to $`O(m + n)`$. If the grid were `[][]byte`, the same code would make every row share one backing array, and changing one row would change all of them; then each row would need its own allocation.

**Transferable Insight:** Immutable values can be shared freely; for mutable slices and arrays, first check whether a deep copy is needed.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Constructive Algorithms**: Here, instead of searching for a valid grid, we directly give a grid whose structure is simple and easy to prove. This fits problems that ask for any object satisfying a condition. The limitation is that the construction takes an observation, and its correctness must be proved separately, not just checked on examples.
2. **Uniqueness by First Deviation**: Here, comparing any path with $`P^*`$ at their first difference shows that step must land on an obstacle. This fits proofs that a path, sequence, or plan is unique. It requires showing that the deviating step always violates a constraint.
3. **Validator-Based Testing**: Here, `checkGrid` checks the size, the character set, and the path count, independent of any specific answer. This fits problems with non-unique answers. It requires that the condition can be checked efficiently and that the validator itself is tested.
4. **Grid Path Counting DP**: Used here only in the tests, $`dp[i][j] = dp[i-1][j] + dp[i][j-1]`$. This fits counting right/down paths in a grid. Watch for overflow when counts can be large; here the grid is at most $`25 \times 25`$ and the path counts in the tests are small.
5. **Sharing Immutable Values**: Here, the $`m-1`$ rows after the first share one string `row`, bringing time and space from $`O(mn)`$ down to $`O(m + n)`$. This fits problems whose output contains a lot of repeated content. It requires the shared value to be immutable, such as a Go `string`; slices like `[]byte` and `[]int` share their backing array and cannot be shared this way.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. Relation is either "Direct" or "Related idea"; if no section fits, write "No matching section" and keep only the concept name.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Counting with binomial coefficients | C.1 Counting | Related idea | The test uses $`\binom{4}{2} = 6`$ paths for the all-free $`3 \times 3`$ grid to check the validator `countPaths` |
| Constructive proof | No matching section | Related idea | Existence gives the path $`P^*`$ directly; uniqueness looks at the column of the first down move |

## 7. Review Log

| Date | Independent Completion | Understanding Change | To Improve |
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

- [62. Unique Paths](https://leetcode.cn/problems/unique-paths/) — Counting right/down paths in an obstacle-free grid, which is exactly `countPaths` on an all-free grid; the answer is $`\binom{m+n-2}{m-1}`$.
- [63. Unique Paths II](https://leetcode.cn/problems/unique-paths-ii/) — Path counting with obstacles, using the same transition as `countPaths`; this problem can be seen as its inverse: given the answer 1, construct the obstacle layout.

---

**One-Sentence Summary**

Turn "exactly one path" into "leave only one corridor": the first row plus the last column form an L, any path's first down move can only happen in the last column, so the path is unique; when answers are not unique, test with a validator instead of comparing a fixed output.
