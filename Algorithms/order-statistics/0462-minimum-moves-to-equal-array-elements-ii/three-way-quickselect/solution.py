import random


def partition3(a: list[int], l: int, r: int) -> tuple[int, int]:
    """Rearrange a[l..r] around the pivot x = a[r] into three parts:
    a[l..lt-1] < x, a[lt..gt] == x, a[gt+1..r] > x, and return (lt, gt)."""
    x = a[r]
    lt, i, gt = l, l, r
    while i <= gt:
        if a[i] < x:
            a[lt], a[i] = a[i], a[lt]
            lt += 1
            i += 1
        elif a[i] > x:
            a[i], a[gt] = a[gt], a[i]
            gt -= 1  # the element swapped in from gt has not been examined yet
        else:
            i += 1
    return lt, gt


def random_partition3(a: list[int], l: int, r: int) -> tuple[int, int]:
    """Move a uniformly random element of a[l..r] to a[r] and partition."""
    p = random.randint(l, r)
    a[p], a[r] = a[r], a[p]
    return partition3(a, l, r)


def quick_select(a: list[int], l: int, r: int, k: int) -> int:
    """Return the element that would be at index k if a[l..r] were sorted.
    Reorders a[l..r] in place.

    The Go version recurses; here the tail recursion is a loop, because
    CPython does not eliminate tail calls and limits recursion depth.
    """
    while True:
        lt, gt = random_partition3(a, l, r)
        if k < lt:
            r = lt - 1
        elif k > gt:
            l = gt + 1
        else:
            return a[k]


class Solution:
    def minMoves2(self, nums: list[int]) -> int:
        """Move every element to a median found by quick_select.
        Reorders nums in place."""
        x = quick_select(nums, 0, len(nums) - 1, len(nums) // 2)
        return sum(abs(num - x) for num in nums)
