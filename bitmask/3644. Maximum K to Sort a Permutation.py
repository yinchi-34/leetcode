class Solution:
    def sortPermutation(self, nums: List[int]) -> int:
        k = -1
        n = len(nums)

        for i in range(n):
            if nums[i] == i:
                continue
            k &= nums[i]

        return 0 if k == -1 else k