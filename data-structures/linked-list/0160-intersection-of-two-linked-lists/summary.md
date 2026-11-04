# 160. Intersection of Two Linked Lists

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [160. Intersection of Two Linked Lists](https://leetcode.cn/problems/intersection-of-two-linked-lists/)
- **Official Difficulty**: Easy
- **Primary Category**: Linked List
- **Secondary Category**: Two Pointers
- **Tags**: `Linked List`, `Two Pointers`, `Reference Equality`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Easy

**Four-Dimensional Rating — Two Pointers with Head Switching**

- Evaluation Scope: list lengths $`1 \le m, n \le 3 \times 10^4`$, $`1 \le \mathrm{Node.val} \le 10^5`$, no cycles; time $`O(m + n)`$, auxiliary space $`\Theta(1)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 3 | The key observation is that "A's private part + shared part + B's private part" has the same length as "B's private part + shared part + A's private part", so each pointer walks the other list too to cancel the length difference |
| Knowledge Prerequisites | 2 | Needs singly linked lists with `Next` pointers and a `nil` tail, and identifying "the same node" by address rather than by value |
| Correctness Reasoning Difficulty | 3 | First prove the intersection is a shared suffix, then split on "are the private prefixes equally long" and "do the lists intersect" to show the first meeting point is exactly the intersection start or `nil` |
| Implementation Difficulty | 2 | The code is short, but the boundary "stepping onto `nil` counts as a step before switching to the other head" is easy to get wrong, and getting it wrong loops forever when the lists do not intersect |

Ratings follow [RATING.md](../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

Given the heads `headA` and `headB` of two singly linked lists, return the node at which the two lists intersect, or `nil` if they do not. "Intersect" means that from some node on, the two lists share the very same nodes (the same memory), not merely equal values. The lists must keep their original structure after the function returns.

The judge describes the input with `intersectVal`, `listA`, `listB`, `skipA`, `skipB`: after its first `skipB` nodes, `listB` is attached to node `skipA` (0-based) of `listA`. The function sees only `headA` and `headB`.

**Examples**

- Input: `intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3`
- Output: the node with value `8`

The first `1` in A and the `1` in B are two different nodes: equal values, but not the intersection. The intersection starts at the node with value `8`.

- Input: `intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2`
- Output: `null`

**Key Requirements**

1. Input constraints: $`1 \le m, n \le 3 \times 10^4`$, $`1 \le \mathrm{Node.val} \le 10^5`$; the linked structure has no cycles.
2. Intersection is decided by node identity (pointer equality), not by node value.
3. The lists must not be modified. Follow-up: $`O(m + n)`$ time and $`O(1)`$ extra space.
4. Edge cases: the lists may not intersect at all; they may intersect at the head (both are the same list), or one may be a suffix of the other.

## 4. Algorithm Specification

**Algorithm Name:** Two Pointers with Head Switching

### 4.1 Input

Two acyclic singly linked lists $`A`$ and $`B`$ with heads `headA` and `headB` and $`m`$ and $`n`$ nodes. Write them as node sequences (0-based):

```math
A = \langle A_0, A_1, \ldots, A_{m-1} \rangle, \qquad B = \langle B_0, B_1, \ldots, B_{n-1} \rangle
```

where $`A_{i+1} = A_i.\mathit{next}`$ and $`A_{m-1}.\mathit{next} = \mathrm{NIL}`$, and likewise for $`B`$. They satisfy:

- $`1 \le m, n \le 3 \times 10^4`$;
- the linked structure is acyclic, so the nodes within one list are pairwise distinct.

### 4.2 Output

If $`A`$ and $`B`$ share a node, output the first node of $`A`$ that also appears in $`B`$; otherwise output $`\mathrm{NIL}`$. The lists are unchanged.

### 4.3 Pseudocode

```text
GET-INTERSECTION-NODE(headA, headB)
    a = headA
    b = headB
    while a ≠ b                       ▷ compares the nodes themselves, not their values
        if a == NIL
            a = headB
        else a = a.next
        if b == NIL
            b = headA
        else b = b.next
    return a
```

When the loop ends, `a == b`, and `a` is returned directly; there is no separate result variable.

**State Definitions**

- `a`: a pointer starting at `headA`; it walks all of $`A`$, spends one step on $`\mathrm{NIL}`$, then walks $`B`$ from `headB`.
- `b`: a pointer starting at `headB`; it walks all of $`B`$, spends one step on $`\mathrm{NIL}`$, then walks $`A`$ from `headA`.
- Each iteration advances both pointers by one step, so at the start of iteration $`t`$ both have taken $`t`$ steps.

### 4.4 Correctness

**Proof method: fix the shape of the lists (Lemma 1), write down both pointers' positions at step $`t`$ (Lemma 2), then do a case analysis on the first meeting**

#### I. Key Properties Required for Correctness

1. **Each node has exactly one `next`** (definition of a linked list): the sequence of nodes after a given node is uniquely determined.
2. **No cycles** (problem condition): the nodes of each list are pairwise distinct, and each list reaches $`\mathrm{NIL}`$ in finitely many steps.
3. **Lemma 1 (the intersection is a shared suffix)**: there exist $`\alpha, \beta, c \ge 0`$ with $`m = \alpha + c`$ and $`n = \beta + c`$ such that
   - $`A_{\alpha + k} = B_{\beta + k}`$ for $`k < c`$, written $`C_k`$;
   - no node of the private prefix $`P_A = \langle A_0, \ldots, A_{\alpha-1} \rangle`$ is in $`B`$, and no node of $`P_B = \langle B_0, \ldots, B_{\beta-1} \rangle`$ is in $`A`$.

   So the correct answer is $`C_0`$ (when $`c > 0`$) or $`\mathrm{NIL}`$ (when $`c = 0`$).

**Proof of Lemma 1.** If $`A`$ and $`B`$ share no node, take $`\alpha = m`$, $`\beta = n`$, $`c = 0`$. Otherwise let $`\alpha`$ be the index in $`A`$ of the first shared node, and $`\beta`$ its index in $`B`$ (unique, since the nodes of $`B`$ are pairwise distinct). By property 1, the sequences after $`A_\alpha = B_\beta`$ are identical on both sides until both reach $`\mathrm{NIL}`$, so $`m - \alpha = n - \beta`$, call it $`c \ge 1`$, and $`A_{\alpha+k} = B_{\beta+k}`$. No node of $`P_A`$ is in $`B`$ by the definition of $`\alpha`$. If some $`B_j`$ with $`j < \beta`$ were in $`A`$, its index in $`A`$ would be at least $`\alpha`$ ($`\alpha`$ is the first shared node), so $`B_j = C_k = B_{\beta+k}`$ with $`j < \beta + k`$, contradicting that the nodes of $`B`$ are pairwise distinct. ∎

**Lemma 2 (position at step $`t`$).** Define the two sequences

```math
S_A = \langle A_0, \ldots, A_{m-1}, \mathrm{NIL}, B_0, \ldots, B_{n-1}, \mathrm{NIL} \rangle, \qquad
S_B = \langle B_0, \ldots, B_{n-1}, \mathrm{NIL}, A_0, \ldots, A_{m-1}, \mathrm{NIL} \rangle
```

each of length $`m + n + 2`$. For $`0 \le t \le m + n + 1`$, if the loop has not ended before iteration $`t`$, then `a` $`= S_A[t]`$ and `b` $`= S_B[t]`$.

*Proof*: at $`t = 0`$, `a = headA` $`= A_0`$. In each iteration, `a` follows `next` from a non-$`\mathrm{NIL}`$ node and switches to `headB` from $`\mathrm{NIL}`$; by the definition of $`A_i.\mathit{next}`$ this turns $`S_A[t]`$ into exactly $`S_A[t+1]`$ ($`\mathrm{NIL}`$ follows $`A_{m-1}`$, $`B_0`$ follows $`\mathrm{NIL}`$, and $`\mathrm{NIL}`$ follows $`B_{n-1}`$). The same holds for `b`. ∎

Let $`T = \alpha + \beta + c + 1`$. Since $`T = m + 1 + \beta = n + 1 + \alpha`$,

```math
S_A[T] = B_\beta, \qquad S_B[T] = A_\alpha
```

where for $`c = 0`$, $`B_\beta = B_n`$ means the $`\mathrm{NIL}`$ that follows $`B`$ in the sequence, and likewise for $`A_\alpha`$.

#### II. Correctness Proof

**Theorem.** The loop of `GET-INTERSECTION-NODE` ends after at most $`T`$ iterations and returns $`C_0`$ (if $`c > 0`$) or $`\mathrm{NIL}`$ (if $`c = 0`$).

*Proof*: the loop ends at the first $`t`$ with $`S_A[t] = S_B[t]`$ and returns $`S_A[t]`$ (Lemma 2). Within $`0 \le t \le T`$, each node appears in $`S_A`$ and $`S_B`$ at these indices:

| Node | Index in $`S_A[0..T]`$ | Index in $`S_B[0..T]`$ |
|---|---|---|
| $`A_i \in P_A`$ | $`i`$ | $`n + 1 + i`$ |
| $`B_j \in P_B`$ | $`m + 1 + j`$ | $`j`$ |
| $`C_k`$ ($`c > 0`$) | $`\alpha + k`$; also $`T`$ when $`k = 0`$ | $`\beta + k`$; also $`T`$ when $`k = 0`$ |
| $`\mathrm{NIL}`$ | $`m`$; also $`T`$ when $`c = 0`$ | $`n`$; also $`T`$ when $`c = 0`$ |

The table holds because $`S_A[0..T]`$ is exactly $`A`$, $`\mathrm{NIL}`$, $`P_B`$, $`S_A[T]`$, and by Lemma 1, $`P_B`$ contains no node of $`A`$ and $`P_A`$ contains no node of $`B`$; $`S_B`$ is symmetric.

By the table, a private node can never make $`S_A[t] = S_B[t]`$: $`i \ne n + 1 + i`$ and $`m + 1 + j \ne j`$. So a meeting can only happen at a shared node or at $`\mathrm{NIL}`$. Two cases.

- **Case 1: $`\alpha = \beta`$ (equal private prefixes, hence $`m = n`$).** For $`t < \alpha`$ both sides are private nodes, so there is no meeting.
  - If $`c > 0`$: at $`t = \alpha`$, $`S_A[\alpha] = C_0 = S_B[\alpha]`$, the first meeting; $`C_0`$ is returned.
  - If $`c = 0`$: at $`t = \alpha = m = n`$ both are $`\mathrm{NIL}`$, the first meeting; $`\mathrm{NIL}`$ is returned.
- **Case 2: $`\alpha \ne \beta`$ (hence $`m \ne n`$).** By the table, $`C_k`$ sits at indices $`\alpha + k \ne \beta + k`$ and $`\mathrm{NIL}`$ at $`m \ne n`$, so there is no meeting for $`t < T`$. At $`t = T`$, $`S_A[T] = B_\beta`$ and $`S_B[T] = A_\alpha`$:
  - if $`c > 0`$, both are $`C_0`$, and $`C_0`$ is returned;
  - if $`c = 0`$, both are $`\mathrm{NIL}`$, and $`\mathrm{NIL}`$ is returned.

In both cases the first meeting happens at $`t \le T \le m + n + 1`$, within the range where Lemma 2 applies, so the loop ends and the return value is correct. The algorithm only reads `next` and writes no node, so the lists are unchanged. ∎

#### III. Conclusion

Lemma 1 splits the two lists into "private prefixes + a shared suffix", and the answer is the first node of the shared suffix. Each pointer walks its own list and then the other list, so at step $`T`$ both have taken $`\alpha + \beta + c + 1`$ steps and arrive at the start of the shared suffix together (or at $`\mathrm{NIL}`$ together); private nodes always sit at different positions in the two sequences, so they never cause an early, wrong meeting.

### 4.5 Complexity Analysis

Let the list lengths be $`m`$ and $`n`$, with $`\alpha`$, $`\beta`$, $`c`$ as in Lemma 1.

**Time Complexity: $`O(m + n)`$**

By the proof of the theorem, the loop runs

```math
t^* =
\begin{cases}
\alpha, & \alpha = \beta,\ c > 0 \\
m, & \alpha = \beta,\ c = 0 \\
\alpha + \beta + c + 1, & \alpha \ne \beta
\end{cases}
```

iterations, each doing constant work, so the total time is $`\Theta(t^*) = O(m + n)`$. The worst case is lists of different lengths that do not intersect, with $`t^* = m + n + 1`$; the test `max_length_unequal_no_intersection` is that case. The best case is `headA == headB`, with $`t^* = 0`$.

**Space Complexity: $`\Theta(1)`$**

Only the two pointers `a` and `b` are used: no heap allocation, no recursion, and no modification of the lists.

### 4.6 Test Design

The tests use `buildIntersecting` to build, from the judge's format (`listA`, `listB`, `skipA`, `skipB`), two lists that really share nodes, and to return the expected intersection node; results are compared by pointer equality. Every case also uses `sameStructure` to check that after the call each list still visits exactly the nodes it was built with, matching the requirement that the lists keep their original structure.

| Case | Expected | What It Checks |
|---|---|---|
| `example_1` | node with value 8 | Official example 1; both private parts contain a node with value 1, so comparing values would go wrong |
| `example_2` | node with value 2 | Official example 2 |
| `example_3_no_intersection` | `nil` | Official example 3: different lengths, no intersection |
| `same_list` | the head | `headA == headB`; zero iterations |
| `b_is_suffix_of_a` | node 2 of A | $`\beta = 0`$: B has no private prefix |
| `a_is_suffix_of_b` | head of A | $`\alpha = 0`$: A has one node |
| `equal_prefixes` | first node with value 9 | Case 1: equal private prefixes, meeting during the first pass |
| `intersect_at_last_node` | last node of A | The shared suffix has one node |
| `equal_length_no_intersection` | `nil` | Case 1 with $`c = 0`$: both reach `nil` together at the end of the first pass |
| `single_nodes_no_intersection` | `nil` | Two one-node lists with equal values but different nodes |
| `max_length_long_vs_short` | last node of A | $`m = 3 \times 10^4`$, $`n = 1`$ |
| `max_length_equal_no_intersection` | `nil` | $`m = n = 3 \times 10^4`$, no intersection |
| `max_length_unequal_no_intersection` | `nil` | $`m = 3 \times 10^4`$, $`n = m - 1`$, no intersection; the loop runs the full $`m + n + 1`$ iterations |

Three mutants of the implementation were tried (each case run separately with a 3-second timeout):

- Skipping the `nil` step and switching to the other head directly from the last node (`if a.Next == nil { a = headB }`): every non-intersecting case times out. The two pointers are never `nil` at the same time and private nodes never meet, so the loop never ends.
- Advancing together without switching heads (`for a != nil && b != nil && a != b`): `example_1`, `example_3_no_intersection`, `a_is_suffix_of_b`, `intersect_at_last_node`, `max_length_long_vs_short`, and `max_length_unequal_no_intersection` fail.
- Switching back to the pointer's **own** head after `nil`: all cases pass; see Challenge 3.

## 5. Learning Notes

### Challenge 1: Why does switching heads "align" the pointers?

When the lists have different lengths, two pointers that start at the heads together reach the shared suffix at different times. A common approach is to compute both lengths first and let the longer list advance $`\lvert m - n \rvert`$ steps.

**Understanding Process**

Head switching aligns the pointers without computing lengths. Pointer `a` walks $`P_A`$, the shared suffix, $`\mathrm{NIL}`$, $`P_B`$, and then reaches the start of the shared suffix; pointer `b` walks $`P_B`$, the shared suffix, $`\mathrm{NIL}`$, $`P_A`$, and reaches the same start. Both paths have length $`\alpha + c + 1 + \beta`$, since reordering a sum does not change it, so they arrive in the same step. The length-first approach is also $`O(m + n)`$ time and $`O(1)`$ space, with two extra traversals; both are the same idea: cancel the difference $`\alpha - \beta`$.

**Transferable Insight:** when two processes of different lengths need to line up, let each one also walk the other's segment, and the totals become equal.

### Challenge 2: Why must `nil` count as a step?

In the pseudocode, a pointer that reaches $`\mathrm{NIL}`$ does not jump to the other list immediately: it stays on $`\mathrm{NIL}`$ for one iteration and switches to the other head on the next.

**Understanding Process**

This step is what ends the loop when the lists do not intersect. If $`c = 0`$, the lists share no node, and the only place the pointers can be equal is when both are $`\mathrm{NIL}`$. The $`\mathrm{NIL}`$ entries in $`S_A`$ and $`S_B`$ are exactly what make step $`T`$ of Lemma 2 land on two $`\mathrm{NIL}`$s. The first mutant drops this step, so the pointers are never $`\mathrm{NIL}`$ and the loop runs forever when the lists do not intersect. When they do intersect, the step costs nothing: both pointers take one extra step and still reach $`C_0`$ together.

**Transferable Insight:** when a loop stops on "equal", make sure the "no answer" case also has an end point both sides reach (here `nil`); otherwise the loop may not terminate.

### Challenge 3: Why does switching back to the pointer's own head also "pass"?

In the mutation tests, changing `a = headB` to `a = headA` (and likewise for `b`) still passes every case.

**Understanding Process**

With that change, `a` cycles through $`A`$ with period $`m + 1`$ and `b` cycles through $`B`$ with period $`n + 1`$. By the same reasoning as the theorem, private nodes never meet; if the two meet at $`C_k`$ or at $`\mathrm{NIL}`$, stepping back $`k`$ (or $`c`$) steps puts both at $`C_0`$, so the first meeting is still the correct answer. Both are $`\mathrm{NIL}`$ at step $`\mathrm{lcm}(m + 1, n + 1) - 1`$, so the loop always ends. But the meeting can take on the order of $`\mathrm{lcm}(m + 1, n + 1)`$ steps: in `max_length_unequal_no_intersection`, $`m + 1 = 30001`$ and $`n + 1 = 30000`$ are coprime, so it takes about $`9 \times 10^8`$ steps; in the environment where these notes were written, that one case took 1.66 seconds, while the correct version takes only $`m + n + 1 \approx 6 \times 10^4`$ steps. The tests check results only, so they cannot catch it; the difference is in complexity.

**Transferable Insight:** "the result is right" is not the same as "the algorithm is right". Switching to the other head is what bounds the meeting at $`m + n + 1`$ steps; when analyzing a variant, argue the result and the complexity separately.

### Challenge 4: Compare nodes, not values

**Understanding Process**

In example 1, the second node of A and the third node of B both have value 1, but they are different nodes; the intersection only starts at the node with value 8. In Go, `a != b` compares whether two `*ListNode` pointers point to the same node, which is exactly "the same node" in the problem. The "shared node" of Lemma 1 means the same: only sharing the same memory forces all later nodes to coincide (property 1); equal values do not.

**Transferable Insight:** first decide whether the problem wants value equality or reference equality. In linked lists, trees, and graphs, "the same node" usually means reference equality.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Shared Suffix Structure**: here, Lemma 1 proves that once two acyclic singly linked lists meet, they coincide until the tail, so the answer is the first node of the shared suffix. Applies to any problem where two chains merge somewhere, such as the paths from two tree nodes to the root. It requires each node to have a single successor; with several successors (general graphs), paths can split again after merging and the conclusion fails.
2. **Length Equalization by Concatenation**: here, the pointers walk "A then B" and "B then A", whose total lengths are equal, so they line up without computing lengths. Applies to problems that need two sequences of different lengths to reach a position together. It needs a definite switching point (here `nil`); otherwise the loop does not end when there is no shared node.
3. **Two Pointers on Linked Lists**: here, two pointers advancing in lockstep replace a hash set of visited nodes with $`\Theta(1)`$ extra space. Applies to finding the middle, the $`k`$-th node from the end, cycle detection, and similar problems. Limitation: it relies on the list's special structure (here, the shared suffix) and is not a general set-intersection method.
4. **Reference Equality**: here, `a != b` compares node addresses, not node values. Applies to deciding "is this the same node" in linked lists, trees, and graphs. When values can repeat, comparing values gives wrong answers.

Storing all nodes of A in a hash set and then scanning B for the first node in the set is also correct, in expected $`O(m + n)`$ time, but it needs $`\Theta(m)`$ extra space and does not meet the follow-up.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. Relation is either "Direct" or "Related idea"; if no section fits, write "No matching section" and keep only the concept name.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Singly linked list | §10.2 Linked lists | Direct | `Next` of `ListNode` is the next attribute of a CLRS node, and the `nil` tail corresponds to NIL; this problem only reads along `next`, with no insertion or deletion, and Lemma 1 relies on "each node has exactly one next" |

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

- [141. Linked List Cycle](https://leetcode.cn/problems/linked-list-cycle/) — Also two pointers advancing on a linked list with "pointers equal" as the signal, but at different speeds (fast and slow pointers), to detect a cycle.
- [142. Linked List Cycle II](https://leetcode.cn/problems/linked-list-cycle-ii/) — Finds the cycle entrance: after the fast and slow pointers meet, one returns to the head and both advance at the same speed, meeting at the entrance; like this problem, the alignment comes from two paths of equal length.
- [19. Remove Nth Node From End of List](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/) — One pointer advances $`N`$ steps first and then both move together, a direct use of the "advance by the length difference" alignment from Challenge 1.
- [236. Lowest Common Ancestor of a Binary Tree](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/) — The paths from two nodes to the root are two merging "chains", and the lowest common ancestor is the start of their shared suffix; with parent pointers, this problem's method applies directly.

---

**One-Sentence Summary**

Two acyclic singly linked lists that intersect share every node from the intersection to the tail, so letting each pointer walk "its own list + `nil` + the other list" makes the totals equal: the first meeting is the intersection start, and without an intersection both reach `nil` together.
