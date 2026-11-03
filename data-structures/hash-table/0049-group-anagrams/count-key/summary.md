# 49. Group Anagrams — Count Key

English | [简体中文](./summary.zh-CN.md)

Terminology follows the [shared glossary](../../../../docs/GLOSSARY.md), and math notation follows the repository's [GitHub math writing conventions](../../../../docs/GITHUB_MATH.md). The hash table itself (its definition, Go `map` behavior, and common usage patterns) is covered in the [hash table overview](../../README.md).

This problem has two solutions in two directories that share one set of test cases:

- [**Sorted key**](../sorted-key/summary.md): sort the letters of each string and use the result as the key. The problem statement, the formal definition of anagrams (Lemma 0), and the general correctness proof for grouping by key (Lemma 2) are on that page.
- **Count key** (this page): use the 26 letter counts of each string as the key, with no sorting.

This page only covers how to compute a valid key without sorting.

## 1. Problem Information

- **Problem**: [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)
- **Official Difficulty**: Medium
- **Primary Category**: Hash Table
- **Secondary Category**: Counting
- **Tags**: `Hash Table`, `Counting`, `Direct-Address Table`, `Canonical Form`, `Group by Key`, `Equivalence Classes`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](../solution_test.go) (in the parent directory, shared by both solutions)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Count-Key Grouping**

- Evaluation Scope: $`1 \le n \le 10^4`$, $`0 \le |s_i| \le 100`$, lowercase English letters only; expected $`O(L + n|\Sigma|)`$ time, $`O(n|\Sigma|)`$ extra space (notation in §4.5); Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2 | One direct transformation: turn "being anagrams" into "equal count vectors over the 26 letters" and look up by the count vector |
| Knowledge Prerequisites | 2 | A hash table plus counting letters in an array (a direct-address table), both common tools |
| Correctness Reasoning Difficulty | 2 | The count vector is exactly the quantity in Lemma 0, so the only addition is an intuitive invariant showing the counting loop computes the counts; the grouping part cites Lemma 2 |
| Implementation Difficulty | 2 | One must know that Go arrays are comparable and can be `map` keys; concatenating counts into a string without separators would let different counts collide on one key |

