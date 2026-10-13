class Solution:
    def maxAlternatingSum(self, nums: List[int]) -> int:
        # deal = [abs(num) for num in nums]
        # plus = []
        # reduces = []
        # n = len(nums)
        # deal.sort(reverse=True)

        # if len(deal) % 2 == 0:
        #     plus, reduces = deal[0:n//2], deal[n//2:]
        # else:
        #     plus, reduces = deal[0:n//2 + 1], deal[n//2 + 1:]

        # return  sum(i**2 for i in plus) - sum(j**2 for j in reduces)
        sq = sorted(x * x for x in nums)
        k = len(nums) // 2
        return sum(sq[k:]) - sum(sq[:k])