# 49. Group Anagrams — Sorted Key

English | [简体中文](./summary.zh-CN.md)

Terminology follows the [shared glossary](../../../../docs/GLOSSARY.md), and math notation follows the repository's [GitHub math writing conventions](../../../../docs/GITHUB_MATH.md). The hash table itself (its definition, Go `map` behavior, and common usage patterns) is covered in the [hash table overview](../../README.md); this note covers only how this problem uses a hash table.

This problem has two solutions in two directories that share one set of test cases:

- **Sorted key** (this page): sort the letters of each string and use the result as the hash-table key. This page gives the problem statement, a formal definition of "anagram", and the general correctness proof for grouping by key; the other page cites it.
- [**Count key**](../count-key/summary.md): use the 26 letter counts of each string as the key, with no sorting.

## 1. Problem Information

- **Problem**: [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)
- **Official Difficulty**: Medium
- **Primary Category**: Hash Table
- **Secondary Category**: Sorting
- **Tags**: `Hash Table`, `Sorting`, `Canonical Form`, `Group by Key`, `Equivalence Classes`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](../solution_test.go) (in the parent directory, shared by both solutions)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Sorted-Key Grouping**

- Evaluation Scope: $`1 \le n \le 10^4`$, $`0 \le |s_i| \le 100`$, lowercase English letters only; expected $`O(n k \log k)`$ time, $`O(n + L)`$ extra space (notation in §4.5); Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2 | One direct transformation: seeing that "being anagrams" is the same as "equal after sorting" turns grouping into a lookup by the sorted string |
| Knowledge Prerequisites | 2 | A hash table (key to list) and sorting a byte slice, both common tools |
| Correctness Reasoning Difficulty | 2 | The key point is that the sorted key satisfies both directions, "same group implies same key" and "same key implies same group"; one intuitive loop invariant then shows every string lands in exactly one group |
| Implementation Difficulty | 1 | One loop and one `map` append; the only catch is that a slice cannot be a key and must be converted to `string` first |

