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
| [49. 字母异位词分组（排序键）](../data-structures/hash-table/0049-group-anagrams/sorted-key/summary.zh-CN.md) | 直接对应 | 引理 2 用“键集合与各组内容恰好对应前 $`i`$ 个字符串”的不变式，按初始化、保持、终止三步证明分组不重不漏 |
| [49. 字母异位词分组（计数键）](../data-structures/hash-table/0049-group-anagrams/count-key/summary.zh-CN.md) | 直接对应 | 引理 3 用“`cnt` 是前缀中各字母的出现次数”的不变式，按初始化、保持、终止三步证明计数正确 |
| [128. 最长连续序列](../data-structures/hash-table/0128-longest-consecutive-sequence/summary.zh-CN.md) | 直接对应 | 引理 1 的内层不变式与外层“`res` 是已处理段首的最大段长”的不变式，都按初始化、保持、终止三步证明 |
| [453. 最小操作次数使数组元素相等（成批模拟）](../Algorithms/simulation/0453-minimum-moves-to-equal-array-elements/batched-simulation/summary.zh-CN.md) | 直接对应 | 成批模拟用“前缀相等、后缀从未被跳过”的不变式，按初始化、保持、终止三步证明 |

## 第 3 章 函数的增长

### §3.2 标准记号与常用函数

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 相关思想 | 斐波那契递推与本题递推同类，且杨辉三角浅对角线之和等于斐波那契数 |

## 第 7 章 快速排序

### §7.1 快速排序的描述

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [462. 最小操作次数使数组元素相等 II（两路划分）](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/lomuto-quickselect/summary.zh-CN.md) | 直接对应 | `partition` 就是 PARTITION，正确性按四个区域的循环不变式证明 |
| [462. 最小操作次数使数组元素相等 II（三路划分）](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/three-way-quickselect/summary.zh-CN.md) | 相关思想 | 三路划分沿用 PARTITION “按区域写不变式”的证明结构，把区域从两类扩展到三类 |

### §7.3 快速排序的随机化版本

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [462. 最小操作次数使数组元素相等 II（两路划分）](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/lomuto-quickselect/summary.zh-CN.md) | 直接对应 | `randomPartition` 就是 RANDOMIZED-PARTITION |
| [462. 最小操作次数使数组元素相等 II（三路划分）](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/three-way-quickselect/summary.zh-CN.md) | 直接对应 | `randomPartition3` 与 RANDOMIZED-PARTITION 相同：先把随机元素换到 `a[r]` 再划分 |

### 思考题 7-2 针对相同元素值的快速排序

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [462. 最小操作次数使数组元素相等 II（两路划分）](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/lomuto-quickselect/summary.zh-CN.md) | 直接对应 | 全相等数组让本解法退化为 $`\Theta(n^2)`$，正是该思考题讨论的情形 |
| [462. 最小操作次数使数组元素相等 II（三路划分）](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/three-way-quickselect/summary.zh-CN.md) | 直接对应 | 三路划分把等于主元的元素单独成段，避免全相等数组上的 $`\Theta(n^2)`$ |

## 第 8 章 线性时间排序

### §8.2 计数排序

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [49. 字母异位词分组（计数键）](../data-structures/hash-table/0049-group-anagrams/count-key/summary.zh-CN.md) | 相关思想 | `cnt` 的计算与 COUNTING-SORT 中统计每个值出现次数的那一步相同；本解法到此为止，不再把计数展开成有序序列，因为计数向量本身就能作键 |

## 第 9 章 中位数和顺序统计量

### §9.2 期望为线性时间的选择算法

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [462. 最小操作次数使数组元素相等 II（两路划分）](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/lomuto-quickselect/summary.zh-CN.md) | 直接对应 | `QuickSelect` 就是 RANDOMIZED-SELECT；期望 $`O(n)`$ 的分析假设元素互不相同 |
| [462. 最小操作次数使数组元素相等 II（三路划分）](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/three-way-quickselect/summary.zh-CN.md) | 直接对应 | `QuickSelect` 即 RANDOMIZED-SELECT：只递归包含第 $`k`$ 小的一侧，期望 $`O(n)`$ |

### §9.3 最坏情况为线性时间的选择算法

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [462. 最小操作次数使数组元素相等 II（三路划分）](../Algorithms/order-statistics/0462-minimum-moves-to-equal-array-elements-ii/three-way-quickselect/summary.zh-CN.md) | 相关思想 | 本实现最坏 $`O(n^2)`$；若需要最坏 $`O(n)`$，可改用中位数的中位数选主元 |

## 第 10 章 基本数据结构

