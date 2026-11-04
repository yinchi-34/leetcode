# 226. Invert Binary Tree — Explicit-Stack DFS

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../../docs/GITHUB_MATH.md).

This problem has two solutions in two directories, with the same test cases (each directory has its own copy of `solution_test.go`):

- **Explicit-stack DFS** (this note): a slice serves as a stack; when a node is popped, its children are pushed and then its left and right pointers are swapped.
- [**Recursive DFS**](../recursive-dfs/summary.md): invert both subtrees first, then swap at the root.

## 1. Problem Information

- **Problem**: [226. Invert Binary Tree](https://leetcode.cn/problems/invert-binary-tree/)
- **Official Difficulty**: Easy
- **Primary Category**: Binary Tree
- **Secondary Category**: Depth-First Search, Stack
- **Tags**: `Binary Tree`, `Depth-First Search`, `Stack`, `Iteration`, `In-Place Modification`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Easy

**Four-Dimensional Rating — Explicit-Stack Depth-First Search**

- Evaluation Scope: number of nodes $`0 \le n \le 100`$, $`-100 \le \mathrm{Node.val} \le 100`$; time $`\Theta(n)`$, auxiliary space $`O(h)`$ (at most $`h + 1`$ nodes on the stack, where $`h`$ is the tree height; only $`O(1)`$ on a chain); Go implementation, in place.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2 | One must first see that "mirror = swap each node's left and right pointers once, in any order" before the recursive definition can be replaced by per-node processing in an arbitrary order |
| Knowledge Prerequisites | 2 | Requires implementing a stack with a slice (`append` to push, truncation to pop) and understanding how a stack replaces the call stack |
| Correctness Reasoning | 3 | Structural induction does not apply directly; it needs a loop invariant over "processed / on the stack / untouched" nodes plus a lemma that swaps are order-independent, and the stack-size bound needs its own argument |
| Implementation Difficulty | 2 | The loop body is only pop, push children, swap, but the empty tree and empty children must be kept off the stack |

Ratings follow the repository's [RATING.md](../../../../docs/RATING.md). Each dimension is rated independently, with no combined total.

## 2. Problem Understanding

Given the root `root` of a binary tree, invert the tree (produce its mirror: swap the left and right children of every node, at every level) and return its root.

**Example**

- Input: `root = [4,2,7,1,3,6,9]`, output: `[4,7,2,9,6,3,1]`
- Input: `root = [2,1,3]`, output: `[2,3,1]`
- Input: `root = []`, output: `[]`

**Key Requirements**

1. Input constraints: the number of nodes is in $`[0, 100]`$ and $`-100 \le \mathrm{Node.val} \le 100`$.
2. The swap must happen at every node, not just the root.
3. Edge cases: the empty tree; nodes with one child; trees that are chains.

## 4. Algorithm Specification

**Algorithm Name:** Explicit-Stack Depth-First Search, swapping on pop

### 4.1 Input

A binary tree $`T`$: either the empty tree $`\varnothing`$ or $`(r, L, R)`$, where $`L`$ and $`R`$ are disjoint left and right subtrees, neither containing $`r`$. $`0 \le \lvert T \rvert \le 100`$, and the input really is a tree (every non-root node has exactly one parent).

### 4.2 Output

The root of the mirror $`M(T)`$ of $`T`$, where

```math
M(T) =
\begin{cases}
\varnothing, & T = \varnothing \\
(r,\ M(R),\ M(L)), & T = (r, L, R)
\end{cases}
```

The implementation rewires the input nodes' pointers in place and returns the original root.

### 4.3 Pseudocode

The stack is a 0-based slice whose top is its last element.

```text
INVERT-TREE(root)
    if root == NIL
        return NIL
    S = a stack containing only root
    while S is not empty
        x = POP(S)
        if x.left ≠ NIL
            PUSH(S, x.left)
        if x.right ≠ NIL
            PUSH(S, x.right)
        swap x.left and x.right
    return root
```

The algorithm returns `root` directly and has no result variable. The code matches the pseudocode line by line: reading `stack[len(stack)-1]` and truncating is POP, `append` is PUSH, and `node.Left, node.Right = node.Right, node.Left` is the swap.

The line `stack[0] = root` in the code is redundant: the slice literal `[]*TreeNode{root}` already puts `root` at index 0, and this line writes the same value again. Removing it changes nothing (with that line deleted, every test still passes). The `root == nil` check at the top cannot be dropped, however: without it a `nil` would be pushed, and reading `node.Left` after popping it would panic with a nil pointer dereference.

**State Definitions**

- `stack`: nodes that have been discovered but not yet swapped.
- `node`: the node popped and processed in this iteration.

### 4.4 Correctness

**Proof Method:** Loop invariant plus an order-independence lemma

#### I. Key Properties Required for Correctness

1. **The input is a tree** (given): every non-root node has exactly one parent, so it is pushed exactly once, when its parent is processed.
2. **A swap changes only the current node's own two fields** (from the code): processing `x` changes no other node's pointers, and it does not change the **set** of `x`'s children, only their left/right positions.
3. **Lemma 1 (swap order does not matter)**: If every node of $`T`$ swaps its left and right pointers exactly once, the final structure is $`M(T)`$, whatever the order.

**Proof of Lemma 1.** By Key Property 2, the two fields of node $`v`$ change only when $`v`$ itself is processed, so in the end $`v`$'s left pointer is its original right child and its right pointer is its original left child, regardless of when other nodes are processed. Then show by induction on the number of nodes that "a tree in which every node has swapped" is $`M(T)`$: the empty tree is trivial; for a non-empty tree, the root $`r`$'s left pointer points to the root of the original $`R`$, every node of $`R`$ has also swapped, and by the induction hypothesis that subtree is $`M(R)`$; likewise the right side is $`M(L)`$. So the whole is $`(r, M(R), M(L)) = M(T)`$. ∎

#### II. Loop Invariant

Below, "child" and "descendant" refer to the **original** tree $`T`$; by Key Property 2, swaps do not change child sets, so these relations stay fixed during the run. Assume $`n \ge 1`$ (when $`n = 0`$, line 1 returns NIL $`= M(\varnothing)`$). Each time the `while` condition is checked, split the nodes into three classes:

- $`P`$: processed, that is, popped and swapped exactly once;
- $`S`$: on the stack, each exactly once, not yet swapped;
- $`U`$: untouched, not yet swapped and not on the stack.

**Invariant**: $`P`$, $`S`$, $`U`$ are disjoint and cover all nodes, and $`U`$ is exactly the union of the **proper descendants** of the nodes in $`S`$.

**Initialization**: $`S = \{r\}`$, $`P = \varnothing`$, and $`U`$ is every other node, that is, all proper descendants of $`r`$. It holds.

**Maintenance**: Pop $`x \in S`$. $`x`$ has not been swapped, so `x.left` and `x.right` are its original children. They are proper descendants of $`x`$, hence in $`U`$ by the invariant; and since the subtrees of different stack nodes are disjoint (every node has one parent), they belong to no other stack node's subtree. Pushing the non-empty children moves them from $`U`$ to $`S`$, once each; swapping $`x`$ then moves $`x`$ to $`P`$. The new $`U`$ is the old $`U`$ minus $`x`$'s children: the proper descendants of $`x`$ minus its children are exactly the union of the children's proper descendants, and the parts belonging to other stack nodes are unchanged. The invariant is maintained.

**Termination**: Each iteration moves exactly one node from $`S`$ into $`P`$, where it stays, so the loop stops after exactly $`n`$ iterations. Then $`S = \varnothing`$, so $`U = \varnothing`$ by the invariant, and $`P`$ is all nodes: every node has swapped exactly once. By Lemma 1 the structure is $`M(T)`$, and the returned `root` is its root. ∎

#### III. Conclusion

The stack ensures that every node is discovered once and processed once, and Lemma 1 ensures that the processing order does not matter, so the algorithm turns every valid input into its mirror in place. It also shows that "swap first, then push" is equally correct: the set of children pushed is the same, only their order differs (that change was tried in the tests and everything passes).

### 4.5 Complexity Analysis

Let $`n = \lvert T \rvert`$ and let $`h`$ be the tree's maximum depth (counting nodes).

**Time Complexity: $`\Theta(n)`$**

By 4.4, the loop runs exactly $`n`$ times. Each iteration does one pop, at most two pushes, and one swap. `append` grows capacity geometrically when it runs out, so total copying is $`O(n)`$, amortized $`O(1)`$ per push. Total time is $`\Theta(n)`$.

**Space Complexity: $`O(h)`$, $`\Theta(n)`$ in the worst case**

Let a node's depth be the number of edges from the root, with the root at 0. **Claim**: each time the `while` condition is checked, the depths of the stack's nodes are non-decreasing from bottom to top; every depth smaller than the top's appears at most once; the top's depth appears at most twice.

- Initially only the root is on the stack, so it holds.
- After popping the top $`x`$ (depth $`k`$), all remaining depths are $`\le k`$, with at most one node of depth $`k`$ left. The pushed children all have depth $`k + 1`$, strictly greater than every remaining depth, and there are at most two. The claim is maintained.

There are at most $`h`$ distinct depths, and only the top one can appear twice, so the stack holds at most $`h + 1`$ nodes. Truncating a slice does not free its backing array, but the capacity is at most a constant times the largest length ever reached, still $`O(h)`$.

- **Chain**: each iteration pops one node and pushes one, so the stack always holds a single node and extra space is $`O(1)`$. This beats the recursive solution, which needs $`\Theta(n)`$ levels of call stack on a chain.
- **Worst case**: a spine going down the right side where each spine node also has a left leaf. Each iteration pushes the left leaf first and the next spine node second, and the spine node is always popped, so the left leaves pile up; the stack reaches about $`h \approx n/2`$ nodes, that is, $`\Theta(n)`$.

### 4.6 Test Design

The two solution directories have identical test cases in `solution_test.go`; see the [recursive DFS test design](../recursive-dfs/summary.md#46-test-design). In short: `levelOrder` fails if a node is reached twice; `TestInvertTree` checks that the input root is returned; `TestInvertTreeChain` uses left and right chains of 100 nodes with a 1-second timeout; `TestInvertTreeTwice` checks that inverting twice gives back the original tree.

Changes tried on this implementation, and the results:

| Change | Result |
|---|---|
| Delete the swap | Every `TestInvertTree` case except `example_3_empty` and `single_node` fails, and `TestInvertTreeChain` fails |
| Do not push the right child | `example_1`, `deeper_levels_must_swap`, `asymmetric_with_gaps`, `negative_and_duplicate_values`, `TestInvertTreeChain`, and `TestInvertTreeTwice` fail |
| Swap only when both children are non-empty | `left_child_only`, `right_child_only`, `asymmetric_with_gaps`, `negative_and_duplicate_values`, and `TestInvertTreeChain` fail |
| Push the left child twice (a node processed twice, so its swaps cancel) | The same failures as "do not push the right child"; the left chain times out because repeated processing grows exponentially |
| Delete `stack[0] = root` | Everything passes (the line is redundant) |
| Swap first, then push | Everything passes (Lemma 1: order does not matter) |
| Take nodes from the front instead (level order) | Everything passes (same reason) |

## 5. Learning Notes

### Challenge 1: Why can the order differ from the recursion's?

The recursive version is "invert the subtrees first, then swap at the root", which looks order-dependent.

**Understanding Process**

Look at the mirror one node at a time: in the end, every node's left pointer is its original right child and its right pointer is its original left child (Lemma 1). Each swap touches only its own two fields, and swaps at different nodes do not interact, so as long as every node swaps exactly once, a stack, a queue, or recursion, preorder or postorder, all give the same result. Here the stack's only job is to visit every node exactly once.

Conversely, swapping the same node twice is the same as not swapping it. That is why "push the left child twice" fails in the tests, and it is the essence of the sequential-assignment bug in Challenge 1 of the recursive note: some subtree is processed twice.

**Transferable Insight:** When per-node operations do not interfere and each node needs exactly one, the traversal order can be chosen freely; correctness then reduces to "every node exactly once", which an invariant proves.

### Challenge 2: Read the children before or after the swap?

**Understanding Process**

This implementation reads `node.Left` and `node.Right` and pushes them, then swaps. Swapping first and then pushing the (now repositioned) two children is also correct: both children are pushed either way, the set is the same, and only the pop order changes. The one thing to watch is that both children are pushed; neither may be skipped because its position changed.

Another detail is the pop order: left is pushed first and right second, so the right subtree is processed first. Since each node swaps when it is popped, the processing order "root, original right subtree, original left subtree" is exactly the preorder of the **inverted** tree.

**Transferable Insight:** When simulating a traversal with an explicit stack, the push order decides the visit order, but as long as every node is pushed once, operations whose result does not depend on order are unaffected; only when a specific output order is required (such as a preorder sequence) must the right child be pushed before the left.

### Challenge 3: How is the explicit stack's space better than recursion's?

**Understanding Process**

The recursive version's call stack is as deep as the current path, $`\Theta(n)`$ on a chain. This implementation's stack holds nodes that are "discovered, waiting to be processed", and a node's children are pushed only after it is popped, so a chain keeps just one node on the stack. The claim in 4.5 bounds the stack by $`h + 1`$ nodes; the worst case is still $`\Theta(n)`$, but it occurs on trees that leave a pending sibling at every level, not on chains.

Here $`n \le 100`$, so neither version risks a stack overflow. The value of an explicit stack is that, when trees are large and possibly deep, it is not limited by the call stack's size.

**Transferable Insight:** To estimate an explicit stack's space, count how many pending nodes can be on it at once, rather than copying the recursion's $`O(h)`$; the two can have their worst cases on differently shaped trees.

## 6. Key Takeaways

This problem illustrates the following reusable concepts:

1. **Explicit Stack Replacing Recursion**: here, the stack holds "discovered, unprocessed" nodes; each iteration pops one and pushes its children. Applies to depth-first traversal of trees and graphs and to avoiding deep recursion. Condition: processing a node does not require its subtrees to be finished; postorder results (such as height) need extra state.
2. **Discovered / Processed Invariant**: here, the classes $`P`$, $`S`$, $`U`$ and "$`U`$ is the proper descendants of $`S`$" prove that every node is visited exactly once. Applies to proving that stack- or queue-based traversals (DFS, BFS) visit each node exactly once. Precondition: the graph is a tree, or a visited mark prevents repeated pushes.
3. **Order-Independent Local Updates**: here, Lemma 1 shows that one swap per node suffices, in any order. Applies when deciding whether a per-node operation can use any traversal order. Limitation: each node must be processed exactly once; operations like swapping cancel when done twice.
4. **Stack Size Bound by Depth Profile**: here, stack depths are non-decreasing and only the top depth can repeat, so at most $`h + 1`$ nodes are on the stack. Applies to analyzing the space of "pop one, push its children" DFS. Note that the worst case is not a chain but a tree that leaves a pending sibling at every level.

### CLRS Connections

Section numbers follow the third edition of CLRS. "Relation" is either "Direct" or "Related idea"; when no specific section matches, the "CLRS Section" column says "No matching section" and only the concept name is kept.

| Topic | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Stacks | §10.1 Stacks and queues | Direct | `stack` is a last-in, first-out stack: `append` is PUSH, reading the last element and truncating is POP, and `len(stack) > 0` is the negation of STACK-EMPTY |
| Linked representation and nonrecursive traversal of binary trees | §10.4 Representing rooted trees | Direct | Inversion rewrites every node's left and right attributes; Exercise 10.4-3 asks for an $`O(n)`$-time nonrecursive traversal of a binary tree using a stack, which is exactly this traversal with "print the key" replaced by "swap the pointers" |
| Binary trees distinguish left and right children | Appendix B.5 Trees (B.5.3 Binary and positional trees) | Direct | $`M(T)`$ in 4.2 is defined on the recursive definition of binary trees; `left_child_only` and `right_child_only` reflect B.5.3's distinction that "only a left child" and "only a right child" are different binary trees |

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

- [144. Binary Tree Preorder Traversal](https://leetcode.cn/problems/binary-tree-preorder-traversal/) — The same "pop one, push its children" explicit stack; to output preorder, the right child must be pushed before the left, whereas the swap here does not depend on order, so the push order is free.
- [101. Symmetric Tree](https://leetcode.cn/problems/symmetric-tree/) — A tree is symmetric exactly when its left subtree is the mirror of its right subtree; an iterative version pushes or enqueues pairs of nodes that should mirror each other and compares them.
- [3619. Count Islands With Total Value Divisible by K (explicit-stack DFS)](../../../../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/iterative-dfs/summary.md) — Explicit-stack DFS on a grid graph; a cell can be discovered by several neighbors, so a visited mark is needed, whereas here the input is a tree and each node is pushed only by its parent.
- [226. Invert Binary Tree (recursive DFS)](../recursive-dfs/summary.md) — The recursive version of the same problem, proved by structural induction on the mirror's definition; compare the two versions' space on a chain.

---

**One-Sentence Summary**

The mirror equals "swap each node's left and right pointers once", in any order, so using a stack to take out every node exactly once and swap it is enough; correctness rests on the "processed / on the stack / untouched" invariant, and the stack holds at most $`h + 1`$ nodes, only one on a chain.
