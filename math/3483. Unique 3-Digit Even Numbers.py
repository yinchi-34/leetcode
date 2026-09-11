class Solution:
    def totalNumbers(self, digits: List[int]) -> int:
        st = set()
        for a, b, c in permutations(digits, 3):
            if c and a % 2 == 0:
                st.add(c * 100 + b * 10 + a)

        return len(st)