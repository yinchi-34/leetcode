# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def nodesBetweenCriticalPoints(self, head: Optional[ListNode]) -> List[int]:
        res = [-1, -1]

        nums = []
        idx = 0
        ima = []
        while head:
            nums.append(head.val)
            # print(nums,ima, idx)

            if idx >= 2:
                print(nums[idx - 2], nums[idx - 1], nums[idx])
                if nums[idx - 2] > nums[idx - 1] and nums[idx] > nums[idx - 1]:
                    ima.append(idx - 1)
                elif nums[idx - 2] < nums[idx - 1] and nums[idx] < nums[idx - 1]:
                    ima.append(idx - 1)
            idx += 1
            head = head.next

        if len(ima) >= 2:
            maxDistance = ima[-1] - ima[0]
            minDistance = min(abs(ima[i] - ima[i - 1]) for i in range(1, len(ima)))
            res = [minDistance, maxDistance]

        return res