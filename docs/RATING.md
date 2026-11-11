# Problem Difficulty Rating

This repository uses three models to describe problem difficulty: the **zerotrac rating**, **LeetCode official difficulty**, and a **four-dimensional rating**.

The three results are recorded independently. Each model retains its own meaning; the scores are neither converted into one another nor combined into a single overall score.

## Rating Models

| Model | Result | Evaluates | Source |
| --- | --- | --- | --- |
| zerotrac rating | Numerical difficulty rating | Problem | LeetCode Problem Rating |
| LeetCode official difficulty | Easy / Medium / Hard | Problem | Official LeetCode problem page |
| Four-dimensional rating | Four independent scores from 1 to 5 | A specific solution under defined constraints | Criteria defined in this document |

The zerotrac rating and official difficulty are recorded once for each problem.

The four-dimensional rating must specify the solution being evaluated. If a problem contains multiple solutions, each solution may be rated separately.

## zerotrac Rating

This repository uses the difficulty ratings published by [LeetCode Problem Rating][zerotrac]. A higher numerical rating indicates that the model estimates the problem to be more difficult.

According to the author's [explanation of the rating method][rating-method], the model uses contestants' contest ratings and problem-solving results to estimate problem difficulty through an Elo-based model and maximum likelihood estimation.

Let the contestant rating be $A$ and the problem difficulty rating be $B$. The predicted probability of solving the problem is:

$$
P(\text{solve}) = \frac{1}{1 + 10^{(B-A)/400}}
$$

The value of $B$ is estimated from actual contest results by finding the difficulty value that best explains the observed outcomes.

When the contestant rating and problem difficulty rating are equal, the model predicts a 50% probability of solving the problem.

This repository directly references the published ratings and preserves links to the original data source. Only two states are recorded: if the data contains the problem, record the score rounded to two decimal places; otherwise, write "No data available".

The rating backend of the project is not open source. The estimated ratings are intended for comparing relative problem difficulty and do not represent an individual's actual probability of solving a problem.

## LeetCode Official Difficulty

This repository records the official difficulty label displayed on each LeetCode problem page:

| Chinese Label | English |
| --- | --- |
| 简单 | Easy |
| 中等 | Medium |
| 困难 | Hard |

The problem link is preserved, and the difficulty label is recorded exactly as provided by LeetCode.

If multiple solutions are included for the same problem, the official difficulty remains unchanged.

## Four-Dimensional Rating

The four-dimensional rating describes where the difficulty of a specific solution comes from.

It evaluates four aspects:

- modeling,
- knowledge prerequisites,
- correctness reasoning,
- implementation.

This rating system is defined specifically for this repository.

### Rating Conventions

- **Scope**: Clearly specify the problem, solution, input constraints, target complexity, and implementation language.
- **Baseline knowledge**: The reader is assumed to understand basic Go and Python syntax, arrays, slices and lists, loops, and functions.
- **Score range**: Each dimension receives an integer score from 1 to 5. Higher scores indicate greater difficulty. The scores are recorded independently and are not summed.
- **Rating basis**: Choose the level that best matches the actual requirements of the solution and provide one concrete sentence explaining the rating.
- **Consistency**: Ratings are based on the requirements of the solution rather than the evaluator's personal familiarity with the technique.

An algorithm name, source-code length, or time complexity does not determine the rating by itself.

For example, using dynamic programming does not necessarily imply a high knowledge requirement, and the possibility of integer overflow does not necessarily imply high implementation difficulty.

### Modeling Difficulty

How much abstraction or transformation is required to turn the problem statement into an algorithmic model?

| Score | Criteria |
| --- | --- |
| 1 | The problem can be solved by directly simulating the description without additional transformation |
| 2 | Requires one straightforward transformation, such as counting, sorting, or representing data as intervals |
| 3 | Requires discovering a key relationship or designing a state representation that captures the problem |
| 4 | Requires multiple transformations or combining several properties or models |
| 5 | Requires a non-obvious construction, equivalence transformation, or complex state design |

