class Solution:
    def minOperations(self, s: str) -> int:
        """
        这个问题：
        - 最小值是否已经位于左边界？
        - 最大值是否已经位于右边界？
        - 极值是否出现在内部？

        """
        n = len(s)
        mx, mn = max(s), min(s)

        if all(a <= b for a, b in zip(s, s[1:])):
            return 0
        if n == 2:
            return -1
        if s[0] == mn or s[-1] == mx:
            return 1
        for i in range(1, n - 1):
            if s[i] == mn or s[i] == mx:
                return 2

        return 3