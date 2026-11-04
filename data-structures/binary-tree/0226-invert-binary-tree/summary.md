# 226. Invert Binary Tree

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [226. Invert Binary Tree](https://leetcode.cn/problems/invert-binary-tree/)
- **Official Difficulty**: Easy
- **Primary Category**: Binary Tree
- **Secondary Category**: Depth-First Search
- **Tags**: `Binary Tree`, `Depth-First Search`, `Recursion`, `Postorder Traversal`, `In-Place Modification`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Easy

**Four-Dimensional Rating — Recursive Depth-First Search**

- Evaluation Scope: number of nodes $`0 \le n \le 100`$, $`-100 \le \mathrm{Node.val} \le 100`$; time $`\Theta(n)`$, auxiliary space $`\Theta(h)`$ (recursion stack, where $`h`$ is the tree height, $`h = n`$ in the worst case); Go implementation, in place.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 1 | The mirror has a recursive definition: a non-empty tree's mirror keeps the root and puts the mirror of the right subtree on the left and the mirror of the left subtree on the right, which can be written as recursion straight from the definition |
| Knowledge Prerequisites | 2 | Requires understanding binary trees as left/right pointers, recursion, and in-place modification that rewires pointers instead of copying nodes |
| Correctness Reasoning | 2 | One strong induction on the number of nodes, plus an argument that the two recursive calls act on disjoint subtrees, so neither their order nor the in-place writes interfere |
| Implementation Difficulty | 1 | One `nil` base case and one multiple assignment; the only trap is splitting the swap into two sequential assignments (see Challenge 1) |

Ratings follow the repository's [RATING.md](../../../docs/RATING.md). Each dimension is rated independently, with no combined total.

## 2. Problem Understanding

Given the root `root` of a binary tree, invert the tree and return its root. "Invert" means producing its **mirror**: the left and right children of every node are swapped, at every level.

**Example**

- Input: `root = [4,2,7,1,3,6,9]`
- Output: `[4,7,2,9,6,3,1]`

The subtrees 2 and 7 of root 4 are swapped; 1 and 3 under 2, and 6 and 9 under 7, are swapped too. Level by level, each level's order is exactly reversed.

- Input: `root = [2,1,3]`
- Output: `[2,3,1]`

- Input: `root = []`
- Output: `[]`

**Key Requirements**

1. Input constraints: the number of nodes is in $`[0, 100]`$ and $`-100 \le \mathrm{Node.val} \le 100`$. Node values do not affect the algorithm; they are only used to check results.
2. The swap must happen at **every** node, not just the root: swapping only the root's two children leaves the left/right order below the second level unchanged.
3. Edge cases: the empty tree stays empty; a node with one child has that child moved to the other side; the tree may be a chain.

## 4. Algorithm Specification

**Algorithm Name:** Recursive Depth-First Search, inverting both subtrees first and then swapping at the root

### 4.1 Input

A binary tree $`T`$, defined recursively: $`T`$ is either the empty tree $`\varnothing`$ or a triple $`(r, L, R)`$, where $`r`$ is the root node and $`L`$, $`R`$ are its left and right subtrees (also binary trees).

Let $`\lvert T \rvert`$ be the number of nodes in $`T`$, with:

- $`0 \le \lvert T \rvert \le 100`$;
- every node except the root has exactly one parent, and following child pointers never returns to a node already visited, so the input really is a tree.

In a non-empty tree $`(r, L, R)`$, $`L`$ and $`R`$ are disjoint, neither contains $`r`$, and $`\lvert T \rvert = 1 + \lvert L \rvert + \lvert R \rvert`$.

### 4.2 Output

Define the **mirror** $`M(T)`$ of $`T`$:

```math
M(T) =
\begin{cases}
\varnothing, & T = \varnothing \\
(r,\ M(R),\ M(L)), & T = (r, L, R)
\end{cases}
```

Output the root of $`M(T)`$ (`nil` for the empty tree). The implementation rewrites the `Left` and `Right` pointers of the input nodes directly, so after it returns the input tree itself has become $`M(T)`$, built from exactly the same nodes.

### 4.3 Pseudocode

```text
INVERT-TREE(x)                        ▷ x is the root of a subtree, possibly NIL
    if x == NIL
        return NIL
    newLeft  = INVERT-TREE(x.right)
    newRight = INVERT-TREE(x.left)
    x.left  = newLeft
    x.right = newRight
    return x
```

The algorithm returns `x` directly and has no result variable. The Go code does lines 4–7 with one multiple assignment, `root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)`: the Go specification evaluates all right-hand expressions in order before assigning any of them, so both recursive calls read the **original** `Right` and `Left`, which is equivalent to storing them in `newLeft` and `newRight` first as the pseudocode does.

**State Definitions**

- `x`: the root of the current subtree; `x == NIL` means an empty subtree.
- Return value: `x` itself (NIL for an empty tree); on return, the subtree rooted at `x` has been rewritten into the mirror of the original subtree.

### 4.4 Correctness

**Proof Method:** Strong induction on the number of nodes (structural induction)

#### I. Key Properties Required for Correctness

1. **The input is a tree** (given): $`L`$ and $`R`$ are disjoint and neither contains $`r`$, so $`\lvert L \rvert, \lvert R \rvert < \lvert T \rvert`$. This makes the recursive size strictly smaller, which guarantees termination and makes the induction possible.
2. **A call rewrites only pointers inside its own subtree** (to be proved, folded into the theorem's induction): `INVERT-TREE(x)` modifies only the `left` and `right` fields of nodes in the subtree rooted at `x`, and touches no node outside it.
3. **Both recursive calls read the original pointers** (guaranteed by the order of lines 4 and 5, or by Go's evaluation rule for multiple assignment): before the call on line 5, `x.left` has not been rewritten.

#### II. Correctness Proof

**Theorem.** For every binary tree $`T`$ satisfying the conditions in 4.1, `INVERT-TREE(T)` terminates; it modifies only child pointers of nodes in $`T`$; and on return, the structure reached from the return value is exactly $`M(T)`$, using the same set of nodes as $`T`$.

*Proof*: Strong induction on $`n = \lvert T \rvert`$.

**Base case ($`n = 0`$)**: $`T = \varnothing`$, so `x == NIL` holds and the call returns NIL, which equals $`M(\varnothing) = \varnothing`$; no pointer is modified.

**Inductive step ($`n \ge 1`$)**: Assume the claim holds for every tree with fewer than $`n`$ nodes. $`T = (r, L, R)`$ is non-empty, so the base branch is skipped.

- **Line 4**: $`\lvert R \rvert < n`$, so by the induction hypothesis `INVERT-TREE(x.right)` terminates, rewrites only pointers inside $`R`$, and returns the root of $`M(R)`$, which is stored in `newLeft`. Because $`L`$ and $`R`$ are disjoint and $`r \notin R`$, this step changes neither the structure of $`L`$ nor the values of `x.left` and `x.right`.
- **Line 5**: `x.left` still points to the original $`L`$ (Key Property 3), whose structure is unchanged. In the same way we get the root of $`M(L)`$ in `newRight`, and this step does not affect the $`M(R)`$ already obtained (it is made of nodes of $`R`$, disjoint from $`L`$).
- **Lines 6 and 7**: only $`r`$'s own two pointers are rewritten, so the structure rooted at $`r`$ is $`(r, M(R), M(L)) = M(T)`$.

Every pointer modified by the call belongs to $`\{r\} \cup L \cup R`$, and the nodes used are still those of $`T`$. ∎

The two recursive calls may be done in either order: they act on disjoint subtrees and never read each other's nodes. What must not happen is rewriting `x.left` before line 5 (see Challenge 1).

**Corollary (the mirror is an involution).** For every binary tree $`T`$, $`M(M(T)) = T`$. By induction on the number of nodes: $`M(M(\varnothing)) = \varnothing`$; for a non-empty tree, $`M(M((r, L, R))) = M((r, M(R), M(L))) = (r, M(M(L)), M(M(R))) = (r, L, R)`$. The test `TestInvertTreeTwice` checks exactly this.

**Corollary (inorder sequence is reversed).** Let $`\mathrm{In}(T)`$ be the inorder sequence of $`T`$. Then $`\mathrm{In}(M(T))`$ is the reverse of $`\mathrm{In}(T)`$. By induction: the empty tree is trivial; for a non-empty tree, $`\mathrm{In}(M(T)) = \mathrm{In}(M(R)) \cdot r \cdot \mathrm{In}(M(L))`$, which by the induction hypothesis equals $`\mathrm{rev}(\mathrm{In}(R)) \cdot r \cdot \mathrm{rev}(\mathrm{In}(L)) = \mathrm{rev}(\mathrm{In}(L) \cdot r \cdot \mathrm{In}(R))`$. So inverting a binary search tree gives a tree whose inorder walk lists the keys from **largest to smallest**.

#### III. Conclusion

The mirror of the empty tree is the empty tree; the mirror of a non-empty tree reduces by definition to the mirrors of two smaller subtrees, and the induction hypothesis guarantees that each step is correct and that they do not interfere, after which swapping two pointers at the root finishes the job. Therefore the algorithm turns every valid input into its mirror in place and returns the original root.

### 4.5 Complexity Analysis

Let $`n = \lvert T \rvert`$ and let $`h`$ be the tree's maximum depth (counting nodes, 0 for the empty tree).

**Time Complexity: $`\Theta(n)`$**

Count the calls. Each node is the argument of exactly one non-empty call: the root is called from outside, and every other node is called by its unique parent. The number of empty calls equals the number of empty child pointers: when $`n \ge 1`$ there are $`2n`$ child pointers, $`n - 1`$ of which point to nodes, leaving $`n + 1`$ empty; when $`n = 0`$ there is a single empty call. So the total number of calls is

```math
n + (n + 1) = 2n + 1
```

Apart from recursion, each call does a constant amount of work, so the total time is $`\Theta(n)`$. Any correct algorithm must rewrite the pointers of every node with at least one child, and it cannot know whether a node has children without reading it, so it must visit all $`n`$ nodes; $`\Omega(n)`$ is unavoidable.

**Space Complexity: $`\Theta(h)`$**

The algorithm allocates no new nodes; the result is the rewritten input tree and is not counted as extra space. All extra space is the recursion stack, whose maximum depth is $`h + 1`$ (reaching the deepest leaf and then making one more call on an empty child; 1 when $`n = 0`$), that is, $`\Theta(h)`$.

- **Worst case**: the tree is a chain, $`h = n`$, and the stack depth is $`\Theta(n)`$. The test `TestInvertTreeChain` uses left and right chains of 100 nodes.
- **Best case**: a complete binary tree, $`h = \lfloor \log_2 n \rfloor + 1`$, and the stack depth is $`\Theta(\log n)`$.

### 4.6 Test Design

The tests build trees with `buildTree` from LeetCode's level-order format, using the constant `null` for empty positions. It is $`-2^{31}`$, outside the node-value range $`[-100, 100]`$. `levelOrder` is its inverse: it serializes the result tree back into a level-order array for comparison. If it reaches the same node twice, the result is no longer a tree (both children point to the same subtree), and the test fails immediately. `TestInvertTree` also checks that the returned pointer is the input root, confirming the change is in place.

| Case | Expected | What It Checks |
|---|---|---|
| `example_1` | `[4,7,2,9,6,3,1]` | Official example 1, a perfect tree of three levels |
| `example_2` | `[2,3,1]` | Official example 2 |
| `example_3_empty` | `[]` | Official example 3, the empty tree, the base case |
| `single_node` | `[1]` | Only a root, unchanged by inversion |
| `left_child_only` | `[1,null,2]` | The only child moves from the left to the right |
| `right_child_only` | `[1,2]` | The only child moves from the right to the left |
| `deeper_levels_must_swap` | `[1,3,2,7,6,5,4]` | Swapping only at the root without recursing gives `[1,3,2,6,7,4,5]` |
| `asymmetric_with_gaps` | `[1,3,2,null,5,4,null,6]` | An asymmetric four-level tree with gaps |
| `negative_and_duplicate_values` | `[0,100,-100,100,null,null,-100]` | Boundary values and duplicates |
| `TestInvertTreeChain` | Mirrored chain | A left chain of the maximum 100 nodes becomes a right chain and vice versa; with a 1-second timeout |
| `TestInvertTreeTwice` | Original tree | Inverting twice returns the original tree (involution) |

Four mutations of the implementation were checked, and each makes some case fail:

- Recursing without swapping (`Left = invert(Left)`, `Right = invert(Right)`): every `TestInvertTree` case except `example_3_empty` and `single_node` fails, and `TestInvertTreeChain` fails.
- Swapping only at the root without recursing: `example_1`, `deeper_levels_must_swap`, `asymmetric_with_gaps`, `negative_and_duplicate_values`, and `TestInvertTreeChain` fail.
- Recursing into one side only (`Left, Right = Right, invert(Left)`): the same cases fail, and `TestInvertTreeTwice` fails as well.
- Splitting into two sequential assignments (see Challenge 1): every `TestInvertTree` case except `example_3_empty` and `single_node` fails (`left_child_only` loses its only child; the others report a node reached twice), `TestInvertTreeChain` fails by timeout, and `TestInvertTreeTwice` fails.

## 5. Learning Notes

### Challenge 1: Can the swap be split into two assignments?

A natural way to write it is

```go
root.Left = invertTree(root.Right)
root.Right = invertTree(root.Left) // wrong
```

**Understanding Process**

The `root.Left` read on the second line has already been rewritten by the first line and points to the **already inverted original right subtree**. The second line inverts the original right subtree once more (back to its original shape), and the original left subtree is dropped entirely, so `Left` and `Right` end up pointing to the same subtree. The result is no longer a tree, and every node of the original left subtree is lost.

This bug also makes the running time exponential: along a right chain, every level recurses twice into the same subtree, so a right chain of 100 levels needs about $`2^{100}`$ calls. The timeout on the chain test exists so that this bug shows up as a failure instead of a hang.

The fix is to make sure both recursive calls read the original pointers: store them in temporaries as the pseudocode does, use Go's multiple assignment (all right-hand sides are evaluated before any assignment), or swap first and then recurse (see Challenge 2).

**Transferable Insight:** When rewiring pointers in place, ask whether the field a statement reads was just rewritten by an earlier statement. The standard way to swap two fields is read first, then write; in Go, that is one multiple assignment.

### Challenge 2: Preorder or postorder?

**Understanding Process**

This implementation is postorder: both subtrees are inverted first, then the root swaps. Preorder is also correct: swap `x.left` and `x.right` first, then recursively invert the two new children. Both satisfy "every node swaps exactly once", and a node $`v`$'s final position depends only on whether each node on its path from the root (including itself) has swapped, not on the order of the swaps. Inorder does not work: it recurses into the left subtree, swaps, and then recurses into the "right subtree", but the right pointer now points to the original left subtree that was just inverted, which flips it back, while the original right subtree is never processed.

For the same reason, a queue-based level-order traversal that swaps the children of each dequeued node is a correct non-recursive version; its extra space is the width of the widest level rather than the height.

**Transferable Insight:** When the per-node operation is a local swap whose result does not depend on the order of operations, preorder, postorder, and level order all work, as long as each node is processed exactly once. Inorder puts a node's two subtree visits on opposite sides of the swap and is the easiest to get wrong.

### Challenge 3: The mirror seen through paths

**Understanding Process**

Write the path from the root to a node as a string of L and R. By induction on the definition of $`M`$, a node with path $`s`$ in the original tree has, in the mirror, the path obtained by changing every L in $`s`$ to R and every R to L. So the mirror changes no node's depth, every level appears in reverse order in level-order output (the third level `1,3,6,9` of example 1 becomes `9,6,3,1`), and the inorder sequence is reversed (the corollary in 4.4).

This is also why binary trees distinguish left from right children: CLRS Appendix B.5.3 stresses that in a binary tree, "only a left child" and "only a right child" are two different trees, and the mirror transformation is exactly a swap of that positional information.

**Transferable Insight:** To decide whether two trees are mirrors of each other or whether a tree is symmetric, translate the question into "nodes at corresponding paths are equal, with L and R swapped", then write a comparison that recurses on both trees at once.

## 6. Key Takeaways

This problem illustrates the following reusable concepts:

1. **Recursive Definition to Recursive Code**: here, write down the recursive definition of $`M(T)`$ first; the code matches it line by line. Applies to mirroring, copying, serializing, and other "structural transformation of the whole tree" problems. Condition: at the root, the transformation depends only on the transformations of the two subtrees.
2. **Structural Induction with Disjoint Subtrees**: here, the induction also proves that "a call modifies only pointers inside its own subtree", so the two recursive calls do not interfere. Applies to proving any recursive algorithm that modifies a tree in place. Precondition: the input really is a tree, with no nodes shared between subtrees.
3. **Read Before Write When Swapping**: here, both recursive calls must read the original `Left` and `Right`. Applies wherever fields or pointers are swapped in place (reversing a linked list, swapping array elements). Go's multiple assignment reads before it writes; splitting it into two statements may read a value that was just overwritten.
4. **Order-Independent Local Updates**: here, swapping once at each node suffices, so preorder, postorder, and level order are all correct, but inorder is not. Applies when deciding whether a per-node operation can be done iteratively or level by level. Limitation: when the operation needs the subtrees to be finished first (such as computing height), only postorder works.

### CLRS Connections

Section numbers follow the third edition of CLRS. "Relation" is either "Direct" or "Related idea"; when no specific section matches, the "CLRS Section" column says "No matching section" and only the concept name is kept.

| Topic | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Linked representation of binary trees | §10.4 Representing rooted trees | Direct | Inversion rewrites the left and right attributes of every node; an empty child `nil` corresponds to NIL, and the algorithm creates no nodes, only rearranges existing pointers |
| Binary trees distinguish left and right children | Appendix B.5 Trees (B.5.3 Binary and positional trees) | Direct | $`M(T)`$ in 4.2 is defined on the recursive definition of binary trees; the cases `left_child_only` and `right_child_only` reflect B.5.3's distinction that "only a left child" and "only a right child" are different binary trees |
| Recursive traversal of binary trees | §12.1 What is a binary search tree? | Related idea | INORDER-TREE-WALK makes one call per node, and Theorem 12.1 proves it takes $`\Theta(n)`$ time; 4.5 counts $`2n + 1`$ calls the same way. The corollary in 4.4 shows the inverted tree's inorder sequence is the reverse of the original |

## 7. Review Log

| Date | Independent Completion | Understanding Changes | To Improve |
|---|---|---|---|
| To be recorded | To be recorded | To be recorded | To be recorded |

### Next Review Checklist

- [ ] Can I implement it independently without looking at the code?
- [ ] Can I describe the algorithm's input and output precisely?
- [ ] Can I write the pseudocode independently?
- [ ] Can I explain why the algorithm is correct?
- [ ] Can I derive the time and space complexity?
- [ ] Can I recognize the same algorithmic idea in other problems?

## 8. Related Problems

- [101. Symmetric Tree](https://leetcode.cn/problems/symmetric-tree/) — A tree is symmetric exactly when its left subtree is the mirror of its right subtree; using the path view from Challenge 3, recurse on both at once, comparing "left of left" with "right of right" and "right of left" with "left of right".
- [100. Same Tree](https://leetcode.cn/problems/same-tree/) — The skeleton for recursing on two trees at once; swapping left and right on one side turns it into the mirror comparison of 101.
- [951. Flip Equivalent Binary Trees](https://leetcode.cn/problems/flip-equivalent-binary-trees/) — This problem swaps at every node; 951 allows swapping at any subset of nodes and asks whether one tree can be turned into the other, so the recursion must try both the "swapped" and "not swapped" correspondence.
- [104. Maximum Depth of Binary Tree](../0104-maximum-depth-of-binary-tree/summary.md) — Also postorder recursion, with the same $`2n + 1`$ calls and a stack as deep as the tree; the difference is that 104 combines subtree return values, while this problem rewires pointers in place.

---

**One-Sentence Summary**

The mirror satisfies "the mirror of the empty tree is empty, and the mirror of a non-empty tree keeps the root and swaps the two subtrees, each mirrored", so recursion follows the definition directly; disjoint subtrees keep the two recursive calls from interfering, and the swap must read before it writes, or it inverts the same subtree twice and loses the other one.
