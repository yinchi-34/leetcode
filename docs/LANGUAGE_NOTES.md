# Go and Python Implementation Notes

English | [简体中文](LANGUAGE_NOTES.zh-CN.md)

Every problem in this repository has a Go and a Python implementation of the same algorithm. This page lists, by topic, the differences that cause bugs. Each entry gives the difference, a minimal example, and the problems where it mattered. Section 4.6 of each summary covers only the differences that problem touches and links here.

## Integer Range

Go's `int` is 64 bits on 64-bit platforms and overflows silently. Python's `int` is unbounded and never overflows.

- Sums of squares, products, and prefix sums need a range estimate in Go, and `int64` or a modulus where necessary. Python needs neither.
- Conversely, when a problem asks for 32-bit overflow behavior, Python has to simulate it.

## Division and Modulo of Negative Numbers

Go truncates toward zero; Python rounds toward negative infinity.

| Expression | Go | Python |
| --- | --- | --- |
| `-7 / 2` | `-3` | `-7 // 2 == -4` |
| `-7 % 2` | `-1` | `1` |

For a non-negative remainder, Go needs `((x % m) + m) % m`; Python's `x % m` already is one. For truncation in Python, `int(a / b)` works only while the values fit in a float exactly.

## Slices and Lists: Aliasing and Copying

Go slices and Python lists both refer to underlying data, so assignment does not copy.

- Copy with `append([]int(nil), a...)` in Go and `a.copy()` or `a[:]` in Python.
- 2D initialization: Python's `[[0] * n] * m` makes all m rows the same list. Write `[[0] * n for _ in range(m)]`.
- Go's `append` may write into the caller's backing array; Python's `append` only changes that one list.
- Example: [0462 sorting](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/sorting/). Go copies by hand, while Python's `sorted()` already returns a new list.

## Hash Tables

- Go randomizes `map` iteration order; Python's `dict` keeps insertion order. Tests must not depend on Go map order.
- Keys: a Go array `[26]int` can be a key but a slice cannot. A Python `tuple` can be a key but a `list` cannot. Example: [0049 count key](../data-structures/hash-table/0049-group-anagrams/count-key/).
- Counting: a missing key in Go yields the zero value. In Python use `dict.get(k, 0)`, `collections.Counter`, or `defaultdict`.

## Strings

- In Go, `s[i]` is a `byte` and `for _, r := range s` yields `rune`s. In Python, `s[i]` is a `str` of length 1.
- Letter to index: `s[i] - 'a'` in Go, `ord(c) - ord('a')` in Python.
- Strings are immutable in both. To build one character by character, use `strings.Builder` or `[]byte` in Go, and a list plus `''.join()` in Python.

## Recursion Depth

Go stacks grow on demand, so hundreds of thousands of levels usually work. CPython's default recursion limit is 1000, and it does not eliminate tail calls.

- When the depth can reach n (a path-shaped tree, a degenerate quickselect, a grid DFS), rewrite the Python version as a loop or with an explicit stack, or raise `sys.setrecursionlimit`.
- Example: the Python [0462 quickselect](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/lomuto-quickselect/) turns the tail recursion into a `while` loop.

## Sorting

- Go: `sort.Ints`, `slices.Sort`, or `sort.Slice(a, less)`. Python: `a.sort()` in place or `sorted(a)` for a new list, with `key=lambda x: (x[0], -x[1])` for several keys.
- Python's sort is stable. Go's `sort.Slice` is not; use `sort.SliceStable` when stability matters.

## Heaps, Queues, and Stacks

- Heap: Go's `container/heap` needs an interface implementation. Python's `heapq` is a min-heap only; store negated values for a max-heap.
- Queue: Go uses a slice with `q = q[1:]`. Python uses `collections.deque` and `popleft()`, never `list.pop(0)`, which is $`O(n)`$.
- Stack: the end of a slice or list in both.

## Performance

The same algorithm is often tens of times slower in Python than in Go. When the complexity fits but the constant is large (say $`10^7`$ loop iterations), Python may time out. Prefer built-ins (`sum`, `sorted`, `collections`), which run in C.
