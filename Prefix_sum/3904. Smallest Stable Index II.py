class Solution:
    def firstStableIndex(self, nums: list[int], k: int) -> int:
        """
        与3903的变化在于数量的变化，从时间复杂度考虑，则O(n^2)已经太多，故不用暴力版，需要考虑前缀预处理
        即昨天的写法

        """

        n = len(nums)

        suf_nums = [0] * n
        suf_nums[-1] = nums[-1]
        for i in range(n - 2, -1, -1):
            suf_nums[i] = min(suf_nums[i + 1], nums[i])

        per_max = nums[0]
        for i in range(n):
            if nums[i] > per_max:
                per_max = nums[i]
            if per_max - suf_nums[i] <= k:
                return i

        return -1