# CLRS 反查索引

[English](CLRS_INDEX.md) | 简体中文

章节号以《算法导论》第三版为准。本索引汇总各题题解「CLRS 关联」表中写了章节号的条目，按章节排序，只列出已有题目关联的章节。写“无对应章节”的知识点不收录。

每节一张表，格式如下：

```markdown
## 第 16 章 贪心算法

### §16.2 贪心策略的要素

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [题号. 题名](../Algorithms/<分类>/<题目录>/summary.zh-CN.md) | 直接对应 / 相关思想 | 具体到本题的哪一步 |
```

## 第 2 章 算法基础

### §2.1 插入排序

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 直接对应 | 用“前 $`i`$ 行已正确”的外层循环不变式，按初始化、保持、终止三步证明 |
| [3644. 排序排列](../Algorithms/bit-manipulation/3644-maximum-k-to-sort-a-permutation/summary.zh-CN.md) | 直接对应 | 用“`res` 是前缀中支撑集元素的按位与”的不变式，按初始化、保持、终止三步证明代码算出 $`M`$ |
| [128. 最长连续序列](../data-structures/hash-table/0128-longest-consecutive-sequence/summary.zh-CN.md) | 直接对应 | 引理 1 的内层不变式与外层“`res` 是已处理段首的最大段长”的不变式，都按初始化、保持、终止三步证明 |
| [453. 最小操作次数使数组元素相等](../Algorithms/simulation/0453-minimum-moves-to-equal-array-elements/summary.zh-CN.md) | 直接对应 | 解法一用“前缀相等、后缀从未被跳过”的不变式，按初始化、保持、终止三步证明 |

## 第 3 章 函数的增长

### §3.2 标准记号与常用函数

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 相关思想 | 斐波那契递推与本题递推同类，且杨辉三角浅对角线之和等于斐波那契数 |

## 第 11 章 散列表

### §11.2 散列表

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [128. 最长连续序列](../data-structures/hash-table/0128-longest-consecutive-sequence/summary.zh-CN.md) | 直接对应 | `numSet` 的插入与存在性查询期望 $`O(1)`$，是整体期望 $`O(n)`$ 的前提 |

## 第 15 章 动态规划

### §15.3 动态规划原理

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 相关思想 | 上一行每个数被下一行两个位置复用，按行号递增填表保证先算后用 |

## 第 17 章 摊还分析

### §17.1 聚合分析

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [128. 最长连续序列](../data-structures/hash-table/0128-longest-consecutive-sequence/summary.zh-CN.md) | 直接对应 | 内层 `while` 的总执行次数为 $`m - (\text{段首个数}) \le m`$，不逐次估计 |

### §17.3 势能法

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [453. 最小操作次数使数组元素相等](../Algorithms/simulation/0453-minimum-moves-to-equal-array-elements/summary.zh-CN.md) | 相关思想 | 用 $`\Phi = \sum x_i - n \cdot \min x`$ 刻画状态，界定每步 $`\Delta\Phi \ge -1`$ 得到步数下界；这里用于证明下界，不是做摊还分析 |

## 附录 C 计数与概率

### C.1 计数

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 相关思想 | 证明 $`P(i, j) = \binom{i}{j}`$，并据此得出 $`\binom{29}{14}`$ 的溢出上界 |
| [3963. 构造恰好一条路径的网格](../Algorithms/constructive-algorithms/3963-create-grid-with-exactly-one-path/summary.zh-CN.md) | 相关思想 | 测试中全空 $`3 \times 3`$ 网格的路径数取 $`\binom{4}{2} = 6`$，用来检验验证器 `countPaths` |
