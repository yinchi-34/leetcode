# 104. Maximum Depth of Binary Tree

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [104. Maximum Depth of Binary Tree](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)
- **Official Difficulty**: Easy
- **Primary Category**: Binary Tree
- **Secondary Category**: Depth-First Search
- **Tags**: `Binary Tree`, `Depth-First Search`, `Recursion`, `Postorder Traversal`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Easy

**Four-Dimensional Rating — Recursive Depth-First Search**

- Evaluation Scope: number of nodes $`0 \le n \le 10^4`$, $`-100 \le \mathrm{Node.val} \le 100`$; time $`\Theta(n)`$, auxiliary space $`\Theta(h)`$ (recursion stack, where $`h`$ is the maximum depth, $`h = n`$ in the worst case); Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 1 | Maximum depth is recursive by nature: a non-empty tree's depth is the larger child depth plus 1, which can be written as recursion straight from the definition |
| Knowledge Prerequisites | 2 | Needs binary trees represented with left/right pointers, and recursion that combines subproblem results into the answer |
| Correctness Reasoning Difficulty | 2 | One strong induction on the number of nodes suffices; the only case split is showing that an empty child returning 0 does not affect the maximum |
| Implementation Difficulty | 1 | Only a `nil` base case and a one-line recursive expression, with no extra state |

