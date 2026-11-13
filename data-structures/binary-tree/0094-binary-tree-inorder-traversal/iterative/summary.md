# 94. Binary Tree Inorder Traversal — Explicit Stack

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../../docs/GITHUB_MATH.md). The binary tree itself (its definition, representation in Go, and common traversal patterns) is covered in the [binary tree overview](../../README.md).

This problem has two solutions, each in its own directory, sharing one test file:

- [**Recursive**](../recursive/summary.md): recurse directly on the definition "left subtree → root → right subtree". The problem statement, the formal definition of the inorder sequence $`\mathrm{in}(T)`$, and the test design are there.
- **Explicit stack** (this note): keep the ancestors "whose left subtree is unfinished and whose own value is not yet output" on a stack, answering the follow-up "can you do it iteratively?".

This note covers only how to produce the same sequence without recursion.

## 1. Problem Information

- **Problem**: [94. Binary Tree Inorder Traversal](https://leetcode.cn/problems/binary-tree-inorder-traversal/)
- **Official Difficulty**: Easy
- **Primary Category**: Binary Tree
- **Secondary Category**: Stack
- **Tags**: `Binary Tree`, `Stack`, `Inorder Traversal`, `Iteration`, `Loop Invariant`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](../solution_test.go) (in the parent directory, shared by both solutions)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Easy

**Four-Dimensional Rating — Explicit-Stack Inorder Traversal**

- Evaluation Scope: $`0 \le n \le 100`$ nodes, $`-100 \le \mathrm{Node.val} \le 100`$; time $`\Theta(n)`$; extra space excluding the output $`O(h)`$ ($`h`$ is the maximum depth, worst case $`h = n`$); Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2 | One must work out what the stack holds: ancestors already passed whose left subtree is unfinished, which is the suspended calls of the recursion made explicit |
| Knowledge Prerequisites | 2 | Needs a stack (the end of a Go slice as the top) and the pointer representation of binary trees |
| Correctness Reasoning Difficulty | 3 | With no recursion to apply induction to directly, one must construct a loop invariant describing the "remaining sequence to output", and separately argue that each node is pushed at most once so the loop terminates |
| Implementation Difficulty | 2 | Both halves of the loop condition `len(stack) > 0 \|\| node != nil` are needed; after a pop the code must turn right, not keep going left |

