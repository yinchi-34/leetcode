# 全局术语表（Shared Glossary）

| 中文 | 英文 | 定义 |
| --- | --- | --- |
| 需求计数 | demand counting | 以未满足需求的数量表示状态的计数方法。 |
| 半完成匹配 | half-completed match | 由两步组成的匹配中，仅第一步完成的状态。 |
| 循环不变式 | loop invariant | 在每次循环的指定边界均成立的命题。 |
| 下界函数 | lower-bound function | 在定义域内逐点不超过目标函数的函数。 |
| 闭合 | closure | 下界函数 $`L`$ 对任意代价为 $`c`$ 的合法插入转移 $`x \to y`$ 满足 $`L(y) \le L(x) + c`$ 的性质。 |
| 二项式系数 | binomial coefficient | 从 $`n`$ 个不同元素中选出 $`k`$ 个的方案数，记作 $`\binom{n}{k}`$。 |
| 帕斯卡法则 | Pascal's rule | 对 $`0 < k < n`$，$`\binom{n}{k} = \binom{n-1}{k-1} + \binom{n-1}{k}`$。 |
| 自底向上填表 | bottom-up tabulation | 按依赖关系的拓扑序计算并存储子问题的值，使每个值在被使用前已经算出。 |
| 滚动数组 | rolling array | 当每层状态只依赖上一层时，只保留一层存储并原地更新的空间优化方法。 |
| 斐波那契数列 | Fibonacci sequence | 满足 $`F(1) = F(2) = 1`$、$`F(n) = F(n-1) + F(n-2)`$ 的整数序列。 |
| 子掩码 | submask | 若 $`x \mathbin{\&} y = x`$，即 $`x`$ 的每个为 1 的二进制位在 $`y`$ 中也为 1，则称 $`x`$ 是 $`y`$ 的子掩码。 |
| 星形对换 | star transposition | 都包含同一个固定元素（枢纽）的对换；只用这类对换就能实现任意排列。 |
| 单位元 | identity element | 对二元运算 $`\circ`$，满足 $`e \circ x = x \circ e = x`$ 的元素 $`e`$，例如按位与的全 1。 |
| 不动点 | fixed point | 置换 $`\pi`$ 中满足 $`\pi(i) = i`$ 的元素 $`i`$。 |
| 支撑集 | support | 置换 $`\pi`$ 中所有非不动点组成的集合，$`\mathrm{supp}(\pi) = \{i : \pi(i) \ne i\}`$。 |
| 散列集合 | hash set | 只存关键字、不带附加数据的散列表，支持期望 $`O(1)`$ 的插入与存在性查询。 |
| 极大连续段 | maximal run | 整数集合 $`S`$ 的子集 $`\{h, \ldots, h+L-1\}`$，满足 $`h-1 \notin S`$ 且 $`h+L \notin S`$，两端都不能再延长。 |
| 段首 | run head | 满足 $`h \in S`$ 且 $`h-1 \notin S`$ 的值 $`h`$，即某个极大连续段的最小元素。 |
| 聚合分析 | aggregate analysis | 先界定 $`n`$ 个操作的总代价 $`T(n)`$，再以 $`T(n)/n`$ 作为每个操作摊还代价的分析方法。 |
| 势函数 | potential function | 把状态映射为实数的函数；通过界定每步操作引起的变化量来证明步数或代价的界。 |
| 平移不变性 | translation invariance | 性质 $`P`$ 满足：对任意常数 $`t`$，$`P(x)`$ 成立当且仅当 $`P(x + t\mathbf{1})`$ 成立，其中 $`\mathbf{1}`$ 为全 1 向量。 |
| 顺序统计量 | order statistic | $`n`$ 个元素的集合中第 $`i`$ 小的元素称为第 $`i`$ 个顺序统计量。 |
| 中位数 | median | 位于排序后中间位置的顺序统计量；$`n`$ 为偶数时有下中位数与上中位数两个，0-based 下标分别为 $`n/2 - 1`$ 与 $`n/2`$。 |
| 三路划分 | three-way partitioning | 以主元 $`x`$ 把区间原地重排为小于 $`x`$、等于 $`x`$、大于 $`x`$ 三段的划分方法。 |
| 树高 | height (of a tree) | 根到叶的最长简单向下路径上的边数（CLRS 附录 B.5 的定义）；只有根的树高为 0。LeetCode 104 的“最大深度”按结点数计，非空树上等于树高加 1。 |
| 后序遍历 | postorder traversal | 先递归处理左、右子树，再处理根结点的遍历顺序。 |
| 结构归纳 | structural induction | 对递归定义的对象（如树）证明命题的方法：先证基础对象成立，再假设各组成部分成立，证明由它们构成的对象成立。 |
| 字母异位词 | anagram | 字符串 $`t`$ 是 $`s`$ 的字母异位词，当且仅当 $`t`$ 由 $`s`$ 的字母重新排列得到，等价地，两者每个字母的出现次数都相同。 |
| 等价关系 | equivalence relation | 满足自反、对称、传递三条性质的二元关系。 |
| 等价类 | equivalence class | 在等价关系下与某元素等价的全部元素组成的集合；一个集合上的全部等价类构成它的一个划分。 |
| 规范形式 | canonical form | 为每个等价类选定的唯一代表；两个元素等价当且仅当它们的规范形式相同。 |
| 完全不变量 | complete invariant | 对等价关系 $`\sim`$，满足 $`f(s) = f(t) \iff s \sim t`$ 的函数 $`f`$：等价的元素取值相同，取值相同的元素等价。 |
| 直接寻址表 | direct-address table | 关键字取自小全域 $`\{0, \ldots, m-1\}`$ 时，以关键字本身为数组下标存取元素的表，各操作最坏 $`O(1)`$。 |
| 中序遍历 | inorder traversal | 先递归处理左子树，再处理根结点，最后递归处理右子树的遍历顺序；二叉搜索树的中序序列按关键字非降序排列。 |
| 前序遍历 | preorder traversal | 先处理根结点，再递归处理左、右子树的遍历顺序。 |
| 显式栈 | explicit stack | 由程序自己维护、用来保存原本由递归调用栈保存的挂起状态的栈，用于把递归改写成迭代。 |
