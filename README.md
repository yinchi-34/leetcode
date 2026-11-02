# leetcode-go

[![Go](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go&logoColor=white)](go.mod)
[![Code: MIT](https://img.shields.io/badge/code-MIT-blue)](LICENSE)
[![Docs: CC BY 4.0](https://img.shields.io/badge/docs-CC%20BY%204.0-lightgrey)](LICENSE-docs)

English | [简体中文](README.zh-CN.md)

Go solutions and learning notes for LeetCode, organized by algorithm and data structure. The problem notes record:

- **How the approach emerged**: the first idea, where it fell short, and the key change in thinking.
- **Why the solution works**: correctness proofs using loop invariants, induction, exchange arguments, or other suitable methods.
- **Where the difficulty lies**: separate ratings for modeling, knowledge prerequisites, correctness reasoning, and implementation, with reasons for each score.
- **What transfers to other problems**: reusable methods and the conditions under which they apply.

See the [problem index](docs/PROBLEMS.md) for the full list.

## Reading a Problem

The [problem-note templates](templates/) use the following structure. Read the sections that match your current needs:

| Section | When to read it |
| --- | --- |
| 1. Problem information and difficulty ratings | To understand the problem and the requirements of the solution |
| 2. Problem understanding | When the goal or constraints are still unclear |
| 3. Initial approach | To compare your reasoning with the recorded thought process |
| 4. Algorithm specification: pseudocode, correctness, and complexity | To check the steps, proof, and computational cost |
| 5. Learning notes | When you get stuck or want to understand a mistake |
| 6. Key takeaways | To identify methods you can reuse |
| 7. Review log | To test your understanding after some time |
| 8. Related problems | To find further practice |

The four-dimensional rating evaluates **a specific solution under stated constraints**. Each dimension receives an independent score from 1 to 5; the scores are not summed. LeetCode official difficulty and the zerotrac rating are recorded separately. See the [rating criteria](docs/RATING.md).

## Repository Structure

Browse [algorithms](Algorithms/), [data structures](data-structures/), or [introductory exercises](primers/).

<details>
<summary>Show the full directory tree</summary>

```text
.
├── Algorithms/                  # Problems grouped by algorithm
│   └── <category>/<number>-<english-title>/
│       ├── solution.go
│       ├── solution_test.go
│       ├── summary.md           # English notes, when available
│       └── summary.zh-CN.md     # Chinese notes, when available
├── data-structures/             # Problems grouped by data structure
│   └── <category>/<number>-<english-title>/
├── primers/                     # Introductory exercises: code and tests
├── templates/                   # English and Chinese problem-note templates
├── docs/                        # Rating criteria and documentation conventions
│   ├── RATING.md                # Rating criteria in English
│   ├── RATING.zh-CN.md          # Rating criteria in Chinese
│   ├── CLRS_INDEX.md            # CLRS reverse index in English
│   ├── CLRS_INDEX.zh-CN.md      # CLRS reverse index in Chinese
│   ├── PROBLEMS.md              # Problem index in English
│   ├── PROBLEMS.zh-CN.md        # Problem index in Chinese
│   ├── GLOSSARY.md              # Shared Chinese–English glossary
│   ├── GITHUB_MATH.md           # GitHub math conventions
│   ├── PROMPT.zh-CN.md          # Prompt for writing summaries (Chinese)
│   └── investigations/          # Investigation records
│       ├── README.md
│       ├── README.zh-CN.md
│       └── github-math-rendering-investigation.md
├── README.md
└── README.zh-CN.md
```

</details>

## Running Tests

Requires Go 1.22 or later, as declared in [go.mod](go.mod). Run the following commands from the repository root:

```bash
go test ./...
go test ./Algorithms/greedy-algorithms/...
```

The first command tests all packages; the second tests one category.

## Adding a Problem

1. Create a directory under `Algorithms/<category>/` or `data-structures/<category>/`, named `<number>-<english-title>` in lowercase with hyphens. Introductory exercises go under `primers/`.
2. Add `solution.go` and `solution_test.go`, both using `package solution`. Test the examples, boundary cases, and cases that exercise the key branches of the algorithm.
3. When writing notes, copy the [English template](templates/summary.md) and [Chinese template](templates/summary.zh-CN.md) into the problem directory. Change `../docs/` to `../../../docs/` in the copied files, then replace the placeholders and remove the copying instructions in the opening comment.
4. Rate the solution using the [rating criteria](docs/RATING.md), stating its constraints and target complexity. Keep the facts, scores, and mathematical meaning consistent between the two languages.
5. Check terminology against the [shared glossary](docs/GLOSSARY.md) and mathematical notation against the [GitHub math conventions](docs/GITHUB_MATH.md). Update the [problem index](docs/PROBLEMS.md) and the [CLRS index](docs/CLRS_INDEX.md), verify the relative links, and run `go test ./...`.

## Documentation Conventions

- All CLRS section numbers cited in this repository refer to the **3rd edition** (Cormen, Leiserson, Rivest, Stein, *Introduction to Algorithms*, 3rd ed., MIT Press, 2009).
- [CLRS index](docs/CLRS_INDEX.md): problems grouped by CLRS section.
- [Shared glossary](docs/GLOSSARY.md): canonical Chinese–English terminology and definitions.
- [GitHub math conventions](docs/GITHUB_MATH.md): notation, layout, and rendering verification.

## Investigation Records

[Investigation records](docs/investigations/README.md) document problem diagnosis, comparison experiments, and validation results. The collection currently includes the [GitHub math rendering investigation](docs/investigations/github-math-rendering-investigation.md).

## License

- Source code (`*.go` files, `go.mod`, and code snippets inside the documents) is licensed under the [MIT License](LICENSE).
- Documentation, solution write-ups, and notes (all `*.md` files) are licensed under [CC BY 4.0](LICENSE-docs). You may share and adapt them, as long as you give attribution and link to this repository.
