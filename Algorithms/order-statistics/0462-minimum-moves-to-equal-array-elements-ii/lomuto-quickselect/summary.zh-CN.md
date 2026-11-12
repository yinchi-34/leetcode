# 462. 最小操作次数使数组元素相等 II — 两路（Lomuto）划分快速选择

[English](./summary.md) | 简体中文

术语遵循[共享术语表](../../../../docs/GLOSSARY.md)，数学记号遵循仓库的 [GitHub 数学书写规范](../../../../docs/GITHUB_MATH.md)。

本题有三种解法，分别放在三个目录中，共用同一份测试用例：

- [**排序**](../sorting/summary.zh-CN.md)：排序后取中位数。题意和中位数最优性（引理 1、2）的证明见该文。
- [**三路划分快速选择**](../three-way-quickselect/summary.zh-CN.md)：随机快速选择，划分为小于、等于、大于三段。快速选择的归纳证明和期望时间分析见该文。
- **两路（Lomuto）划分快速选择**（本文）：CLRS 的 PARTITION，即最初写的版本。

本文只讨论它与三路划分版本的不同：划分方式，以及重复元素多时为什么会退化。

## 1. 题目信息

- **题目**：[462. Minimum Moves to Equal Array Elements II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/)
- **官方难度**：Medium（中等）
- **主分类**：Order Statistics（顺序统计量）
- **辅助分类**：Randomized Algorithms（随机化算法）
- **技术标签**：`顺序统计量`、`随机化算法`、`中位数`、`快速选择`、`Lomuto 划分`
- **代码**：[solution.go](./solution.go) · [solution.py](./solution.py)
- **测试**：[solution_test.go](../solution_test.go) · [solution_test.py](../solution_test.py)（放在上一级目录，三种解法共用）

### 难度评分

**外部评分**

- zerotrac 评分：暂无数据（[数据来源](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt)）
- 力扣官方难度：Medium

**四维评分 — 两路（Lomuto）划分随机快速选择**

- 评价范围：$`1 \le n \le 10^5`$，$`-10^9 \le a_i \le 10^9`$，答案在 32 位整数范围内；元素互不相同时期望时间 $`O(n)`$，存在大量重复元素时退化为 $`\Theta(n^2)`$；除递归栈外额外空间 $`O(1)`$；Go 实现。

| 维度 | 分数 | 评分理由 |
|---|---|---|
| 建模难度 | 2 | 需要把“最少操作次数”转化为“绝对偏差之和的最小值”，再转化为求第 $`\lfloor n/2 \rfloor`$ 小的数 |
| 知识门槛 | 3 | 需要掌握 CLRS 的 PARTITION、RANDOMIZED-PARTITION 与 RANDOMIZED-SELECT |
| 正确性论证难度 | 4 | 正确性同样需要中位数最优、划分不变式、对区间长度的归纳；时间分析还要指出期望线性的前提是元素互不相同 |
| 实现难度 | 2 | 照 CLRS 的写法即可；细节是 `i` 从 `l-1` 开始、最后把主元换到 `i+1` |

评分遵循仓库 [RATING.zh-CN.md](../../../../docs/RATING.zh-CN.md)，各维度独立评价，不合成总分。

## 2. 题意理解（Problem Understanding）

