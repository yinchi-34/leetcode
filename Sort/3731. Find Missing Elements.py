class Solution:
    def findMissingElements(self, nums: List[int]) -> List[int]:
        res = []
        nums.sort()

        max_num = max(nums)
        min_num = min(nums)
        for i in range(min_num, max_num + 1):
            if i not in nums:
                res.append(i)

        return res