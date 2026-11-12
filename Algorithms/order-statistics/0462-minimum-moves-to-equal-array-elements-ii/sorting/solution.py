class Solution:
    def minMoves2(self, nums: list[int]) -> int:
        """Find a median by sorting and return the total distance of all
        elements to it. sorted() returns a new list, so nums is unchanged."""
        a = sorted(nums)
        x = a[len(a) // 2]
        return sum(abs(num - x) for num in a)