The explanation should identify the key transformation.

For example:

> Transform a continuous range-sum query into the difference between two prefix sums.

### Knowledge Prerequisites

What prior knowledge is required to understand and use the solution?

| Score | Criteria |
| --- | --- |
| 1 | Basic syntax, arrays, loops, and functions are sufficient |
| 2 | Requires one common data structure or technique, such as a hash map, stack, or prefix sum |
| 3 | Requires a systematic understanding of one algorithmic technique or a small combination of common techniques |
| 4 | Requires multiple algorithms working together, or relatively specialized data structures or mathematical knowledge |
| 5 | Requires a long chain of prerequisites or deep understanding of multiple specialized topics |

The explanation should list the actual knowledge required rather than relying on broad labels such as "graph theory" or "dynamic programming".

### Correctness Reasoning Difficulty

How difficult is it to explain why the solution is correct for all valid inputs?

| Score | Criteria |
| --- | --- |
| 1 | Correctness follows directly from the definition or exhaustive enumeration |
| 2 | Requires simple case analysis, a local relationship, or an intuitive invariant |
| 3 | Requires an explicit invariant, induction relation, or proof that the solution is complete |
| 4 | Requires multiple lemmas, a relatively complex exchange argument, or multi-level induction |
| 5 | Requires a long proof chain, complex case analysis, or a global property that is difficult to observe directly |

The score is based on the complexity of the reasoning rather than the name of the proof technique.

Direct proof, contradiction, induction, loop invariants, and exchange arguments may all be used.

### Implementation Difficulty

Assuming the algorithm is already understood, how difficult is it to implement correctly in Go?

Every problem also has a Python implementation, but this score is rated for the Go version only, so that scores stay comparable across problems. Section 4.6 of each note describes where Python is easier or harder, for example no overflow but a recursion limit.

| Score | Criteria |
| --- | --- |
| 1 | Few states, straightforward control flow, and simple boundary conditions |
| 2 | Requires handling a small number of indices, branches, or boundary cases |
| 3 | Multiple states must work together, and update order, deduplication, or recursion boundaries are error-prone |
| 4 | Multiple data structures interact, state transitions are complex, and several types of boundary conditions must be coordinated |
| 5 | The implementation has a long dependency chain, multiple layers of interacting state, and local changes can easily break the overall logic |

The explanation should identify concrete implementation challenges, such as slice sharing, state update order, interval boundary semantics, or integer range handling.

## Missing Values

| Status | Meaning |
| --- | --- |
| No data available | The zerotrac data has been checked, but it does not contain the problem |
| To be rated | A four-dimensional rating has not yet been assigned |

Every problem should retain fields for all three rating models.

Missing data should use one of the statuses above. Do not use `0`, and do not infer a missing value from another rating model.

## Template

The following template can be copied directly into a problem document.

If multiple solutions are included, repeat the **Four-Dimensional Rating** section for each solution.

```markdown
## Difficulty Rating

- Problem: [Problem Number and Name](problem-link)
- zerotrac rating: Score / No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode official difficulty: Easy / Medium / Hard

### Four-Dimensional Rating

- Solution:
- Input constraints:
- Target complexity: Time O(...), extra space O(...)
- Implementation language: Go

| Dimension | Score | Reason |
| --- | --- | --- |
| Modeling difficulty | To be rated | |
| Knowledge prerequisites | To be rated | |
| Correctness reasoning difficulty | To be rated | |
| Implementation difficulty | To be rated | |
```

## References

- [LeetCode Problem Rating][zerotrac]: zerotrac rating data and project documentation.
- [LeetCode Contest Problem Difficulty Ranking][rating-method]: explanation of the rating model.

[zerotrac]: https://github.com/zerotrac/leetcode_problem_rating
[rating-method]: https://leetcode.cn/discuss/post/3157139/li-kou-jing-sai-ti-nan-du-pai-xing-bang-39wgi/