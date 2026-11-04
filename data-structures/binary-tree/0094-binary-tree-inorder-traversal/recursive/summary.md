# 94. Binary Tree Inorder Traversal — Recursive

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../../docs/GITHUB_MATH.md). The binary tree itself (its definition, representation in Go, and common traversal patterns) is covered in the [binary tree overview](../../README.md); this note covers only how this problem traverses it.

This problem has two solutions, each in its own directory, sharing one test file:

- **Recursive** (this note): recurse directly on the definition "left subtree → root → right subtree". This note gives the problem statement, the formal definition of the inorder sequence, and the test design; the other note refers to them.
- [**Explicit stack**](../iterative/summary.md): simulate the call stack with an explicit stack, answering the problem's follow-up "can you do it iteratively?".

## 1. Problem Information

- **Problem**: [94. Binary Tree Inorder Traversal](https://leetcode.cn/problems/binary-tree-inorder-traversal/)
- **Official Difficulty**: Easy
- **Primary Category**: Binary Tree
- **Secondary Category**: Depth-First Search
- **Tags**: `Binary Tree`, `Depth-First Search`, `Recursion`, `Inorder Traversal`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](../solution_test.go) (in the parent directory, shared by both solutions)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Easy

**Four-Dimensional Rating — Recursive Inorder Traversal**

- Evaluation Scope: $`0 \le n \le 100`$ nodes, $`-100 \le \mathrm{Node.val} \le 100`$; time $`\Theta(n + \sum_u d(u))`$, worst case $`\Theta(n^2)`$ (notation in §4.5); extra space $`O(n)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 1 | Inorder traversal is defined recursively; "left → root → right" translates word for word into three lines of code |
| Knowledge Prerequisites | 2 | Needs a binary tree represented by left and right pointers, and recursion that combines subproblem results into the full result |
| Correctness Reasoning Difficulty | 2 | One strong induction on the number of nodes, where the induction hypothesis is exactly the two recursive calls' return values |
| Implementation Difficulty | 1 | One `nil` base case and three `append` calls; the only catch is writing `...` to spread a child slice |

Ratings follow [RATING.md](../../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

Given the root `root` of a binary tree, return the values of all nodes in **inorder**: for every node, first output the values of its whole left subtree, then its own value, then the values of its whole right subtree.

**Example**

- Input: `root = [1,null,2,3]`
- Output: `[1,3,2]`

Root 1 has no left child, so 1 comes first; then the right subtree rooted at 2 is processed: its left child 3 is output before 2.

- Input: `root = []`, output: `[]`
- Input: `root = [1]`, output: `[1]`

**Key Requirements**

1. Constraints: the number of nodes is in $`[0, 100]`$, and $`-100 \le \mathrm{Node.val} \le 100`$.
2. The output order depends only on the tree's shape, not on the node values; only when the tree is a binary search tree is the inorder sequence sorted.
3. Edge cases: an empty tree returns an empty list; the tree may degenerate into a chain.
4. Follow-up: the statement notes that the recursive solution is trivial and asks whether it can be done iteratively. See the [explicit-stack solution](../iterative/summary.md).

## 4. Algorithm Specification

**Algorithm Name:** Recursive Inorder Traversal

### 4.1 Input

A binary tree $`T`$, defined recursively: $`T`$ is either the empty tree $`\varnothing`$ or a triple $`(r, L, R)`$, where $`r`$ is the root and $`L`$, $`R`$ are the left and right subtrees (also binary trees).

Let $`\lvert T \rvert`$ be the number of nodes in $`T`$, where:

- $`0 \le \lvert T \rvert \le 100`$;
- the input really is a tree: every node except the root has exactly one parent, and following child pointers never returns to a node already visited. So $`L`$ and $`R`$ are disjoint, neither contains $`r`$, and $`\lvert T \rvert = 1 + \lvert L \rvert + \lvert R \rvert`$.

### 4.2 Output

Write $`\cdot`$ for sequence concatenation and $`\langle\rangle`$ for the empty sequence. The **inorder sequence** of $`T`$ is

```math
\mathrm{in}(T) =
\begin{cases}
\langle\rangle, & T = \varnothing \\
\mathrm{in}(L) \cdot \langle r.\mathrm{val} \rangle \cdot \mathrm{in}(R), & T = (r, L, R)
\end{cases}
```

The output is $`\mathrm{in}(T)`$. From the definition and $`\lvert T \rvert = 1 + \lvert L \rvert + \lvert R \rvert`$ in 4.1, induction on the number of nodes shows that $`\mathrm{in}(T)`$ contains each node's value exactly once and has length $`\lvert T \rvert`$.

### 4.3 Pseudocode

```text
INORDER(x)                              ▷ x is a subtree root, possibly NIL
    if x == NIL
        return ⟨⟩
    res ← INORDER(x.left)
    res ← res · ⟨x.val⟩
    res ← res · INORDER(x.right)
    return res
