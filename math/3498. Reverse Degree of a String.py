class Solution:
    def reverseDegree(self, s: str) -> int:
        return sum((ord('z') - ord(word) + 1) * (i + 1) for i, word in enumerate(s))