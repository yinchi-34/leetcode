class Solution:
    def removeOuterParentheses(self, s: str) -> str:
        # 这道题第一次解法是stack，重新读题目，不需求知道每一个元素，只是深度去除外面一层，故修改用变量depth
        # 第二次写的时候判断逻辑是线性，从看到第一个左括号处理
        # 可以直接归纳总结为
        # 这是总原则：层数不为0则需要输出
        # 剩下的是depth的变化
        # 状态变化代表外层的变化，再
        res = ''
        depth = 0

        for c in s:
            if c == ')':
                depth -= 1

            if depth > 0:
                res += c

            if c == '(':
                depth += 1

        return res