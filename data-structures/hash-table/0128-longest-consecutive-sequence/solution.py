class Solution:
    def longestConsecutive(self, nums: list[int]) -> int:
        res = 0
        num_set = set(nums)
        for num in num_set:
            if num - 1 not in num_set:  # num starts a run
                current = num
                while current + 1 in num_set:
                    current += 1
                res = max(res, current - num + 1)
        return res
