class Solution:
    def minAddToMakeValid(self, s: str) -> int:
        left = 0
        res = 0
        for c in s:
            if c == '(':
                left += 1
            elif left > 0:
                left -= 1
            else:
                res += 1

        return left + res