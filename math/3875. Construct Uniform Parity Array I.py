class Solution:
    """
    nums1分为两个情况：
    · 全部奇数或者偶数——不用操作
    · 偶数奇数都有——以其中一个奇数作为转换器，所有偶数变为奇数
    所以都是true
    """
    def uniformArray(self, nums1: list[int]) -> bool:
        return True