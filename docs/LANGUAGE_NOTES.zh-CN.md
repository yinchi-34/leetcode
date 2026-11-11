# Go 与 Python 实现差异

[English](LANGUAGE_NOTES.md) | 简体中文

本仓库每道题都有 Go 和 Python 两种实现，算法相同，写法不同。本文按主题列出容易出错的差异。每条都给出差异本身、一个最小例子，以及踩过这个坑的题目。题解第 4.6 节只写本题涉及的差异，并链接到这里。

## 整数范围

Go 的 `int` 在 64 位平台上是 64 位，会静默溢出。Python 的 `int` 没有上限，不会溢出。

- 平方和、乘积、前缀和在 Go 中要先估算范围，必要时用 `int64` 或取模；Python 不用。
- 反过来，题目要求“按 32 位溢出”的行为时，Python 要手动模拟。

## 负数的除法与取模

Go 向零截断，Python 向下取整。

| 表达式 | Go | Python |
| --- | --- | --- |
| `-7 / 2` | `-3` | `-7 // 2 == -4` |
| `-7 % 2` | `-1` | `1` |

需要非负余数时，Go 写 `((x % m) + m) % m`，Python 直接写 `x % m`。需要向零截断时，Python 写 `int(a / b)`，但这只在不超出浮点精度时正确。

## 切片与列表：别名和拷贝

Go 切片和 Python 列表都是对底层数据的引用，赋值不会复制。

- Go 复制用 `append([]int(nil), a...)`，Python 用 `a.copy()` 或 `a[:]`。
- 二维初始化：Python 的 `[[0] * n] * m` 让 m 行共享同一个列表，要写成 `[[0] * n for _ in range(m)]`。
- Go 的 `append` 可能写入调用方切片的底层数组，Python 的 `append` 只改这一个列表。
- 例：[0462 排序解法](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/sorting/)，Go 要手动复制，Python 的 `sorted()` 本身返回新列表。

## 哈希表

- Go `map` 的遍历顺序是随机的，Python `dict` 按插入顺序遍历。测试不能依赖 Go map 的顺序。
- 键：Go 的数组 `[26]int` 可以作键，切片不行。Python 的 `tuple` 可以作键，`list` 不行。例：[0049 计数键](../data-structures/hash-table/0049-group-anagrams/count-key/)。
- 计数：Go 访问不存在的键得到零值，Python 用 `dict.get(k, 0)`、`collections.Counter` 或 `defaultdict`。

## 字符串

- Go 的 `s[i]` 是 `byte`，`for _, r := range s` 得到 `rune`。Python 的 `s[i]` 是长度为 1 的 `str`。
- 字母转下标：Go 写 `s[i] - 'a'`，Python 写 `ord(c) - ord('a')`。
- 两者的字符串都不可变。逐字符拼接时，Go 用 `strings.Builder` 或 `[]byte`，Python 用列表收集再 `''.join()`。

## 递归深度

Go 的栈会自动增长，递归几十万层通常没问题。CPython 默认递归上限是 1000 层，并且不做尾调用优化。

- 递归深度可能达到 n 时（链状的树、退化的快速选择、网格 DFS），Python 改写成循环或显式栈，或者调用 `sys.setrecursionlimit`。
- 例：[0462 快速选择](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/lomuto-quickselect/)的 Python 版把尾递归改成了 `while` 循环。

## 排序

- Go 用 `sort.Ints`、`slices.Sort` 或 `sort.Slice(a, less)`。Python 用 `a.sort()` 原地排序或 `sorted(a)` 返回新列表，多关键字时写 `key=lambda x: (x[0], -x[1])`。
- Python 的排序是稳定的。Go 的 `sort.Slice` 不稳定，需要稳定时用 `sort.SliceStable`。

## 堆、队列与栈

- 堆：Go 用 `container/heap`，要自己实现接口。Python 用 `heapq`，只有最小堆，最大堆可以存相反数。
- 队列：Go 用切片 `q = q[1:]`。Python 用 `collections.deque` 的 `popleft()`，不要用 `list.pop(0)`，它是 $`O(n)`$。
- 栈：两者都用切片或列表的末尾。

## 性能

同样的算法，Python 通常比 Go 慢几十倍。复杂度够用但常数大时（例如 $`10^7`$ 次循环），Python 可能超时。优先使用内置函数（`sum`、`sorted`、`collections`），它们在 C 层实现。
