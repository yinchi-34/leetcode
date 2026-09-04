class Solution:
    def firstStableIndex(self, nums: list[int], k: int) -> int:
        """
        暴力版本：完全按照题意，循环两次，逐个减少。

        在累计的过程中，有重复的计算，即记录当前状态的前半段的最大值以及最小值——前缀预处理

        其中最大值的前缀预处理和第二次循环找符合题意的最小稳定下标的顺序是一致的。
        故只需要记录从右往左的最小值。
        """
        n = len(nums)

        suf_min = [0] * n
        suf_min[n - 1] = nums[-1]
        for i in range(n - 2, -1, -1):
            suf_min[i] = min(suf_min[i + 1], nums[i])

        per_max = nums[0]
        for i, x in enumerate(nums):
            if x > per_max:
                per_max = x
            if per_max - suf_min[i] <= k:
                return i

        return -1