class Solution:
    def maxSum(self, nums: List[int], k: int) -> int:
        MOD = 10 ** 9 + 7
        counts = [0] * 30

        for num in nums:
            for i in range(30):
                if num & (1 << i):
                    counts[i] += 1

        res = 0
        for _ in range(k):
            cur = 0
            for j in range(29, -1, -1):
                if counts[j] > 0:
                    cur |= (1 << j)
                    counts[j] -= 1

            res = (res + cur * cur) % MOD

        return res