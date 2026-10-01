class Solution:
    def isValid(self, s: str) -> bool:
        parentheses = {
            '{': '}',
            '[': ']',
            '(': ')'
        }

        stack = []

        for a in s:
            # print(stack, a)
            if a in ['(', '{', '[']:
                stack.append(a)
            elif a in [')', '}', ']']:
                if not stack or parentheses[stack[-1]] != a:
                    # 情况一：为空，没有可匹配的；情况二：有，但是匹配失败
                    return False
                elif parentheses[stack[-1]] == a:
                    # print(stack, parentheses[stack[-1]], a)
                    stack.pop()
                    # print(stack)

        return True if stack == [] else False
