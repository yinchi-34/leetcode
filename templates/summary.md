# {{Problem Number}}. {{Problem Title}}

English | [简体中文](./summary.zh-CN.md)

<!--
Template maintenance: copy this file and summary.zh-CN.md into
Algorithms/<category>/<problem>/ or data-structures/<category>/<problem>/.
In the copied files, change ../docs/ to ../../../docs/.
Code and test links refer to files in the
destination problem directory. Remove this comment from completed notes.
-->

Use the [shared glossary](../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [{{Problem Number and Title}}]({{Problem URL}})
- **Official Difficulty**: {{Easy / Medium / Hard}}
- **Primary Category**: {{Algorithm Category}}
- **Secondary Category**: {{Algorithm Category, optional}}
- **Tags**: `{{Tag 1}}`, `{{Tag 2}}`
- **Code**: [solution.go](./solution.go) · [solution.py](./solution.py)
- **Tests**: [solution_test.go](./solution_test.go) · [solution_test.py](./solution_test.py)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: {{Score, two decimal places / No data available}} ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: {{Easy / Medium / Hard}}

**Four-Dimensional Rating — {{Algorithm Name}}**

- Evaluation Scope: {{Input constraints}}; {{Target complexity}}; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | {{1–5}} | {{Reason}} |
| Knowledge Prerequisites | {{1–5}} | {{Reason}} |
| Correctness Reasoning Difficulty | {{1–5}} | {{Reason}} |
| Implementation Difficulty | {{1–5}} | {{Reason}} |

Ratings follow [RATING.md](../docs/RATING.md). Each dimension is evaluated independently; no overall score is calculated.

## 2. Problem Understanding

{{Describe the problem in your own words.}}

**Example**

- Input: `{{Example Input}}`
- Output: `{{Example Output}}`

**Key Requirements**

1. {{Input constraint}}
2. {{Problem requirement}}
3. {{Boundary condition}}

## 3. Initial Approach

### First Idea: {{Initial Approach}}

{{Describe your initial idea and why it seemed appropriate.}}

### Difficulties Encountered

{{Explain the limitations, mistakes, or unnecessary complexity of the initial approach.}}

### Thought Transformation

From:

> {{Original way of thinking}}

To:

> {{New way of thinking}}

{{Explain the key insight that led to this transformation and why it simplifies or improves the solution.}}

## 4. Algorithm Specification

**Algorithm Name:** {{Algorithm Name}}

### 4.1 Input

{{Formally define the input using mathematical notation, such as a sequence, set, graph, or another mathematical object.}}

For example:

$`A = \langle a_1, a_2, \ldots, a_n \rangle`$

Subject to:

- {{Input condition 1}}
- {{Input condition 2}}

### 4.2 Output

{{Formally define the expected output.}}

For example:

$`B = \langle b_1, b_2, \ldots, b_m \rangle`$

The output must satisfy:

- {{Postcondition 1}}
- {{Postcondition 2}}

### 4.3 Pseudocode

```text
ALGORITHM(input)
    {{Initialization}}

    {{Algorithm steps}}

    return res
```

**State Definitions**

- `{{Variable 1}}`: {{Meaning}}
- `{{Variable 2}}`: {{Meaning}}

### 4.4 Correctness

**Proof Method:** {{Choose an appropriate proof method}}

#### I. Key Properties Required for Correctness

The correctness of the algorithm depends on the following properties:

1. **{{Property 1}}**: {{Explanation}}
2. **{{Property 2}}**: {{Explanation}}
3. **{{Property 3}}**: {{Explanation}}

{{Explain how these properties support the correctness of the algorithm.}}

#### II. Correctness Proof

> Choose a proof method appropriate for the algorithm. The following structure is intended for algorithms proved using loop invariants. Replace it when another method is more suitable.

**Loop Invariant**

At the {{beginning / end}} of each iteration, the following properties hold:

**Invariant I: {{Name}}**

{{Precisely describe the property that remains unchanged throughout the iterations.}}

**Invariant II: {{Name, optional}}**

{{Describe the correctness property of the partial result or an additional state.}}

**Initialization**

{{Prove that the invariant holds before the first iteration.}}

**Maintenance**

Assume the invariant holds before the current iteration.

Consider the possible cases:

- **Case 1:** {{Explain why the invariant is preserved.}}
- **Case 2:** {{Explain why the invariant is preserved.}}

Therefore, the invariant continues to hold after the iteration.

**Termination**

{{Explain why the algorithm terminates.}}

{{Use the invariant and termination condition to prove that the final output satisfies the required postconditions.}}

#### III. Conclusion

{{Summarize why the algorithm produces the correct result for every valid input.}}

### 4.5 Complexity Analysis

Let $`n`$ denote the input size.

**Time Complexity: O({{...}})**

{{Derive the time complexity based on the number of iterations, recursive calls, or fundamental operations.}}

**Space Complexity: O({{...}})**

{{Analyze the memory required for variables, auxiliary data structures, recursion stacks, and output storage.}}

{{If necessary, distinguish auxiliary space from total space including the output.}}

### 4.6 Implementation Notes: Go and Python

{{Only what genuinely differs between the two languages in this problem, such as integer range, negative modulo, copying and aliasing, recursion depth, or library differences; link each point to the matching section of the [Go and Python implementation notes](../docs/LANGUAGE_NOTES.md). If nothing differs in a way that matters, write "The two implementations correspond line by line."}}

## 5. Learning Notes

### Challenge 1: {{Specific Difficulty}}

{{Describe the original confusion, incorrect assumption, or reasoning difficulty.}}

**Understanding Process**

{{Explain how you resolved the difficulty and what changed in your understanding.}}

**Transferable Insight:** {{Extract a generalizable lesson.}}

### Challenge 2: {{Specific Difficulty, Optional}}

{{Describe another meaningful challenge encountered while solving the problem.}}

**Understanding Process**

{{Explain the reasoning that led to a better understanding.}}

**Transferable Insight:** {{Summarize the reusable idea.}}

## 6. Key Takeaways

This problem illustrates the following reusable concepts:

1. **{{Concept 1}}**: {{Explanation}}
2. **{{Concept 2}}**: {{Explanation}}
3. **{{Concept 3}}**: {{Explanation}}

{{Explain when these concepts can be applied to other problems, including any important assumptions or limitations.}}

### CLRS Connections

Section numbers follow CLRS, 3rd edition. Relation is either "Direct" or "Related idea"; if no section fits, write "No matching section" and keep only the concept name.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| {{Concept}} | {{e.g. §16.2 Elements of the greedy strategy}} | {{Direct / Related idea}} | {{Which step of this problem}} |

## 7. Review Log

| Date | Independent Completion | Understanding Gained | Areas for Improvement |
|---|---|---|---|
| {{YYYY-MM-DD}} | {{Status}} | {{New Insight}} | {{Remaining Question}} |

### Next Review Checklist

- [ ] Can I implement the algorithm without looking at the solution?
- [ ] Can I accurately define the input and output?
- [ ] Can I write the pseudocode independently?
- [ ] Can I explain why the algorithm is correct?
- [ ] Can I derive its time and space complexity?
- [ ] Can I recognize the same algorithmic pattern in other problems?

## 8. Related Problems

- [{{Problem Number and Title}}]({{Problem URL}}) — {{Reason for the connection}}

---

**One-Sentence Summary**

{{Summarize the most important algorithmic insight gained from this problem.}}
