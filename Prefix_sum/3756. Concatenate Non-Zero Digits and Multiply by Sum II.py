"""
前缀和
三前缀和：
    1.预处理取模
    2.维护前缀数组
    3.维护字符串前缀包含的非零数字的数量

利用前缀和求差
"""
from typing import List

# 设置取模初始值
MOD = 10 ** 9 + 7
# 预处理取模
pow10 = [1] * 100001
for i in range(1, 100001):
    pow10[i] = pow10[i - 1] * 10 % MOD


class Solution:
    def sumAndMultiply(self, s: str, queries: List[List[int]]) -> List[int]:
        # 获取前缀数组
        n = len(s)
        sum = [0] * (n + 1)
        x = [0] * (n + 1)
        cnt = [0] * (n + 1)

        for i, c in enumerate(s):
            d = int(c)
            sum[i + 1] = sum[i] + d
            x[i + 1] = (x[i] * 10 + d) % MOD if d > 0 else x[i]
            cnt[i + 1] = cnt[i] + (d > 0)

        m = len(queries)
        res = [0] * m
        for i in range(m):
            l = queries[i][0]
            r = queries[i][1] + 1
            length = cnt[r] - cnt[l]
            res[i] = (x[r] - x[l] * pow10[length]) * (sum[r] - sum[l]) % MOD

        return res