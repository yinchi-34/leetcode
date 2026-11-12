# 462. 最小操作次数使数组元素相等 II — 排序

[English](./summary.md) | 简体中文

术语遵循[共享术语表](../../../../docs/GLOSSARY.md)，数学记号遵循仓库的 [GitHub 数学书写规范](../../../../docs/GITHUB_MATH.md)。

本题有三种解法，分别放在三个目录中，共用同一份测试用例：

- **排序**（本文）：排序后取中位数。本文给出题意和中位数最优性的完整证明，另外两篇直接引用。
- [**三路划分快速选择**](../three-way-quickselect/summary.zh-CN.md)：随机快速选择，划分为小于、等于、大于三段。
- [**两路（Lomuto）划分快速选择**](../lomuto-quickselect/summary.zh-CN.md)：CLRS 的 PARTITION，重复元素多时会退化。

## 1. 题目信息

- **题目**：[462. Minimum Moves to Equal Array Elements II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/)
- **官方难度**：Medium（中等）
- **主分类**：Order Statistics（顺序统计量）
- **辅助分类**：Sorting（排序）
- **技术标签**：`顺序统计量`、`排序`、`中位数`
- **代码**：[solution.go](./solution.go) · [solution.py](./solution.py)
- **测试**：[solution_test.go](../solution_test.go) · [solution_test.py](../solution_test.py)（放在上一级目录，三种解法共用）

### 难度评分

**外部评分**

- zerotrac 评分：暂无数据（[数据来源](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt)）
- 力扣官方难度：Medium

**四维评分 — 排序取中位数**

- 评价范围：$`1 \le n \le 10^5`$，$`-10^9 \le a_i \le 10^9`$，答案在 32 位整数范围内；时间 $`O(n \log n)`$，额外空间 $`\Theta(n)`$；Go 实现。

| 维度 | 分数 | 评分理由 |
|---|---|---|
| 建模难度 | 2 | 需要把“最少操作次数”转化为“绝对偏差之和的最小值”，并想到目标值取中位数 |
| 知识门槛 | 1 | 只需要排序和绝对值 |
| 正确性论证难度 | 2 | 只需中位数最优性：用“目标每右移 1，代价变化 $`2L - n`$”做一次分类即可 |
| 实现难度 | 1 | 排序后取 `a[n/2]` 求和；复制输入以免修改调用方的切片 |

评分遵循仓库 [RATING.zh-CN.md](../../../../docs/RATING.zh-CN.md)，各维度独立评价，不合成总分。

## 2. 题意理解（Problem Understanding）

给定长度为 $`n`$ 的整数数组 `nums`。一次操作选择一个元素，让它加 1 或减 1。求使所有元素相等的最少操作次数。

**示例**

- 输入：`nums = [1,2,3]`
- 输出：`2`（都变成 2：1 加 1，3 减 1）
- 输入：`nums = [1,10,2,9]`
- 输出：`16`

**关键要求**

1. 输入约束：$`1 \le n \le 10^5`$，$`-10^9 \le a_i \le 10^9`$，答案保证在 32 位整数范围内。
2. 每次只改变一个元素，改变量是 $`\pm 1`$；最终值不指定。
3. 边界：$`n = 1`$ 或全部相等时答案为 0；数组中可能有大量重复元素。

## 3. 初始思路（Initial Approach）

### 第一次想到：取中位数

两种实现都写过：Python 版先排序，取 `nums[len(nums)//2]` 作为目标；Go 版用快速选择找中位数（见[两路划分快速选择](../lomuto-quickselect/summary.zh-CN.md)）。本文是 Python 版的 Go 写法。

### 遇到的问题

当时的疑问是：如果中位数离两边的数都很远（有离群值），取中位数还对吗？

### 思维转变

从：

> 目标值应该离所有数都“近”，离群值会把它拉过去。

转变为：

> 目标值每移动 1，代价的变化只取决于它左右各有几个数，与距离无关。

