# 3411. Maximum Subarray With Equal Products

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [3411. Maximum Subarray With Equal Products](https://leetcode.com/problems/maximum-subarray-with-equal-products/)
- **Official Difficulty**: Easy
- **Primary Category**: Sliding Window
- **Secondary Category**: Number Theory
- **Tags**: `sliding window`, `number theory`, `prime factorization`, `greatest common divisor`, `pairwise coprime`
- **Code**: [solution.go](./solution.go) · [solution.py](./solution.py)
- **Tests**: [solution_test.go](./solution_test.go) · [solution_test.py](./solution_test.py)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: 1443.19 ([data source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Easy

**Four-Dimensional Rating — Pairwise-Coprime Sliding Window**

- Evaluation Scope: $`2 \le n \le 100`$, $`1 \le \mathrm{nums}[i] \le 10`$; time $`O(n \log C)`$ ($`C`$ bounds the window product), extra space $`O(1)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 4 | First split $`\mathrm{prod} = \mathrm{lcm} \cdot \gcd`$ by prime into "sum of exponents = max + min", then derive that for length at least 3 this is equivalent to pairwise coprimality, then notice that pairwise coprimality is closed under taking subarrays, which is what makes a sliding window work; all three steps are needed |
| Knowledge Prerequisites | 3 | The fundamental theorem of arithmetic, the exponent form of $`\gcd`$ and $`\mathrm{lcm}`$, Euclid's algorithm, and the sliding window; none is hard alone, but they have to be combined |
| Correctness Reasoning Difficulty | 4 | Two number-theory lemmas (length 2 always holds; for length at least 3 the condition equals pairwise coprimality), then a loop invariant showing the left pointer is always the left end of the longest pairwise coprime subarray ending at the right pointer, plus arguments that the division is exact and nothing overflows |
| Implementation Difficulty | 1 | Under 20 lines, one window and one product variable; starting `res` at 2 covers the length-2 case |

Ratings follow the repository's [RATING.md](../../../docs/RATING.md). Each dimension is rated independently and no total is computed.

## 2. Problem Understanding

Given an array `nums` of positive integers, a subarray `arr` is **product equivalent** if

```math
\mathrm{prod}(\mathrm{arr}) = \mathrm{lcm}(\mathrm{arr}) \cdot \gcd(\mathrm{arr})
```

where $`\mathrm{prod}`$ is the product of all elements and $`\gcd`$ and $`\mathrm{lcm}`$ are the greatest common divisor and least common multiple of all elements. Return the length of the longest product equivalent subarray. A subarray is a contiguous non-empty sequence.

**Example**

- Input: `nums = [1,2,1,2,1,1,1]`
- Output: `5` (subarray `[1,2,1,1,1]`: product 2, $`\gcd = 1`$, $`\mathrm{lcm} = 2`$)

- Input: `nums = [2,3,4,5,6]`
- Output: `3` (subarray `[3,4,5]`: product 60, $`\gcd = 1`$, $`\mathrm{lcm} = 60`$)

**Key Requirements**

1. $`2 \le n \le 100`$, $`1 \le \mathrm{nums}[i] \le 10`$.
2. The condition is one equation about the whole subarray. It cannot be checked pair by pair for adjacent elements, so a structural characterization is needed first.
3. A length-1 subarray $`[x]`$ qualifies exactly when $`x = x^2`$, that is, $`x = 1`$. Since $`n \ge 2`$, it is never the only optimum.

## 4. Algorithm Specification

**Algorithm Name:** Pairwise-Coprime Sliding Window

### 4.1 Input

$`A = \langle a_0, a_1, \ldots, a_{n-1} \rangle`$

such that:

- $`2 \le n \le 100`$
- $`1 \le a_i \le 10`$

Write $`A[l..r] = \langle a_l, \ldots, a_r \rangle`$ for $`0 \le l \le r < n`$ and $`P(l, r) = \prod_{i=l}^{r} a_i`$, with the empty product $`P(l, l-1) = 1`$.

### 4.2 Output

```math
\mathrm{res} = \max\bigl\{\, r - l + 1 \;:\; 0 \le l \le r < n,\ \mathrm{prod}(A[l..r]) = \mathrm{lcm}(A[l..r]) \cdot \gcd(A[l..r]) \,\bigr\}
```

### 4.3 Pseudocode

Indices are 0-based.

```text
MAX-LENGTH(A)
    res ← 2                              // every length-2 subarray qualifies (Lemma 1)
    prod ← 1                             // prod = P(left, r-1)
    left ← 0
    for r ← 0 to n - 1
        while GCD(prod, A[r]) > 1        // A[r] shares a prime with some window element
            prod ← prod / A[left]
            left ← left + 1
        prod ← prod · A[r]
        res ← max(res, r - left + 1)
    return res

GCD(a, b)
    while b ≠ 0
        (a, b) ← (b, a mod b)
    return a
```

**State Definitions**

- `left`, `r`: the current window is $`A[\mathrm{left}..r]`$. The code calls the right end `right`.
- `prod`: the product of all elements in the current window.
- `res`: the longest product equivalent subarray found so far.

### 4.4 Correctness

**Proof Method:** comparing prime exponents + loop invariant

#### I. Key Properties Required for Correctness

1. **Fundamental theorem of arithmetic** (a known theorem, not given by the problem): every positive integer is a unique product of prime powers, and two positive integers are equal exactly when every prime has the same exponent in both.
2. **Length 2 always qualifies** (Lemma 1): the answer is at least 2, so `res` can start at 2.
3. **For length at least 3 the condition equals pairwise coprimality** (Lemma 2): the whole-array equation becomes a pairwise relation.
4. **Pairwise coprimality is closed under taking subarrays**: every subarray of a pairwise coprime array is pairwise coprime. So the valid left ends for a fixed right end $`r`$ form a suffix interval, and the smallest valid left end never decreases as $`r`$ grows, which is what a sliding window needs.

#### II. Correctness Proof

For a positive integer $`m`$ and a prime $`p`$, let $`v_p(m)`$ be the exponent of $`p`$ in the prime factorization of $`m`$ (0 when $`p \nmid m`$).

**Lemma 1 (two numbers).** For all positive integers $`a, b`$, $`\gcd(a,b) \cdot \mathrm{lcm}(a,b) = a \cdot b`$.

*Proof.*

Step 1, factor into primes. Let $`a = \prod_p p^{x_p}`$ and $`b = \prod_p p^{y_p}`$, where $`p`$ ranges over all primes and exponents may be 0.

Step 2, write out $`\gcd`$ and $`\mathrm{lcm}`$. By definition:

```math
\gcd(a,b) = \prod_p p^{\min(x_p, y_p)}, \qquad \mathrm{lcm}(a,b) = \prod_p p^{\max(x_p, y_p)}
```

Step 3, multiply them. Powers of the same base multiply by adding exponents:

```math
\gcd(a,b) \cdot \mathrm{lcm}(a,b) = \prod_p p^{\min(x_p,y_p) + \max(x_p,y_p)}
```

Step 4, apply an identity. For any two real numbers, $`\min(x,y) + \max(x,y) = x + y`$: one term takes the smaller and the other the larger, so each number is used exactly once, only in a different order.

Conclusion: $`\gcd(a,b) \cdot \mathrm{lcm}(a,b) = \prod_p p^{x_p + y_p} = a \cdot b`$.

So every length-2 subarray is product equivalent. Since $`n \ge 2`$, one exists, and the answer is at least 2.

**Lemma 2 (three or more numbers).** Let $`k \ge 3`$ and $`b_1, \ldots, b_k`$ be positive integers. Then $`\prod_i b_i = \mathrm{lcm}(b_1, \ldots, b_k) \cdot \gcd(b_1, \ldots, b_k)`$ if and only if $`b_1, \ldots, b_k`$ are pairwise coprime.

*Proof.* Follow the four steps of Lemma 1. For each prime $`p`$ let $`e_i = v_p(b_i)`$. The exponents of $`p`$ in the product, the $`\mathrm{lcm}`$, and the $`\gcd`$ are $`\sum_i e_i`$, $`\max_i e_i`$, and $`\min_i e_i`$. By the fundamental theorem of arithmetic, the equation holds exactly when, for every prime $`p`$,

```math
\sum_{i=1}^{k} e_i = \max_i e_i + \min_i e_i
```

Fix $`p`$ and sort the exponents as $`e_{(1)} \le e_{(2)} \le \cdots \le e_{(k)}`$. Subtracting the right side from the left leaves the sum of the middle $`k-2`$ terms:

```math
\sum_{i=1}^{k} e_i - e_{(k)} - e_{(1)} = \sum_{j=2}^{k-1} e_{(j)}
```

So the equation is equivalent to $`\sum_{j=2}^{k-1} e_{(j)} = 0`$.

- **Necessity**: if the equation holds, the middle terms are non-negative and sum to 0, so all are 0. Because $`k \ge 3`$ there is at least one middle term, $`e_{(2)} = 0`$, hence $`e_{(1)} \le e_{(2)} = 0`$. So at most one of $`e_1, \ldots, e_k`$ (namely $`e_{(k)}`$) is positive: $`p`$ divides at most one $`b_i`$. Holding for every prime, this is pairwise coprimality.
- **Sufficiency**: if the numbers are pairwise coprime, each prime divides at most one $`b_i`$, so $`e_{(1)} = \cdots = e_{(k-1)} = 0`$, the middle sum is 0, and the equation holds.

When $`k = 2`$ there are no middle terms and the equation holds unconditionally, which is Lemma 1. The assumption $`k \ge 3`$ is used only in the necessity direction, where it forces the minimum to be 0.

**Corollary (equivalent form of the problem).** Let $`L^*`$ be the length of the longest pairwise coprime subarray of $`A`$ (a single element counts as pairwise coprime, so $`L^* \ge 1`$). The answer is $`\max(2, L^*)`$.

*Proof.* A product equivalent subarray of length 1 has length 1; all length-2 subarrays qualify by Lemma 1; and those of length at least 3 are exactly the pairwise coprime ones by Lemma 2. So the maximum length is $`\max(2, L^*)`$.

It remains to show the algorithm computes $`\max(2, L^*)`$. For each $`r`$ define

```math
\ell(r) = \min\{\, l \le r : A[l..r] \text{ is pairwise coprime} \,\}
```

The set is non-empty because $`A[r..r]`$ is pairwise coprime. By closure under subarrays, $`A[l..r]`$ is pairwise coprime exactly when $`\ell(r) \le l \le r`$. Also $`\ell(r) \le \ell(r+1)`$: $`A[\ell(r+1)..r+1]`$ is pairwise coprime, so dropping its last element leaves $`A[\ell(r+1)..r]`$ pairwise coprime.

**Loop Invariant**

At the start of outer iteration $`r`$ ($`0 \le r \le n`$, with $`r = n`$ meaning after the loop):

**Invariant I: Window**

$`\mathrm{left} = \ell(r-1)`$ (with $`\ell(-1) = 0`$), and $`\mathrm{prod} = P(\mathrm{left}, r-1)`$.

**Invariant II: Result**

$`\mathrm{res} = \max\bigl(2,\ \max_{0 \le j < r} (j - \ell(j) + 1)\bigr)`$.

**Initialization**

At $`r = 0`$: $`\mathrm{left} = 0 = \ell(-1)`$ and $`\mathrm{prod} = 1 = P(0, -1)`$; $`\mathrm{res} = 2`$ and the inner maximum is over the empty set. Both invariants hold.

**Maintenance**

Assume the invariants hold at the start of iteration $`r`$, and let $`x = a_r`$.

*Inner `while` loop.* Its invariant is "$`\mathrm{prod} = P(\mathrm{left}, r-1)`$, $`A[\mathrm{left}..r-1]`$ is pairwise coprime, and $`\mathrm{left} \le \ell(r)`$".

- It holds on entry by Invariant I, and $`\ell(r-1) \le \ell(r)`$ gives the last part.
- Key test: when $`A[\mathrm{left}..r-1]`$ is pairwise coprime, $`A[\mathrm{left}..r]`$ is pairwise coprime exactly when $`\gcd(\mathrm{prod}, x) = 1`$. The primes dividing $`\mathrm{prod}`$ are exactly the union of the primes dividing the window elements, so $`x`$ shares a prime with $`\mathrm{prod}`$ exactly when it shares a prime with some window element.
- If $`\gcd(\mathrm{prod}, x) > 1`$, then $`A[\mathrm{left}..r]`$ is not pairwise coprime, so $`\mathrm{left} < \ell(r)`$, and after incrementing $`\mathrm{left}`$ we still have $`\mathrm{left} \le \ell(r)`$. $`\mathrm{prod}`$ is $`a_{\mathrm{left}}`$ times the other elements, so `prod / A[left]` is exact and gives $`P(\mathrm{left}+1, r-1)`$; the shorter window is still pairwise coprime.
- Termination: when the window is empty ($`\mathrm{left} = r`$), $`\mathrm{prod} = 1`$ and $`\gcd(1, x) = 1`$, so the loop must exit. Hence $`\mathrm{left} \le r`$ and the loop runs at most $`r - \ell(r-1)`$ times. On exit $`\gcd(\mathrm{prod}, x) = 1`$, so by the key test $`A[\mathrm{left}..r]`$ is pairwise coprime and $`\mathrm{left} \ge \ell(r)`$. Together with $`\mathrm{left} \le \ell(r)`$, $`\mathrm{left} = \ell(r)`$.

*Rest of the body.* After `prod ← prod · x`, $`\mathrm{prod} = P(\ell(r), r)`$, so Invariant I holds at the next iteration. `res ← max(res, r - left + 1)` adds $`r - \ell(r) + 1`$ to the maximum, so Invariant II holds.

**Termination**

The outer loop ends after $`n`$ iterations. By Invariant II,

```math
\mathrm{res} = \max\Bigl(2,\ \max_{0 \le j < n} \bigl(j - \ell(j) + 1\bigr)\Bigr)
```

Every pairwise coprime subarray $`A[l..j]`$ has $`l \ge \ell(j)`$ and so length at most $`j - \ell(j) + 1`$, and $`A[\ell(j)..j]`$ itself is pairwise coprime. So the inner maximum is $`L^*`$, and $`\mathrm{res} = \max(2, L^*)`$, which is the answer by the corollary.

**No overflow.** The window is pairwise coprime, so each prime appears in at most one element and $`v_p(\mathrm{prod})`$ equals $`v_p(a_i)`$ for some $`a_i \le 10`$. The largest powers of 2, 3, 5, 7 in 1 to 10 are $`2^3, 3^2, 5, 7`$, so $`\mathrm{prod}`$ divides $`2^3 \cdot 3^2 \cdot 5 \cdot 7 = 2520`$. Before a multiplication $`\mathrm{prod} \le 2520`$, and after it at most $`25200`$, far inside Go's `int`.

#### III. Conclusion

Lemmas 1 and 2 turn "the product equals lcm times gcd" into "length 2, or pairwise coprime". Pairwise coprimality is closed under subarrays, so the longest pairwise coprime subarray ending at each position can be tracked with a left pointer that only moves forward, and one $`\gcd`$ between the window product and the new element decides whether it can join. The algorithm returns $`\max(2, L^*)`$ on every valid input, which is the length of the longest product equivalent subarray.

### 4.5 Complexity Analysis

Let $`n`$ be the array length and $`C = 25200`$ the bound on the numbers passed to $`\gcd`$ (see "No overflow" above).

**Time Complexity: O(n log C)**

The outer loop runs $`n`$ times. Each successful pass of the inner `while` increments `left`, which never decreases and never exceeds $`n`$, so by aggregate analysis the inner body runs at most $`n`$ times in total; each outer iteration also makes one failing test. That is at most $`2n`$ calls to $`\gcd`$, each taking $`O(\log C)`$ steps of Euclid's algorithm. Total time is $`O(n \log C)`$, which is $`O(n)`$ under this problem's constraints since $`C`$ is a constant.

**Space Complexity: O(1)**

Only the integers `res`, `prod`, `left`, and `right` are used; the input is neither modified nor copied.

### 4.6 Implementation Notes: Go and Python

- **Greatest common divisor**: Go's standard library has no integer $`\gcd`$, so `solution.go` writes Euclid's algorithm; Python uses `math.gcd`. See [Greatest Common Divisor](../../../docs/LANGUAGE_NOTES.md#greatest-common-divisor).
- **Exact division**: Go's `prod /= nums[left]` is integer division for positive integers. Python must use `prod //= nums[left]`; `/=` would turn `prod` into a float, and `math.gcd` then raises an error because its argument is not an integer. See [Division and Modulo of Negative Numbers](../../../docs/LANGUAGE_NOTES.md#division-and-modulo-of-negative-numbers).
- **Integer range**: by the "No overflow" argument the window product is at most 25200, so Go's `int` is enough and the two languages do not differ here.
- Everything else corresponds line by line.

## 5. Learning Notes

### Challenge 1: How does one whole-array equation become a pairwise coprime relation?

The condition involves three whole-array quantities at once, the product, the lcm, and the gcd, and it is not obvious what a qualifying subarray looks like.

**Understanding Process**

All three are per-prime statistics of the exponents: the product takes their sum, the lcm their maximum, the gcd their minimum. Split by prime, the single equation becomes independent exponent equations $`\sum e_i = \max e_i + \min e_i`$. For two numbers this is the identity $`\min + \max = x + y`$ (Lemma 1). For three or more, the extra middle terms must all be 0 and the minimum is forced to 0, so each prime may appear in only one element (Lemma 2).

**Transferable Insight:** When a condition mixes $`\gcd`$, $`\mathrm{lcm}`$, products, and divisibility, split it by prime into exponents. Multiplication becomes addition, and $`\gcd`$ and $`\mathrm{lcm}`$ become minimum and maximum, which can then be analyzed one prime at a time.

### Challenge 2: Why does a sliding window work?

Not every "longest subarray" problem admits two pointers.

**Understanding Process**

A sliding window needs the property to be closed under subarrays: if $`A[l..r]`$ has it, so does every $`A[l'..r']`$ with $`l \le l' \le r' \le r`$. Pairwise coprimality has this, because removing elements cannot create a new common prime. So the valid left ends for a right end $`r`$ form a suffix, and the smallest one, $`\ell(r)`$, never decreases as $`r`$ grows; the left pointer never moves back. The original "product equivalent" property is **not** closed under subarrays: `[2,3]` qualifies but its subarray `[2]` does not ($`2 \ne 2 \cdot 2`$); `[2,4]` qualifies but is not pairwise coprime, so extending it to the right only produces arrays that fail. That is why the transformation of Difficulty 1 has to come first.

**Transferable Insight:** Before using a sliding window, check that the property is closed under subarrays. If the original condition is not, see whether it can be turned into an equivalent closed condition, and handle the leftover special cases (here, length 2) separately.

### Challenge 3: How can the window test a new element in O(1)?

Checking the new element against every window element costs $`O(\text{window length})`$.

**Understanding Process**

When the window is pairwise coprime, the primes of the window product are exactly the union of the primes of its elements, so "$`x`$ is coprime to every window element" is the same as $`\gcd(\mathrm{prod}, x) = 1`$. Removing the left element is an exact division, because that element is coprime to the rest. The product is also naturally bounded (it divides 2520), so it cannot overflow.

**Transferable Insight:** If the "union" information of a window can be summarized by an invertible aggregate (here the product, inverted by exact division), elements can be added and removed in $`O(1)`$. The inverse must be valid: here exact division relies on the window being pairwise coprime.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Prime Exponent Decomposition**: turns $`\mathrm{prod} = \mathrm{lcm} \cdot \gcd`$ into $`\sum e_i = \max e_i + \min e_i`$ for each prime. Useful when a condition is built from products, $`\gcd`$, $`\mathrm{lcm}`$, and divisibility. It relies on the fundamental theorem of arithmetic and applies only to positive integers.
2. **Sliding Window over a Hereditary Property**: here it tracks the longest pairwise coprime subarray ending at $`r`$. Useful for finding the longest, or counting, subarrays with a property that is closed under subarrays; without closure the window is not monotone.
3. **Invertible Window Aggregate**: here the window product `prod` and one $`\gcd`$ decide whether a new element is coprime to the whole window, and removal is exact division. Useful when the window can be summarized by a quantity that supports both "add" and "undo", much like a prefix sum. The inverse must be exact (here guaranteed by pairwise coprimality) and the aggregate must not overflow.
4. **Small-Case Separation**: Lemma 2 only holds for length at least 3; length 2 is handled by Lemma 1, so `res` starts at 2. Useful whenever the general characterization breaks down or degenerates for small sizes.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. Relation is either "Direct" or "Related idea"; if no section fits, write "No matching section" and keep only the concept name.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Fundamental theorem of arithmetic, coprimality | §31.1 Elementary number-theoretic notions | Direct | Lemmas 1 and 2 compare the exponent of every prime on both sides via unique factorization, and characterize pairwise coprimality as "each prime divides at most one element" |
| Euclid's algorithm | §31.2 Greatest common divisor | Direct | `gcd(prod, x)` uses repeated remainders to decide whether a new element can join the window, in $`O(\log C)`$ steps |
| Loop invariant | §2.1 Insertion sort | Direct | The invariant "`left` is the left end of the longest pairwise coprime subarray ending at $`r-1`$" is proved by initialization, maintenance, and termination |
| Aggregate analysis | §17.1 Aggregate analysis | Direct | `left` never decreases, so the inner `while` body runs at most $`n`$ times in total instead of being bounded per iteration |
| Sliding window | No matching section | — | Left and right pointers move forward monotonically to track the longest valid window ending at $`r`$ |

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

- [2654. Minimum Number of Operations to Make All Array Elements Equal to 1](https://leetcode.com/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1/) — also maintains $`\gcd`$ over subarrays, but looks for the shortest subarray with $`\gcd = 1`$, a property closed under taking **superarrays** rather than subarrays; a useful contrast in the direction of monotonicity.
- [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) — the same sliding-window structure: "all elements distinct" and "all elements pairwise coprime" are both pairwise constraints closed under subarrays, and the left end shrinks until the new element no longer conflicts.

---

**One-Sentence Summary**

Split by prime, $`\mathrm{prod} = \mathrm{lcm} \cdot \gcd`$ for length at least 3 is exactly pairwise coprimality, which is closed under subarrays, so a sliding window with the window product and one $`\gcd`$ per step solves it; length 2 always qualifies, so the answer is at least 2.
