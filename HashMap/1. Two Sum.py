class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        counts = defaultdict(int)

        for i in range(len(nums)):
            remains = target - nums[i]
            if remains in counts:
                return [counts[remains], i]
            counts[nums[i]] = i
        return []