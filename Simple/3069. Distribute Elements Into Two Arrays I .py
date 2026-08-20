class Solution:
    def resultArray(self, nums: List[int]) -> List[int]:
        a1, a2 = [nums[0]], [nums[1]]

        for n in nums[2:]:
            if a1[-1] > a2[-1]:
                a1.append(n)
            else:
                a2.append(n)

        return a1 + a2
