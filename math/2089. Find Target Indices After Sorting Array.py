class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        less = 0
        equal = 0

        for num in nums:
            if num < target:
                less += 1
            elif num == target:
                equal += 1

        return list(range(less, less + equal))