```

The code starts from an empty slice `res` and `append`s the left result, the root's value, and the right result in turn, which is equivalent to the pseudocode. When `x == NIL` the code returns the empty slice `[]int{}` rather than `nil`; both have length 0 and concatenate identically.

**State**

- `x`: the root of the current subtree; `x == NIL` means an empty subtree.
- `res`: the sequence the current call is building; on return it equals the inorder sequence of the subtree rooted at `x`.

### 4.4 Correctness

**Proof Method: Strong Induction on the Number of Nodes (Structural Induction)**

#### 1. Key Conditions for Correctness

1. **The input is a tree** (given): $`\lvert L \rvert, \lvert R \rvert < \lvert T \rvert`$, so the recursion strictly shrinks, which guarantees termination and makes the induction possible.
2. **The algorithm mirrors the definition**: the two branches of the pseudocode match the two branches of $`\mathrm{in}`$ one to one: the empty tree gives $`\langle\rangle`$, and a nonempty tree gives the concatenation "left · root · right".
3. **The recursive calls do not interfere**: each call reads only nodes in its own subtree, modifies no node, and returns a freshly built slice, so the order of the left and right calls does not affect the result.

#### 2. Correctness Proof

**Theorem.** For every binary tree $`T`$ satisfying 4.1, `INORDER(T)` terminates and returns $`\mathrm{in}(T)`$.

*Proof*: strong induction on $`n = \lvert T \rvert`$.

**Base ($`n = 0`$)**: $`T = \varnothing`$, so `x == NIL` holds and the call returns $`\langle\rangle = \mathrm{in}(\varnothing)`$.

**Inductive step ($`n \ge 1`$)**: assume the claim for every tree with fewer than $`n`$ nodes. $`T = (r, L, R)`$ is nonempty, so the base branch is skipped. By key condition 1, $`\lvert L \rvert, \lvert R \rvert \le n - 1`$, so by the induction hypothesis `INORDER(x.left)` terminates and returns $`\mathrm{in}(L)`$, and `INORDER(x.right)` terminates and returns $`\mathrm{in}(R)`$. After the three concatenations,

```math
\mathrm{res} = \mathrm{in}(L) \cdot \langle r.\mathrm{val} \rangle \cdot \mathrm{in}(R) = \mathrm{in}(T)
```

where the last equality is the definition of $`\mathrm{in}`$. ∎

#### 3. Conclusion

The empty tree returns the empty sequence; for a nonempty tree, the definition reduces its inorder sequence to those of two smaller subtrees, which the induction hypothesis guarantees are already correct, and the concatenation order matches the definition. So the algorithm returns the correct inorder sequence for every valid input.

### 4.5 Complexity Analysis

Let $`n = \lvert T \rvert`$. For a node $`u`$, let $`d(u)`$ be its depth (edges from the root to $`u`$; the root has depth 0), let $`T_u`$ be the subtree rooted at $`u`$, and let $`h`$ be the tree's maximum depth (counted in nodes, as in problem 104; 0 for the empty tree).

**Number of calls: $`2n + 1`$**

Each node is the argument of exactly one nonempty call (the root from outside, every other node from its unique parent); empty calls correspond to empty child pointers, of which there are $`n + 1`$ when $`n \ge 1`$, and there is 1 empty call when $`n = 0`$. That is $`2n + 1`$ calls in total.

**Time complexity: $`\Theta(n + \sum_u d(u))`$, between $`\Theta(n \log n)`$ and $`\Theta(n^2)`$**

Apart from constant overhead, the main cost of each call is **copying**: the call at node $`u`$ copies the slices returned by its left and right subtrees element by element into its own `res`, then appends its own value, writing $`\lvert T_u \rvert`$ elements in total. So the total number of elements written by all calls is

```math
\sum_{u} \lvert T_u \rvert = \sum_{v} \bigl(d(v) + 1\bigr) = n + \sum_{v} d(v)
```

The first equality swaps the order of summation: node $`v`$ belongs to $`T_u`$ exactly when $`u`$ is $`v`$ itself or an ancestor of $`v`$, and there are $`d(v) + 1`$ such $`u`$. In other words, each value is copied once at itself and once at each of its ancestors. The extra copying from `append` growth adds only a constant factor, by the amortized analysis of dynamic tables in CLRS §17.4.

- **Worst case**: the tree is a chain with depths $`0, 1, \ldots, n-1`$, summing to $`n(n-1)/2`$, so the time is $`\Theta(n^2)`$.
- **Best case**: a complete binary tree, whose depths sum to $`\Theta(n \log n)`$, so the time is $`\Theta(n \log n)`$.

Here $`n \le 100`$, so this causes no problem in practice; but the same code on a $`10^5`$-node chain would copy about $`5 \times 10^9`$ elements. Passing the result slice down as a parameter (or a closure variable) and appending once per visited node restores $`\Theta(n)`$. That is what INORDER-TREE-WALK in CLRS §12.1 does: it prints each node when visiting it instead of returning subsequences.

**Space complexity: $`O(n)`$ extra, $`\Theta(n)`$ output**

- **Recursion stack**: at its deepest it goes from the root down to a leaf plus one empty call, depth $`h + 1`$, so $`\Theta(h)`$, worst case $`\Theta(n)`$.
- **Intermediate slices alive at the same time**: at any moment only the calls on the current recursion path hold an unreturned `res`, and they hold values of disjoint subtrees, so their total length is at most $`n`$.

The extra space is $`O(n)`$ in total (not counting old slices awaiting garbage collection). The output itself takes $`\Theta(n)`$.

### 4.6 Test Design

The tests live in [solution_test.go](../solution_test.go) in the parent directory and are shared by both solutions. `buildTree` builds a tree from LeetCode's level-order form, written as `[]any` with `nil` for a missing child. Each solution package defines its own `TreeNode` (as on LeetCode), so the tests build one test-local tree and convert it to each package's type with `toRecursive` and `toIterative`.

| Case | Expected | What it checks |
|---|---|---|
| `Example 1` | `[1,3,2]` | Official example 1: the root has no left child, and its right child has a left child |
| `Example 2` | `[4,2,6,5,7,1,3,9,8]` | Example 2 from leetcode.com; both subtrees have several levels, covering node shapes such as "empty left child, nonempty right child" |
| `Empty tree` | `[]` | The empty tree, the base case |
| `Single node` | `[1]` | Only the root |
| `Left-skewed chain` | `[1,2,3]` | Everything hangs to the left; the stack solution's stack grows to $`n`$ |
| `Binary search tree gives sorted order` | `[1,...,7]` | A 7-node perfect BST, whose inorder sequence is sorted |
| `Negative values` | `[-1,0,1]` | Negative values, confirming the output does not depend on value signs |
| `DeepChain` (separate test function) | `[1,...,10000]` | A $`10^4`$-node chain hanging entirely to the right, checking deep recursion and long chains |

Mutation testing on both implementations made at least one case fail for every mutation:

- Putting the root's value before the left subtree (making it preorder): `Example 1`, `Example 2`, `Left-skewed chain`, the BST case, and the negative-values case fail.
- Dropping the recursion on the right subtree: `Example 1`, `Example 2`, the BST case, the negative-values case, and `DeepChain` fail.

Mutation testing for the explicit-stack solution is in [its §4.6](../iterative/summary.md).

## 5. Learning Notes

### Challenge 1: Where does returning and concatenating slices hide its cost?

The code seems to do just three `append`s per node, so it is easy to estimate it as $`O(n)`$.

**Understanding Process**

`append(res, child...)` is not constant time; it copies every element of `child`. A value returned from a leaf all the way up to the root is copied once at every level, so the total copying is the sum of "depth plus 1" over all values (§4.5). On a chain that is $`1 + 2 + \cdots + n = \Theta(n^2)`$.

**Transferable Insight:** When estimating the time of a recursion, count not just the calls but also the cost of combining subresults in each call. For a recursion that returns new containers and merges them, the total cost is the sum over elements of how many times each is merged; to avoid it, pass one shared result container down so each element is written once.

### Challenge 2: How does the inorder sequence relate to sorted order?

**Understanding Process**

The inorder sequence depends only on the tree's shape. For a general binary tree it can be in any order (for example `[1,null,2,3]` gives `[1,3,2]`). Only when the tree satisfies the binary-search-tree property (values in the left subtree are no greater than the root, values in the right subtree no smaller) is the inorder sequence in nondecreasing order, as discussed in CLRS §12.1 around Theorem 12.1. The BST test case checks exactly this.

**Transferable Insight:** "Inorder means sorted" is a property of binary search trees, not of inorder traversal. Problems such as validating a BST or finding its $`k`$-th smallest element rely on it; it does not apply to general binary trees.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Inorder Traversal**: here, output recursively following the definition "left → root → right". It applies to every BST problem that relies on order (validating a BST, the $`k`$-th smallest element, flattening a BST into a sorted list). Limit: on a general binary tree the inorder sequence carries no meaning about value order.
2. **Structural Induction**: here, strong induction on the number of nodes, with the induction hypothesis matching the two recursive calls' return values. It applies to proving almost any recursive algorithm on trees. It requires each subtree to be strictly smaller than the whole tree.
3. **Cost of Merging Returned Results**: here, each value is copied once at each of its ancestors, for $`\Theta(n + \sum d(u))`$ total time. It applies to analyzing any recursion in which subproblems return containers that the parent merges. The fix is to pass a shared result container so each element is written once.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. "Relation" is either "Direct" or "Related idea"; when no specific section matches, the "CLRS Section" column says "No matching section" and only the concept name is kept.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Linked representation of binary trees | §10.4 Representing rooted trees | Direct | `TreeNode`'s `Left` and `Right` pointers are the left and right attributes of CLRS nodes, and a `nil` child is NIL |
| Inorder tree walk | §12.1 What is a binary search tree? | Direct | This solution is INORDER-TREE-WALK, except that CLRS prints on visit while this solution returns and concatenates subsequences, which causes the copying cost in §4.5; the BST case matches the discussion around Theorem 12.1 that an inorder walk prints keys in sorted order |
| Amortized analysis of dynamic tables | §17.4 Dynamic tables | Related idea | In §4.5, the extra copying from `append` growth adds only a constant factor |
| Recursive definition of binary trees | Appendix B.5 Trees | Direct | 4.1 writes the input as either the empty tree or $`(r, L, R)`$, and both the definition of $`\mathrm{in}`$ and the proof follow this recursive structure |

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

- [144. Binary Tree Preorder Traversal](https://leetcode.cn/problems/binary-tree-preorder-traversal/) — Moves the root's output before the left subtree; the recursion and its proof are otherwise identical. It is also the "preorder" mutation in this problem's mutation testing.
- [145. Binary Tree Postorder Traversal](https://leetcode.cn/problems/binary-tree-postorder-traversal/) — Moves the root's output to the end; the iterative version is harder than inorder because the root can be output only after both subtrees are done.
- [98. Validate Binary Search Tree](https://leetcode.cn/problems/validate-binary-search-tree/) — Uses "a BST's inorder sequence is strictly increasing" and compares adjacent values during an inorder traversal.
- [230. Kth Smallest Element in a BST](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/) — Stops at the $`k`$-th node of the inorder traversal; the explicit-stack version can stop early.

---

**One-Sentence Summary**

Inorder traversal follows the recursive definition "left → root → right" directly and is proved by structural induction; but returning a new slice per call and concatenating copies each value once at every ancestor, which degrades to $`\Theta(n^2)`$ on a chain.
