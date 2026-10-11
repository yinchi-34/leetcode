class Solution:
    def sumOfSquares(self, nums: List[int]) -> int:
        # 任何因子的方法都可以考虑枚举的方法，减少时间复杂度
        n = len(nums)
        res = 0

        for i in range(1, isqrt(n) + 1):
            if n % i == 0:
               res += nums[i - 1] ** 2
               if i * i < n:
                res += nums[n // i - 1] ** 2
        return res