Ratings follow [RATING.md](../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

Given the root `root` of a binary tree, return its maximum depth: the number of **nodes** on the longest path from the root down to the farthest leaf. A leaf is a node with no children.

**Example**

- Input: `root = [3,9,20,null,null,15,7]`
- Output: `3`

The root 3 has a leaf 9 on the left and 20 on the right, which has children 15 and 7. The longest root-to-leaf path is 3 → 20 → 15 (or 3 → 20 → 7), with 3 nodes.

- Input: `root = [1,null,2]`
- Output: `2`

**Key Requirements**

1. Input constraints: the number of nodes is in $`[0, 10^4]`$, and $`-100 \le \mathrm{Node.val} \le 100`$. Node values do not affect the answer.
2. Depth counts nodes, not edges: a tree with only a root has depth 1.
3. Boundaries: the empty tree (`root == nil`) has depth 0; the tree may degenerate into a chain whose depth equals the node count $`10^4`$.

## 4. Algorithm Specification

**Algorithm Name:** Recursive Depth-First Search, combining subtree results in postorder

### 4.1 Input

A binary tree $`T`$, defined recursively: $`T`$ is either the empty tree $`\varnothing`$ or a triple $`(r, L, R)`$, where $`r`$ is the root and $`L`$, $`R`$ are the left and right subtrees (themselves binary trees).

Let $`\lvert T \rvert`$ be the number of nodes of $`T`$. Subject to:

- $`0 \le \lvert T \rvert \le 10^4`$;
- every node except the root has exactly one parent, and following child pointers never returns to a node already visited, so the input really is a tree.

For a non-empty tree $`(r, L, R)`$, $`\lvert T \rvert = 1 + \lvert L \rvert + \lvert R \rvert`$. A node whose two children are both empty is a leaf.

### 4.2 Output

A **root-to-leaf path** is a sequence of nodes that starts at the root, moves to a child of the current node at each step, and ends at a leaf. Define

```math
D(T) =
\begin{cases}
0, & T = \varnothing \\
\max \{\, \text{number of nodes on } P : P \text{ is a root-to-leaf path of } T \,\}, & T \ne \varnothing
\end{cases}
```

The output is $`D(T)`$.

### 4.3 Pseudocode

```text
MAX-DEPTH(x)                          ▷ x is the root of a subtree, possibly NIL
    if x == NIL
        return 0
    return MAX(MAX-DEPTH(x.left), MAX-DEPTH(x.right)) + 1
```

The algorithm returns an expression directly and has no result variable. The helper `max` in the code corresponds to `MAX` in the pseudocode and returns the larger of two integers.

**State Definitions**

- `x`: the root of the current subtree; `x == NIL` means the subtree is empty.
- Return value: the value $`D`$ of the subtree rooted at `x`.

### 4.4 Correctness

**Proof Method:** Strong induction on the number of nodes (structural induction)

#### I. Key Properties Required for Correctness

1. **The input is a tree** (given): the subtrees $`L`$ and $`R`$ are disjoint and neither contains $`r`$, so $`\lvert L \rvert, \lvert R \rvert < \lvert T \rvert`$. This makes every recursive call strictly smaller, which guarantees termination and makes the induction possible.
2. **The empty tree has depth 0, and a non-empty tree has depth at least 1** (by definition): every non-empty tree has at least one root-to-leaf path, and that path contains at least the root.
3. **Lemma 1 (recursive structure of depth)**: for a non-empty tree $`T = (r, L, R)`$, $`D(T) = 1 + \max(D(L), D(R))`$.

**Proof of Lemma 1.** There are two cases.

- **Case 1: $`L = R = \varnothing`$.** $`r`$ is a leaf and the only root-to-leaf path is $`\langle r \rangle`$, so $`D(T) = 1 = 1 + \max(0, 0)`$.
- **Case 2: at least one child is non-empty.** Then $`r`$ is not a leaf. The second node of any root-to-leaf path is the root of some non-empty child subtree $`C`$, and removing $`r`$ from the path leaves a root-to-leaf path of $`C`$ (nodes of $`C`$ have the same children in $`C`$ as in $`T`$, so they have the same leaves). Conversely, prepending $`r`$ to any root-to-leaf path of $`C`$ gives a root-to-leaf path of $`T`$. Therefore

```math
D(T) = 1 + \max \{\, D(C) : C \in \{L, R\},\ C \ne \varnothing \,\}
```

If one child is empty, its $`D`$ value is 0, while the other, non-empty child has $`D`$ value at least 1 (Key Property 2). Adding 0 to the set being maximized does not change the result, so $`D(T) = 1 + \max(D(L), D(R))`$. ∎

#### II. Correctness Proof

**Theorem.** For every binary tree $`T`$ satisfying the conditions in 4.1, `MAX-DEPTH(T)` terminates and returns $`D(T)`$.

*Proof*: Strong induction on $`n = \lvert T \rvert`$.

**Base case ($`n = 0`$)**: $`T = \varnothing`$, so `x == NIL` holds and the call returns 0, which equals $`D(\varnothing)`$.

**Inductive step ($`n \ge 1`$)**: Assume the claim holds for every tree with fewer than $`n`$ nodes. $`T = (r, L, R)`$ is non-empty, so the base branch is skipped. By Key Property 1, $`\lvert L \rvert + \lvert R \rvert = n - 1`$ and both are less than $`n`$. By the induction hypothesis, both recursive calls terminate and return $`D(L)`$ and $`D(R)`$. `MAX` returns the larger one, and adding 1 gives $`1 + \max(D(L), D(R))`$, which equals $`D(T)`$ by Lemma 1. ∎

The order of the two recursive calls does not matter: each only reads its own subtree and modifies no node.

#### III. Conclusion

The empty tree returns 0 by definition. For a non-empty tree, Lemma 1 reduces the maximum depth to the maximum depths of two smaller subtrees, which the induction hypothesis guarantees are computed correctly. Therefore the algorithm returns the correct maximum depth for every valid input.

### 4.5 Complexity Analysis

Let $`n = \lvert T \rvert`$ and $`h = D(T)`$.

**Time Complexity: $`\Theta(n)`$**

Count the calls. Every node is the argument of exactly one non-empty call: the root is called from outside, and every other node is called by its unique parent. The number of empty calls equals the number of empty child pointers: for $`n \ge 1`$ there are $`2n`$ child pointers, $`n - 1`$ of which point to nodes, leaving $`n + 1`$ empty ones; for $`n = 0`$ there is a single empty call. So the total number of calls is

```math
n + (n + 1) = 2n + 1
```

Each call does a constant amount of work apart from its recursive calls, so the total time is $`\Theta(n)`$. Any correct algorithm must examine every node (skipping one could miss a deeper path hanging below it), so $`\Omega(n)`$ cannot be avoided.

**Space Complexity: $`\Theta(h)`$**

The algorithm allocates no heap memory; all auxiliary space is the recursion stack. The chain of active calls starts at the root and moves one step down a path at each level. At its deepest it reaches a leaf and makes one more call on an empty child, so the maximum stack depth is $`h + 1`$ (1 when $`n = 0`$), that is, $`\Theta(h)`$.

- **Worst case**: the tree is a chain, $`h = n`$, and the stack depth is $`\Theta(n)`$. The tests `left_chain_max_nodes` and `right_chain_max_nodes` are chains of $`10^4`$ nodes.
- **Best case**: a complete binary tree, $`h = \lfloor \log_2 n \rfloor + 1`$, and the stack depth is $`\Theta(\log n)`$.

Goroutine stacks in Go grow on demand, so $`10^4`$ levels of recursion run without problems.

### 4.6 Test Design

The tests build trees with `buildTree` from LeetCode's level-order format, using the constant `null` for empty positions. It is $`-2^{31}`$, outside the node-value range $`[-100, 100]`$, so it cannot be confused with a real node. `chain` builds a chain of $`10^4`$ nodes directly.

| Case | Expected | What It Checks |
|---|---|---|
| `example_1` | 3 | Official example 1: shallow left subtree, deeper right subtree |
| `example_2_right_only` | 2 | Official example 2: the root has only a right child |
| `empty_tree` | 0 | The empty tree, the base case |
| `single_node` | 1 | Only a root; depth counts nodes, so it is 1, not 0 |
| `left_deeper` | 4 | The left subtree is deeper and the right subtree is a single leaf |
| `one_side_empty` | 4 | Every node has only a right child; checks that the 0 from an empty subtree is never chosen as the result |
| `perfect_tree_depth_3` | 3 | A perfect binary tree of 7 nodes, all leaves at the same depth |
| `left_chain_max_nodes` | 10000 | The maximum node count, all hanging to the left; recursion depth reaches $`10^4`$ |
| `right_chain_max_nodes` | 10000 | Same, all hanging to the right |

Four mutations of the implementation were checked, and each makes some case fail:

- Recursing only into the left subtree: `example_1`, `example_2_right_only`, `one_side_empty`, and `right_chain_max_nodes` fail.
- Recursing only into the right subtree: `left_deeper` and `left_chain_max_nodes` fail.
- Changing `max` to return the smaller value: every case except `empty_tree`, `single_node`, and `perfect_tree_depth_3` fails.
- Returning 1 for the empty tree: every case fails.

## 5. Learning Notes

### Challenge 1: Why is no special case needed when one child is empty?

A node with one child is not a leaf, so any root-to-leaf path from it must go into the non-empty side. It seems as if the code should say "if the left child is empty, use the right; if the right is empty, use the left."

**Understanding Process**

Case 2 of Lemma 1 shows that the true formula maximizes over the non-empty children only. An empty child returns 0, while a non-empty child returns at least 1, so putting the empty child into `max` as well never changes the result. In other words, 0 acts as a neutral value for maximum here: for any non-negative $`a`$, $`\max(0, a) = a`$.

Compare [111. Minimum Depth of Binary Tree](https://leetcode.cn/problems/minimum-depth-of-binary-tree/): copying the pattern as `min(minDepth(L), minDepth(R)) + 1` lets `min` pick the 0 from an empty child, treating "go into the empty subtree" as a path, which gives the wrong answer. That problem must handle "one child is empty" separately.

**Transferable Insight:** When a recursion combines subproblem results, check that the base case's return value is neutral for the combining operation. For maximum, 0 is neutral (when values are non-negative); for minimum, it is not. When it is not neutral, exclude empty subtrees explicitly.

### Challenge 2: Does depth count nodes or edges?

**Understanding Process**

This problem's "maximum depth" counts the nodes on a path, so a tree with only a root has depth 1. CLRS Appendix B.5 defines node depth and tree height by counting edges, so a tree with only a root has height 0. For a non-empty tree the two differ by exactly 1: this problem's answer is the CLRS height plus 1.

In code, the difference shows up in the base case and where 1 is added: here the empty tree returns 0 and each level adds 1, so a leaf gets 1. If depth counted edges, a common choice is to return $`-1`$ for the empty tree so that a leaf gets 0.

**Transferable Insight:** When a problem says "depth", "height", or "levels", first confirm whether it counts nodes or edges and what value the empty tree takes. The base case's return value must match that definition.

### Challenge 3: Recursion depth and tree height

**Understanding Process**

The recursion stack is as deep as the tree is tall, not $`\log n`$. The problem does not promise a balanced tree, and $`10^4`$ nodes can form a chain with recursion depth $`10^4`$. Go grows its stacks on demand, so this is fine here; in an environment with a small default stack, switch to an explicit-stack DFS, or a queue-based BFS that counts one level each time a whole level is processed.

Also, the code defines its own `max`. Since Go 1.21, `max` is a built-in function, and the repository's `go.mod` requires Go 1.22, so the helper could be dropped in favor of the built-in. Keeping it is also legal: a package-level definition shadows the built-in `max` within the package, with the same behavior.

**Transferable Insight:** When estimating the space of a recursive solution, use the actual depth of the recursion, which on a binary tree is its height. Write $`O(\log n)`$ only when balance is guaranteed.

## 6. Key Takeaways

This problem illustrates the following reusable concepts:

1. **Postorder Aggregation**: here, compute the depths of the left and right subtrees first, then combine them at the root as $`1 + \max`$. Applies to height, node count, subtree sums, balance checks, diameter, and other problems where the whole tree's result is computed from the subtrees' results. Condition: the root's answer depends only on each subtree's answer (plus a small amount of extra information).
2. **Structural Induction**: here, strong induction on the number of nodes, with the induction hypothesis matching the return values of the two recursive calls. Applies to proving almost every recursive algorithm on trees. Precondition: subtrees are strictly smaller than the tree, which holds because the input is a tree.
3. **Neutral Base Case**: here, the empty tree returns 0, which is neutral for maximum, so no special case is needed when one child is empty. Applies to designing base cases of recursions. Limitation: for operations such as minimum, 0 is no longer neutral and empty subtrees must be excluded (see problem 111).
4. **Recursion Depth Equals Tree Height**: here, the auxiliary space is $`\Theta(h)`$, which is $`\Theta(n)`$ in the worst case. Applies to estimating the stack space of any recursion on a tree. Write $`O(\log n)`$ only when the height is guaranteed (as in balanced trees).

### CLRS Connections

Section numbers follow CLRS, 3rd edition. Relation is either "Direct" or "Related idea"; if no section fits, write "No matching section" and keep only the concept name.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Linked representation of binary trees | §10.4 Representing rooted trees | Direct | The `Left` and `Right` pointers of `TreeNode` are the left and right attributes of CLRS nodes, and an empty child `nil` corresponds to NIL |
| Node depth and tree height | Appendix B.5 Trees | Direct | Challenge 2: CLRS defines depth and height by edges, while this problem counts nodes, so on a non-empty tree the answer is the CLRS height plus 1 |
| Recursive traversal of binary trees | §12.1 What is a binary search tree? | Related idea | INORDER-TREE-WALK makes one call per node, and Theorem 12.1 proves it takes $`\Theta(n)`$ time; 4.5 counts calls the same way ("one non-empty call per node"), except that this problem combines subtree results in postorder |

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

- [111. Minimum Depth of Binary Tree](https://leetcode.cn/problems/minimum-depth-of-binary-tree/) — The same recursive structure, but once `max` becomes `min`, the 0 from an empty subtree is no longer neutral, so "one child is empty" must be handled separately; a direct test of Challenge 1.
- [110. Balanced Binary Tree](https://leetcode.cn/problems/balanced-binary-tree/) — Builds on this problem's postorder height recursion and additionally checks that the two subtree heights differ by at most 1.
- [543. Diameter of Binary Tree](https://leetcode.cn/problems/diameter-of-binary-tree/) — Still computes subtree depths in postorder, while updating a global maximum with "left depth + right depth" at each node; the diameter counts edges.
- [559. Maximum Depth of N-ary Tree](https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/) — The same recurrence generalized to any number of children: depth is the maximum child depth plus 1.

---

**One-Sentence Summary**

Maximum depth satisfies "0 for the empty tree, and the larger subtree depth plus 1 for a non-empty tree", so recursion straight from the definition works; the 0 returned by the empty tree is neutral for maximum, so an empty child needs no special case, and the recursion stack is as deep as the tree is tall.
