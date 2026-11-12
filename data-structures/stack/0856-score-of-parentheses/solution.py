class Solution:
    def scoreOfParentheses(self, s: str) -> int:
        stack = [0]  # stack[-1] is the score of the innermost open group
        for c in s:
            if c == "(":
                stack.append(0)
            else:
                top = stack.pop()
                stack[-1] += max(2 * top, 1)
        return stack[0]
