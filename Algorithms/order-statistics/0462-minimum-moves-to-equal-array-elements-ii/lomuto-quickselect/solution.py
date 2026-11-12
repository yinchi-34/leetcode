import random


def partition(a: list[int], l: int, r: int) -> int:
    """The two-way partition of CLRS §7.1 (Lomuto) with pivot x = a[r]:
    return q with a[l..q-1] <= x, a[q] = x, and a[q+1..r] > x."""
    x = a[r]
    i = l - 1
    for j in range(l, r):
        if a[j] <= x:
            i += 1
            a[i], a[j] = a[j], a[i]
    a[i + 1], a[r] = a[r], a[i + 1]
    return i + 1


def random_partition(a: list[int], l: int, r: int) -> int:
    p = random.randint(l, r)
    a[p], a[r] = a[r], a[p]
    return partition(a, l, r)


def quick_select(a: list[int], l: int, r: int, k: int) -> int:
    """Return the element that would be at index k if a[l..r] were sorted.
    Reorders a[l..r] in place. Every element equal to the pivot goes left,
    so many duplicates make it quadratic.

    The tail recursion of the Go version is a loop here: with many
    duplicates the recursion would be about n levels deep, past CPython's
    default limit of 1000.
    """
    while True:
        q = random_partition(a, l, r)
        if q == k:
            return a[q]
        if q < k:
            l = q + 1
        else:
            r = q - 1


class Solution:
    def minMoves2(self, nums: list[int]) -> int:
        """Move every element to a median found by quick_select.
        Reorders nums in place."""
        x = quick_select(nums, 0, len(nums) - 1, len(nums) // 2)
        return sum(abs(num - x) for num in nums)