### §10.4 有根树的表示

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [104. 二叉树的最大深度](../data-structures/binary-tree/0104-maximum-depth-of-binary-tree/summary.zh-CN.md) | 直接对应 | `TreeNode` 的 `Left`、`Right` 指针就是 CLRS 中结点的 left、right 属性，空孩子 `nil` 对应 NIL |
| [226. 翻转二叉树](../data-structures/binary-tree/0226-invert-binary-tree/summary.zh-CN.md) | 直接对应 | 翻转就是改写每个结点的 left、right 属性；空孩子 `nil` 对应 NIL，算法不新建结点，只重排已有指针 |

## 第 11 章 散列表

### §11.1 直接寻址表

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [49. 字母异位词分组（计数键）](../data-structures/hash-table/0049-group-anagrams/count-key/summary.zh-CN.md) | 直接对应 | `cnt` 是以字母编号 $`0 \ldots 25`$ 为关键字的直接寻址表，每次计数最坏 $`O(1)`$ |

### §11.2 散列表

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [49. 字母异位词分组（排序键）](../data-structures/hash-table/0049-group-anagrams/sorted-key/summary.zh-CN.md) | 直接对应 | `groups` 以排序串为键，每个字符串一次查表与追加，期望 $`O(\ell_i)`$ |
| [49. 字母异位词分组（计数键）](../data-structures/hash-table/0049-group-anagrams/count-key/summary.zh-CN.md) | 直接对应 | `groups` 以 `[26]int` 为键，每个字符串一次查表与追加，期望 $`O(\lvert\Sigma\rvert)`$ |
| [128. 最长连续序列](../data-structures/hash-table/0128-longest-consecutive-sequence/summary.zh-CN.md) | 直接对应 | `numSet` 的插入与存在性查询期望 $`O(1)`$，是整体期望 $`O(n)`$ 的前提 |

## 第 12 章 二叉搜索树

### §12.1 什么是二叉搜索树

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [104. 二叉树的最大深度](../data-structures/binary-tree/0104-maximum-depth-of-binary-tree/summary.zh-CN.md) | 相关思想 | INORDER-TREE-WALK 对每个结点调用一次，定理 12.1 证明其时间为 $`\Theta(n)`$；4.5 用同样的“每个结点一次非空调用”的方式数调用次数，只是本题按后序合并子树结果 |
| [226. 翻转二叉树](../data-structures/binary-tree/0226-invert-binary-tree/summary.zh-CN.md) | 相关思想 | INORDER-TREE-WALK 对每个结点调用一次，定理 12.1 证明其时间为 $`\Theta(n)`$；4.5 用同样的方式数出 $`2n + 1`$ 次调用。4.4 的推论说明翻转后的中序序列是原序列的逆序 |

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
| [453. 最小操作次数使数组元素相等（成批模拟）](../Algorithms/simulation/0453-minimum-moves-to-equal-array-elements/batched-simulation/summary.zh-CN.md) | 相关思想 | 用 $`\Phi = \sum x_i - n \cdot \min x`$ 刻画状态，界定每步 $`\Delta\Phi \ge -1`$ 得到步数下界；这里用于证明下界，不是做摊还分析 |

## 附录 B 集合等离散数学内容

### B.5 树

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [104. 二叉树的最大深度](../data-structures/binary-tree/0104-maximum-depth-of-binary-tree/summary.zh-CN.md) | 直接对应 | 难点二：CLRS 按边数定义深度和树高，本题按结点数计，非空树上答案等于 CLRS 树高加 1 |
| [226. 翻转二叉树](../data-structures/binary-tree/0226-invert-binary-tree/summary.zh-CN.md) | 直接对应 | B.5.3 二叉树和位置树：4.2 的 $`M(T)`$ 按二叉树的递归定义给出；`left_child_only` 与 `right_child_only` 两个用例对应 B.5.3 中“只有左孩子”和“只有右孩子”是不同二叉树的区分 |
### B.2 关系

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [49. 字母异位词分组（排序键）](../data-structures/hash-table/0049-group-anagrams/sorted-key/summary.zh-CN.md) | 直接对应 | 引理 0 证明“互为异位词”是等价关系，由定理 B.1 其等价类构成划分，输出的组就是这些等价类 |
| [49. 字母异位词分组（计数键）](../data-structures/hash-table/0049-group-anagrams/count-key/summary.zh-CN.md) | 直接对应 | 引用排序键解法的引理 0：“互为异位词”是等价关系，由定理 B.1 其等价类构成划分，输出的组就是这些等价类 |

## 附录 C 计数与概率

### C.1 计数

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 相关思想 | 证明 $`P(i, j) = \binom{i}{j}`$，并据此得出 $`\binom{29}{14}`$ 的溢出上界 |
| [3963. 构造恰好一条路径的网格](../Algorithms/constructive-algorithms/3963-create-grid-with-exactly-one-path/summary.zh-CN.md) | 相关思想 | 测试中全空 $`3 \times 3`$ 网格的路径数取 $`\binom{4}{2} = 6`$，用来检验验证器 `countPaths` |