题意、示例与约束见[排序解法](../sorting/summary.zh-CN.md#2-题意理解problem-understanding)。与三路划分版本一样，核心子问题是求升序排列后下标 $`m = \lfloor n/2 \rfloor`$ 的元素 $`b_m`$。

## 3. 初始思路（Initial Approach）

### 第一次想到：照 CLRS 写随机快速选择

这是最初写的 Go 版本：`partition`、`randomPartition`、`quickSelect` 三个函数，分别对应 CLRS 的 PARTITION、RANDOMIZED-PARTITION 与 RANDOMIZED-SELECT。仓库中的代码与原版相同，只删去了 `rand.Seed`：它从 Go 1.20 起已废弃，而仓库使用 Go 1.22，全局随机源会自动播种。

### 遇到的问题

当时不知道这种划分叫 Lomuto 划分，也只想到了最坏 $`n^2`$。用 $`n = 10^5`$ 实测：随机数组约 36 万次比较、1 ms；**全相等数组约 37 亿次比较、7.7 秒**。

### 思维转变

从：

> 随机选主元就能避免最坏情况。

转变为：

> 两路划分把所有等于主元的元素都放进左边，重复值多时，随机化也救不了。

修正后的版本见[三路划分快速选择](../three-way-quickselect/summary.zh-CN.md)。本解法保留在仓库中，作为对照。

## 4. 算法表达（Algorithm Specification）

**算法名称：** 两路（Lomuto）划分随机快速选择（Randomized Quickselect with Lomuto Partitioning）

### 4.1 输入（Input）

$`A = \langle a_0, a_1, \ldots, a_{n-1} \rangle`$

满足：

- $`1 \le n \le 10^5`$
- $`a_i \in \mathbb{Z}`$，$`-10^9 \le a_i \le 10^9`$

### 4.2 输出（Output）

```math
\mathrm{res} = \sum_{i=0}^{n-1} \lvert a_i - b_m \rvert, \qquad m = \lfloor n/2 \rfloor .
```

### 4.3 伪代码（Pseudocode）

下标从 0 开始。

```text
PARTITION(A, l, r)                   // CLRS §7.1，主元 x = A[r]
    x ← A[r]
    i ← l - 1
    for j ← l to r - 1
        if A[j] ≤ x
            i ← i + 1
            swap A[i], A[j]
    swap A[i+1], A[r]
    return i + 1                     // A[l..q-1] ≤ x，A[q] = x，A[q+1..r] > x

QUICK-SELECT(A, l, r, k)             // 要求 l ≤ k ≤ r
    p ← 均匀随机取自 [l, r]
    swap A[p], A[r]
    q ← PARTITION(A, l, r)
    if q = k
        return A[q]
    if q < k
        return QUICK-SELECT(A, q + 1, r, k)
    return QUICK-SELECT(A, l, q - 1, k)

MIN-MOVES-2(A)
    x ← QUICK-SELECT(A, 0, n - 1, ⌊n/2⌋)
    res ← Σ |A[i] - x|
    return res
```

**状态定义**

- `i`：“$`\le x`$ 区”的最后一个位置。
- `j`：下一个待检查的位置。
- `q`：划分后主元的位置。

### 4.4 正确性（Correctness）

**证明方法：PARTITION 用循环不变式；快速选择沿用三路划分版本的强归纳。**

#### 一、正确性依赖的关键条件

1. **引理 1、2**（见[排序解法](../sorting/summary.zh-CN.md#44-正确性correctness)）：答案为 $`f(b_m) = \sum_i \lvert a_i - b_m \rvert`$。
2. **引理 3′（Lomuto 划分）**：`PARTITION` 返回 $`q`$ 时，$`A[l..r]`$ 是原区间的一个排列，且 $`A[l..q-1] \le x`$、$`A[q] = x`$、$`A[q+1..r] > x`$。
3. **引理 4（快速选择）**：`QUICK-SELECT(A, l, r, k)` 返回升序排列后位于下标 $`k`$ 的元素。

#### 二、正确性证明

**引理 3′：PARTITION 的循环不变式**（CLRS §7.1）

在 `for` 循环每次检查条件时：

- $`A[l..i]`$ 的元素都 $`\le x`$；
- $`A[i+1..j-1]`$ 的元素都 $`> x`$；
- $`A[j..r-1]`$ 尚未检查；
- $`A[r] = x`$。

**初始化**：$`i = l - 1`$、$`j = l`$，前两个区域为空。

**保持**：

- 若 $`A[j] > x`$，只需 $`j`$ 加 1，第二个区域向右扩展。
- 若 $`A[j] \le x`$，$`i`$ 加 1 后交换 $`A[i]`$ 与 $`A[j]`$。第一个区域多了一个 $`\le x`$ 的元素；被换到 $`j`$ 的是原第二个区域的第一个元素（或是自交换），第二个区域整体右移一格。

**终止**：$`j = r`$ 时未检查区为空。最后交换 $`A[i+1]`$ 与 $`A[r]`$，把主元放在两个区域之间，返回 $`q = i + 1`$。只做交换，所以结果是排列。$`\square`$

**引理 4**：证明与[三路划分版本的引理 4](../three-way-quickselect/summary.zh-CN.md#44-正确性correctness)相同，把三段换成 $`A[l..q-1]`$、$`A[q]`$、$`A[q+1..r]`$ 即可。由引理 3′，它们依次占据升序下标 $`l..q-1`$、$`q`$、$`q+1..r`$；$`A[l..q-1]`$ 中可能有等于 $`x`$ 的元素，但这不影响它们在升序中的位置。主元本身被排除，所以子区间严格更短。$`\square`$

#### 三、结论

由引理 4，`QUICK-SELECT` 返回 $`b_m`$；重排不改变多重集合（引理 3′），所以求和得到 $`f(b_m)`$，由引理 1、2 即为答案。结果总是正确的，问题只在运行时间。

### 4.5 复杂度分析（Complexity Analysis）

设输入规模为 $`n`$。

**时间复杂度：元素互不相同时期望 O(n)；大量重复时 Θ(n²)**

**元素互不相同时**：三路划分版本的分阶段论证要求“下一层长度不超过 $`\max(\rho, s - 1 - \rho)`$”，其中 $`\rho`$ 是主元的名次。元素互不相同时，Lomuto 划分满足这一点，期望仍为 $`O(n)`$，与 CLRS §9.2 的结论一致。

**有重复元素时**：Lomuto 划分把所有等于主元的元素放进左侧，上面的要求不再成立。全相等时，无论选中哪个元素，`A[j] ≤ x` 总成立，$`q`$ 总等于 $`r`$，每层长度只减 1。以 $`k = \lfloor n/2 \rfloor`$ 为例，区间要从 $`n`$ 缩到约 $`n/2`$，比较次数约为

```math
\sum_{s = n/2}^{n} s \approx \frac{3n^2}{8} .
```

$`n = 10^5`$ 时约 $`3.75 \times 10^9`$，实测约 $`3.7 \times 10^9`$ 次比较、7.7 秒，与估计一致。

求和一次遍历 $`\Theta(n)`$。

**空间复杂度：O(1)（不计递归栈）**

原地重排。递归深度期望 $`O(\log n)`$，最坏 $`O(n)`$；全相等时深度为 $`\Theta(n)`$。

### 4.6 实现细节：Go 与 Python（Implementation Notes）

- **递归改循环**：这一点在本解法中是必需的。重复元素多时，每次划分只排除一个元素，递归深度接近 $`n`$。Go 的栈会自动增长，可以撑住；CPython 默认 1000 层就会抛出 `RecursionError`。所以 Python 版把尾递归改写成 `while` 循环。见[递归深度](./../../../../docs/LANGUAGE_NOTES.zh-CN.md#递归深度)。
- **退化时的耗时**：同样的 $`\Theta(n^2)`$ 退化，Python 比 Go 慢得多，所以两种语言的测试都不对本解法跑全等的大数组。
- **随机主元**：Go 的 `rand.Intn(r-l+1) + l` 对应 Python 的 `random.randint(l, r)`，后者两端都包含。

## 5. 思维难点（Learning Notes）

### 难点一：为什么随机选主元在这里没用

**理解过程**

随机化保证的是：主元在升序中的名次均匀分布。但 `A[j] ≤ x` 把每个等于主元的元素都放到左边，相当于把主元当成“相同值里最大的那个”。全相等时，无论选中哪个，名次都是最后一名，每次都是最坏的划分。详细讨论见[三路划分版本的难点一](../three-way-quickselect/summary.zh-CN.md#难点一随机选主元为什么挡不住重复元素)。

**可迁移思维：** 期望复杂度的证明往往默认元素互不相同；遇到大量重复值时，要重新检查这个假设。

## 6. 可迁移知识（Key Takeaways）

本题总结出以下可复用的方法：

1. **Lomuto Partition（Lomuto 划分）**：CLRS 的 PARTITION，单指针扫描，把 $`\le x`$ 的元素依次换到前面。写法简单，是快速排序与快速选择的常见起点。限制：重复值多时划分严重不平衡，需要改用三路划分或 Hoare 划分。
2. **Measuring the Worst Case（实测最坏情况）**：本题通过全相等数组的实测，发现了理论分析中“元素互不相同”的隐含假设。适用于检验任何声称“期望”复杂度的算法。

### CLRS 关联

章节号以 CLRS 第三版为准。

| 知识点 | CLRS 章节 | 关联程度 | 本题中的体现 |
|---|---|---|---|
| 划分及其循环不变式 | §7.1 快速排序的描述 | 直接对应 | `partition` 就是 PARTITION，正确性按四个区域的循环不变式证明 |
| 随机选主元 | §7.3 快速排序的随机化版本 | 直接对应 | `randomPartition` 就是 RANDOMIZED-PARTITION |
| 重复元素的划分 | 思考题 7-2 针对相同元素值的快速排序 | 直接对应 | 全相等数组让本解法退化为 $`\Theta(n^2)`$，正是该思考题讨论的情形 |
| 期望线性时间的选择 | §9.2 期望为线性时间的选择算法 | 直接对应 | `QuickSelect` 就是 RANDOMIZED-SELECT；期望 $`O(n)`$ 的分析假设元素互不相同 |

## 7. 复习记录（Review Log）

| 日期 | 独立完成情况 | 理解变化 | 待改进 |
|---|---|---|---|
| 待记录 | 待记录 | 待记录 | 待记录 |

### 下次复习检查

- [ ] 能否不看代码独立实现？
- [ ] 能否准确描述算法输入与输出？
- [ ] 能否独立写出伪代码？
- [ ] 能否解释算法正确性？
- [ ] 能否推导时间与空间复杂度？
- [ ] 能否识别其他题目中的相同算法思想？

## 8. 关联题目（Related Problems）

- [215. Kth Largest Element in an Array](https://leetcode.cn/problems/kth-largest-element-in-an-array/) — 直接求第 $`k`$ 大；测试数据中有大量重复值时，两路划分同样会退化。

---

**一句话总结**

照 CLRS 写的随机快速选择结果总是对的，但 `A[j] ≤ x` 让重复值全部挤到一边，全相等时退化为 $`\Theta(n^2)`$。