Ratings follow [RATING.md](../../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

Given an array of strings `strs`, put strings that are **anagrams** of each other into the same group and return all groups. An anagram is a string formed by rearranging the letters of another string, using each letter exactly once, such as `"eat"`, `"tea"`, and `"ate"`. The groups and the strings within a group may be returned in any order.

**Example**

- Input: `strs = ["eat","tea","tan","ate","nat","bat"]`
- Output: `[["bat"],["nat","tan"],["ate","eat","tea"]]`

**Key Requirements**

1. $`1 \le \mathrm{strs.length} \le 10^4`$, $`0 \le \mathrm{strs}[i].\mathrm{length} \le 100`$, and `strs[i]` consists of lowercase English letters.
2. Every input string appears in exactly one group; repeated strings (such as two copies of `"abc"`) are separate elements and both go into the same group.
3. The empty string `""` is valid input; all empty strings are anagrams of each other and form one group. Strings of different lengths are never in the same group.

## 4. Algorithm Specification

**Algorithm Name:** Sorted-Key Grouping

### 4.1 Input

$`S = \langle s_0, s_1, \ldots, s_{n-1} \rangle`$

subject to:

- $`1 \le n \le 10^4`$
- each $`s_i`$ is a string over the alphabet $`\Sigma = \{\mathtt{a}, \ldots, \mathtt{z}\}`$ with $`0 \le |s_i| \le 100`$

For a string $`s`$ and a letter $`c \in \Sigma`$, let $`\mathrm{cnt}_s(c)`$ be the number of occurrences of $`c`$ in $`s`$.

### 4.2 Output

Say $`t`$ is an **anagram** of $`s`$, written $`s \sim t`$, if $`|t| = |s|`$ and there is a permutation $`\pi`$ of $`\{0, \ldots, |s|-1\}`$ with $`t[j] = s[\pi(j)]`$ for all $`j`$.

On the index set $`\{0, \ldots, n-1\}`$, define $`i \equiv j`$ if and only if $`s_i \sim s_j`$. The output `res` is a list of groups and must satisfy:

- **Exactly once**: every index $`i`$ appears in exactly one group (the group stores the string $`s_i`$).
- **Same group implies equivalent**: any two strings in the same group are anagrams.
- **Equivalent implies same group**: any two anagrams are in the same group.

In other words, `res` is exactly the set of equivalence classes of $`\equiv`$ (Lemma 0 proves that $`\equiv`$ is an equivalence relation).

### 4.3 Pseudocode

The pseudocode is 0-indexed.

```text
GROUP-ANAGRAMS-SORTED-KEY(S)
    T ← empty hash table            // key → list of strings
    for i ← 0 to n - 1
        key ← SORT(s_i)             // letters of s_i in ascending order
        if key ∉ T
            T[key] ← empty list
        APPEND(T[key], s_i)

    res ← empty list
    for each key in T
        APPEND(res, T[key])
    return res
```

**State Definitions**

- `T`: `groups` in the code, of type `map[string][]string`, mapping each sorted key to the list of strings with that key.
- `key`: `key` in the code, the string $`s_i`$ with its letters sorted. The code copies `s` into a byte slice `b` with `[]byte(s)`, sorts it in place with `slices.Sort(b)`, and converts it back with `string(b)` to use as the key; the input strings are never modified.
- `res`: the list of all groups.

The code has no explicit `if key ∉ T` branch: reading a missing key gives `groups[key] == nil`, and `append(nil, s)` creates a new slice, which has the same effect as the pseudocode.

### 4.4 Correctness

**Proof Method: Partition into Equivalence Classes + Complete Invariant + Loop Invariant**

#### 1. Key Conditions for Correctness

1. **Anagrams are exactly strings with equal letter counts** (Lemma 0): $`s \sim t \iff \mathrm{cnt}_s = \mathrm{cnt}_t`$. Hence $`\equiv`$ is an equivalence relation, and its classes partition the index set into disjoint groups.
2. **The sorted key is a complete invariant** (Lemma 1): $`\mathrm{SORT}(s) = \mathrm{SORT}(t) \iff s \sim t`$. The "$`\Rightarrow`$" direction keeps different classes from being merged; the "$`\Leftarrow`$" direction keeps one class from being split.
3. **Grouping by key in a hash table is exact** (Lemma 2): for any key function $`f`$ with $`f(s) = f(t) \iff s \sim t`$, grouping by $`f`$ yields exactly the equivalence classes. This does not depend on the form of the key, and the [count-key solution](../count-key/summary.md) cites it directly.

#### 2. Correctness Proof

**Lemma 0**: $`s \sim t \iff \mathrm{cnt}_s = \mathrm{cnt}_t`$; hence $`\sim`$ (and therefore $`\equiv`$) is an equivalence relation.

*Proof*:

- ($`\Rightarrow`$) Since $`t[j] = s[\pi(j)]`$ and $`\pi`$ is a bijection, for each letter $`c`$ the positions $`j`$ where $`t`$ has $`c`$ correspond one to one with the positions $`\pi(j)`$ where $`s`$ has $`c`$, so $`\mathrm{cnt}_t(c) = \mathrm{cnt}_s(c)`$.
- ($`\Leftarrow`$) Equal counts imply equal lengths: $`|s| = \sum_c \mathrm{cnt}_s(c) = \sum_c \mathrm{cnt}_t(c) = |t|`$. For each letter $`c`$, map the position $`j`$ of the $`r`$-th occurrence of $`c`$ in $`t`$ to the position of the $`r`$-th occurrence of $`c`$ in $`s`$, and call it $`\pi(j)`$. Because $`c`$ occurs equally often in both strings, this is a bijection within each letter; the position sets of different letters are disjoint and cover all positions, so $`\pi`$ is a permutation, and $`t[j] = c = s[\pi(j)]`$.

"Equal counts" is clearly reflexive, symmetric, and transitive, so $`\sim`$ is an equivalence relation, and so is $`\equiv`$. By the correspondence between equivalence relations and partitions (CLRS Appendix B.2, Theorem B.1), the classes of $`\equiv`$ partition $`\{0, \ldots, n-1\}`$ into disjoint groups whose union is the whole set.

**Lemma 1**: $`\mathrm{SORT}(s) = \mathrm{SORT}(t) \iff s \sim t`$.

*Proof*: Sorting only rearranges letters and does not change counts, so $`\mathrm{cnt}_{\mathrm{SORT}(s)} = \mathrm{cnt}_s`$. On the other hand, a string in ascending letter order is determined by its counts:

```math
\mathrm{SORT}(s) = \mathtt{a}^{\mathrm{cnt}_s(\mathtt{a})}\, \mathtt{b}^{\mathrm{cnt}_s(\mathtt{b})} \cdots \mathtt{z}^{\mathrm{cnt}_s(\mathtt{z})}
```

where $`c^m`$ means $`m`$ copies of $`c`$ written in a row.

- ($`\Leftarrow`$) If $`s \sim t`$, Lemma 0 gives $`\mathrm{cnt}_s = \mathrm{cnt}_t`$, and substituting into the formula gives $`\mathrm{SORT}(s) = \mathrm{SORT}(t)`$.
- ($`\Rightarrow`$) If $`\mathrm{SORT}(s) = \mathrm{SORT}(t)`$, then $`\mathrm{cnt}_s = \mathrm{cnt}_{\mathrm{SORT}(s)} = \mathrm{cnt}_{\mathrm{SORT}(t)} = \mathrm{cnt}_t`$, and Lemma 0 gives $`s \sim t`$.

**Lemma 2 (grouping by key)**: Let the key function $`f`$ satisfy $`f(s) = f(t) \iff s \sim t`$, and let the algorithm group by $`f`$. Then the output is exactly the set of equivalence classes of $`\equiv`$, and each group lists its strings in input order.

**Loop Invariant**

At the start of each iteration of the first `for` loop (after processing $`s_0, \ldots, s_{i-1}`$):

**Invariant I: Key Set**

The key set of $`T`$ is exactly $`\{f(s_j) : 0 \le j < i\}`$.

**Invariant II: Group Contents**

For every key $`k`$ in $`T`$, $`T[k]`$ is $`\langle s_j : 0 \le j < i,\ f(s_j) = k \rangle`$, the strings among the first $`i`$ whose key is $`k`$, in increasing index order.

**Initialization**

With $`i = 0`$, $`T`$ is empty and both sets are empty, so the invariants hold.

**Maintenance**

Assume the invariants hold before processing $`s_i`$, and let $`k = f(s_i)`$.

- **Case 1: $`k \notin T`$**. The algorithm creates $`T[k] = \langle s_i \rangle`$. By Invariant I, none of the first $`i`$ strings has key $`k`$, so $`\langle s_i \rangle`$ is exactly the sequence of strings with key $`k`$ among the first $`i+1`$; the other lists are unchanged. The key set gains $`f(s_i)`$, and both invariants hold.
- **Case 2: $`k \in T`$**. The algorithm appends $`s_i`$ to the end of $`T[k]`$. The old list holds the strings with key $`k`$ among the first $`i`$, in increasing index order; appending $`s_i`$, whose index is larger, gives the strings with key $`k`$ among the first $`i+1`$, still in increasing index order. The other lists and the key set are unchanged, so the invariants hold.

**Termination**

The loop ends with $`i = n`$, when each key $`k`$ of $`T`$ maps to $`\langle s_j : f(s_j) = k \rangle`$. The second loop copies every list into `res` unchanged.

- **Exactly once**: each index $`j`$ belongs to the list of key $`f(s_j)`$ and to no other list.
- **Same group implies equivalent**: strings $`s_j, s_{j'}`$ in one group have $`f(s_j) = f(s_{j'})`$, so the "$`\Rightarrow`$" direction of $`f`$ gives $`s_j \sim s_{j'}`$.
- **Equivalent implies same group**: if $`s_j \sim s_{j'}`$, the "$`\Leftarrow`$" direction gives $`f(s_j) = f(s_{j'})`$, so both are in the same list.

Thus every list is one whole equivalence class, and each class corresponds to exactly one key.

#### 3. Conclusion

By Lemma 1 the sorted key satisfies the condition of Lemma 2, so on every valid input the algorithm outputs exactly the equivalence classes of $`\equiv`$, as the problem requires. Go `map` iteration order is random, which only affects the order of the groups, and the problem accepts any order; Invariant II guarantees that each group keeps input order.

### 4.5 Complexity Analysis

Let $`n`$ be the number of strings, $`\ell_i = |s_i|`$, $`L = \sum_i \ell_i`$ the total length, $`k = \max_i \ell_i`$ (here $`k \le 100`$), and $`g`$ the number of groups ($`g \le n`$). "Expected" below refers to the expected cost of a single hash-table operation (see the [overview §2.1](../../README.md)).

**Time Complexity: expected O(n k log k)**

Steps for one string $`s_i`$:

- Copying with `[]byte(s)` and converting with `string(b)`: $`O(\ell_i)`$ each.
- `slices.Sort(b)`: Go's `slices.Sort` is pattern-defeating quicksort (pdqsort), which switches to heapsort when recursion gets too deep, so the worst case is $`O(\ell_i \log \ell_i)`$.
- Lookup and append: hashing a string key reads the whole key, $`O(\ell_i)`$; an expected $`O(1)`$ key comparisons cost $`O(\ell_i)`$ each; `append` is amortized $`O(1)`$. Together, expected $`O(\ell_i)`$.

The second loop visits $`g`$ keys in $`O(g)`$. The expected total is

```math
O\Bigl(n + \sum_{i} \ell_i \log \ell_i\Bigr) \le O(n + L \log k) \le O(n k \log k)
```

The separate $`n`$ term is there because an empty string has $`\ell_i = 0`$ yet still costs one lookup.

**Space Complexity: O(n + L)**

- Hash-table keys: $`g`$ distinct sorted strings whose total length is at most $`L`$, so $`O(L)`$.
- Hash-table entries and lists: $`g`$ entries holding $`n`$ string headers in total, $`O(n)`$. These lists are the output; the string contents are shared with the input, not copied.
- Temporary slice `b`: $`O(k)`$, allocated anew each iteration, so the old one can be collected.
- `res`: $`g`$ slice headers, $`O(g)`$.

The total space including the output is $`O(n + L)`$. Excluding the output (the $`n`$ string headers and `res`), the keys still need $`O(g + L)`$, which is $`O(n + L)`$ in the worst case where no two strings are anagrams.

## 5. Learning Notes

### Challenge 1: What makes a key valid for grouping?

Grouping means computing a key for each string and putting strings with equal keys together. A wrong key fails silently, and typical test cases may not catch it.

**Understanding Process**

The key function $`f`$ must satisfy both directions. If only "$`s \sim t \Rightarrow f(s) = f(t)`$" holds, no group is split but different groups may merge: using the **set** of letters as the key wrongly merges `"aab"` and `"abb"`. If only "$`f(s) = f(t) \Rightarrow s \sim t`$" holds, nothing is wrongly merged but a group may be split: using the string itself as the key separates `"eat"` and `"tea"`. The sorted string satisfies both directions (Lemma 1), so it is a valid key. The test case `Same letter set, different multiplicities` checks for the first mistake.

**Transferable Insight:** When designing a grouping key, check it as a **complete invariant** of the equivalence relation: equivalent objects must have equal keys, and objects with equal keys must be equivalent. Keys checked in only one direction are a common source of bugs.

### Challenge 2: Why can't Go use the sorted `[]byte` as the key directly?

A Go `map` key must be of a comparable type. Slices are not comparable, so `map[[]byte][]string` does not compile.

**Understanding Process**

That is why the code converts `b` with `string(b)` before using it as a key. The conversion copies the bytes, so the key does not share memory with `b`, and later changes to `b` cannot affect keys in the table. Likewise, `[]byte(s)` copies, so sorting never touches the input string. Another option is a fixed-length array as the key, since arrays are comparable; that is how the [count-key solution](../count-key/summary.md) works.

**Transferable Insight:** Before using compound data as a Go `map` key, check that it is comparable: strings, arrays, and structs with only comparable fields are; slices, maps, and functions are not. Non-comparable data must first be encoded as a string or a fixed-length array.

### Challenge 3: Why not compare strings pairwise?

The most direct approach is to compare each string against a representative of every existing group, which is exactly the reference implementation `bruteForce` in the tests.

**Understanding Process**

Each string may be compared with up to $`g`$ representatives at $`O(k)`$ each, for $`O(n g k)`$ total, or $`O(n^2 k)`$ in the worst case. Once each string is reduced to a canonical form, finding its group changes from "compare with every representative" to "one table lookup", which finds the group in expected constant lookups.

**Transferable Insight:** To classify objects under an equivalence relation, find a canonical form and look it up in a hash table; this replaces quadratic pairwise comparisons with a linear number of lookups.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Canonical Form**: here, the sorted string is the unique representative of each anagram class. It applies to testing equality and grouping of objects considered "the same under some transformation", such as permutations (sort), cyclic shifts (take the lexicographically smallest rotation), or uniform translation (subtract the first element). The canonical form must be a complete invariant and affordable to compute.
2. **Group by Key**: here, one pass with `map[key][]element` builds the groups while keeping input order inside each group. It applies to any equivalence classification where a complete invariant can be computed; it only groups, cannot answer ordered questions such as "which group is larger", and returns groups in no particular order.
3. **Equivalence Classes**: here, proving first that "being anagrams" is an equivalence relation is what makes the grouping a well-defined partition. It applies whenever one must check that "grouping" makes sense; if the relation is not transitive (for example, "differ in one letter"), grouping by key does not work, and union-find or graph connected components are needed instead.
4. **Sorting as Normalization**: here, sorting erases letter positions and keeps only counts. It applies to comparisons that care about multisets, not order; it costs $`O(k \log k)`$, and with a small fixed alphabet, counting can replace it (see the [count-key solution](../count-key/summary.md)).

### CLRS Connections

Section numbers follow the third edition of CLRS. "Connection" is either "Direct" or "Related idea"; when no specific section applies, the "CLRS Section" column says "No matching section" and only the concept name is kept.

| Concept | CLRS Section | Connection | Where It Appears |
|---|---|---|---|
| Loop invariant | §2.1 Insertion sort | Direct | Lemma 2 uses the invariant "the key set and each group match exactly the first $`i`$ strings", proved by initialization, maintenance, and termination, to show each string lands in exactly one group |
| Hash tables | §11.2 Hash tables | Direct | `groups` is keyed by the sorted string; each string costs one lookup and one append, expected $`O(\ell_i)`$ |
| Equivalence relations and partitions | Appendix B.2 Relations | Direct | Lemma 0 proves "being anagrams" is an equivalence relation; by Theorem B.1 its classes form a partition, and the output groups are exactly those classes |

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

- [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/) — Decides whether two strings are anagrams, the single-pair version of Lemmas 0 and 1; this problem turns that test into a key for grouping everything at once.
- [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/) — Also characterizes anagrams by letter counts, but must maintain the counts incrementally over a sliding window instead of re-sorting each window.
- [249. Group Shifted Strings](https://leetcode.com/problems/group-shifted-strings/) (premium) — The same "canonical form + group by key" pattern with a different relation, "every letter shifted cyclically by the same amount"; the canonical form is the string shifted so that its first letter becomes `a`.

---

**One-Sentence Summary**

Anagrams are identical after sorting and strings that differ after sorting are never anagrams, so using the sorted string as a hash-table key puts every string into its unique group in one pass.
