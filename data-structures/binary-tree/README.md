# Binary Tree

English | [简体中文](./README.zh-CN.md)

Use the [shared glossary](../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../docs/GITHUB_MATH.md).

This is the overview note for the binary tree: its definition, representation, traversals and their costs, and the usage patterns that recur across problems. Each application problem's note covers only how that problem uses a binary tree; knowledge about the binary tree itself lives here.

## 1. Definition

CLRS Appendix B.5.3 defines a binary tree recursively: a binary tree either **contains no nodes** (the empty tree) or is made of three disjoint sets of nodes: a **root**, a binary tree called the **left subtree**, and a binary tree called the **right subtree**.

Common terms:

| Term | Meaning |
|---|---|
| Leaf | A node whose left and right subtrees are both empty |
| Depth | The number of edges on the path from the root to the node; the root has depth 0 |
| Height | The number of edges on the longest path from the node down to a leaf; the tree's height is the root's height |
| Full binary tree | Every node is either a leaf or has exactly two children |
| Complete binary tree | In CLRS, a tree whose leaves all have the same depth and whose internal nodes all have two children; LeetCode's "complete binary tree" usually means "every level full except possibly the last, which is filled from the left", which differs from CLRS |

The recursive definition is the core of binary-tree problems: almost every algorithm is "handle the root, then recurse on both subtrees", and almost every correctness proof is **induction on subtree size**.

## 2. Representation and Cost

### 2.1 The linked representation in CLRS (§10.4)

Each node `x` stores a key `x.key` and three pointers: `x.p` (parent), `x.left`, and `x.right`. A missing child or parent is NIL; `T.root` points to the root, and `T.root = NIL` means the tree is empty.

### 2.2 In Go

LeetCode's Go problems all use this node type, with no parent pointer:

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```

- The empty tree is a `nil` pointer. A recursive function usually starts with `if root == nil { return ... }`, and that line is also the base case of the induction proof.
- With no parent pointer, "moving up" means either carrying information back to the parent through return values, or recording parents yourself with a stack or a hash table.
- Local tests need to rebuild a tree from LeetCode's level-order array (such as `[1,null,2,3]`). The repository's tests write it as `[]any`, with `nil` for a missing child, and attach children level by level with a queue.

### 2.3 Cost of traversal

Let the tree have $`n`$ nodes and height $`h`$.

- **Time**: each node is visited exactly once, so a traversal takes $`\Theta(n)`$ (CLRS Theorem 12.1 proves this for inorder; preorder and postorder are the same).
- **Space**: the recursion depth is $`h + 1`$, so the call stack uses $`O(h)`$. Since $`h`$ ranges from $`\lfloor \lg n \rfloor`$ (balanced) to $`n - 1`$ (a chain), the worst case is $`O(n)`$. Write space as $`O(h)`$ and state the worst case; do not write $`O(\lg n)`$ directly.
- **Slice concatenation in Go**: if each recursive call returns a fresh slice that the caller joins with `append(res, child...)`, every value is copied once at each of its ancestors, for $`\sum_v \mathrm{depth}(v) = O(nh)`$ copies in total, which is $`\Theta(n^2)`$ on a chain. Passing the result slice down as a parameter (or a closure variable) and appending once per visited node restores $`\Theta(n)`$.

## 3. Common Usage Patterns

Each pattern lists what the recursive function does and when it applies. New problems are added under the matching pattern.

### 3.1 The three depth-first orders

- **Preorder**: root → left → right. Fits problems that handle the parent first and pass information to the children, and serializing a tree.
- **Inorder**: left → root → right. On a **binary search tree** it outputs keys in sorted order (CLRS §12.1), so validating a BST and finding the $`k`$-th smallest element both use inorder.
- **Postorder**: left → right → root. Fits problems where a parent's answer depends on both children's answers, such as height, diameter, or freeing a tree.
- **Proof method**: strong induction on the number of nodes in the subtree: assume the recursive calls are correct on both subtrees, then show that combining them is correct for the whole tree.

### 3.2 Returning information bottom-up (postorder)

- **What the recursive function does**: returns a summary of the subtree (height, node count, whether it is balanced), and the parent computes its own value from its two children's values.
- **When it applies**: the answer can be merged from the answers of the left and right subtrees. When the final answer differs from the return value (the diameter is not the height), keep an outer variable and update the global best at each node.

### 3.3 Passing information top-down (preorder)

- **What the recursive function does**: passes information about the path from the root to the current node (path sum, depth, allowed value range) to the children as parameters.
- **When it applies**: the condition depends on ancestors, such as passing value bounds down to validate a BST, or path-sum problems.

### 3.4 Level-order traversal (breadth-first)

- **What it uses**: a queue, processed level by level; recording the queue length at the start of each round groups the nodes of one level together.
- **When it applies**: level-by-level output, the right side view, minimum depth, and other level-based questions. It is the tree special case of breadth-first search in CLRS §22.2.

### 3.5 Replacing recursion with an explicit stack

- **Method**: a stack holds the ancestors that are not yet finished. Iterative inorder pushes all the way left, visits on pop, then turns to the right subtree.
- **When it applies**: when deep recursion may overflow the stack, or when nodes must be produced one at a time like an iterator. CLRS Exercise 10.4-3 asks for a nonrecursive traversal with a stack, and Exercise 12.1-3 asks for a nonrecursive inorder walk.
- **Further**: Morris traversal temporarily threads empty right pointers to reach $`O(1)`$ extra space, at the cost of modifying the tree during the walk.

## 4. Application Problems

| Problem | Location | Pattern | Notes |
|---|---|---|---|
| 94. Binary Tree Inorder Traversal: [recursive](./0094-binary-tree-inorder-traversal/recursive/summary.md), [explicit stack](./0094-binary-tree-inorder-traversal/iterative/summary.md) | This directory | 3.1 Inorder; 3.5 Explicit stack | Recursive "left → root → right" (a fresh slice per call, for $`O(nh)`$ total copying, see Section 2.3); explicit stack that pushes all the way left, visits on pop, then turns right |
| [104. Maximum Depth of Binary Tree](./0104-maximum-depth-of-binary-tree/summary.md) | This directory | 3.2 Bottom-up return values | Postorder: the root's depth is the larger subtree depth plus 1 |

## 5. CLRS Connections

Section numbers follow the third edition.

| Topic | Section | Relation | Where it appears here |
|---|---|---|---|
| Recursive definition of binary trees | Appendix B.5.3 Binary and positional trees | Direct | The definition and terms in Section 1 |
| Linked representation of rooted trees | §10.4 Representing rooted trees | Direct | The `left`, `right`, `p` pointers in Section 2.1; Exercise 10.4-3 in Section 3.5 |
| Inorder walk and the BST property | §12.1 What is a binary search tree? | Direct | $`\Theta(n)`$ traversal in Section 2.3 (Theorem 12.1); sorted inorder in Section 3.1; Exercise 12.1-3 in Section 3.5 |
| Breadth-first search | §22.2 Breadth-first search | Related idea | Level-order traversal in Section 3.4 is BFS on a tree |
