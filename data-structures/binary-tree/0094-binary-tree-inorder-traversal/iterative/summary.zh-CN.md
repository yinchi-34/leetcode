# 94. 二叉树的中序遍历 — 显式栈

[English](./summary.md) | 简体中文

术语遵循[共享术语表](../../../../docs/GLOSSARY.md)，数学记号遵循仓库的 [GitHub 数学书写规范](../../../../docs/GITHUB_MATH.md)。二叉树本身的定义、Go 中的表示和常见遍历模式见[二叉树总体笔记](../../README.zh-CN.md)。

本题有两种解法，分别放在两个目录中，共用同一份测试用例：

- [**递归**](../recursive/summary.zh-CN.md)：按“左子树 → 根 → 右子树”的定义直接递归。题意、中序序列 $`\mathrm{in}(T)`$ 的形式化定义和测试设计见该文。
- **显式栈**（本文）：用栈保存“左子树还没处理完、自己还没输出”的祖先，回答题目的进阶要求“用迭代算法完成”。

本文只讨论“怎样不用递归得到同一个序列”。

## 1. 题目信息

- **题目**：[94. Binary Tree Inorder Traversal](https://leetcode.cn/problems/binary-tree-inorder-traversal/)
- **官方难度**：Easy（简单）
- **主分类**：Binary Tree（二叉树）
- **辅助分类**：Stack（栈）
- **技术标签**：`二叉树`、`栈`、`中序遍历`、`迭代`、`循环不变式`
- **代码**：[solution.go](./solution.go)
- **测试**：[solution_test.go](../solution_test.go)（放在上一级目录，两种解法共用）

### 难度评分

**外部评分**

- zerotrac 评分：暂无数据（[数据来源](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt)）
- 力扣官方难度：Easy

**四维评分 — 显式栈中序遍历**

- 评价范围：结点数 $`0 \le n \le 100`$，$`-100 \le \mathrm{Node.val} \le 100`$；时间 $`\Theta(n)`$；不计输出的额外空间 $`O(h)`$（$`h`$ 为最大深度，最坏 $`h = n`$）；Go 实现。

| 维度 | 分数 | 评分理由 |
|---|---|---|
| 建模难度 | 2 | 要想清楚栈里存什么：是“已经走过、但左子树还没处理完”的祖先，相当于把递归调用栈中挂起的调用显式保存下来 |
| 知识门槛 | 2 | 需要栈（Go 中用切片的末尾做栈顶）和二叉树的指针表示 |
| 正确性论证难度 | 3 | 循环里没有递归可以直接套归纳，要自己构造一个描述“剩余待输出序列”的循环不变式，并单独论证每个结点至多入栈一次，从而循环终止 |
| 实现难度 | 2 | 循环条件 `len(stack) > 0 \|\| node != nil` 的两半缺一不可；弹出后要转向右孩子而不是继续向左 |

评分遵循仓库 [RATING.zh-CN.md](../../../../docs/RATING.zh-CN.md)，各维度独立评价，不合成总分。

## 2. 题意理解（Problem Understanding）

题意、示例与关键要求见[递归解法 §2](../recursive/summary.zh-CN.md#2-题意理解problem-understanding)。本解法回答其中的进阶要求：不用递归，只用循环和一个栈。

## 4. 算法表达（Algorithm Specification）

**算法名称：** 显式栈中序遍历（Iterative Inorder Traversal with an Explicit Stack）

### 4.1 输入（Input）

与[递归解法 §4.1](../recursive/summary.zh-CN.md#41-输入input) 相同：一棵结点数 $`0 \le n \le 100`$ 的二叉树 $`T`$，根为 `root`。对结点 $`x`$，记 $`T_x`$ 为以 $`x`$ 为根的子树，并约定 $`T_{\mathrm{NIL}} = \varnothing`$。

### 4.2 输出（Output）

输出 $`\mathrm{in}(T)`$，定义见[递归解法 §4.2](../recursive/summary.zh-CN.md#42-输出output)：

```math
\mathrm{in}(\varnothing) = \langle\rangle, \qquad \mathrm{in}\bigl((r, L, R)\bigr) = \mathrm{in}(L) \cdot \langle r.\mathrm{val} \rangle \cdot \mathrm{in}(R)
```

### 4.3 伪代码（Pseudocode）

栈 $`S`$ 自底向上记为 $`s_1, s_2, \ldots, s_k`$，$`s_k`$ 是栈顶。

```text
INORDER-ITERATIVE(root)
    res ← ⟨⟩
    S ← empty stack
    node ← root
    while S is not empty or node ≠ NIL
        if node ≠ NIL
            PUSH(S, node)                   ▷ 先挂起 node，去处理它的左子树
            node ← node.left
        else
            tmp ← POP(S)                    ▷ tmp 的左子树已处理完
            res ← res · ⟨tmp.val⟩
            node ← tmp.right                ▷ 转去处理 tmp 的右子树
    return res
```

**状态定义**

- `node`：下一棵要处理的子树的根；`NIL` 表示这棵子树为空。
- `S`（代码中的 `stack`）：被挂起的祖先。每个栈中结点的左子树正在处理或已处理完，它自己的值和它的右子树还没输出。
- `res`：已经确定的中序序列前缀。

### 4.4 正确性（Correctness）

**证明方法：循环不变式（初始化、保持、终止）+ 单独的终止性论证**

#### 一、正确性依赖的关键条件

1. **中序序列的递归分解**（定义）：对非空子树 $`T_x`$，$`\mathrm{in}(T_x) = \mathrm{in}(T_{x.\mathrm{left}}) \cdot \langle x.\mathrm{val} \rangle \cdot \mathrm{in}(T_{x.\mathrm{right}})`$；对空子树，$`\mathrm{in}(\varnothing) = \langle\rangle`$。
2. **输入是一棵树**（题目条件）：除根外每个结点有唯一的父结点，且只从父结点的某一个孩子指针到达。这保证每个结点至多被压栈一次（引理 1）。
3. **栈的后进先出**：最后挂起的祖先最先恢复，与递归中最内层调用最先返回一致。

#### 二、正确性证明

**循环不变式（Loop Invariant）**

每次检查 `while` 条件时：

**不变式 I：输出分解**

```math
\mathrm{in}(T) = \mathrm{res} \cdot \mathrm{in}(T_{\mathrm{node}}) \cdot \langle s_k.\mathrm{val} \rangle \cdot \mathrm{in}(T_{s_k.\mathrm{right}}) \cdot \langle s_{k-1}.\mathrm{val} \rangle \cdot \mathrm{in}(T_{s_{k-1}.\mathrm{right}}) \cdots \langle s_1.\mathrm{val} \rangle \cdot \mathrm{in}(T_{s_1.\mathrm{right}})
```

即：已输出的部分是 `res`；接下来要输出整棵 $`T_{\mathrm{node}}`$；然后从栈顶到栈底，依次输出每个挂起结点的值和它的右子树。

**不变式 II：栈是一条向左的路径**

对 $`1 \le i < k`$，$`s_{i+1}`$ 在 $`s_i`$ 的左子树中；`node` 非空时，它在 $`s_k`$ 的左子树中（栈空时在 $`T`$ 中）。因此栈中结点与 `node` 都位于同一条从根向下的路径上。

**初始化（Initialization）**

`res` 为空，栈为空，`node = root`。不变式 I 变为 $`\mathrm{in}(T) = \mathrm{in}(T_{\mathrm{root}})`$，成立；不变式 II 对空栈平凡成立。

**保持（Maintenance）**

假设某次迭代前两个不变式成立。

- **情况一：`node` 非空**。由关键条件 1，把 $`\mathrm{in}(T_{\mathrm{node}})`$ 展开为 $`\mathrm{in}(T_{\mathrm{node.left}}) \cdot \langle \mathrm{node.val} \rangle \cdot \mathrm{in}(T_{\mathrm{node.right}})`$。算法把 `node` 压栈（成为新的 $`s_{k+1}`$），再令 `node ← node.left`。新状态下不变式 I 的右边正是上述展开式接上原来的剩余部分，与原来相等。不变式 II：新栈顶是旧的 `node`，它在旧栈顶的左子树中；新的 `node` 是它的左孩子，在它的左子树中。
- **情况二：`node` 为空**。由循环条件，此时栈非空。$`\mathrm{in}(T_{\mathrm{NIL}}) = \langle\rangle`$，不变式 I 化为 $`\mathrm{in}(T) = \mathrm{res} \cdot \langle s_k.\mathrm{val} \rangle \cdot \mathrm{in}(T_{s_k.\mathrm{right}}) \cdots`$。算法弹出 $`s_k`$，把它的值接到 `res` 末尾，令 `node ← s_k.right`。新状态下右边是 $`(\mathrm{res} \cdot \langle s_k.\mathrm{val} \rangle) \cdot \mathrm{in}(T_{s_k.\mathrm{right}}) \cdot (\text{栈中其余部分})`$，与原来相等。不变式 II：剩余的栈不变；新的 `node` 在 $`s_k`$ 的右子树中，而 $`s_k`$ 在 $`s_{k-1}`$ 的左子树中，所以新的 `node` 也在 $`s_{k-1}`$ 的左子树中。

因此，两个不变式在每次迭代后继续成立。

**终止（Termination）**

*引理 1：每个结点至多被压栈一次。* 结点 $`x`$ 只能以两种方式成为 `node`：$`x`$ 是根（初始化时），或 $`x`$ 是某结点 $`p`$ 的孩子，在 $`p`$ 被压栈后（若 $`x = p.\mathrm{left}`$）或 $`p`$ 被弹栈后（若 $`x = p.\mathrm{right}`$）被赋值。由关键条件 2，$`x`$ 只有一个父结点，只对应其中一个时机。对深度归纳：根至多成为 `node` 一次；若 $`p`$ 至多被压栈一次，那么它也至多被弹栈一次（弹出后就不在栈中，且不会再被压入），所以它的每个孩子至多成为 `node` 一次。`node` 非空时会立刻被压栈，所以每个结点至多被压栈一次。

*计数*：每次迭代恰好做一次压栈或一次弹栈。压栈至多 $`n`$ 次，弹栈次数不超过压栈次数，所以循环至多执行 $`2n`$ 次后终止。

终止时循环条件为假，即栈为空且 `node` 为 NIL。不变式 I 化为

```math
\mathrm{in}(T) = \mathrm{res} \cdot \mathrm{in}(\varnothing) = \mathrm{res}
```

所以返回的 `res` 正是中序序列。由此也得到：$`\mathrm{in}(T)`$ 有 $`n`$ 个值，每次弹栈恰好追加一个，所以弹栈恰好 $`n`$ 次；弹出的结点都压过栈，所以压栈也恰好 $`n`$ 次，循环恰好执行 $`2n`$ 次。

#### 三、结论

不变式 I 把“还没输出的部分”精确写成 `node` 的整棵子树加上栈中每个结点的“值 + 右子树”，两种分支都只是按中序定义把其中一项展开或输出一个值，不改变整体序列；循环至多 $`2n`$ 次后栈空且 `node` 为空，剩余部分为空，`res` 就是 $`\mathrm{in}(T)`$。

### 4.5 复杂度分析（Complexity Analysis）

设 $`n = \lvert T \rvert`$，$`h`$ 为树的最大深度（按结点数计，空树为 0）。

**时间复杂度：$`\Theta(n)`$**

由 §4.4 的计数，循环恰好执行 $`2n`$ 次，每次是常数次操作加一次 `append`。`res` 共追加 $`n`$ 次，栈共压入 $`n`$ 次，切片扩容的总代价按 CLRS §17.4 动态表的摊还分析是 $`O(n)`$。合计 $`\Theta(n)`$。与递归解法不同，每个值只写入 `res` 一次，不会被反复复制。

这里用的是 CLRS §17.1 的聚合分析：不逐次估计“一次弹栈前会连续压栈多少次”，而是统计整个算法中压栈和弹栈的总次数。

**空间复杂度：不计输出 $`O(h)`$，最坏 $`\Theta(n)`$**

由不变式 II，栈中结点都位于同一条从根向下的路径上，所以栈的大小不超过 $`h`$。更准确地说，栈里只有“当前位置是其左子树中的结点”的祖先，所以这个上界不一定能达到：开始时沿最左路径一路压栈，栈至少会长到最左路径的长度。

- **最坏情况**：全部挂在左边的链，栈会长到 $`n`$。
- **最好情况**：全部挂在右边的链，每个结点压栈后立刻因为没有左孩子而被弹出，栈大小始终不超过 1。测试中的 `DeepChain` 就是这种情况。

最坏情况下，显式栈与递归调用栈占用的空间是同一个量级 $`\Theta(n)`$；在右偏的树上显式栈更省，因为递归在向右走时仍要保留调用帧（Go 不做尾调用优化）。它的主要好处是不依赖语言运行时的调用栈深度限制，并且可以在任意一步暂停（例如 230 题找到第 $`k`$ 个后提前停止），而不是降低空间复杂度。输出 `res` 另占 $`\Theta(n)`$。

### 4.6 测试设计（Test Design）

用例与[递归解法 §4.6](../recursive/summary.zh-CN.md#46-测试设计test-design) 共用，每个用例对两种解法各跑一次。对本解法额外关注的是：`Left-skewed chain` 让栈长到 $`n`$，`DeepChain` 让栈始终不超过 1，二者覆盖了栈大小的两个极端。

对本解法做过四种变异检验，每种都有用例失败：

- 压栈时也把值追加到 `res`（混入前序）：除 `Empty tree` 外全部失败。
- 弹栈后不追加值：除 `Empty tree` 外全部失败。
- 循环条件只写 `len(stack) > 0`：初始栈为空，循环一次也不执行，除 `Empty tree` 外全部失败。
- 弹栈后不转向右孩子（令 `node` 为空）：`Example 1`、`Example 2`、BST、负值、`DeepChain` 失败。

## 5. 思维难点（Learning Notes）

### 难点一：循环条件为什么要写 `stack` 非空 **或** `node` 非空？

**理解过程**

两半分别对应不变式 I 中剩余部分的两块：`node` 非空表示还有一整棵子树没处理，栈非空表示还有挂起的祖先没输出。只要其中一块非空，就还有值没输出。

- 只写“栈非空”：开始时栈为空，循环一次也不执行（变异检验中的第三种）。
- 只写“`node` 非空”：遇到第一个没有左孩子的结点就会停下，栈中挂起的祖先永远不会输出。

**可迁移思维：** 写循环条件时，先写出“还剩什么没做”的不变式；循环条件就是“剩余部分非空”，而剩余部分由几块组成，条件就是这几块的“或”。

### 难点二：栈里存的到底是什么？

**理解过程**

对比递归解法：递归调用 `INORDER(x)` 在调用 `INORDER(x.left)` 时被挂起，挂起的调用记住了“回来以后要输出 `x.val`，再处理 `x.right`”。显式栈存的就是这些被挂起的 `x`：弹出 `x` 时，恰好是它的左子树处理完、该“回来”的时刻。所以弹出后先输出 `x.val`，再令 `node = x.right`，与递归的后两步一一对应。

前序遍历在压栈时就输出值，后序遍历要等右子树也处理完才输出，所以后序的迭代写法需要额外记录“右子树是否已处理”。中序正好介于两者之间：在弹栈时输出。

**可迁移思维：** 把递归改写成迭代时，先问“递归调用被挂起时需要记住什么”，栈里就存这些信息；再问“恢复时要做什么”，那就是弹栈后的操作。

## 6. 可迁移知识（Key Takeaways）

本题总结出以下可复用的方法：

1. **Recursion to Explicit Stack（用显式栈代替递归）**：在本题中，栈保存“左子树还没处理完”的祖先，弹栈时相当于递归返回。适用于递归可能太深、或需要逐步产出结果（迭代器）的场景，例如 BST 迭代器、提前停止的遍历。限制：最坏情况的空间量级不变，仍是 $`O(h)`$。
2. **Loop Invariant over the Remaining Work（描述剩余工作的循环不变式）**：在本题中，不变式 I 把“已输出 + 剩余待输出”写成整个答案。适用于证明一切模拟递归、工作列表（worklist）式的算法：只要每一步都只是把剩余部分展开一项或输出一项，整体就保持不变。
3. **Aggregate Analysis of Stack Operations（栈操作的聚合分析）**：在本题中，每个结点恰好压栈、弹栈各一次，循环总次数 $`2n`$。适用于单调栈、括号匹配等“每个元素至多进出栈一次”的算法。

### CLRS 关联

章节号以 CLRS 第三版为准。“关联程度”只写“直接对应”或“相关思想”；对不上具体章节时，“CLRS 章节”一栏写“无对应章节”，只保留概念名。

| 知识点 | CLRS 章节 | 关联程度 | 本题中的体现 |
|---|---|---|---|
| 循环不变式 | §2.1 插入排序 | 直接对应 | 不变式 I（输出分解）与不变式 II（栈是一条向左的路径）都按初始化、保持、终止三步证明 |
| 栈 | §10.1 栈和队列 | 直接对应 | `stack` 用切片末尾做栈顶，`append` 是 PUSH，截断是 POP，后进先出保证最内层的祖先最先恢复 |
| 非递归遍历二叉树 | §10.4 有根树的表示 | 直接对应 | 练习 10.4-3 要求用栈在 $`O(n)`$ 时间内非递归地输出二叉树的每个关键字，本解法就是按中序完成的一种写法 |
| 非递归中序遍历 | §12.1 什么是二叉搜索树 | 直接对应 | 练习 12.1-3 要求给出非递归的中序遍历算法，提示之一是用栈作为辅助数据结构 |
| 聚合分析 | §17.1 聚合分析 | 直接对应 | 每个结点压栈、弹栈各一次，循环总次数为 $`2n`$，与 MULTIPOP 例子中“弹栈总次数不超过压栈总次数”是同一个论证 |

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

- [173. 二叉搜索树迭代器](https://leetcode.cn/problems/binary-search-tree-iterator/) — 把本解法的循环拆成一次一次的 `next()` 调用：栈就是迭代器的状态，均摊每次 $`O(1)`$，空间 $`O(h)`$。
- [230. 二叉搜索树中第 K 小的元素](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/) — 第 $`k`$ 次弹栈时直接返回，体现显式栈“可以随时停下”的好处。
- [144. 二叉树的前序遍历](https://leetcode.cn/problems/binary-tree-preorder-traversal/) — 同样的栈结构，只把输出时机从弹栈改到压栈。
- [145. 二叉树的后序遍历](https://leetcode.cn/problems/binary-tree-postorder-traversal/) — 输出时机要推迟到右子树处理完之后，需要额外记录上一个输出的结点，是本解法的进一步练习。

---

**一句话总结**

用栈挂起“左子树还没处理完”的祖先：一路向左压栈，弹栈时输出并转向右子树；不变式“已输出 · 当前子树 · 栈中各结点的值与右子树 = 完整中序序列”保证正确，每个结点进出栈各一次所以是 $`\Theta(n)`$。
