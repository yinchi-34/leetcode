class Solution:
    def smallestIndex(self, nums: List[int]) -> int:
        def get_digit_num(num: int) -> int:
            total = 0
            while num:
                total += num % 10
                num = num // 10
            return total

        return next((i for i in range(len(nums)) if i == get_digit_num(nums[i])), -1) 