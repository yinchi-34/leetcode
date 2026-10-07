class Solution:
    def removeInvalidParentheses(self, s: str) -> list[str]:
        def is_valid(s: str) -> bool:
            left = 0
            for ch in s:
                if ch == "(":
                    left += 1
                elif ch == ")":
                    if left == 0:
                        return False
                    left -= 1
            return left == 0

        cur = {s}
        while True:
            res = [t for t in cur if is_valid(t)]
            if res:
                return res

            nxt = set()
            for t in cur:
                for i, ch in enumerate(t):
                    if ch == '(' or ch == ')':
                        nxt.add(t[:i] + t[i + 1:])

            cur = nxt