这个观察就是下面引理 2 的证明。

## 4. 算法表达（Algorithm Specification）

**算法名称：** 排序取中位数（Median by Sorting）

### 4.1 输入（Input）

$`A = \langle a_0, a_1, \ldots, a_{n-1} \rangle`$

满足：

- $`1 \le n \le 10^5`$
- $`a_i \in \mathbb{Z}`$，$`-10^9 \le a_i \le 10^9`$

### 4.2 输出（Output）

最少操作次数

```math
\mathrm{res} = \min_{x \in \mathbb{Z}} \sum_{i=0}^{n-1} \lvert a_i - x \rvert .
```

### 4.3 伪代码（Pseudocode）

下标从 0 开始。记 $`m = \lfloor n/2 \rfloor`$。

```text
MIN-MOVES-2-SORT(A)
    B ← SORT(copy of A)              // 升序
    x ← B[m]
    res ← 0
    for i ← 0 to n - 1
        res ← res + |B[i] - x|
    return res
```

**状态定义**

- `x`：选定的目标值，即升序后下标为 $`m`$ 的元素。
- `res`：所有元素到 `x` 的距离之和。

### 4.4 正确性（Correctness）

**证明方法：先把问题化为最小化绝对偏差之和，再用“代价差分”的分类论证证明中位数最优。**

记 $`b_0 \le b_1 \le \cdots \le b_{n-1}`$ 为 $`A`$ 升序排列后的结果，$`f(x) = \sum_i \lvert a_i - x \rvert`$。

#### 一、正确性依赖的关键条件

1. **引理 1（目标固定时的代价）**：若最终值为 $`x`$，最少需要 $`f(x)`$ 次操作。
2. **引理 2（中位数最优）**：$`f`$ 在整数上的最小值在 $`x = b_m`$ 处取到。

两条引理合起来说明答案是 $`f(b_m)`$；排序后 `B[m]` 就是 $`b_m`$。另外两种解法同样依赖这两条引理。

#### 二、正确性证明

**引理 1**：每次操作只让一个元素变化 1，所以元素 $`i`$ 至少要被操作 $`\lvert a_i - x \rvert`$ 次，总次数至少 $`f(x)`$。让每个元素朝 $`x`$ 直接走 $`\lvert a_i - x \rvert`$ 步，恰好用 $`f(x)`$ 次。因为每步改变 1，最终值只能是整数，所以

```math
\mathrm{res} = \min_{x \in \mathbb{Z}} f(x) .
```

**引理 2**：对整数 $`x`$，记 $`L(x) = \#\{ i : a_i \le x \}`$。目标从 $`x`$ 移到 $`x + 1`$ 时，$`a_i \le x`$ 的每一项增加 1，$`a_i \ge x + 1`$ 的每一项减少 1，所以

```math
f(x+1) - f(x) = L(x) - \bigl(n - L(x)\bigr) = 2L(x) - n .
```

- **当 $`x < b_m`$**：$`b_m > x`$，所以 $`b_m, \ldots, b_{n-1}`$ 都大于 $`x`$，$`L(x) \le m`$，于是 $`2L(x) - n \le 2m - n \le 0`$，$`f`$ 不增。
- **当 $`x \ge b_m`$**：$`b_0, \ldots, b_m`$ 都不超过 $`x`$，$`L(x) \ge m + 1`$，于是 $`2L(x) - n \ge 2m + 2 - n \ge 1`$，$`f`$ 严格增。

所以 $`f`$ 在 $`b_m`$ 左侧不增、右侧严格增，最小值为 $`f(b_m)`$。$`\square`$

这个论证只用到 $`L(x)`$，即“左边有几个数”，与离群值有多远无关。$`n`$ 为偶数时，$`[b_{m-1}, b_m]`$ 内的整数代价都相同，取 $`b_m`$ 即可。

#### 三、结论