Ratings follow [RATING.md](../../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

The statement, examples, and key requirements are in [the recursive solution's §2](../recursive/summary.md#2-problem-understanding). This solution answers its follow-up: no recursion, only a loop and a stack.

## 4. Algorithm Specification

**Algorithm Name:** Iterative Inorder Traversal with an Explicit Stack

### 4.1 Input

Same as [the recursive solution's §4.1](../recursive/summary.md#41-input): a binary tree $`T`$ with $`0 \le n \le 100`$ nodes and root `root`. For a node $`x`$, let $`T_x`$ be the subtree rooted at $`x`$, with $`T_{\mathrm{NIL}} = \varnothing`$.

### 4.2 Output

The output is $`\mathrm{in}(T)`$, defined in [the recursive solution's §4.2](../recursive/summary.md#42-output):

```math
\mathrm{in}(\varnothing) = \langle\rangle, \qquad \mathrm{in}\bigl((r, L, R)\bigr) = \mathrm{in}(L) \cdot \langle r.\mathrm{val} \rangle \cdot \mathrm{in}(R)
```

### 4.3 Pseudocode

Write the stack $`S`$ from bottom to top as $`s_1, s_2, \ldots, s_k`$, with $`s_k`$ on top.

```text
INORDER-ITERATIVE(root)
    res ← ⟨⟩
    S ← empty stack
    node ← root
    while S is not empty or node ≠ NIL
        if node ≠ NIL
            PUSH(S, node)                   ▷ suspend node and go handle its left subtree
            node ← node.left
        else
            tmp ← POP(S)                    ▷ tmp's left subtree is done
            res ← res · ⟨tmp.val⟩
            node ← tmp.right                ▷ go handle tmp's right subtree
    return res
```

**State**

- `node`: the root of the next subtree to process; `NIL` means that subtree is empty.
- `S` (`stack` in the code): suspended ancestors. For each node on the stack, its left subtree is in progress or done, while its own value and its right subtree have not been output.
- `res`: the prefix of the inorder sequence already determined.

### 4.4 Correctness

**Proof Method: Loop Invariant (Initialization, Maintenance, Termination) + a Separate Termination Argument**

#### 1. Key Conditions for Correctness

1. **Recursive decomposition of the inorder sequence** (definition): for a nonempty subtree $`T_x`$, $`\mathrm{in}(T_x) = \mathrm{in}(T_{x.\mathrm{left}}) \cdot \langle x.\mathrm{val} \rangle \cdot \mathrm{in}(T_{x.\mathrm{right}})`$; for an empty subtree, $`\mathrm{in}(\varnothing) = \langle\rangle`$.
2. **The input is a tree** (given): every node except the root has a unique parent and is reached only through one child pointer of that parent. This guarantees each node is pushed at most once (Lemma 1).
3. **The stack is last-in first-out**: the most recently suspended ancestor resumes first, just as the innermost recursive call returns first.

#### 2. Correctness Proof

**Loop Invariant**

Each time the `while` condition is tested:

**Invariant I: Output decomposition**

```math
\mathrm{in}(T) = \mathrm{res} \cdot \mathrm{in}(T_{\mathrm{node}}) \cdot \langle s_k.\mathrm{val} \rangle \cdot \mathrm{in}(T_{s_k.\mathrm{right}}) \cdot \langle s_{k-1}.\mathrm{val} \rangle \cdot \mathrm{in}(T_{s_{k-1}.\mathrm{right}}) \cdots \langle s_1.\mathrm{val} \rangle \cdot \mathrm{in}(T_{s_1.\mathrm{right}})
```

That is: what has been output is `res`; next comes the whole of $`T_{\mathrm{node}}`$; then, from the top of the stack to the bottom, each suspended node's value followed by its right subtree.

**Invariant II: The stack is a leftward path**

For $`1 \le i < k`$, $`s_{i+1}`$ is in the left subtree of $`s_i`$; when `node` is nonempty, it is in the left subtree of $`s_k`$ (or in $`T`$ when the stack is empty). So the stack's nodes and `node` all lie on one downward path from the root.

**Initialization**

`res` is empty, the stack is empty, and `node = root`. Invariant I becomes $`\mathrm{in}(T) = \mathrm{in}(T_{\mathrm{root}})`$, which holds; Invariant II holds trivially for an empty stack.

**Maintenance**

Assume both invariants hold before an iteration.

- **Case 1: `node` is nonempty**. By key condition 1, expand $`\mathrm{in}(T_{\mathrm{node}})`$ into $`\mathrm{in}(T_{\mathrm{node.left}}) \cdot \langle \mathrm{node.val} \rangle \cdot \mathrm{in}(T_{\mathrm{node.right}})`$. The algorithm pushes `node` (as the new $`s_{k+1}`$) and sets `node ← node.left`. In the new state the right-hand side of Invariant I is exactly this expansion followed by the old remainder, so it is unchanged. Invariant II: the new top is the old `node`, which is in the old top's left subtree; the new `node` is its left child, in its left subtree.
- **Case 2: `node` is empty**. By the loop condition, the stack is nonempty. Since $`\mathrm{in}(T_{\mathrm{NIL}}) = \langle\rangle`$, Invariant I reads $`\mathrm{in}(T) = \mathrm{res} \cdot \langle s_k.\mathrm{val} \rangle \cdot \mathrm{in}(T_{s_k.\mathrm{right}}) \cdots`$. The algorithm pops $`s_k`$, appends its value to `res`, and sets `node ← s_k.right`. In the new state the right-hand side is $`(\mathrm{res} \cdot \langle s_k.\mathrm{val} \rangle) \cdot \mathrm{in}(T_{s_k.\mathrm{right}}) \cdot (\text{rest of the stack})`$, unchanged. Invariant II: the rest of the stack is unchanged; the new `node` is in $`s_k`$'s right subtree, and $`s_k`$ is in $`s_{k-1}`$'s left subtree, so the new `node` is also in $`s_{k-1}`$'s left subtree.

So both invariants still hold after each iteration.

**Termination**

*Lemma 1: each node is pushed at most once.* A node $`x`$ can become `node` in only two ways: $`x`$ is the root (at initialization), or $`x`$ is a child of some node $`p`$ and is assigned after $`p`$ is pushed (if $`x = p.\mathrm{left}`$) or after $`p`$ is popped (if $`x = p.\mathrm{right}`$). By key condition 2, $`x`$ has one parent and corresponds to only one of these moments. By induction on depth: the root becomes `node` at most once; if $`p`$ is pushed at most once, it is also popped at most once (once popped it is off the stack and never pushed again), so each of its children becomes `node` at most once. A nonempty `node` is pushed immediately, so each node is pushed at most once.

*Counting*: each iteration does exactly one push or one pop. There are at most $`n`$ pushes and no more pops than pushes, so the loop terminates after at most $`2n`$ iterations.

On exit the loop condition is false: the stack is empty and `node` is NIL. Invariant I becomes

```math
\mathrm{in}(T) = \mathrm{res} \cdot \mathrm{in}(\varnothing) = \mathrm{res}
```

so the returned `res` is exactly the inorder sequence. It also follows that $`\mathrm{in}(T)`$ has $`n`$ values and each pop appends exactly one, so there are exactly $`n`$ pops; every popped node was pushed, so there are also exactly $`n`$ pushes, and the loop runs exactly $`2n`$ times.

#### 3. Conclusion

Invariant I writes "what is left to output" exactly as the whole subtree of `node` followed by each stacked node's "value + right subtree". Each branch only expands one term by the inorder definition or outputs one value, leaving the whole sequence unchanged. After at most $`2n`$ iterations the stack is empty and `node` is NIL, the remainder is empty, and `res` is $`\mathrm{in}(T)`$.

### 4.5 Complexity Analysis

Let $`n = \lvert T \rvert`$ and let $`h`$ be the tree's maximum depth (counted in nodes; 0 for the empty tree).

**Time complexity: $`\Theta(n)`$**

By the counting in §4.4, the loop runs exactly $`2n`$ times, each a constant number of operations plus one `append`. `res` receives $`n`$ appends and the stack $`n`$ pushes; the total cost of slice growth is $`O(n)`$ by the amortized analysis of dynamic tables in CLRS §17.4. The total is $`\Theta(n)`$. Unlike the recursive solution, each value is written to `res` once and never copied again.

This is the aggregate analysis of CLRS §17.1: instead of bounding how many pushes can happen before one pop, count the pushes and pops over the whole algorithm.

**Space complexity: $`O(h)`$ excluding the output, worst case $`\Theta(n)`$**

By Invariant II, the stack's nodes all lie on one downward path from the root, so the stack holds at most $`h`$ nodes. More precisely, it holds only the ancestors in whose left subtree the current position lies, so the bound is not always reached: at the start, pushing along the leftmost path makes the stack at least as long as that path.

- **Worst case**: a chain hanging entirely to the left; the stack grows to $`n`$.
- **Best case**: a chain hanging entirely to the right; each node is popped right after being pushed because it has no left child, so the stack never holds more than 1 node. The `DeepChain` test is this case.

In the worst case the explicit stack and the recursion stack take the same $`\Theta(n)`$ space; on right-leaning trees the explicit stack is smaller, because the recursion still keeps a frame when it goes right (Go does not eliminate tail calls). Its main advantages are that it does not depend on the runtime's call-stack depth limit and that it can pause at any step (for example, stopping after the $`k`$-th node in problem 230), not a lower space bound. The output `res` takes another $`\Theta(n)`$.

### 4.6 Test Design

The cases are shared with [the recursive solution's §4.6](../recursive/summary.md#46-test-design), and each case runs against both solutions. For this solution, `Left-skewed chain` makes the stack grow to $`n`$ and `DeepChain` keeps it at most 1, covering both extremes of the stack size.

Mutation testing on this solution made at least one case fail for each of four mutations:

- Also appending the value to `res` on push (mixing in preorder): every case except `Empty tree` fails.
- Not appending the value after a pop: every case except `Empty tree` fails.
- Looping only while `len(stack) > 0`: the stack starts empty, so the loop never runs, and every case except `Empty tree` fails.
- Not turning to the right child after a pop (setting `node` to nil): `Example 1`, `Example 2`, the BST case, the negative-values case, and `DeepChain` fail.

## 5. Learning Notes

### Challenge 1: Why must the loop condition be stack nonempty **or** `node` nonempty?

**Understanding Process**

The two halves match the two parts of the remainder in Invariant I: a nonempty `node` means a whole subtree is still unprocessed, and a nonempty stack means suspended ancestors have not been output. As long as either part is nonempty, some value has not been output.

- Only "stack nonempty": the stack starts empty, so the loop never runs (the third mutation).
- Only "`node` nonempty": the loop stops at the first node with no left child, and the suspended ancestors on the stack are never output.

**Transferable Insight:** When writing a loop condition, first write the invariant for "what is left to do"; the loop condition is "the remainder is nonempty", and if the remainder has several parts, the condition is the "or" of those parts.

### Challenge 2: What exactly does the stack hold?

**Understanding Process**

Compare with the recursive solution: the call `INORDER(x)` is suspended while it calls `INORDER(x.left)`, and the suspended call remembers "when I come back, output `x.val`, then handle `x.right`". The explicit stack holds exactly these suspended `x`: popping `x` is the moment its left subtree is done and it is time to "come back". So after the pop the code outputs `x.val` and sets `node = x.right`, matching the last two steps of the recursion one to one.

Preorder outputs on push, and postorder must wait until the right subtree is also done, so the iterative postorder needs to remember whether the right subtree has been processed. Inorder sits between the two: it outputs on pop.

**Transferable Insight:** When turning a recursion into iteration, first ask "what does a suspended call need to remember", and store that on the stack; then ask "what happens on resumption", which becomes the work after a pop.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Recursion to Explicit Stack**: here, the stack holds ancestors whose left subtree is unfinished, and a pop corresponds to a recursive return. It applies when recursion may be too deep or when results must be produced step by step (an iterator), such as a BST iterator or a traversal that stops early. Limit: the worst-case space bound is unchanged, still $`O(h)`$.
2. **Loop Invariant over the Remaining Work**: here, Invariant I writes "already output + still to output" as the whole answer. It applies to proving any algorithm that simulates recursion or works through a worklist: as long as each step only expands one term of the remainder or outputs one item, the whole is preserved.
3. **Aggregate Analysis of Stack Operations**: here, each node is pushed and popped exactly once, for $`2n`$ iterations in total. It applies to monotonic stacks, bracket matching, and other algorithms where each element enters and leaves the stack at most once.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. "Relation" is either "Direct" or "Related idea"; when no specific section matches, the "CLRS Section" column says "No matching section" and only the concept name is kept.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Loop invariants | §2.1 Insertion sort | Direct | Invariant I (output decomposition) and Invariant II (the stack is a leftward path) are each proved by initialization, maintenance, and termination |
| Stacks | §10.1 Stacks and queues | Direct | `stack` uses the end of a slice as its top, with `append` as PUSH and truncation as POP; last-in first-out makes the innermost ancestor resume first |
| Nonrecursive traversal of a binary tree | §10.4 Representing rooted trees | Direct | Exercise 10.4-3 asks for an $`O(n)`$-time nonrecursive procedure that prints every key of a binary tree using a stack; this solution is one such procedure, in inorder |
| Nonrecursive inorder tree walk | §12.1 What is a binary search tree? | Direct | Exercise 12.1-3 asks for a nonrecursive inorder tree walk, with a hint that an easy solution uses a stack |
| Aggregate analysis | §17.1 Aggregate analysis | Direct | Each node is pushed and popped once, for $`2n`$ iterations in total; the same argument as the MULTIPOP example's "pops never exceed pushes" |

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

- [173. Binary Search Tree Iterator](https://leetcode.cn/problems/binary-search-tree-iterator/) — Splits this solution's loop into separate `next()` calls: the stack is the iterator's state, giving amortized $`O(1)`$ per call and $`O(h)`$ space.
- [230. Kth Smallest Element in a BST](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/) — Returns at the $`k`$-th pop, showing the explicit stack's ability to stop at any time.
- [144. Binary Tree Preorder Traversal](https://leetcode.cn/problems/binary-tree-preorder-traversal/) — The same stack structure, with output moved from pop to push.
- [145. Binary Tree Postorder Traversal](https://leetcode.cn/problems/binary-tree-postorder-traversal/) — Output must wait until the right subtree is done, which needs a record of the last node output; a natural next exercise after this solution.

---

**One-Sentence Summary**

Use a stack to suspend ancestors whose left subtree is unfinished: push all the way left, output on pop, then turn right; the invariant "output so far · current subtree · each stacked node's value and right subtree = the whole inorder sequence" proves correctness, and each node enters and leaves the stack once, so the time is $`\Theta(n)`$.
