# leetcode-go

[![Go](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go&logoColor=white)](go.mod)
[![Code: MIT](https://img.shields.io/badge/code-MIT-blue)](LICENSE)
[![Docs: CC BY 4.0](https://img.shields.io/badge/docs-CC%20BY%204.0-lightgrey)](LICENSE-docs)

[English](README.md) | 简体中文

用 Go 刷 LeetCode 的代码与学习笔记，按算法和数据结构分类整理。题解记录：

- **思路是怎么来的**：最初的想法、卡在哪里，以及关键的思维转变。
- **为什么是对的**：用循环不变式、归纳、交换论证等适合的方法给出正确性证明。
- **难在哪里**：分别评价建模、知识门槛、正确性论证和实现的难度，说明各项评分的理由。
- **能迁移到哪里**：提炼可以用于其他题目的方法，以及方法适用的条件。

完整列表见[题目总表](docs/PROBLEMS.zh-CN.md)。

## 如何阅读一道题

[题解模板](templates/)采用以下结构，可以按需跳读：

| 章节 | 适合什么时候看 |
| --- | --- |
| 1. 题目信息与难度评分 | 了解题目与解法的要求 |
| 2. 题意理解 | 仍不确定题目目标或约束 |
| 3. 初始思路 | 想对照自己的思考过程 |
| 4. 算法表达：伪代码、正确性、复杂度 | 想确认算法步骤、证明与计算成本 |
| 5. 思维难点 | 做错了或卡住了 |
| 6. 可迁移知识 | 做完后提炼可复用的方法 |
| 7. 复习记录 | 隔一段时间回来自测 |
| 8. 关联题目 | 想找同类题巩固 |

四维评分评价的是**指定约束下的具体解法**，每项 1～5 分，独立评价、不求和。力扣官方难度和 zerotrac 评分另行记录。详细标准见[难度评分标准](docs/RATING.zh-CN.md)。

## 目录结构

可以从[算法分类](Algorithms/)、[数据结构分类](data-structures/)或[入门练习](primers/)开始浏览。

<details>
<summary>展开完整目录树</summary>

```text
.
├── Algorithms/                  # 按算法分类的题目
│   └── <分类>/<题号>-<英文题名>/
│       ├── solution.go
│       ├── solution_test.go
│       ├── summary.md           # 英文题解，按整理进度补充
│       └── summary.zh-CN.md     # 中文题解，按整理进度补充
├── data-structures/             # 按数据结构分类的题目
│   └── <分类>/<题号>-<英文题名>/
├── primers/                     # 入门练习，只有代码和测试
├── templates/                   # 中英文题解模板
├── docs/                        # 评分标准与文档规范
│   ├── RATING.md                # 英文难度评分标准
│   ├── RATING.zh-CN.md          # 中文难度评分标准
│   ├── CLRS_INDEX.md            # 英文 CLRS 反查索引
│   ├── CLRS_INDEX.zh-CN.md      # 中文 CLRS 反查索引
│   ├── PROBLEMS.md              # 英文题目总表
│   ├── PROBLEMS.zh-CN.md        # 中文题目总表
│   ├── GLOSSARY.md              # 中英文全局术语表
│   ├── GITHUB_MATH.md           # GitHub 数学书写规范
│   ├── PROMPT.zh-CN.md          # 题解写作提示词
│   └── investigations/          # 排查记录
│       ├── README.md
│       ├── README.zh-CN.md
│       └── github-math-rendering-investigation.md
├── README.md
└── README.zh-CN.md
```

</details>

## 运行测试

需要 Go 1.22 或更高版本，版本要求见 [go.mod](go.mod)。在仓库根目录执行：

```bash
go test ./...
go test ./Algorithms/greedy-algorithms/...
```

第一条命令测试所有包，第二条命令测试一个分类。

## 新增一道题

1. 在 `Algorithms/<分类>/` 或 `data-structures/<分类>/` 下建目录，命名为 `<题号>-<英文题名>`，全部小写，用连字符分隔。入门练习放在 `primers/` 下。
2. 编写 `solution.go` 和 `solution_test.go`，两者统一使用 `package solution`，测试题目示例、边界条件与算法的关键分支。
3. 整理题解时，将[中文模板](templates/summary.zh-CN.md)和[英文模板](templates/summary.md)复制到题目目录。将其中的 `../docs/` 改为 `../../../docs/`，再替换占位内容，并删除开头注释中的复制说明。
4. 按[评分标准](docs/RATING.zh-CN.md)评价解法，写明约束和目标复杂度。中英文版本的事实、评分与数学含义应保持一致。
5. 对照[全局术语表](docs/GLOSSARY.md)检查术语，按 [GitHub 数学书写规范](docs/GITHUB_MATH.md)检查数学记号，同步更新[题目总表](docs/PROBLEMS.zh-CN.md)和 [CLRS 反查索引](docs/CLRS_INDEX.zh-CN.md)，确认相对链接有效，并运行 `go test ./...`。

## 文档约定

- 本仓库引用的《算法导论》（CLRS）章节号均以**第三版**（Cormen, Leiserson, Rivest, Stein, *Introduction to Algorithms*, 3rd ed., MIT Press, 2009）为准。
- [CLRS 反查索引](docs/CLRS_INDEX.zh-CN.md)：按 CLRS 章节列出关联的题目。
- [全局术语表](docs/GLOSSARY.md)：统一中英文术语及定义。
- [GitHub 数学书写规范](docs/GITHUB_MATH.md)：公式写法、排版与渲染验收方式。

## 排查记录

[排查记录](docs/investigations/README.zh-CN.md)收集问题定位、对照实验与验证结果，目前包括 [GitHub 数学渲染排查记录](docs/investigations/github-math-rendering-investigation.md)。

## 许可证

- 源代码（`*.go` 文件、`go.mod`，以及文档中的代码片段）采用 [MIT 许可证](LICENSE)。
- 文档、题解与笔记（所有 `*.md` 文件）采用 [CC BY 4.0](LICENSE-docs) 许可。可以自由转载和改编，但需署名并注明来源仓库链接。