由引理 1、2，答案为 $`f(b_m)`$。算法排序后取 $`B[m] = b_m`$，对复制并排序后的数组求和；排序只改变顺序，不改变多重集合，所以和等于 $`f(b_m)`$。算法对所有合法输入都返回正确答案。

### 4.5 复杂度分析（Complexity Analysis）

设输入规模为 $`n`$。

**时间复杂度：O(n log n)**

排序 $`O(n \log n)`$，求和一次遍历 $`\Theta(n)`$。

**空间复杂度：Θ(n)**

复制数组 $`\Theta(n)`$，加上 `sort.Ints` 的 $`O(\log n)`$ 栈空间。若允许修改输入，可以原地排序，省去复制。

**数值范围**：$`f(b_m) \le n \cdot 2 \cdot 10^9 = 2 \times 10^{14}`$，在 64 位 `int` 范围内；题目还保证答案在 32 位范围内。

### 4.6 实现细节：Go 与 Python（Implementation Notes）

- **不修改输入**：Go 的 `sort.Ints` 原地排序，所以先用 `append([]int(nil), nums...)` 复制。Python 的 `sorted(nums)` 本身返回新列表，不需要手动复制。见[切片与列表](./../../../../docs/LANGUAGE_NOTES.zh-CN.md#切片与列表别名和拷贝)。
- **整数范围**：差的绝对值之和最大约 $`10^5 \times 2 \times 10^9 = 2 \times 10^{14}`$，在 Go 的 64 位 `int` 内。Python 没有溢出问题。

## 5. 思维难点（Learning Notes）

### 难点一：中位数离两边很远时还对吗

**理解过程**

引理 2 的差分是 $`2L(x) - n`$，只取决于 $`x`$ 左边有几个数，与距离无关。目标往离群值那边移动 1，离群值的代价降 1，但更多的数代价升 1。例如 `[1,2,3,1000,1001]`：取 3 代价 1998，取 1000 代价 2995。

**可迁移思维：** 绝对值之和的最小化只看“个数”，所以中位数对离群值稳健；平方和的最小化看“距离”，所以平均数会被离群值拉偏。

## 6. 可迁移知识（Key Takeaways）

本题总结出以下可复用的方法：

1. **Median Minimizes Absolute Deviation（中位数最小化绝对偏差之和）**：本题把最少操作次数化为 $`\min_x \sum \lvert a_i - x \rvert`$，取中位数。适用于“每步改变一个元素 1、求全部相等的最少步数”“一维选址”等问题。条件：代价是距离的绝对值之和；若每次改变量固定为 $`d`$，还要先检查所有数模 $`d`$ 同余。
2. **Cost Difference Argument（代价差分）**：比较 $`f(x+1)`$ 与 $`f(x)`$，找出差分由负变正的位置。适用于目标函数是若干个分段线性凸函数之和的问题。

排序是求中位数最直接的办法；只需要一个顺序统计量时，可以改用快速选择，见另外两篇。

### CLRS 关联

章节号以 CLRS 第三版为准。“关联程度”只写“直接对应”或“相关思想”；对不上具体章节时，“CLRS 章节”一栏写“无对应章节”，只保留概念名。

| 知识点 | CLRS 章节 | 关联程度 | 本题中的体现 |
|---|---|---|---|
| 中位数最小化绝对偏差之和 | 无对应章节 | 相关思想 | 引理 2 的代价差分论证 |

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

- [453. Minimum Moves to Equal Array Elements](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements/) — 同样要求全部相等，但一次操作让 $`n-1`$ 个数加 1；等价于“一个数减 1”后只能往下走，目标被固定在最小值，而不是中位数。
- [2033. Minimum Operations to Make a Uni-Value Grid](https://leetcode.cn/problems/minimum-operations-to-make-a-uni-value-grid/) — 每次加减固定值 $`x`$，先判断所有数模 $`x`$ 同余，再同样取中位数。

---

**一句话总结**

绝对偏差之和只看左右个数，所以目标取中位数；排序后取下标 $`\lfloor n/2 \rfloor`$ 的元素即可。