Ratings follow [RATING.md](../../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

See [sorted key §2](../sorted-key/summary.md#2-problem-understanding) for the statement, example, and key requirements. This solution also relies on one constraint: `strs[i]` contains only lowercase English letters, so the alphabet $`\Sigma`$ has 26 letters and an array of length 26 indexed by `s[i]-'a'` can hold the counts.

## 4. Algorithm Specification

**Algorithm Name:** Count-Key Grouping

### 4.1 Input

Same as [sorted key §4.1](../sorted-key/summary.md#41-input): $`S = \langle s_0, \ldots, s_{n-1} \rangle`$, where each $`s_i`$ is a string over $`\Sigma = \{\mathtt{a}, \ldots, \mathtt{z}\}`$ of length at most 100. Number the letter $`c`$ as $`\mathrm{idx}(c) = c - \mathtt{a} \in \{0, \ldots, 25\}`$.

### 4.2 Output

Same as [sorted key §4.2](../sorted-key/summary.md#42-output): the output is exactly the set of equivalence classes of the relation "being anagrams".

### 4.3 Pseudocode

The pseudocode is 0-indexed.

```text
GROUP-ANAGRAMS-COUNT-KEY(S)
    T ← empty hash table            // count vector → list of strings
    for i ← 0 to n - 1
        cnt[0..25] ← all 0
        for j ← 0 to |s_i| - 1
            cnt[s_i[j] - 'a'] ← cnt[s_i[j] - 'a'] + 1
        if cnt ∉ T
            T[cnt] ← empty list
        APPEND(T[cnt], s_i)

    res ← empty list
    for each key in T
        APPEND(res, T[key])
    return res
```

**State Definitions**

- `T`: `groups` in the code, of type `map[[26]int][]string`, mapping each count vector to the list of strings with that key.
- `cnt`: `cnt` in the code, of type `[26]int`; `cnt[c - 'a']` is the number of times letter $`c`$ occurs in the current string. It is redeclared with `var cnt [26]int` in each outer iteration, which zeroes it.
- `res`: the list of all groups.

Go arrays are values and comparable, so `groups[cnt]` copies the current 26 counts as the key, and the next iteration's `cnt` is independent of keys already in the table. As in the sorted-key solution, the code relies on a missing key reading as `nil` to skip the explicit `if cnt ∉ T` branch.

### 4.4 Correctness

**Proof Method: Counting-Loop Invariant + Lemmas 0 and 2 from the Sorted-Key Solution**

#### 1. Key Conditions for Correctness

1. **Anagrams are exactly strings with equal counts** ([sorted key, Lemma 0](../sorted-key/summary.md#2-correctness-proof)): $`s \sim t \iff \mathrm{cnt}_s = \mathrm{cnt}_t`$.
2. **The counting loop computes $`\mathrm{cnt}_s`$** (Lemma 3, proved here).
3. **Grouping by a complete invariant yields all equivalence classes** ([sorted key, Lemma 2](../sorted-key/summary.md#2-correctness-proof)), provided the key function $`f`$ satisfies $`f(s) = f(t) \iff s \sim t`$.
4. **The input has only lowercase letters** (given): this keeps the index `s[i]-'a'` within $`[0, 25]`$.

#### 2. Correctness Proof

**Lemma 3 (counting loop)**: For a string $`s`$, when the inner loop ends, $`\mathrm{cnt}[\mathrm{idx}(c)] = \mathrm{cnt}_s(c)`$ for every letter $`c`$.

**Loop Invariant**

At the start of each inner iteration (after processing $`s[0..j-1]`$):

**Invariant: Prefix Counts**

For every letter $`c`$, $`\mathrm{cnt}[\mathrm{idx}(c)]`$ equals the number of occurrences of $`c`$ in the prefix $`s[0..j-1]`$.

**Initialization**

With $`j = 0`$ the prefix is empty and every count is 0, matching the zeroed `cnt`.

**Maintenance**

Assume the invariant holds before processing $`s[j]`$, and let $`c' = s[j]`$.

- **Letter $`c'`$**: extending the prefix from $`s[0..j-1]`$ to $`s[0..j]`$ adds one occurrence of $`c'`$, and the algorithm executes exactly $`\mathrm{cnt}[\mathrm{idx}(c')] \leftarrow \mathrm{cnt}[\mathrm{idx}(c')] + 1`$. By condition 4, $`\mathrm{idx}(c') \in [0, 25]`$, so the index is valid.
- **Any other letter $`c \ne c'`$**: its count is unchanged, and since $`\mathrm{idx}`$ is injective, $`\mathrm{cnt}[\mathrm{idx}(c)]`$ is not modified and stays correct.

**Termination**

The loop ends with $`j = |s|`$, when the prefix is all of $`s`$, and the invariant gives $`\mathrm{cnt}[\mathrm{idx}(c)] = \mathrm{cnt}_s(c)`$.

**The count key is a complete invariant**: let $`f(s)`$ be the array computed in Lemma 3. By Lemma 3, $`f(s) = f(t)`$ if and only if $`\mathrm{cnt}_s(c) = \mathrm{cnt}_t(c)`$ for every $`c`$, that is, $`\mathrm{cnt}_s = \mathrm{cnt}_t`$; by Lemma 0 this holds if and only if $`s \sim t`$. Go compares two `[26]int` values component by component, which matches this notion of equality.

**Grouping**: the outer loop has the same structure as in the sorted-key solution, with the key replaced by $`f`$. Since $`f`$ satisfies the condition of Lemma 2, Lemma 2's loop invariant and conclusion carry over unchanged: the output is exactly the set of equivalence classes, and each group keeps input order.

#### 3. Conclusion

The count array is exactly the quantity that decides whether two strings are anagrams, so it neither splits a group nor merges two groups; grouping by it yields exactly the equivalence classes the problem asks for. The order of the groups follows Go's random `map` iteration order, which the problem allows.

### 4.5 Complexity Analysis

Notation follows [sorted key §4.5](../sorted-key/summary.md#45-complexity-analysis): $`n`$ strings, total length $`L = \sum_i \ell_i`$, maximum length $`k`$, and $`g \le n`$ groups; in addition, let $`|\Sigma| = 26`$ be the alphabet size.

**Time Complexity: expected O(L + n|Σ|)**

Steps for one string $`s_i`$:

- Zeroing `cnt`: $`O(|\Sigma|)`$.
- Counting loop: $`\ell_i`$ steps at $`O(1)`$ each, $`O(\ell_i)`$.
- Lookup and append: the key is 26 integers, so hashing and comparing keys cost $`O(|\Sigma|)`$; an expected $`O(1)`$ comparisons and an amortized $`O(1)`$ `append` give expected $`O(|\Sigma|)`$.

The second loop costs $`O(g)`$. The expected total is

```math
O\Bigl(\sum_i \bigl(\ell_i + |\Sigma|\bigr) + g\Bigr) = O(L + n|\Sigma|)
```

Here $`|\Sigma| = 26`$ is a constant, so this is expected $`O(n + L)`$, the same order as reading the input. Compared with the sorted key's $`O(n + \sum_i \ell_i \log \ell_i)`$, the logarithmic factor of sorting is gone. But when strings are short (say $`\ell_i < 26`$), the fixed $`|\Sigma|`$ cost per string can exceed the cost of sorting, so which is faster within this problem's limits must be measured, not read off the asymptotic bounds.

**Space Complexity: O(n|Σ|)**

- Hash-table keys: $`g`$ values of type `[26]int`, $`O(g|\Sigma|)`$. Unlike the sorted key, the key size does not depend on string length.
- Hash-table lists: $`n`$ string headers in total, $`O(n)`$, which are the output; the string contents are shared with the input, not copied.
- `cnt`: $`O(|\Sigma|)`$.
- `res`: $`O(g)`$.

The total space is $`O(n + g|\Sigma|) \le O(n|\Sigma|)`$. Excluding the output, the extra space is $`O(g|\Sigma|)`$, which is $`O(n|\Sigma|)`$ in the worst case.

## 5. Learning Notes

### Challenge 1: Why not join the counts into a string key?

In languages where arrays cannot be keys, a common approach is to join the 26 counts into a string. Joining the numbers directly is wrong.

**Understanding Process**

If the counts are written in decimal and concatenated in order, `"aaaaaaaaaaa"` (11 `a`s) has counts $`(11, 0, 0, \ldots)`$, which becomes `"110"` followed by 24 `0`s; `"abbbbbbbbbb"` (one `a`, ten `b`s) has counts $`(1, 10, 0, \ldots)`$, which also becomes `"110"` followed by 24 `0`s. The two strings are not anagrams but get the same key, violating the "$`f(s) = f(t) \Rightarrow s \sim t`$" direction of a complete invariant. Either put separators between the numbers (such as `"11#0#0..."`) or, as this solution does, use a fixed-length array as the key. The test case `Letter counts that collide if concatenated without separators` uses exactly this pair.

**Transferable Insight:** When encoding several fields into one key, the encoding must be **injective**: different field combinations must give different keys. Directly concatenating variable-length fields is usually not injective; add separators, length prefixes, or use a fixed-length representation.

### Challenge 2: Is the count key always better than the sorted key?

The count key has a lower asymptotic time, but it depends on a small, fixed alphabet.

**Understanding Process**

The count key pays $`O(|\Sigma|)`$ per string to zero, hash, and compare, and each key holds $`|\Sigma|`$ integers. Here $`|\Sigma| = 26`$ is a constant; but if strings could contain arbitrary Unicode characters, $`|\Sigma|`$ would be far too large for array counting. One would then either count with a `map[rune]int` and normalize it, or go back to the sorted key, whose cost depends only on string length.

**Transferable Insight:** Direct addressing (counting by value in an array) suits small value ranges; with large ranges its space and zeroing costs outweigh the benefit. When choosing a canonical form, consider both the string length $`k`$ and the alphabet size $`|\Sigma|`$.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Count Vector as Canonical Form**: here, the count vector over the 26 letters is the unique representative of each anagram class. It applies to comparisons that care about multisets rather than order, such as testing rearrangements or matching anagrams in a sliding window; the alphabet must be small and fixed.
2. **Direct-Address Counting**: here, `cnt[s[i]-'a']++` counts letters in $`O(1)`$ each with no hash collisions. It applies to counting over a small range of values; for large or sparse ranges, count with a hash table instead.
3. **Injective Key Encoding**: here, a fixed-length array as the key avoids different counts colliding as they would in a naively concatenated string. It applies whenever several fields are combined into one hash key; different field combinations must encode to different keys.
4. **Group by Key**: the same as in the [sorted-key solution](../sorted-key/summary.md#6-key-takeaways); correctness depends only on the key being a complete invariant, not on its form.

### CLRS Connections

Section numbers follow the third edition of CLRS. "Connection" is either "Direct" or "Related idea"; when no specific section applies, the "CLRS Section" column says "No matching section" and only the concept name is kept.

| Concept | CLRS Section | Connection | Where It Appears |
|---|---|---|---|
| Loop invariant | §2.1 Insertion sort | Direct | Lemma 3 uses the invariant "`cnt` holds the letter counts of the prefix", proved by initialization, maintenance, and termination, to show the counts are correct |
| Direct-address tables | §11.1 Direct-address tables | Direct | `cnt` is a direct-address table keyed by the letter numbers $`0 \ldots 25`$, with worst-case $`O(1)`$ per count |
| Hash tables | §11.2 Hash tables | Direct | `groups` is keyed by `[26]int`; each string costs one lookup and one append, expected $`O(\lvert\Sigma\rvert)`$ |
| Counting sort | §8.2 Counting sort | Related idea | Computing `cnt` is the same as the step of COUNTING-SORT that counts occurrences of each value; this solution stops there instead of expanding the counts into a sorted sequence, because the count vector itself can serve as the key |
| Equivalence relations and partitions | Appendix B.2 Relations | Direct | Cites Lemma 0 of the sorted-key solution: "being anagrams" is an equivalence relation, so by Theorem B.1 its classes form a partition, and the output groups are exactly those classes |

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

- [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/) — Uses the same length-26 count array to decide whether two strings are anagrams; this problem turns that array from a test into a grouping key.
- [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/) — Maintains the count vector incrementally in a sliding window: each shift adds 1 to one letter and subtracts 1 from another in $`O(1)`$, then compares with the target counts. The sorted key cannot be updated incrementally like this.
- [1347. Minimum Number of Steps to Make Two Strings Anagram](https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/) — Also characterizes anagrams with a length-26 count array; the answer is the sum of the positive componentwise differences of the two count vectors.

---

**One-Sentence Summary**

The letter-count vector is exactly what decides whether strings are anagrams, so using the comparable `[26]int` directly as the hash-table key groups everything in one pass without sorting, at a fixed extra cost of 26 per string.
