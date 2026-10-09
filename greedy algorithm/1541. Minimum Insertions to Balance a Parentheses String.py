class Solution:
    def minInsertions(self, s: str) -> int:
        # 连续两个")"——有顺序性
        # 根据当前字符和已有状态，执行状态转移；在转移过程中修复违反约束的情况，始终维持状态不变量。
        # 通过局部的最小修复来维护全局合法性
        # 维护两个变量：res已经确定要插入的数目，right尚未匹配的右括号需求数量

        """
        贪心 + 状态模拟
        一个 ( 必须对应连续的两个 ))

        状态不变量：
        right 表示尚未匹配的右括号需求，始终 >= 0

        res：已确定必须插入的括号数量
        right：尚未匹配的右括号需求数量

        核心策略：
        仅修复当前必须修复的部分，其余需求留给后续字符满足
        """
        res, right = 0, 0
        for c in s:
            if c == "(":
                if right % 2 == 1:
                    # 代表两个"("中只有一个")"
                    res += 1
                    right -= 1
                right += 2

            else:
                right -= 1

                if right < 0:
                    res += 1
                    right += 2

        return res + right